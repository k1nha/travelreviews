package config

import "github.com/spf13/viper"

var env *envconfig

type envconfig struct {
	DbURL string `mapstructure:"DATABASE_URL"`
	Port  string `mapstructure:"PORT"`
}

func GetEnvConfig() *envconfig {
	return env
}

func LoadEnv(pathUrl string) (*envconfig, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(pathUrl)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		panic(err)
	}

	return env, nil
}
