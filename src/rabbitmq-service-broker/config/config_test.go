package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "rabbitmq-service-broker/config"
)

var _ = Describe("Config", func() {
	It("read a minimal config", func() {
		conf, err := Read(fixture("minimal.yml"))
		Expect(err).NotTo(HaveOccurred())
		Expect(conf).To(Equal(Config{
			Service: ServiceConfig{
				UUID:     "fake-service-uuid",
				Name:     "fake-service-name",
				PlanUUID: "fake-plan-uuid",
				Username: "fake-service-username",
				Password: "fake-service-password",
			},
			RabbitMQ: RabbitMQConfig{
				Hosts:         []string{"fake-host-1", "fake-host-2"},
				Administrator: RabbitMQCredentials{"fake-rmq-user", "fake-rmq-password"},
			},
		}))
	})
})
