package vsphere

import (
	"context"
	"encoding/json"
	"net/http"
)

type DSMSessionResponse struct {
	Token string `json:"token"`
}

// TODO: This is a very simple implementation to test.
// On production we will have TLS, timeout, authentication, etc so this client
// interface should be much more complete
func GetDSMToken(ctx context.Context, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	token := &DSMSessionResponse{}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(token); err != nil {
		return "", err
	}
	return token.Token, nil
}
