package broker

import (
	"context"
	"fmt"
	"net/http"

	"github.com/michaelklishin/rabbit-hole"
	"github.com/pivotal-cf/brokerapi"
)

func (rabbitmqServiceBroker RabbitMQServiceBroker) Provision(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {

	spec := brokerapi.ProvisionedServiceSpec{}

	if err := rabbitmqServiceBroker.createVhost(instanceID); err != nil {
		return spec, err
	}

	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (rabbitmqServiceBroker *RabbitMQServiceBroker) createVhost(vhost string) error {
	//	logger = logger.Session("create-vhost")
	//	logger.Info("get-vhost")
	if _, err := rabbitmqServiceBroker.client.GetVhost(vhost); err == nil {
		err = brokerapi.ErrInstanceAlreadyExists
		//		logger.Error("get-vhost-failed", err)
		return err
	}
	//	logger.Info("get-vhost-succeeded")

	//	logger.Info("put-vhost")
	resp, err := rabbitmqServiceBroker.client.PutVhost(vhost, rabbithole.VhostSettings{})

	if err != nil {
		//		logger.Error("put-vhost-failed", err)
		return err
	}
	//	logger.Info("put-vhost-succeeded")

	if err := validateResponse(resp); err != nil {
		//		logger.Error("put-vhost-failed", err)
		return err
	}

	return nil
}

func validateResponse(resp *http.Response) error {
	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		return fmt.Errorf("http request failed with status code: %v", resp.StatusCode)
	}
	return nil
}
