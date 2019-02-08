package broker

import (
	"context"

	"github.com/pivotal-cf/brokerapi"
)

func (b RabbitMQServiceBroker) Services(ctx context.Context) ([]brokerapi.Service, error) {
	return []brokerapi.Service{
		brokerapi.Service{
			ID:          b.Config.ServiceConfig.UUID,
			Name:        b.Config.ServiceConfig.Name,
			Description: b.Config.ServiceConfig.OfferingDescription,
		},
	}, nil
}
