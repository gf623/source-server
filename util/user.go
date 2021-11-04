package util

import (
	"fmt"
	"path"

	ini "gopkg.in/ini.v1"
)

func LoadConfig(cfgFile string) {
	file := path.Join(path.Dir(cfgFile))
	cfg, err := ini.InsensitiveLoad(file)
	if err != nil {
		fmt.Println(cfg)
	}
}
