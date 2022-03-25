package jpush

import (
	"github.com/goal-web/contracts"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(application contracts.Application) {
	application.Singleton("jpush", func(config contracts.Config) *Client {
		defaultClient = New(config.Get("jpush").(Config))
		return defaultClient
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
