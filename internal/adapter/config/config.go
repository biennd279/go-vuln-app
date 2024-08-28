package config

// Server configuration
type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

// Database configuration
type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

// Redis configuration
type Redis struct {
	Host string
	Port string
}

// JWT configuration
type JWT struct {
	Secret string
}

// Log configuration
type Log struct {
	Level string
}

type Config struct {
	IsDev  bool   `mapstructure:"is_dev"`
	Server Server `mapstructure:"server"`
	//Database Database
	//Redis    Redis
	//JWT      JWT
	//Log      Log `mapstructure:"log"`
}
