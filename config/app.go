package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	facades.Config().Add("app", map[string]any{
		// Application name
		"name": facades.Config().Env("APP_NAME", "Goravel"),

		// Application environment
		"env": facades.Config().Env("APP_ENV", "production"),

		// Application debug mode
		"debug": facades.Config().Env("APP_DEBUG", false),

		// Application URL
		"url": facades.Config().Env("APP_URL", "http://localhost"),

		// Application host
		"host": facades.Config().Env("APP_HOST", "127.0.0.1"),

		// Application port
		"port": facades.Config().Env("APP_PORT", "8000"),

		// Application Timezone
		"timezone": "UTC",

		// Application Locale Configuration
		"locale": facades.Config().Env("APP_LOCALE", "en"),

		// Encryption Key
		"key": facades.Config().Env("APP_KEY", ""),

		// Autoload service providers
		"providers": []any{
			// Framework service providers...
			// Application service providers...
		},
	})
}
