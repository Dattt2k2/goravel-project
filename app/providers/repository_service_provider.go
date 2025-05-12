package providers

import (
	"goravel/app/repositories"
	"goravel/app/services"

	"github.com/goravel/framework/contracts/foundation"
)

type RepositoryServiceProvider struct {
}

func (receiver *RepositoryServiceProvider) Register(app foundation.Application) {
	app.Singleton("user.repository", func(app foundation.Application) (any, error) {
		return repositories.NewUserRepository(), nil
	})

	app.Singleton("user.service", func(app foundation.Application) (any, error) {
		userRepository, _ := app.Make("user.repository")
		return services.NewUserService(userRepository.(interface{})), nil
	})
}
