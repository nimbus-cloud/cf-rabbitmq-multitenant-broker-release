package binding

import (
	"encoding/json"
	"fmt"
)

type protocol struct {
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	VHost     string   `json:"vhost,omitempty"`
	Hostname  string   `json:"host"`
	Hostnames []string `json:"hosts"`
	URI       string   `json:"uri"`
	URIs      []string `json:"uris"`
	Port      int      `json:"port"`
	TLS       bool     `json:"ssl"`
	Path      string   `json:"path,omitempty"`
}

type protocols map[string]protocol

type binding struct {
	DashboardURL string    `json:"dashboard_url"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Hostname     string    `json:"hostname"`
	Hostnames    []string  `json:"hostnames"`
	HTTPAPIURI   string    `json:"http_api_uri"`
	HTTPAPIURIs  []string  `json:"http_api_uris"`
	URI          string    `json:"uri"`
	URIs         []string  `json:"uris"`
	VHost        string    `json:"vhost"`
	TLS          bool      `json:"ssl"`
	Protocols    protocols `json:"protocols"`
}

type Builder struct {
	MgmtDomain    string
	Hostnames     []string
	VHost         string
	Username      string
	Password      string
	TLS           bool
	ProtocolPorts map[string]int // key=protcol, value=port, e.g. "amqp": 4567
}

func (b Builder) AsMap() (output map[string]interface{}, err error) {
	bind := binding{
		VHost:        b.VHost,
		Username:     b.Username,
		Password:     b.Password,
		DashboardURL: b.dashboardURL(),
		Hostname:     b.firstHostname(),
		Hostnames:    b.Hostnames,
		HTTPAPIURI:   b.httpapiuri(),
		HTTPAPIURIs:  b.httpapiuris(),
		URI:          b.uri(b.firstHostname()),
		URIs:         b.uris(),
		TLS:          b.TLS,
		Protocols:    b.protocols(),
	}

	bytes, err := json.Marshal(bind)
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(bytes, &output)
	return output, err
}

func (b Builder) dashboardURL() string {
	return fmt.Sprintf("https://%s/#/login/%s/%s", b.MgmtDomain, b.Username, b.Password)
}

func (b Builder) firstHostname() string {
	return b.Hostnames[0]
}

func (b Builder) uri(hostname string) string {
	return fmt.Sprintf("%s://%s:%s@%s/%s", b.amqpScheme(), b.Username, b.Password, hostname, b.VHost)
}

func (b Builder) uriWithPort(hostname string, port int) string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s", b.amqpScheme(), b.Username, b.Password, hostname, port, b.VHost)
}

func (b Builder) uriWithSchemeAndPort(hostname string, port int, scheme string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s", scheme, b.Username, b.Password, hostname, port, b.VHost)
}

func (b Builder) uris() []string {
	var uris []string
	for _, hostname := range b.Hostnames {
		uris = append(uris, b.uri(hostname))
	}
	return uris
}

func (b Builder) urisWithPort(port int) []string {
	var uris []string
	for _, hostname := range b.Hostnames {
		uris = append(uris, b.uriWithPort(hostname, port))
	}
	return uris
}

func (b Builder) urisWithSchemeAndPort(port int, scheme string) []string {
	var uris []string
	for _, hostname := range b.Hostnames {
		uris = append(uris, b.uriWithSchemeAndPort(hostname, port, scheme))
	}
	return uris
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

func (b Builder) protocols() protocols {
	ps := make(protocols)
	for protocol, port := range b.ProtocolPorts {
		switch protocol {
		case "amqp":
			ps["amqp"] = b.addAMQPProtocol(port)
		case "amqp/ssl":
			ps["amqp+ssl"] = b.addAMQPProtocol(port)
		}
	}
	ps["management"] = b.addMgmtProtocol()
	return ps
}

func (b Builder) addMgmtProtocol() protocol {
	return protocol{
		Username: b.Username,
		Password: b.Password,
		// VHost:     b.VHost,
		Hostname:  b.firstHostname(),
		Hostnames: b.Hostnames,
		URI:       b.uriWithSchemeAndPort(b.firstHostname(), 15672, "http"),
		URIs:      b.urisWithSchemeAndPort(15672, "http"),
		Port:      15672,
		TLS:       false,
		Path:      "/api/",
	}
}

func (b Builder) addAMQPProtocol(port int) protocol {
	return protocol{
		Username:  b.Username,
		Password:  b.Password,
		VHost:     b.VHost,
		Hostname:  b.firstHostname(),
		Hostnames: b.Hostnames,
		URI:       b.uriWithPort(b.firstHostname(), port),
		URIs:      b.urisWithPort(port),
		Port:      port,
		TLS:       b.TLS,
	}
}
