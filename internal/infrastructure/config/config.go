<<<<<<< HEAD
package config

import (
	"github.com/spf13/viper"
)

// Mapea as variáveis de ambiente
type Config struct {
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBName             string `mapstructure:"DB_NAME"`
	UserRepositoryType string `mapstructure:"USER_REPOSITORY_TYPE"`
}

// Carrega as variáveis de ambiente com viper
func LoadConfig() (*Config, error) {
	viper.SetDefault("USER_REPOSITORY_TYPE", "postgres")

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
=======
// package config é responsável por carregar as configurações do projeto.
package config

import (
	"github.com/spf13/viper"
)

// Config é a estrutura que armazena as configurações do projeto.
type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	UserRepositoryType string `mapstructure:"USER_REPOSITORY_TYPE"`
}

// LoadConfig carrega as configurações do projeto.
func LoadConfig() (*Config, error) {
	viper.SetDefault("USER_REPOSITORY_TYPE", "postgres")
	
	// Carrega o arquivo de configuração
	viper.SetConfigFile(".env")

	// Carrega as variáveis de ambiente
	viper.AutomaticEnv()

	// Lê o arquivo de configuração
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Cria uma nova instância de Config
	config := &Config{}

	// Deserializa as configurações para a estrutura Config
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
} 
>>>>>>> develop
