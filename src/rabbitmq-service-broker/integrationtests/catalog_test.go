package integrationtests_test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"rabbitmq-service-broker/broker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi"
	yaml "gopkg.in/yaml.v2"
)

const (
	url      = "http://localhost:8901/v2/catalog"
	username = "admin"
	password = "admin"
)

var _ = Describe("/v2/catalog", func() {
	When("no credentials are provided", func() {
		It("fails with HTTP 401", func() {
			response, err := http.Get(url)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
		})
	})

	When("credentials are provided and they are correct", func() {

		XWhen("an empty config is provided", func() {
			It("succeeds with HTTP 200", func() {
				response, body := doRequest(http.MethodGet, url, nil)
				Expect(response.StatusCode).To(Equal(http.StatusOK))

				catalog := make(map[string][]brokerapi.Service)
				Expect(json.Unmarshal(body, &catalog)).To(Succeed())

				Expect(catalog["services"]).To(HaveLen(1))
				// match against the expectation
				Expect(catalog["services"][0]).To(Equal(brokerapi.Service{
					ID:          "service-id",
					Name:        "service-name",
					Description: "service-description",
				}))
			})

		})

		When("a valid config is provided", func() {
			It("succeeds with HTTP 200", func() {
				response, body := doRequest(http.MethodGet, url, nil)
				Expect(response.StatusCode).To(Equal(http.StatusOK))

				catalog := make(map[string][]brokerapi.Service)
				Expect(json.Unmarshal(body, &catalog)).To(Succeed())

				// read in the expectation
				configFilePath := "./fixtures/config.yml"
				configBytes, err := ioutil.ReadFile(filepath.FromSlash(configFilePath))
				Expect(err).NotTo(HaveOccurred())

				config := broker.Config{}
				Expect(yaml.Unmarshal(configBytes, &config)).To(Succeed())

				Expect(catalog["services"]).To(HaveLen(1))
				// match against the expectation
				Expect(catalog["services"][0]).To(Equal(brokerapi.Service{
					ID:          config.ServiceConfig.UUID,
					Name:        config.ServiceConfig.Name,
					Description: config.ServiceConfig.OfferingDescription,
				}))
			})
		})
	})
})

func doRequest(method, url string, body io.Reader) (*http.Response, []byte) {
	req, err := http.NewRequest(method, url, body)
	Expect(err).ToNot(HaveOccurred())

	req.SetBasicAuth(username, password)
	req.Header.Set("X-Broker-API-Version", "2.14")

	req.Close = true
	resp, err := http.DefaultClient.Do(req)
	Expect(err).ToNot(HaveOccurred())

	bodyContent, err := ioutil.ReadAll(resp.Body)
	Expect(err).NotTo(HaveOccurred())

	Expect(resp.Body.Close()).To(Succeed())
	return resp, bodyContent
}
