package init

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func parseLocalConfig(viper *viper.Viper, configFile string, load loadFunc) error {
	log.Printf("read local config file ...\n -c filename: [%s]", configFile)

	viper.SetConfigName(configFile)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("parse local config failed, err: " + err.Error())
		return err
	}

	if err = load(viper); err != nil {
		log.Println("load local config failed, err: " + err.Error())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("local config changed, event name[%s] value: %s\n", in.Name, in.String())

		if err = load(viper); err != nil {
			log.Println("load local config failed, err: " + err.Error())
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
		log.Println("parse remote config failed, err: " + err.Error())
		return err
	}

	viper.SetConfigType(configType)
	if err = viper.ReadRemoteConfig(); err != nil {
		log.Println("read remote config failed, err: " + err.Error())
		return err
	}

	if err = load(viper); err != nil {
		log.Println("load remote config failed, err: " + err.Error())
		return err
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("WatchRemoteConfig panic: %v\n", err)
			}
		}()

		for {
			time.Sleep(time.Second * 5)

			var err error
			if err = viper.WatchRemoteConfig(); err != nil {
				log.Println("watch remote config failed, err: " + err.Error())
				continue
			}

			if err = load(viper); err != nil {
				log.Println("load remote config failed, err: " + err.Error())
			}
		}
	}()

	return nil
}
