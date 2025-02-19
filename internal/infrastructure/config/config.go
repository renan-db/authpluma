package config

import (
	"log"
	"github.com/spf13/viper"
)

func LoadConfig(env string) error {
	// Carrega o arquivo .env específico do ambiente
	viper.SetConfigName(".env-" + env)
	viper.SetConfigType("env")
	viper.AddConfigPath(".") // Diretório atual

	viper.AutomaticEnv() // Carrega variáveis de ambiente

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
		return err
	}

	log.Printf("Using config file: %s", viper.ConfigFileUsed())
	return nil
} 