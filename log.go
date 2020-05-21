package link

import (
	"github.com/slclub/glog"
	"github.com/spf13/viper"
)

func logInit(conf *viper.Viper) {
	glog.Set("path", GetString("log.abs_path", app_path), conf.GetString("log.rel_path"))
	glog.Set("name", GetString("log.name", "glog"))
	glog.Set("head", GetString("log.head", ""))
	glog.Set("stderr", conf.GetBool("log.stderr"))
	glog.Set("debug", conf.GetBool("server1.debug"))
}

func INFO(args ...interface{}) {
	glog.InfoDepth(1, args...)
}

func DEBUG(args ...interface{}) {
	glog.DebugDepth(1, args...)
}

func WARN(args ...interface{}) {
	glog.WarnningDepth(1, args...)
}

func ERROR(args ...interface{}) {
	glog.ErrorDepth(1, args...)
}

func FATAL(args ...interface{}) {
	glog.FatalDepth(1, args...)
}
