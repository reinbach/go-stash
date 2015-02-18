package oauth1

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"testing"
)

var (
	privateKey  string
	consumerKey string
	c           Consumer
)

func init() {
	consumerKey = os.Getenv("STASH_CONSUMER_KEY")
	privateKey = os.Getenv("STASH_PRIVATE_KEY")

	switch {
	case len(consumerKey) == 0:
		panic(errors.New("must set the STASH_CONSUMER_KEY environment variable"))
	case len(privateKey) == 0:
		panic(errors.New("must set the STASH_PRIVATE_KEY environment variable"))
	}

	c = Consumer{
		ConsumerKey:           consumerKey,
		ConsumerSecret:        "",
		ConsumerPrivateKeyPem: privateKey,
		RequestTokenURL:       OOB,
	}

}

// Test Nonce
func TestNonce(t *testing.T) {
	n1 := nonce()
	n2 := nonce()

	if n1 == n2 {
		t.Error("Nonce not very nonce'y")
	}
}

// Test default headers
func TestHeaders(t *testing.T) {
	consumerKey := "consumerkey"
	h := headers(consumerKey)

	if h["oauth_consumer_key"] != consumerKey {
		t.Error("Wrong oauth_consumer_key is set in headers")
	}
	if h["oauth_signature_method"] != "RSA-SHA1" {
		t.Error("Wrong oauth_signature_method is set in headers")
	}
	if h["oauth_version"] != "1.0" {
		t.Error("Wrong oauth_version is set in headers")
	}
}

// Test generation of RSA SHA1 signature
func TestSign(t *testing.T) {
	msg := "HelloWorld"
	s := sign(msg, privateKey)

	if len(s) == 0 {
		t.Error("RSA SHA1 signature failed")
	}
}

// Test Signing of Params
func TestSignParams(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	oauth_token := "o-token"
	oauth_secret := "o-secret"
	oauth_params := map[string]string{}
	oauth_params["test"] = "world"

	token := NewAccessToken(oauth_token, oauth_secret, oauth_params)
	err := c.SignParams(r, token, nil)
	if err != nil {
		t.Errorf("SignParams error'd out: ", err)
	}

	err = c.SignParams(r, token, oauth_params)
	if err != nil {
		t.Errorf("SignParams error'd out: ", err)
	}
}

func TestSignParamsQueryParams(t *testing.T) {
	r, _ := http.NewRequest("GET", "/?test=world", nil)
	oauth_token := "o-token"
	oauth_secret := "o-secret"

	token := NewAccessToken(oauth_token, oauth_secret, nil)
	err := c.SignParams(r, token, nil)
	if err != nil {
		t.Errorf("SignParams error'd out: ", err)
	}
}

func TestSignParamsNoHeader(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	r.Header = nil
	oauth_token := "o-token"
	oauth_secret := "o-secret"

	token := NewAccessToken(oauth_token, oauth_secret, nil)
	err := c.SignParams(r, token, nil)
	if err != nil {
		t.Errorf("SignParams error'd out: ", err)
	}
}

func TestSignParamsPost(t *testing.T) {
	r, _ := http.NewRequest("POST", "", nil)
	r.Form = url.Values{"test": {"world"}}
	oauth_token := "o-token"
	oauth_secret := "o-secret"

	token := NewAccessToken(oauth_token, oauth_secret, nil)
	err := c.SignParams(r, token, nil)
	if err != nil {
		t.Errorf("SignParams error'd out: ", err)
	}

	if r.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Request Header needs to be application/json, got: %v",
			r.Header.Get("Content-Type"))
	}
}

func TestConsumerSign(t *testing.T) {
	r, _ := http.NewRequest("Get", "", nil)
	oauth_token := "o-token"
	oauth_secret := "o-secret"

	token := NewAccessToken(oauth_token, oauth_secret, nil)
	err := c.Sign(r, token)
	if err != nil {
		t.Errorf("Sign error'd out: ", err)
	}
}

// Test escape function
func TestEscape(t *testing.T) {
	s := escape("HelloWorld")
	if s != "HelloWorld" {
		t.Error("String was not escaped")
	}

	s = escape("Hello World")
	if s != "Hello%20World" {
		t.Error("String was not escaped properly")
	}
}

// Test isEscapable function
func TestIsEscapable(t *testing.T) {
	if isEscapable("s"[0]) == true {
		t.Error("'s' should not be escapable")
	}
	if isEscapable(" "[0]) == false {
		t.Error("' ' should be escapable")
	}
}
