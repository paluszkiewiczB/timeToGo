# airLiveReload

Example of how to use [air](https://github.com/cosmtrek/air) live reload for Go apps

## Prerequisites

1. [Install Air](https://github.com/cosmtrek/air#installation)
2. Create air.toml

 ```shell
 go install github.com/cosmtrek/air@latest
 air init
 air
 ```

### Optional, but recommended

1. Add temporary directory used by air to .gitignore

```shell
stat .gitignore>/dev/null || touch .gitignore
grep -xqF -- "tmp" ".gitignore" || echo "tmp" >> ".gitignore"
```
