package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	facades.Config().Add("database", map[string]any{
		// Default database connection
		"default": facades.Config().Env("DB_CONNECTION", "postgres"),

		// Database connections
		"connections": map[string]any{
			// PostgreSQL configuration
			"postgres": map[string]any{
				"driver":   "postgres",
				"host":     facades.Config().Env("POSTGRES_HOST", "localhost"),
				"port":     facades.Config().Env("POSTGRES_PORT", "5432"),
				"database": facades.Config().Env("POSTGRES_DATABASE", "goravel"),
				"username": facades.Config().Env("POSTGRES_USERNAME", "postgres"),
				"password": facades.Config().Env("POSTGRES_PASSWORD", ""),
				"charset":  "utf8",
				"prefix":   "",
				"sslmode":  "disable",
				"timezone": "UTC",
			},

			// MongoDB configuration
			"mongodb": map[string]any{
				"driver":   "mongodb",
				"host":     facades.Config().Env("MONGODB_HOST", "localhost"),
				"port":     facades.Config().Env("MONGODB_PORT", "27017"),
				"database": facades.Config().Env("MONGODB_DATABASE", "goravel"),
				"username": facades.Config().Env("MONGODB_USERNAME", ""),
				"password": facades.Config().Env("MONGODB_PASSWORD", ""),
				"options": map[string]any{
					"auth_source": facades.Config().Env("MONGODB_AUTH_SOURCE", "admin"),
				},
				"uri": facades.Config().Env("MONGODB_URI", ""), // Hoặc sử dụng URI trực tiếp
			},
		},

		// Migration Repository Table
		"migrations": "migrations",

		// Redis Databases
		"redis": map[string]any{
			"default": map[string]any{
				"host":     facades.Config().Env("REDIS_HOST", "localhost"),
				"password": facades.Config().Env("REDIS_PASSWORD", ""),
				"port":     facades.Config().Env("REDIS_PORT", "6379"),
				"database": facades.Config().Env("REDIS_DB", "0"),
			},
			"cache": map[string]any{
				"host":     facades.Config().Env("REDIS_HOST", "localhost"),
				"password": facades.Config().Env("REDIS_PASSWORD", ""),
				"port":     facades.Config().Env("REDIS_PORT", "6379"),
				"database": facades.Config().Env("REDIS_CACHE_DB", "1"),
			},
		},
	})
}
