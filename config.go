package main

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	CONF_DIR_KEY            = "directory"
	OAUTH_CLIENT_SECRET_KEY = "oauth.ClientSecret"
	OAUTH_TOKEN_KEY         = "oauth.Token"
	SHEETS_KEY              = "sheets"
)

// Configure viper defaults.
func init() {
	viper.SetDefault(OAUTH_CLIENT_SECRET_KEY, "client_secret.json")
	viper.SetDefault(OAUTH_TOKEN_KEY, "client_token.json")

	pflag.StringP("config", "c", "config.yaml", "YAML formatted config file.")
	pflag.StringP(CONF_DIR_KEY, "d", defaultConfDir(), "Directory to search for config file, authentication secret, and save authentication token.")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigType("yaml")
	loaded := false
	if viper.IsSet("config") {
		f, err := os.Open(viper.GetString("config"))
		if err == nil {
			if viper.ReadConfig(f) == nil {
				loaded = true
			}
		}
	}

	if !loaded {
		viper.AddConfigPath(viper.GetString(CONF_DIR_KEY))
		viper.SetConfigName("config")
		viper.ReadInConfig()
	}
}

func defaultConfDir() string {
	xch, s := os.LookupEnv("XDG_CONFIG_HOME")
	if !s {
		usr, _ := user.Current()
		xch = filepath.Join(usr.HomeDir, ".config")
	}
	return filepath.Join(xch, "astrolabe")
}

// Get a config file path. If the passed parameter is a file that's the path, if it's not assume that the file should be in the config directory.
func confFileFind(f string) string {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		f = filepath.Join(viper.GetString(CONF_DIR_KEY), f)
	}

	return f
}

// Get a config file path from a viper key.
func confFilePath(k string) string {
	return confFileFind(viper.GetString(k))
}
