package basicauth

import (
	"encoding/base64"
	"testing"
)

func TestDecodeEmpty(t *testing.T) {
	if _, _, err := Decode(""); err == nil {
		t.Error("Should fail on empty")
	}
}

func TestDecodeMissingPrefix(t *testing.T) {
	if _, _, err := Decode("Bees"); err == nil {
		t.Error("Should fail on missing prefix")
	}
}

func TestDecodeEmptySuffix(t *testing.T) {
	if _, _, err := Decode("Basic "); err == nil {
		t.Error("Should fail on empty suffix")
	}
}

func TestDecodeInvalidSuffix(t *testing.T) {
	if _, _, err := Decode("Basic !*@&#@"); err == nil {
		t.Error("Should fail on invalid suffix")
	}
}

func TestDecodeNoDelimiter(t *testing.T) {
	a := base64.StdEncoding.EncodeToString([]byte("bees"))
	if _, _, err := Decode(a); err == nil {
		t.Error("Should fail on no delmiter")
	}
}

func TestDecodeTooManyDelimiters(t *testing.T) {
	a := base64.StdEncoding.EncodeToString([]byte("b:e:s"))
	if _, _, err := Decode(a); err == nil {
		t.Error("Should fail on too many delimiters")
	}
}

func TestDecodeValidAuth(t *testing.T) {
	u := "ryan"
	p := "clowns"
	a := base64.StdEncoding.EncodeToString([]byte(u + ":" + p))
	username, password, err := Decode("Basic " + a)
	if err != nil {
		t.Error("Error should be nil")
	}
	if username != u {
		t.Errorf("Wrong username. Expected '%s' got '%s'", u, username)
	}
	if password != p {
		t.Errorf("Wrong password. Expected '%s' got '%s'", p, password)
	}
}
