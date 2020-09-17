package link

import (
	"testing"
	//"time"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	conf := Config()
	addr := conf.GetString("server1.addr")
	t.Log("[TEST][CONFIG][ADDR]", addr)

	db_host := conf.GetString("db_main.host")
	db_port := conf.GetString("db_main.host")
	t.Log("[TEST][CONFIG][DB1][HOST]", db_host, "[PORT]", db_port)

	// not exist
	val := GetString("form.not exist", "Not found")
	assert.Equal(t, "Not found", val)

	fmt.Println("app path:", APP_PATH)
}

func TestGetSize(t *testing.T) {

	form_size := GetSize("form.multipart_memory", 0)

	assert.Equal(t, 32<<20, form_size)
	fmt.Println("CONFIG.FORM.MUTILPART_MEMORY", form_size)

	//for cover testing
	size := GetSize("form.size_k", 0)
	assert.Equal(t, 1<<10, size)
	size = GetSize("form.size_kk", 0)
	assert.Equal(t, 1<<10, size)
	size = GetSize("form.size_mm", 0)
	assert.Equal(t, 1<<20, size)
	size = GetSize("form.size_g", 0)
	assert.Equal(t, 1<<30, size)
	size = GetSize("form.size_gg", 0)
	assert.Equal(t, 1<<30, size)

	// not exist
	size = GetSize("form.not exist key", 25)
	assert.Equal(t, 25, size)

	// parse error
	size = GetSize("log.name", 30)
	assert.Equal(t, 30, size)
}
