package link

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var mu sync.Mutex
var conf *viper.Viper
var app_path, _ = os.Getwd()
var BYTE_NUMBER_MAP map[byte]uint8 = map[byte]uint8{'0': 1, '1': 1, '2': 1, '3': 1, '4': 1, '5': 1, '6': 1, '7': 1, '8': 1, '9': 1}

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

// 获取一些配置字符转化成数字
func GetSize(key string, default_value int) int {
	val := conf.GetString(key)
	if val == "" {
		return default_value
	}

	begin := 0
	unit := 0
	for i, _ := range val {
		v := val[i]
		if _, ok := BYTE_NUMBER_MAP[v]; ok && unit == 0 {
			begin++
			continue
		}
		unit = 1
		if v == 'k' || v == 'K' {
			unit = 1024
			break
		}

		if v == 'm' || v == 'M' {
			unit = 1024 * 1024
			break
		}

		if v == 'g' || v == 'G' {
			unit = 1024 * 1024 * 1024
			break
		}
	}

	size, ret := strconv.Atoi(val[:begin])
	if ret != nil {
		return default_value
	}
	size *= unit
	return size
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
