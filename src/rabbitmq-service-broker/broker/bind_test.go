package broker_test

import (
	"context"
	"errors"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"rabbitmq-service-broker/broker/fakes"

	"github.com/pivotal-cf/brokerapi"
)

var _ = Describe("Binding a RMQ service instance", func() {
	var (
		client *fakes.FakeAPIClient
		broker brokerapi.ServiceBroker
		ctx    context.Context
	)

	BeforeEach(func() {
		client = new(fakes.FakeAPIClient)
		broker = defaultServiceBroker(defaultConfig(), client)
		ctx = context.TODO()
	})

	It("creates a user", func() {
		_, err := broker.Bind(ctx, "my-service-instance-id", "binding-id", brokerapi.BindDetails{}, false)
		Expect(err).NotTo(HaveOccurred())

		Expect(client.PutUserCallCount()).To(Equal(1))
		username, info := client.PutUserArgsForCall(0)
		Expect(username).To(Equal("binding-id"))
		Expect(info.Password).To(MatchRegexp(`[a-zA-Z0-9\-_]{24}`))
	})

	It("fails with an error if it cannot create a user", func() {
		client.PutUserReturns(nil, errors.New("foo"))
		_, err := broker.Bind(ctx, "my-service-instance-id", "binding-id", brokerapi.BindDetails{}, false)
		Expect(err).To(MatchError("foo"))
	})

	It("fails with an error if the user already exists", func() {
		client.PutUserReturns(&http.Response{StatusCode: http.StatusNoContent}, nil)
		_, err := broker.Bind(ctx, "my-service-instance-id", "binding-id", brokerapi.BindDetails{}, false)
		Expect(err).To(MatchError(brokerapi.ErrBindingAlreadyExists))
	})
})
