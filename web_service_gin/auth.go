package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
)

// https://stackoverflow.com/questions/41077953/how-to-verify-jwt-signature-with-jwk-in-go
// https://github.com/MicahParks/keyfunc

var auth_key_set jwk.Set

func load_auth_key() bool {

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

	auth_key_set = *set

	return true
}

func auth_key() (string, error) {

	keys := auth_key_set.LookupKeyID("ABCD")
	if len(keys) == 0 {
		return "", errors.New("key is not found")
	}

	log.Println(keys[0])

	return "OK", nil
}

func get_key(token *jwt.Token) (interface{}, error) {

	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("jwt token has no kid")
	}

	keys := auth_key_set.LookupKeyID(kid)
	if len(keys) == 0 {
		return nil, errors.New("jwk is not found")
	}

	key, err := keys[0].Materialize()
	if err != nil {
		return nil, err
	}

	return key, nil
}

func authenticate(c *gin.Context) (bool, error) {
	authorization := c.Request.Header.Get("authorization")
	if strings.TrimSpace(authorization) == "" {
		return false, errors.New("no authorization header")
	}

	token_str := strings.Replace(authorization, "Bearer ", "", 1)
	token, err := jwt.Parse(token_str, get_key)
	if err != nil {
		return false, err
	}

	log.Println(token)

	return true, nil
}

func authenticate_bak(c *gin.Context) (bool, error) {

	authorization := c.Request.Header.Get("authorization")
	if strings.TrimSpace(authorization) == "" {
		return false, errors.New("no authorization header")
	}

	token_str := strings.Replace(authorization, "Bearer ", "", 1)
	token, _, err := new(jwt.Parser).ParseUnverified(token_str, jwt.MapClaims{})

	if err != nil {
		return false, err
	}

	claims := token.Claims.(jwt.MapClaims)
	kid := token.Header["kid"]
	client_id := claims["client_id"]
	scope := claims["scope"]

	if kid == nil || client_id == nil || scope == nil {
		return false, errors.New("incomplete authorization header")
	}

	log.Println(scope)

	return true, nil
}
