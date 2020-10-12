package conf

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config is a configuration.
var Config struct {
	CertPath       string `mapstructure:"cert_path"`
	KeyPath        string `mapstructure:"key_path"`
	ClientCertPath string `mapstructure:"client_cert_path"`
	GRPCPort       string `mapstructure:"grpc_port"`
	GatewayPort    string `mapstructure:"gateway_port"`
	JWTKey         string `mapstructure:"jwt_key"`
	HashKey        string `mapstructure:"hash_key"`
	BlockKey       string `mapstructure:"block_key"`

	DBHost string `mapstructure:"db_host"`
	DBName string `mapstructure:"db_name"`
	DBUser string `mapstructure:"db_user"`
	DBPass string `mapstructure:"db_pass"`

	LDAPHost     string `mapstructure:"ldap_host"`
	LDAPBindDN   string `mapstructure:"ldap_bind_dn"`
	LDAPBindPass string `mapstructure:"ldap_bind_pass"`
	LDAPBaseDN   string `mapstructure:"ldap_base_dn"`
}

// Load parses the config from file or from ENV variables s into a Config.
func Load(confName string) error {
	viper.SetConfigName(confName)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("Failed to read config file: %s", err)
	}

	viper.SetEnvPrefix("logserver")
	if err := viper.BindEnv("cert_path"); err != nil {
		return err
	}
	if err := viper.BindEnv("key_path"); err != nil {
		return err
	}
	if err := viper.BindEnv("client_cert_path"); err != nil {
		return err
	}
	if err := viper.BindEnv("grpc_port"); err != nil {
		return err
	}
	if err := viper.BindEnv("gateway_port"); err != nil {
		return err
	}
	if err := viper.BindEnv("jwt_key"); err != nil {
		return err
	}
	if err := viper.BindEnv("hash_key"); err != nil {
		return err
	}
	if err := viper.BindEnv("block_key"); err != nil {
		return err
	}

	if err := viper.BindEnv("db_host"); err != nil {
		return err
	}
	if err := viper.BindEnv("db_name"); err != nil {
		return err
	}
	if err := viper.BindEnv("db_user"); err != nil {
		return err
	}
	if err := viper.BindEnv("db_pass"); err != nil {
		return err
	}

	if err := viper.BindEnv("ldap_host"); err != nil {
		return err
	}
	if err := viper.BindEnv("ldap_bind_dn"); err != nil {
		return err
	}
	if err := viper.BindEnv("ldap_bind_pass"); err != nil {
		return err
	}
	if err := viper.BindEnv("ldap_base_dn"); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}
