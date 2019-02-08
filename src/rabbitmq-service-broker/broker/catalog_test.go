package broker_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi"

	"rabbitmq-service-broker/broker"
)

var _ = Describe("Service Broker", func() {
	It("returns a valid catalog", func() {
		cfg := broker.Config{
			ServiceConfig: broker.ServiceConfig{
				UUID:                "service-id",
				Name:                "service-name",
				OfferingDescription: "service-description",
			},
		}

		broker := broker.New(cfg)
		services, err := broker.Services(context.Background())
		Expect(err).NotTo(HaveOccurred())

		Expect(services).To(Equal([]brokerapi.Service{brokerapi.Service{
			ID:          cfg.ServiceConfig.UUID,
			Name:        cfg.ServiceConfig.Name,
			Description: cfg.ServiceConfig.OfferingDescription,
		}}))
	})
})
