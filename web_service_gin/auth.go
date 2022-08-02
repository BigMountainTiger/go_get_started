package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
)

// https://stackoverflow.com/questions/41077953/how-to-verify-jwt-signature-with-jwk-in-go
// https://github.com/MicahParks/keyfunc

var expected_client_id string
var expected_scope string
var auth_key_map map[string]interface{}

func init_auth() bool {

	expected_client_id = os.Getenv("CLIENT_ID")
	expected_scope = os.Getenv("SCOPE")

	url := os.Getenv("AUTH_URL")
	resp, err := http.Get(url)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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

func authenticate(c *gin.Context) (bool, error) {
	authorization := c.Request.Header.Get("authorization")
	if strings.TrimSpace(authorization) == "" {
		return false, errors.New("no authorization header")
	}

	token_str := strings.Replace(authorization, "Bearer ", "", 1)
	token, err := jwt.Parse(token_str, _get_key)
	if err != nil {
		return false, err
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims == nil {
		return false, errors.New("no authorization claims")
	}

	client_id := claims["client_id"].(string)
	scope := claims["scope"].([]interface{})

	if client_id != expected_client_id {
		return false, errors.New("no matching client_id")
	}

	scope_matched := false
	for i := 0; i < len(scope); i++ {
		if scope[i] == expected_scope {
			scope_matched = true
			break
		}
	}

	if !scope_matched {
		return false, errors.New("no matching auth scope")
	}

	return true, nil
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
