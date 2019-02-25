package broker

import (
	"context"
	"errors"
	"net/http"

	"github.com/michaelklishin/rabbit-hole"
	"github.com/pivotal-cf/brokerapi"
)

//go:generate counterfeiter -o ./fakes/api_client_fake.go $FILE APIClient

type APIClient interface {
	GetVhost(string) (*rabbithole.VhostInfo, error)
	PutVhost(string, rabbithole.VhostSettings) (*http.Response, error)
}

type RabbitMQServiceBroker struct {
	Config Config
	client APIClient
}

func New(cfg Config, client APIClient) brokerapi.ServiceBroker {
	return RabbitMQServiceBroker{
		Config: cfg,
		client: client,
	}
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) GetInstance(ctx context.Context, instanceID string) (brokerapi.GetInstanceDetailsSpec, error) {
	return brokerapi.GetInstanceDetailsSpec{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) Update(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) LastOperation(ctx context.Context, instanceID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) Bind(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails, asyncAllowed bool) (brokerapi.Binding, error) {
	return brokerapi.Binding{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) Unbind(ctx context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails, asyncAllowed bool) (brokerapi.UnbindSpec, error) {
	return brokerapi.UnbindSpec{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) GetBinding(ctx context.Context, instanceID, bindingID string) (brokerapi.GetBindingSpec, error) {
	return brokerapi.GetBindingSpec{}, errors.New("Not implemented")
}

func (rabbitmqServiceBroker RabbitMQServiceBroker) LastBindingOperation(ctx context.Context, instanceID, bindingID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, errors.New("Not implemented")
}
