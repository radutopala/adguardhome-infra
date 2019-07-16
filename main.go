//`which go` run $0 $@; exit $?;
package main

import (
	"adguardhome/infra"
	_ "adguardhome/infra/recipes/provision/adguardhome"
	_ "adguardhome/infra/recipes/provision/caddy"
	_ "adguardhome/infra/recipes/provision/fail2ban"
	_ "adguardhome/infra/recipes/provision/generic"
	_ "adguardhome/infra/recipes/provision/supervisor"
	_ "adguardhome/infra/recipes/provision/ufw"
	"fmt"
	"github.com/go-exec/exec"
	"time"
)

func main() {
	exec.Task("onStart", func() {
		exec.Set("startTime", time.Now())

		infra.Config.Load(exec.GetOption("config").String())
	}).Private()

	exec.Task("onEnd", func() {
		exec.Println(fmt.Sprintf("Finished in %s!`", time.Now().Sub(exec.Get("startTime").Time()).String()))

		infra.Config.Dump(exec.GetOption("config").String())
	}).Private()

	configOpt := exec.NewOption("config", "Configuration file")
	configOpt.Default = "./config.json"
	exec.AddOption(configOpt)

	exec.TaskGroup("build:infra", "servers:setup", "provision:generic", "provision:ufw", "provision:fail2ban", "provision:supervisor", "provision:adguardhome", "provision:caddy")

	exec.Init()
}
