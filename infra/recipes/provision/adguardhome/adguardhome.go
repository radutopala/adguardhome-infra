package adguardhome

import (
	"adguardhome/infra"
	"github.com/go-exec/exec"
)

func init() {
	updateOpt := exec.NewOption("update", "Update the binary now")
	updateOpt.Type = exec.Bool
	updateOpt.Default = false

	exec.
		Task("provision:adguardhome", func() {
			var condition string

			if exec.TaskContext.GetOption("update").Bool() {
				condition = "[ true ]"
			} else {
				condition = "[ ! -e \"`which ~/adguardhome/bin/adguardhome`\" ]"
			}

			exec.RunIf(condition, []string{
				"mkdir ~/adguardhome",
				"mkdir ~/adguardhome/bin/",
				"mkdir ~/adguardhome/log/",
				"mkdir ~/adguardhome/config/",
				"cd ~/adguardhome; wget https://static.adguard.com/adguardhome/release/AdGuardHome_linux_amd64.tar.gz",
				"tar xvzf AdGuardHome_linux_amd64.tar.gz AdGuardHome/AdGuardHome",
				"mv AdGuardHome/AdGuardHome ~/adguardhome/bin/adguardhome",
				"rm -rf AdGuardHome_linux_amd64.tar.gz",
				"rm -rf AdGuardHome",
			})

			context := infra.Config.Map(infra.Config.Get("adguardhome")).(map[string]interface{})

			exec.UploadTemplateFileSudo("./infra/config/adguardhome/config.yml", "~/adguardhome/config/config.yml", context)
			exec.UploadFileSudo("./infra/config/supervisor/adguardhome.conf", "/etc/supervisor/conf.d/adguardhome.conf")

			exec.RunIf(
				"[ ! \"`sudo supervisorctl status | grep 'no'`\" ]",
				"sudo supervisorctl update; sudo supervisorctl restart all; sudo supervisorctl status; else sudo service supervisor start; sudo supervisorctl status",
			)
		}).
		AddOption(updateOpt).
		OnServers(func() []string {
			return infra.GetServers()
		})

	exec.Before("provision:adguardhome", "servers:setup")
}
