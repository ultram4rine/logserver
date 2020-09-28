package conf

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetConfig function parse a config file to viper.
func GetConfig(confName string) error {
	viper.SetConfigName(confName)
	viper.AddConfigPath("/etc/logserver/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("Failed to read config file: %s", err)
	}

	viper.SetEnvPrefix("logserver")
	if err := viper.BindEnv("cert_path"); err != nil {
		return errors.New("Failed to bind cert_path ENV variable")
	}
	if err := viper.BindEnv("key_path"); err != nil {
		return errors.New("Failed to bind key_path ENV variable")
	}
	if err := viper.BindEnv("client_cert_path"); err != nil {
		return errors.New("Failed to bind client_cert_path ENV variable")
	}
	if err := viper.BindEnv("grpc_port"); err != nil {
		return errors.New("Failed to bind grpc_port ENV variable")
	}
	if err := viper.BindEnv("gateway_port"); err != nil {
		return errors.New("Failed to bind gateway_port ENV variable")
	}
	if err := viper.BindEnv("jwt_key"); err != nil {
		return errors.New("Failed to bind jwt_key ENV variable")
	}
	if err := viper.BindEnv("db_host"); err != nil {
		return errors.New("Failed to bind db_host ENV variable")
	}
	if err := viper.BindEnv("db_name"); err != nil {
		return errors.New("Failed to bind db_name ENV variable")
	}
	if err := viper.BindEnv("db_user"); err != nil {
		return errors.New("Failed to bind db_user ENV variable")
	}
	if err := viper.BindEnv("db_pass"); err != nil {
		return errors.New("Failed to bind db_pass ENV variable")
	}
	if err := viper.BindEnv("ldap_host"); err != nil {
		return errors.New("Failed to bind ldap_host ENV variable")
	}
	if err := viper.BindEnv("ldap_bind_dn"); err != nil {
		return errors.New("Failed to bind ldap_bind_dn ENV variable")
	}
	if err := viper.BindEnv("ldap_bind_pass"); err != nil {
		return errors.New("Failed to bind ldap_bind_pass ENV variable")
	}
	if err := viper.BindEnv("ldap_base_dn"); err != nil {
		return errors.New("Failed to bind ldap_base_dn ENV variable")
	}

	return nil
}
