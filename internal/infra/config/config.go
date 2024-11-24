package config

type (
	Config struct {
		Env     string     `yaml:"env" env-default:"dev"`
		Postgre PostgreSQL `yaml:"postgres"`
	}
	PostgreSQL struct {
		Host     string `yaml:"host" env-default:"localhost"`
		Port     int    `yaml:"port" env-default:"5121"`
		User     string `yaml:"user" env-default:"postgres"`
		Password string `yaml:"password" env-default:"root"`
		DbName   string `yaml:"dbname" env-default:"database"`
	}
)
