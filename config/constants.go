package config

type Environment struct {
	MongodbConnectionString string `mapstructure:"MONGODB_CONNECTION_STRING"`
	MongodbName             string `mapstructure:"MONGODB_DATABASE"`
}
