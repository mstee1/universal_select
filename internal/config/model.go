package config

type Config struct {
	Database Database `mapstructure:"database"`
	Version  string
}

type Database struct {
	DbPort     int    `mapstructure:"dbPort"`
	DbHost     string `mapstructure:"dbHost"`
	DbName     string `mapstructure:"dbName"`
	DbUser     string `mapstructure:"dbUser"`
	DbPassword string `mapstructure:"dbPassword"`
}
