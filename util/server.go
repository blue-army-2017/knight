package util

import (
	"fmt"
	"strings"
)

func GetServerUri() string {
	if strings.Contains(Config.Host, "localhost") {
		return fmt.Sprintf("%s:%s", Config.Host, Config.Port)
	} else {
		return Config.Host
	}
}
