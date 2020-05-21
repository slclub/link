package link

import (
	"fmt"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	conf := Config()
	addr := conf.GetString("server1.addr")
	fmt.Println("[TEST][CONFIG][ADDR]", addr)

	db_host := conf.GetString("db_main.host")
	db_port := conf.GetString("db_main.host")
	fmt.Println("[TEST][CONFIG][DB1][HOST]", db_host, "[PORT]", db_port)

	time.Sleep(2 * time.Second)
}
