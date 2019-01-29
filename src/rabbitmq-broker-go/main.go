package main

import (
	"log"
	"net/http"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
)

func main() {
	brokerLogger := lager.NewLogger("rabbitmq-multitenant-broker")
	brokerCredentials := brokerapi.BrokerCredentials{
		Username: "admin",
		Password: "password",
	}
	broker := &RabbitmqServiceBroker{}

	brokerAPI := brokerapi.New(broker, brokerLogger, brokerCredentials)

	http.Handle("/", brokerAPI)
	log.Fatal(http.ListenAndServe(":8901", nil))
}
