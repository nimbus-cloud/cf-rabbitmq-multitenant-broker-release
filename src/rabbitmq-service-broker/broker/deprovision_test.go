package broker_test

import (
	"context"
	"errors"
	"net/http"

	"rabbitmq-service-broker/broker/fakes"

	rabbithole "github.com/michaelklishin/rabbit-hole"
	"github.com/pivotal-cf/brokerapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Provisioning a RMQ service instance", func() {
	var (
		client *fakes.FakeAPIClient
		broker brokerapi.ServiceBroker
		ctx    context.Context
	)

	When("the instance exists", func() {
		BeforeEach(func() {
			client = new(fakes.FakeAPIClient)
			client.GetVhostReturns(&rabbithole.VhostInfo{}, nil)

			broker = defaultServiceBroker(defaultConfig(), client)
			ctx = context.TODO()
		})

		It("deletes a vhost", func() {
			client.DeleteVhostReturns(&http.Response{StatusCode: http.StatusNoContent}, nil)
			spec, err := broker.Deprovision(ctx, "my-service-id", brokerapi.DeprovisionDetails{}, false)
			Expect(err).NotTo(HaveOccurred())
			Expect(spec.IsAsync).To(BeFalse())

			Expect(client.DeleteVhostCallCount()).To(Equal(1))
			Expect(client.DeleteVhostArgsForCall(0)).To(Equal("my-service-id"))
		})

		It("fails if it cannot delete the vhost", func() {
			client.DeleteVhostReturns(nil, errors.New("oops"))
			_, err := broker.Deprovision(ctx, "my-service-id", brokerapi.DeprovisionDetails{}, false)
			Expect(err).To(HaveOccurred())
		})
	})

	When("the instance does not exist", func() {
		BeforeEach(func() {
			client = new(fakes.FakeAPIClient)
			client.GetVhostReturns(nil, errors.New("vhost does not exist"))

			broker = defaultServiceBroker(defaultConfig(), client)
			ctx = context.TODO()
		})

		It("returns an error if vhost does not exist", func() {
			_, err := broker.Deprovision(ctx, "my-service-id", brokerapi.DeprovisionDetails{}, false)
			Expect(err).To(MatchError(brokerapi.ErrInstanceDoesNotExist))
		})
	})
})
