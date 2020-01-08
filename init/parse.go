package init

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func parseLocalConfig(viper *viper.Viper, configFile string, load loadFunc) error {
	log.Info().Msgf("read local config file, filename: [%s]", configFile)

	viper.SetConfigName(configFile)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Err(err).Msg("parse local config failed")
		return err
	}

	if err = load(viper); err != nil {
		log.Err(err).Msg("load local config failed")
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Info().Msgf("local config changed, event name[%s] value: %s", in.Name, in.String())

		if err = load(viper); err != nil {
			log.Err(err).Msg("load local config failed")
		}
	})

	return nil
}

func parseRemoteConfig(viper *viper.Viper, configType, provider, endpoint, path, secretkeyring string, load loadFunc) error {
	log.Printf(`read remote config ...
 -t configType:    [%s]
 -P provider:      [%s]
 -h endpoint:      [%s]
 -p path:          [%s]
 -s secretkeyring: [%s]`,
		configType, provider, endpoint, path, secretkeyring)

	var err error
	if len(secretkeyring) > 0 {
		err = viper.AddSecureRemoteProvider(provider, endpoint, path, secretkeyring)
	} else {
		err = viper.AddRemoteProvider(provider, endpoint, path)
	}
	if err != nil {
		log.Err(err).Msg("parse remote config failed")
		return err
	}

	viper.SetConfigType(configType)
	if err = viper.ReadRemoteConfig(); err != nil {
		log.Err(err).Msg("read remote config failed")
		return err
	}

	if err = load(viper); err != nil {
		log.Err(err).Msg("load remote config failed")
		return err
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Panic().Msg("WatchRemoteConfig panic")
			}
		}()

		for {
			time.Sleep(time.Second * 5)

			var err error
			if err = viper.WatchRemoteConfig(); err != nil {
				log.Err(err).Msg("watch remote config failed")
				continue
			}

			if err = load(viper); err != nil {
				log.Err(err).Msg("load remote config failed")
			}
		}
	}()

	return nil
}
