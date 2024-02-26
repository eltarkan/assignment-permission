package config

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

var (
	basePath, _  = os.Getwd()
	Config       = Environment{}
	DBConnection *mongo.Client
)

func LoadConfig() error {
	viper.AddConfigPath(basePath)
	//viper.SetConfigName(".env")
	//viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.BindEnv("MONGODB_CONNECTION_STRING")
	viper.BindEnv("MONGODB_DATABASE")

	//err := viper.ReadInConfig()
	//if err != nil {
	//	return err
	//}

	err := viper.Unmarshal(&Config)
	if err != nil {
		return err
	}

	return nil
}
