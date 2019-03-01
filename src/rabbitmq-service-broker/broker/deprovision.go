package broker

import (
	"context"

	"github.com/pivotal-cf/brokerapi"
)

func (b RabbitMQServiceBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (spec brokerapi.DeprovisionServiceSpec, err error) {
	logger := b.logger.Session("deprovision")
	if _, err := b.client.GetVhost(instanceID); err != nil {
		logger.Info("get-vhost-failed")
		return spec, brokerapi.ErrInstanceDoesNotExist
	}
	if err := b.deleteVhost(instanceID); err != nil {
		return spec, err
	}
	logger.Info("deprovision-succeeded")
	return spec, nil
}
