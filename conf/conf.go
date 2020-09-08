package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Conf variable contains app configuration.
var Conf struct {
	App  app  `toml:"app"`
	DB   db   `toml:"db"`
	LDAP ldap `toml:"ldap"`
}

// ParseConfig function parse a config file into Conf variable.
func ParseConfig(confPath string) error {
	if _, err := toml.DecodeFile(confPath, &Conf); err != nil {
		return fmt.Errorf("Failed to decode config file from %s", confPath)
	}
	return nil
}

type app struct {
	CertPath       string `toml:"cert_path"`
	KeyPath        string `toml:"key_path"`
	ClientCertPath string `toml:"client_cert_path"`
	ListenPort     string `toml:"listen_port"`
	GatewayPort    string `toml:"gateway_port"`
}

type db struct {
	Host string `toml:"host"`
	Name string `toml:"name"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
}

type ldap struct {
	Host     string `toml:"host"`
	BindDN   string `toml:"bind_dn"`
	BindPass string `toml:"bind_pass"`
	BaseDN   string `toml:"base_dn"`
}
