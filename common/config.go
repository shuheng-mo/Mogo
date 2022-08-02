package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

// GetConsulConfig set the configuration of Consul
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// the address of consul
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// set prefix, not default prefix
		consul.WithPrefix(prefix),
		// remove prefix or not, can access config without prefix
		consul.StripPrefix(true),
	)

	// init
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}

	// load config
	err = config.Load(consulSource)
	return config, err
}
