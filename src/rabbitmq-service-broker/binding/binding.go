package binding

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type binding struct {
	DashboardURL string `mapstructure:"dashboard_url"`
	// Username     string   `mapstructure:"username"`
	Password    string   `mapstructure:"password"`
	Hostname    string   `mapstructure:"hostname"`
	Hostnames   []string `mapstructure:"hostnames"`
	HTTPAPIURI  string   `mapstructure:"http_api_uri"`
	HTTPAPIURIs []string `mapstructure:"http_api_uris"`
	URI         string   `mapstructure:"uri"`
	// URIs         []string `mapstructure:"uris"`
	// VHOST        string   `mapstructure:"vhost"`
	TLS bool `mapstructure:"ssl"`
	// Protocols    map[string]map[string]interface{} `mapstructure:"protocols"`
}

type Builder struct {
	MgmtDomain string
	Hostnames  []string
	Vhost      string
	Username   string
	Password   string
	AMQPPort   string
	MQTTPort   string
	STOMPPort  string
	TLS        bool
}

func (b Builder) AsMap() (output map[string]interface{}) {
	bind := binding{
		DashboardURL: b.dashboardURL(),
		Hostname:     b.firstHostname(),
		Hostnames:    b.Hostnames,
		HTTPAPIURI:   b.httpapiuri(),
		HTTPAPIURIs:  b.httpapiuris(),
		Password:     b.Password,
		TLS:          b.TLS,
		URI:          b.uri(),
	}
	mapstructure.Decode(bind, &output)
	return output
}

func (b Builder) dashboardURL() string {
	return fmt.Sprintf("https://%s/#/login/%s/%s", b.MgmtDomain, b.Username, b.Password)
}

func (b Builder) firstHostname() string {
	return b.Hostnames[0]
}

// {:uri           (uri-for (cfg/amqp-scheme) (URLEncoder/encode username) password first-node-host vhost)
//  :uris          (map #(uri-for (cfg/amqp-scheme) (URLEncoder/encode username) password % vhost) node-hosts)
//   (format "%s://%s:%s@%s/%s" scheme username password node-host (URLEncoder/encode vhost)))

func (b Builder) uri() string {
	return fmt.Sprintf("%s://%s:%s@%s/%s", b.amqpScheme(), b.Username, b.Password, b.firstHostname(), b.Vhost)
}

func (b Builder) amqpScheme() string {
	if b.TLS {
		return "amqps"
	}
	return "amqp"
}

func (b Builder) httpapiuri() string {
	return fmt.Sprintf("https://%s:%s@%s/api/", b.Username, b.Password, b.MgmtDomain)
}

func (b Builder) httpapiuris() []string {
	return []string{b.httpapiuri()}
}
