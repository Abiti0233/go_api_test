package config

type Config struct {
	Port       string `envconfig:"APP_PORT"     default:"8081"`
	DBHost     string `envconfig:"DB_HOST"      default:"db"`
	DBPort     string `envconfig:"DB_PORT"      default:"5432"`
	DBUser     string `envconfig:"DB_USER"      default:"postgres"`
	DBPassword string `envconfig:"DB_PASSWORD"  default:"secret"`
	DBName     string `envconfig:"DB_NAME"      default:"mytodo"`
}
