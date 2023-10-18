package auth

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"song.com/go_get_started/oauth2_example/utils"
)

var mapUrl = os.Getenv("mapUrl")
var auth_key_map map[string]interface{}

func VerifyByMicrosoft() {
	load_keys()
	token, err := load_token()
	if err != nil {
		log.Fatalln(err)
	}

	claims, err := verify(token)
	if err != nil {
		log.Fatalln(err)
	}

	field := "onpremisessamaccountname"
	ntid := (*claims)[field]
	if ntid == nil {
		log.Fatalln("no ntid")
	}

	log.Println(ntid.(string))

}

func load_token() (string, error) {

	tb, err := utils.Read_file_from_home("token")
	if err != nil {
		return "", err
	}
	if len(tb) == 0 {
		return "", errors.New("no token received")
	}

	return string(tb), nil
}

func load_keys() bool {

	resp, err := http.Get(mapUrl)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	auth_key_string := string(body)
	set, err := jwk.ParseString(auth_key_string)
	if err != nil {
		return false
	}

	key_map := make(map[string]interface{})
	for i := 0; i < set.Len(); i++ {
		key, _ := set.Get(i)

		var pubkey interface{}
		err := key.Raw(&pubkey)
		if err == nil {
			key_map[key.KeyID()] = pubkey
		}
	}

	auth_key_map = key_map

	return true
}

func verify(token_str string) (*jwt.MapClaims, error) {

	token, err := jwt.Parse(token_str, _get_key)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || len(claims) == 0 {
		return nil, errors.New("no authorization claims")
	}

	return &claims, nil
}

func _get_key(token *jwt.Token) (interface{}, error) {

	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("jwt token has no kid")
	}

	pubkey := auth_key_map[kid]
	if pubkey == nil {
		return nil, errors.New("jwk is not found")
	}

	return pubkey, nil
}
