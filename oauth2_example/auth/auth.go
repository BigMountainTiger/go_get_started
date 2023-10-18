package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/pkg/browser"
	"song.com/go_get_started/oauth2_example/utils"
)

var oauthTokenURL string = os.Getenv("oauthTokenURL")
var oauth2URL string = os.Getenv("oauth2URL")
var callbackProtocol string = os.Getenv("callbackProtocol")
var callbackHost string = os.Getenv("callbackHost")
var callbackPort string = os.Getenv("callbackPort")

func AuthByMicrosoft() {
	token, _, err := AuthenticateSSO()
	if err != nil {
		log.Fatal(err)
	}

	err = utils.Save_file_to_home("token", []byte(token))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Token saved")
}

func AuthenticateSSO() (string, string, error) {

	code, err := get_code()
	if err != nil {
		return "", "", err
	}
	if code == "" {
		return "", "", errors.New("get an empty code")
	}

	token, err := get_token(code)
	if err != nil {
		return "", "", err
	}

	payload, err := parse_token(token)
	if err != nil {
		return "", "", err
	}

	return token, payload, nil
}

func parse_token(token string) (string, error) {

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid auth token received")
	}

	_, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}

	payload, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	return string(payload), nil
}

func get_token(code string) (string, error) {
	tokenURL := fmt.Sprintf("%s/users/token", oauthTokenURL)
	client := &http.Client{}

	jsonData, _ := json.Marshal(map[string]string{
		"code": code,
	})

	request, err := http.NewRequest("POST", tokenURL, strings.NewReader(string(jsonData)))
	if err != nil {
		return "", err
	}

	request.Header.Add("content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("SSO token exchanged failed: %s", err)
	}
	defer response.Body.Close()

	type tokenResponse struct {
		Token string `json:"token"`
	}

	var tr tokenResponse
	err = json.NewDecoder(response.Body).Decode(&tr)
	if err != nil {
		return "", err
	}

	return tr.Token, nil
}

func get_code() (string, error) {

	respCh := make(chan url.Values)
	sigintCh := make(chan os.Signal, 1)
	signal.Notify(sigintCh, os.Interrupt)
	defer signal.Stop(sigintCh)

	callbackUrl := fmt.Sprintf("%s:%s", callbackHost, callbackPort)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		respCh <- query

		http.Redirect(w, r, "https://autobahn.comcast.com/welcome", http.StatusFound)
	})

	listener, err := net.Listen(callbackProtocol, callbackUrl)
	if err != nil {
		return "", fmt.Errorf("could not listen for callback: %s", err)
	}
	defer listener.Close()

	go func() {
		err := http.Serve(listener, nil)
		if err != nil && err != http.ErrServerClosed {
			respCh <- nil
		}
	}()

	err = browser.OpenURL(oauth2URL)
	if err != nil {
		return "", fmt.Errorf("could not open browser: %s", err)
	}

	var responseValues url.Values
	select {
	case s := <-respCh:
		responseValues = s
	case <-sigintCh:
		return "", errors.New("operation canceled")
	case <-time.After(2 * time.Minute):
		return "", errors.New("time out waiting for SSO callback")
	}

	if error_description := responseValues.Get("error_description"); error_description != "" {
		return "", fmt.Errorf("error performing SSO: %s", error_description)
	}

	return responseValues.Get("code"), nil

}
