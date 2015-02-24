package stash

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"testing"

	"code.google.com/p/go.crypto/ssh"
)

var (
	k = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC5aE+Qmt7Ky+IZlex7VFq1KCiatq2ZFkQPdOLDg/IVPdREPa5DcsacgfrUod6LIGvyKEK89KIpAsHTxjF3+R2sSVE79fWLyvcI+Ulw/jnVu7lldVlm3tvp80KZSsCoZzxoeSjhOKYPrjmExy+ztYQt25Ev3ZF2G5DynyWHU5T038sjir3ooYgqNPIMkkyeumJ8EZTGftL+GpTiefIAp9dyYUBPr6WmB8AugOM+mT4E8HH3Ssmxcgi5fWPMCfdnJFpFXZCcXRuGga/CnTd52e1UhvU5DTGHxWnAMcs/t71AYN1RARP8ASAsGfsKtiYQKB2+1yOR2SzWOpsXoFnmd9CZ"
)

func GenerateNewPublicKey() string {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("Failed to generate new RSA key: ", err)
	}
	pk, err := ssh.NewPublicKey(&key.PublicKey)
	if err != nil {
		log.Fatal("Failed to generate public key: ", err)
	}
	return string(ssh.MarshalAuthorizedKey(pk))
}

func TestKeyCreateUpdate(t *testing.T) {
	key, err := client.Keys.CreateUpdate(k)
	if err != nil {
		t.Errorf("Unexpected error on `client.Keys.Create()`, got %v", err)
	}
	if key.Text == "" {
		t.Error("Expected key text, got nothing")
	}
}

func TestKeyCreateUpdateInvalidKey(t *testing.T) {
	key, err := client.Keys.CreateUpdate("key")
	if err == nil {
		t.Error("Expected error on `client.Keys.CreateUpdate()`")
	}
	if key != nil {
		t.Errorf("Did not expect key, got %v", key)
	}
}

func TestKeyFind(t *testing.T) {
	key, err := client.Keys.Find(k)
	if err != nil {
		t.Errorf("Unexpected error on `client.Keys.Find()`, got %v", err)
	}
	if key.Text == "" {
		t.Error("Expected key text, got nothing")
	}
}

func TestKeyFindNotFound(t *testing.T) {
	key, err := client.Keys.Find("key")
	if err != ErrNotFound {
		t.Error("Expected not found error on `client.Keys.Find()`")
	}
	if key != nil {
		t.Errorf("Did not expect key, got %v", key)
	}
}

func TestKeyCreateInvalidKey(t *testing.T) {
	key, err := client.Keys.Create("key")
	if err == nil {
		t.Error("Expected error on `client.Keys.Create()`")
	}
	if key != nil {
		t.Errorf("Did not expect key, got %v", key)
	}
}

func TestKeyCreate(t *testing.T) {
	newKey := GenerateNewPublicKey()
	key, err := client.Keys.Create(newKey)
	if err != nil {
		t.Error("Unexpected error on `client.Keys.Create()`, got %v", err)
	}
	if key.Text == "" {
		t.Errorf("Expect key text, got nothin")
	}
}
