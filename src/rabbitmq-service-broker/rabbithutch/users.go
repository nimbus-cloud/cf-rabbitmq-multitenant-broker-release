package rabbithutch

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	rabbithole "github.com/michaelklishin/rabbit-hole"
	"github.com/pivotal-cf/brokerapi"
)

func (r *rabbitHutch) CreateUser(username, vhost, tags string) (string, error) {
	if tags == "" {
		tags = "policymaker,management"
	}

	password, err := generatePassword()
	if err != nil {
		return "", err
	}

	userSettings := rabbithole.UserSettings{
		Password: password,
		Tags:     tags,
	}

	response, err := r.client.PutUser(username, userSettings)
	if err != nil {
		return "", err
	}
	if response != nil && response.StatusCode == http.StatusNoContent {
		return "", brokerapi.ErrBindingAlreadyExists
	}

	permissions := rabbithole.Permissions{Configure: ".*", Write: ".*", Read: ".*"}
	err = validateResponse(r.client.UpdatePermissionsIn(vhost, username, permissions))
	if err != nil {
		r.DeleteUser(username)
		return "", err
	}

	return password, nil
}

func validateResponse(resp *http.Response, err error) error {
	if err != nil {
		return err
	}

	if resp != nil && (resp.StatusCode < http.StatusOK || resp.StatusCode > 299) {
		return fmt.Errorf("http request failed with status code: %v", resp.StatusCode)
	}

	return nil
}

func generatePassword() (string, error) {
	rb := make([]byte, 24)
	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(rb), nil
}

func (r *rabbitHutch) DeleteUser(username string) error {
	resp, err := r.client.DeleteUser(username)
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return brokerapi.ErrBindingDoesNotExist
	}
	if rabbitErr, ok := err.(rabbithole.ErrorResponse); ok && rabbitErr.StatusCode == http.StatusNotFound {
		return brokerapi.ErrBindingDoesNotExist
	}
	if err := validateResponse(resp, err); err != nil {
		return err
	}
	return nil
}

func (r *rabbitHutch) Unbind(username string) error {
	defer func() {
		conns, _ := r.client.ListConnections()
		for _, conn := range conns {
			if conn.User == username {
				r.client.CloseConnection(conn.Name)
			}
		}
	}()

	if err := r.DeleteUser(username); err != nil {
		return err
	}

	return nil
}
