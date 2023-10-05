package settings

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var Configs *Config

type Config struct {
	Host     string `mapstructure:"MIKROTIK_HOST"`
	User     string `mapstructure:"MIKROTIK_USER"`
	Password string `mapstructure:"MIKROTIK_PASS"`
	TLS      string `mapstructure:"MIKROTIK_TLS"`
}

func InitEnvConfigs() {

	Configs = loadEnvVariables()

}

func loadEnvVariables() (config *Config) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath(homedir)

	// Tell viper the name of your file
	viper.SetConfigName(".mikrodns")

	// Tell viper the type of your file
	viper.SetConfigType("yml")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
