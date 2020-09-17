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
var APP_PATH, _ = os.Getwd()
var BYTE_NUMBER_MAP map[byte]uint8 = map[byte]uint8{'0': 1, '1': 1, '2': 1, '3': 1, '4': 1, '5': 1, '6': 1, '7': 1, '8': 1, '9': 1}

func init() {
	APP_PATH = get_root_path()
	conf = newConfig()
	logInit(conf)
}
func get_root_path() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return dir
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
func GetSizeInt64(key string, default_value int64) int64 {
	val := conf.GetString(key)
	if val == "" {
		return default_value
	}

	begin := 0
	var unit int64 = 0
	for i, _ := range val {
		v := val[i]
		if _, ok := BYTE_NUMBER_MAP[v]; ok && unit == 0 {
			begin++
			continue
		}
		unit = 1
		switch v {
		case 'k', 'K':
			unit = 1024
			goto WALK
		case 'm', 'M':
			unit = 1024 * 1024
			goto WALK
		case 'g', 'G':
			unit = 1024 * 1024 * 1024
			goto WALK
		}
	}

WALK:
	size, ret := strconv.ParseInt(val[:begin], 10, 64)
	if ret != nil {
		return default_value
	}
	size *= unit
	return size
}

func GetSize(key string, default_value int) int {
	val := GetSizeInt64(key, int64(default_value))
	return int(val)
}

func newConfig() *viper.Viper {
	APP_PATH, err := os.Getwd()
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
	vp.AddConfigPath(filepath.Join(APP_PATH, "etc"))
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
