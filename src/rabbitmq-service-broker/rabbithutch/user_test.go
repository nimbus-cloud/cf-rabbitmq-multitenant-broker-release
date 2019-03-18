package rabbithutch_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "rabbitmq-service-broker/rabbithutch"
	"rabbitmq-service-broker/rabbithutch/fakes"
)

var _ = Describe("Binding a RMQ service instance", func() {
	var (
		rabbitClient *fakes.FakeAPIClient
		rabbithutch  RabbitHutch
	)

	BeforeEach(func() {
		rabbitClient = new(fakes.FakeAPIClient)
		rabbithutch = New(rabbitClient)
	})

	Describe("the user", func() {

		FIt("creates a user", func() {
			rabbitClient.UpdatePermissionsInReturns(&http.Response{StatusCode: http.StatusOK}, nil)

			password, err := rabbithutch.CreateUser("fake-user", "fake-vhost", "")
			Expect(err).NotTo(HaveOccurred())

			Expect(rabbitClient.PutUserCallCount()).To(Equal(1))
			username, info := rabbitClient.PutUserArgsForCall(0)
			Expect(username).To(Equal("fake-user"))
			Expect(info.Tags).To(Equal("policymaker,management"))
			Expect(info.Password).To(MatchRegexp(`[a-zA-Z0-9\-_]{24}`))
			Expect(password).To(Equal(info.Password))
		})

		// It("fails with an error if it cannot create a user", func() {
		// 	rabbitClient.PutUserReturns(nil, errors.New("foo"))
		// 	_, err := broker.Bind(ctx, "my-service-instance-id", "binding-id", brokerapi.BindDetails{}, false)
		// 	Expect(err).To(MatchError("foo"))
		// })
		//
		// 	It("fails with an error if the user already exists", func() {
		// 		rabbitClient.PutUserReturns(&http.Response{StatusCode: http.StatusNoContent}, nil)
		// 		_, err := broker.Bind(ctx, "my-service-instance-id", "binding-id", brokerapi.BindDetails{}, false)
		// 		Expect(err).To(MatchError(brokerapi.ErrBindingAlreadyExists))
		// 	})
		//
		// 	It("grants the user full permissions to the vhost", func() {
		// 		_, err := broker.Bind(ctx, "my-service-instance-id", "binding-id", brokerapi.BindDetails{}, false)
		// 		Expect(err).NotTo(HaveOccurred())
		// 		Expect(rabbitClient.UpdatePermissionsInCallCount()).To(Equal(1))
		// 		vhost, username, permissions := rabbitClient.UpdatePermissionsInArgsForCall(0)
		// 		Expect(vhost).To(Equal("my-service-instance-id"))
		// 		Expect(username).To(Equal("binding-id"))
		// 		Expect(permissions.Configure).To(Equal(".*"))
		// 		Expect(permissions.Read).To(Equal(".*"))
		// 		Expect(permissions.Write).To(Equal(".*"))
		// 	})
		//
		// 	When("user tags are set in the config", func() {
		// 		BeforeEach(func() {
		// 			rabbitClient = new(fakes.FakeAPIClient)
		// 			broker = defaultServiceBroker(defaultConfigWithUserTags(), rabbitClient, rabbithutch)
		// 			ctx = context.TODO()
		// 			rabbitClient.UpdatePermissionsInReturns(&http.Response{StatusCode: http.StatusOK}, nil)
		// 		})
		//
		// 		It("creates a user with the tags", func() {
		// 			_, err := broker.Bind(ctx, "my-service-instance-id", "binding-id", brokerapi.BindDetails{}, false)
		// 			Expect(err).NotTo(HaveOccurred())
		//
		// 			Expect(rabbitClient.PutUserCallCount()).To(Equal(1))
		// 			username, info := rabbitClient.PutUserArgsForCall(0)
		// 			Expect(username).To(Equal("binding-id"))
		// 			Expect(info.Password).To(MatchRegexp(`[a-zA-Z0-9\-_]{24}`))
		// 			Expect(info.Tags).To(Equal("administrator"))
		// 		})
		// 	})
	})
})
