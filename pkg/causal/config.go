package causal

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config contains all the data necessary for causal to run
type Config struct {
	config     *viper.Viper
	WatcherDir string
	AlerterDir string
}

// SetupConfig returns a viper that contains all available config
func SetupConfig() (v *viper.Viper, err error) {

	v = viper.New()

	return v, nil

	v.SetConfigName("causal.conf")    // name of config file (without extension)
	v.AddConfigPath("/etc/causal/")   // path to look for the config file in
	v.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	v.AddConfigPath(".")              // optionally look for config in the working directory
	err = v.ReadInConfig()            // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		return nil, fmt.Errorf("fatal error config file: %s", err)
	}

	return v, nil

}
