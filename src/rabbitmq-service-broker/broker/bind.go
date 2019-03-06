package broker

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	rabbithole "github.com/michaelklishin/rabbit-hole"

	"github.com/pivotal-cf/brokerapi"
)

func (b RabbitMQServiceBroker) Bind(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails, asyncAllowed bool) (brokerapi.Binding, error) {
	tags := b.cfg.RabbitMQ.RegularUserTags
	if tags == "" {
		tags = "policymaker,management"
	}

	userSettings := rabbithole.UserSettings{
		Password: generateString(24),
		Tags:     tags,
	}
	username := bindingID
	err := b.createUser(username, userSettings)
	if err != nil {
		return brokerapi.Binding{}, err
	}

	err = b.assignPermissionsToUser(instanceID, bindingID)
	if err != nil {
		return brokerapi.Binding{}, err
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
