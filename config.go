package link

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

var mu sync.Mutex
var conf *viper.Viper
var app_path, _ = os.Getwd()

func init() {
	conf = newConfig()
	logInit(conf)
}

// From here get viper struct *viper.Viper.
func Config() *viper.Viper {
	return conf
}

func GetString(key, default_value string) string {
	val := conf.GetString(key)
	if val != "" {
		return val
	}

	return default_value
}

func newConfig() *viper.Viper {
	config_path, err := os.Getwd()
	if err != nil {
		panic("[LINK][config] path error:" + err.Error())
	}
	var env string
	if env = os.Getenv("ENV"); env == "" {
		env = "dev"
	}

	vp := viper.New()
	vp.SetConfigType("ini")
	vp.SetConfigName("go.ini")
	vp.AddConfigPath(filepath.Join(config_path, "etc"))
	err = readConfig(vp)
	if err != nil {
		fmt.Println(err)
	}
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("[LINK][VIPER][CONFIG][CHANGED]:", e.Name)
	})

	return vp
}

func readConfig(vp *viper.Viper) error {
	mu.Lock()
	defer mu.Unlock()
	err := vp.ReadInConfig()
	if err != nil {
		return errors.New("[LINK][VIPER][CONFIG][READ]")
	}
	return nil
}
