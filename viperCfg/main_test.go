package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	pflag.Int("server-port", 8080, "set port of the server")
	c := readConfig()

	t.Run("port was read from 'config.yaml' and not overridden", func(t *testing.T) {
		if c.Id != -1 {
			t.Errorf(cmpMsg(-1, c.Id))
		}
	})

	t.Run("server.port was overridden by 'override.yaml'", func(t *testing.T) {
		if c.Server.Port != 1 {
			t.Errorf(cmpMsg(1, c.Server.Port))
		}
	})

	t.Run("redis.keyPrefix was read from config.yaml and overrode default value from defaultRedisCfg()", func(t *testing.T) {
		if c.Redis.KeyPrefix != "overridden" {
			t.Errorf(cmpMsg("overridden", c.Redis.KeyPrefix))
		}
	})

	t.Run("redis.user was set in defaultRedisCfg() and not read from config files", func(t *testing.T) {
		if c.Redis.User != "non-root" {
			t.Errorf(cmpMsg("non-root", c.Redis.User))
		}
	})

	t.Run("redis.password was not set, so it's empty", func(t *testing.T) {
		if c.Redis.Password != "" {
			t.Errorf(cmpMsg("", c.Redis.Password))
		}
	})

	t.Run("redis.keyPrefix was overriden by environment variable", func(t *testing.T) {
		must(os.Setenv("REDIS.KEYPREFIX", "env-prefix"))
		defer os.Unsetenv("REDIS.KEYPREFIX")
		envC := readConfig()
		if envC.Redis.KeyPrefix != "env-prefix" {
			t.Errorf(cmpMsg("env-prefix", envC.Redis.KeyPrefix))
		}
	})
}

func cmpMsg(ex, ac any) string {
	return fmt.Sprintf("expected: %s, actual: %s", ex, ac)
}
