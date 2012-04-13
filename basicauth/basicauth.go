package basicauth

import (
	"encoding/base64"
	"errors"
	"strings"
)

func Decode(h string) (u string, p string, e error) {
	prefix := "Basic "
	if !strings.HasPrefix(h, prefix) {
		e = errors.New("Bad Request")
		return
	}
	h = h[6:]
	auth, err := base64.StdEncoding.DecodeString(h)
	if err != nil {
		e = errors.New("Bad Request")
		return
	}
	fields := strings.Split(string(auth), ":")
	if len(fields) != 2 {
		e = errors.New("Bad Request")
		return
	}
	u = fields[0]
	p = fields[1]
	return
}
