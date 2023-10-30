package injectors

import (
	"github.com/datafabric/gateway/controllers"
	"github.com/datafabric/gateway/services"
)

// PortalInjector portal injector
type PortalInjector struct{}

// Init for interconnection [ controller(App) - Service(Repository) - repository - datastore ] : Dependency Injection
func (PortalInjector) Init(in *Injector) *controllers.PortalController {
	svc, err := services.PortalServiceInitialize(in.Log.Logger, "localhost", 9360)
	if err != nil {
		return nil
	}
	return controllers.PortalControllerInitialize(svc)
}
