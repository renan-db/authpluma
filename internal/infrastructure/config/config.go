package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(_ string) error {
	viper.AutomaticEnv()  // Carrega automaticamente as variáveis de ambiente
	return nil
} 