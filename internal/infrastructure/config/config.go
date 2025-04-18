package config

import (
	"github.com/spf13/viper"
)

// Mapea as variáveis de ambiente
type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

// Carrega as variáveis de ambiente com viper
func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Lê o arquivo de configuração
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	// Deserializa as configurações para a estrutura Config
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
} 