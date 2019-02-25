package broker_test

import (
	"context"
	"fmt"
	"net/http"

	serviceBroker "rabbitmq-service-broker/broker"

	"rabbitmq-service-broker/broker/fakes"

	"github.com/michaelklishin/rabbit-hole"

	"github.com/pivotal-cf/brokerapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Broker Provisioning and Deprovisioning", func() {
	var (
		client *fakes.FakeAPIClient
		config serviceBroker.Config
		broker brokerapi.ServiceBroker
		//	credentialsGenerator *fakes.FakeCredentialsGenerator
		//		logger *lagertest.TestLogger
		ctx context.Context
	)

	BeforeEach(func() {
		//	config = defaultConfig()
		client = new(fakes.FakeAPIClient)
		//	//	credentialsGenerator = new(fakes.FakeCredentialsGenerator)
		//	logger = lagertest.NewTestLogger("test")
		broker = serviceBroker.New(config, client)
		ctx = context.TODO()
	})

	Context("Provision", func() {

		BeforeEach(func() {
			client.GetVhostReturns(nil, fmt.Errorf("vhost does not exist"))
			client.PutVhostReturns(&http.Response{StatusCode: http.StatusNoContent}, nil)
			//		client.DeleteVhostReturns(&http.Response{StatusCode: http.StatusNoContent}, nil)
			//		client.UpdatePermissionsInReturns(&http.Response{StatusCode: http.StatusNoContent}, nil)
			//		client.PutPolicyReturns(&http.Response{StatusCode: http.StatusNoContent}, nil)
			//		client.PutUserReturns(&http.Response{StatusCode: http.StatusNoContent}, nil)
			//		client.GetUserReturns(nil, fmt.Errorf(rabbitbroker.NotFoundIdentifier))

			//		credentialsGenerator.GenerateReturns(rabbitbroker.Credentials{
			//			Username: "test-user",
			//			Password: "test-password",
			//		})
		})

		It("returns no error", func() {
			_, err := broker.Provision(ctx, "my-service-id", brokerapi.ProvisionDetails{}, false)
			Expect(err).NotTo(HaveOccurred())
		})

		It("checks whether the host exist", func() {
			_, err := broker.Provision(ctx, "my-service-id", brokerapi.ProvisionDetails{}, false)
			Expect(client.GetVhostCallCount()).To(Equal(1))
			Expect(client.GetVhostArgsForCall(0)).To(Equal("my-service-id"))
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when the vhost exists", func() {
			It("returns ErrInstanceAlreadyExists error", func() {
				client.GetVhostReturns(&rabbithole.VhostInfo{}, nil)
				_, err := broker.Provision(ctx, "my-service-id", brokerapi.ProvisionDetails{}, false)
				Expect(err).To(Equal(brokerapi.ErrInstanceAlreadyExists))
				Expect(client.PutVhostCallCount()).To(Equal(0))
			})
		})
		It("creates a vhost", func() {
			_, err := broker.Provision(ctx, "my-service-id", brokerapi.ProvisionDetails{}, false)
			Expect(client.PutVhostCallCount()).To(Equal(1))
			vhost, _ := client.PutVhostArgsForCall(0)
			Expect(vhost).To(Equal("my-service-id"))
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when the vhost creation fails", func() {
			It("returns an error", func() {
				client.PutVhostReturns(nil, fmt.Errorf("vhost-creation-failed"))
				_, err := broker.Provision(ctx, "my-service-id", brokerapi.ProvisionDetails{}, false)
				Expect(err).To(MatchError("vhost-creation-failed"))
			})

			Context("with error status code", func() {
				It("returns the status as error", func() {
					client.PutVhostReturns(&http.Response{StatusCode: http.StatusInternalServerError}, nil)
					_, err := broker.Provision(ctx, "my-service-id", brokerapi.ProvisionDetails{}, false)
					Expect(err).To(MatchError("http request failed with status code: 500"))
				})
			})
		})
	})

})
