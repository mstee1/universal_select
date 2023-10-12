package config

type Config struct {
	Database Database `mapstructure:"database"`
	Sql      Sql      `mapstructure:"sql"`
	Version  string
}

type Database struct {
	DbPort     int    `mapstructure:"dbPort"`
	DbHost     string `mapstructure:"dbHost"`
	DbName     string `mapstructure:"dbName"`
	DbUser     string `mapstructure:"dbUser"`
	DbPassword string `mapstructure:"dbPassword"`
}

type Sql struct {
	Select1 string `mapstructure:"select1"`
	Select2 string `mapstructure:"select2"`
	Select3 string `mapstructure:"select3"`
	Select4 string `mapstructure:"select4"`
}
