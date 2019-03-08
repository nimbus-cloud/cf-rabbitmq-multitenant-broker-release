package binding_test

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"rabbitmq-service-broker/binding"
)

var _ = Describe("Binding", func() {
	It("generates the right binding with TLS enabled", func() {
		b := binding.Builder{
			MgmtDomain: "pivotal-rabbitmq.sys.philippinebrown.cf-app.com",
			Hostnames:  []string{"10.0.4.100", "10.0.4.101"},
			Vhost:      "6418d19f-e9e8-4c8b-9c92-5087c89cbc46",
			Username:   "b2a5de47-796b-414d-bab4-eb299c268653",
			Password:   "cfrnvvjtr6t803ilrdhgbe8mn7",
			AMQPPort:   "fake-amqp-port",
			MQTTPort:   "fale=mqtt-port",
			STOMPPort:  "fake-stomp-port",
			TLS:        true,
		}
		creds, err := json.Marshal(b.AsMap())
		Expect(err).NotTo(HaveOccurred())

		expected, err := ioutil.ReadFile("fixtures/binding_tls.json")
		Expect(err).NotTo(HaveOccurred())

		Expect(string(creds)).To(MatchJSON(expected))
	})
})
