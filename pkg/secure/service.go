package secure

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

// Cookie provides gorilla secure cookie functionality
type Cookie struct {
	SecureCookie *securecookie.SecureCookie
}

// CookieConfig provided options for the Cookie
type CookieConfig struct {
	Hmac      string
	HmacBlock string
}

// NewCookie returns a new Secure Cookie
// config is the hmackey and hmacblock key
func NewCookie(config CookieConfig) (*Cookie, error) {
	s, err := createEmptySecureCookie(config.Hmac, config.HmacBlock)
	if err != nil {
		return nil, fmt.Errorf("Cannot create new Secure Cookie: %v", err)
	}
	return &Cookie {
		SecureCookie: s,
	}, nil
}

// Encode encodes the value map and returns an http secure cookie
// name is the cookie name
// value is a map of values to encode
func (c Cookie) Encode(name string, value map[string]string) (*http.Cookie, error) {
	encoded, err := c.SecureCookie.Encode(name, value)
	if err != nil {
		return nil, fmt.Errorf("Value cannot be Encoded in the secure cookie: %v", err)
	}
	return &http.Cookie{
		Name:     name,
		Value:    encoded,
		Secure:   true,
		HttpOnly: true,
	}, nil
}

// Validate compares between the Value in the Secure Cookie vs input value
// name is the cookie name
// encodedValue is the encoded string
// input is a map of values to compare
func (c Cookie) Validate(name string, encodedValue string, input map[string]string) (bool, error) {
	var cookieValue map[string]string
	err := c.SecureCookie.Decode(name, encodedValue, &cookieValue)
	if err != nil {
		return false, fmt.Errorf("Cannot read cookie encoded value: %v", err)
	}
	fmt.Printf("decrypted cookie: %v\n", cookieValue)
	if input == nil || cookieValue == nil || input["payload"] != cookieValue["payload"] {
		return false, fmt.Errorf("Input doesn't match cookie value")
	}
	return true, nil
}

// createEmptySecureCookie creates a secure gorilla cookie
func createEmptySecureCookie(hmac, hmacblock string) (*securecookie.SecureCookie, error) {
	hmacKey, err := base64.URLEncoding.DecodeString(hmac)
	if err != nil {
		return nil, fmt.Errorf("Hmac Key invalid: %v", err)
	}
	hmacBlockKey, err := base64.URLEncoding.DecodeString(hmacblock)
	if err != nil {
		return nil, fmt.Errorf("HmacBlock Key invalid: %v", err)
	}
	secureCookie := securecookie.New(hmacKey, hmacBlockKey)
	return secureCookie, nil
}
