package broker

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	rabbithole "github.com/michaelklishin/rabbit-hole"

	"github.com/pivotal-cf/brokerapi"
)

func (b RabbitMQServiceBroker) Bind(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails, asyncAllowed bool) (brokerapi.Binding, error) {
	response, err := b.client.PutUser(bindingID, rabbithole.UserSettings{
		Password: generateString(24),
	})
	if err != nil {
		return brokerapi.Binding{}, err
	}
	if response != nil && response.StatusCode == http.StatusNoContent {
		return brokerapi.Binding{}, brokerapi.ErrBindingAlreadyExists
	}
	return brokerapi.Binding{}, nil
}

func generateString(size int) string {
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		// TODO ?
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(rb)
}
