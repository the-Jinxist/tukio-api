package config

import (
	"github.com/spf13/viper"
)

const (
	EnvKey          = "ENVIRONMENT"
	DbDriverKey     = "DB_DRIVER"
	DbSourceKey     = "DB_SOURCE"
	FromEmailKey    = "FROM_EMAIL"
	FromPasswordKey = "FROM_PASSWORD"
)

type Config struct {
	Environment   string `mapstructure:"ENVIRONMENT"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	EmailFrom     string `mapstructure:"FROM_EMAIL"`
	EmailPassword string `mapstructure:"FROM_PASSWORD"`
}

func LoadConfigs(path string) (Config, error) {

	var (
		err    error
		config Config
	)

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.AutomaticEnv()

	if viper.GetString(EnvKey) == "" {

		err = viper.ReadInConfig()
		if err != nil {
			return config, err
		}

		err = viper.Unmarshal(&config)
		if err != nil {
			return config, err
		}

		return config, err
	} else {
		config.Environment = viper.GetString(EnvKey)
		config.DBSource = viper.GetString(DbSourceKey)
		config.EmailFrom = viper.GetString(FromEmailKey)
		config.EmailPassword = viper.GetString(FromPasswordKey)

		return config, err

	}

}
