# viperCfg

Example of how to use [viper](https://github.com/spf13/viper) to configure the application.

## Cases covered

- [X] read configuration from .yaml file
- [X] override (patch) the configuration with other file - useful for bundling default config in Docker container and
  mounting patches through Docker volumes or Kubernetes config maps
- [X] override patched config with environment variable (default behaviour of viper after
  calling [viper.AutomaticEnv()](https://pkg.go.dev/github.com/spf13/viper#AutomaticEnv))
  ```shell
  # cannot use 'export REDIS.KEYPREFIX=env because of the dot character
  env "REDIS.KEYPREFIX=setByEnv" go run main.go
  ```
- [X] override value by a flag (**no unit test**). Flags must be read manually (in this example with [pflag](https://github.com/spf13/pflag))
  ```shell
    go run main.go --server-port -17
  ```
  
## Order of precedence

> Flag overrides everything else

Values are loaded in following order, starting from least important (lowest priority):
1. default struct values (see `func defaultRedisCfg()`)
2. file config.yaml
3. file override.yaml
4. environment variable
5. flag 
