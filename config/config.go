package config

type Config struct {
	MongoUrl    string
	PostgresUrl string
}

func LoadConfig() *Config {
	return &Config{
		MongoUrl:    "mongodb://localhost:27017",
		PostgresUrl: "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable",
	}
}
