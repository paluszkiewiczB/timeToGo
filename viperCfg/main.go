package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

type cfg struct {
	Id     int
	Server struct {
		Port int
	}
	Redis *redisCfg
}

func (c *cfg) String() string {
	return fmt.Sprintf("{Id: %d, Server.Port: %d, Redis: %s}", c.Id, c.Server.Port, c.Redis)
}

func defaultCfg() *cfg {
	return &cfg{
		Redis: defaultRedisCfg(),
	}
}

func defaultRedisCfg() *redisCfg {
	return &redisCfg{
		User:      "non-root",
		KeyPrefix: "time_to_go",
	}
}

type redisCfg struct {
	Nodes     []string
	User      string
	Password  string
	KeyPrefix string
}

func (r *redisCfg) String() string {
	return fmt.Sprintf("{Nodes: %s, User: %s, Password: %s, KeyPrefix: %s}", r.Nodes, r.User, r.Password, r.KeyPrefix)
}

func main() {
	pflag.Int("server-port", 8080, "set port of the server")
	readConfig()
}

func readConfig() *cfg {
	v := newViper("config.yaml")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	must(v.ReadInConfig())

	c := defaultCfg()
	log.Printf("default config struct: %+v\n", c)

	must(v.Unmarshal(c))
	log.Printf("after reading config.yaml: %+v\n", c)

	override := newViper("override.yaml")
	// flag is read before reading override.yaml, so it overrides all the other values
	// from default struct, config.yaml, override.yaml and environment variable
	must(override.BindPFlag("server.port", pflag.CommandLine.Lookup("server-port")))
	must(override.ReadInConfig())
	// unmarshalling to c again - there is no merging logic, maps and lists will be simply overridden
	must(override.Unmarshal(c))
	log.Printf("after reading override.yaml: %+v\n", c)

	return c
}

func newViper(cfgFile string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(cfgFile)
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	return v
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
