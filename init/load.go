package init

import (
	"errors"

	"github.com/spf13/viper"
)

type loadFunc func(*viper.Viper) error

func loadConfig(viper *viper.Viper) error {
	if viper == nil {
		return errors.New("viper is nil")
	}

	return nil
}
