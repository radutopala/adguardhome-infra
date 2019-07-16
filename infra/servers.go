package infra

import (
	"fmt"
	"github.com/go-exec/exec"
)

func GetServers() []string {
	return []string{Config.Get("server.name").String()}
}

func init() {
	exec.Task("servers:setup", func() {
		exec.Server(Config.Get("server.name").String(), fmt.Sprintf("%s@%s", Config.Get("server.user").String(), Config.Get("server.ip").String()))
	}).Private().Once()
}
