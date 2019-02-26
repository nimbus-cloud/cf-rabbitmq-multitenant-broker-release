package broker

import (
	"context"
	"fmt"
	"net/http"

	rabbithole "github.com/michaelklishin/rabbit-hole"

	"github.com/pivotal-cf/brokerapi"

	"code.cloudfoundry.org/lager"
)

func (b RabbitMQServiceBroker) Provision(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	logger := b.logger.Session("provision")

	spec := brokerapi.ProvisionedServiceSpec{}
	vhost := instanceID

	if err := b.createVhost(vhost); err != nil {
		return spec, err
	}

	if err := b.assignPermissionsToUser(vhost, b.config.RabbitMQConfig.Administrator.Username); err != nil {
		return spec, err
	}

	if b.config.RabbitMQConfig.Management.Username != "" {
		if err := b.assignPermissionsToUser(vhost, b.config.RabbitMQConfig.Management.Username); err != nil {
			return spec, err
		}
	}

	if b.config.RabbitMQConfig.Policy.Enabled {
		if err := b.createPolicy(vhost); err != nil {
			return spec, err
		}
	}

	logger.Info("provision-succeeded")
	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (b *RabbitMQServiceBroker) createVhost(vhost string) error {
	logger := b.logger.Session("create-vhost")
	logger.Info("get-vhost")
	if _, err := b.client.GetVhost(vhost); err == nil {
		err = brokerapi.ErrInstanceAlreadyExists
		logger.Error("get-vhost-failed", err)
		return err
	}
	logger.Info("get-vhost-succeeded")

	logger.Info("put-vhost")
	resp, err := b.client.PutVhost(vhost, rabbithole.VhostSettings{})

	if err != nil {
		logger.Error("put-vhost-failed", err)
		return err
	}
	logger.Info("put-vhost-succeeded")

	if err := validateResponse(resp); err != nil {
		logger.Error("put-vhost-failed", err)
		return err
	}

	return nil
}

func (b *RabbitMQServiceBroker) assignPermissionsToUser(vhost, username string) error {
	logger := b.logger.Session("assign-persmissions-to-user", lager.Data{"username": username})
	logger.Info("update-permissions")

	permissions := rabbithole.Permissions{Configure: ".*", Write: ".*", Read: ".*"}
	resp, err := b.client.UpdatePermissionsIn(vhost, username, permissions)
	if err != nil {
		logger.Error("update-permissions-failed", err)
		return err
	}

	if err := validateResponse(resp); err != nil {
		logger.Error("update-permission-failed", err)
		return err
	}

	logger.Info("update-permissions-succeeded")
	return nil
}

func (b *RabbitMQServiceBroker) createPolicy(vhost string) error {
	logger := b.logger.Session("create-policy")

	policy := rabbithole.Policy{
		Definition: rabbithole.PolicyDefinition(b.config.RabbitMQConfig.Policy.Definition),
		Priority:   b.config.RabbitMQConfig.Policy.Priority,
		Vhost:      vhost,
		Pattern:    ".*",
		ApplyTo:    "all",
		Name:       b.config.RabbitMQConfig.Policy.Name,
	}

	logger.Info("put-policy", lager.Data{"policy": policy})
	resp, err := b.client.PutPolicy(vhost, b.config.RabbitMQConfig.Policy.Name, policy)
	if err != nil {
		logger.Error("put-policy-failed", err)
		return err
	}

	if err := validateResponse(resp); err != nil {
		logger.Error("put-policy-failed", err)
		return err
	}

	logger.Info("put-policy-succeeded")
	return nil
}

func validateResponse(resp *http.Response) error {
	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		return fmt.Errorf("http request failed with status code: %v", resp.StatusCode)
	}
	return nil
}
