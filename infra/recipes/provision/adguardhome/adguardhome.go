package adguardhome

import (
	"adguardhome/infra"
	"github.com/go-exec/exec"
)

func init() {
	exec.
		Task("provision:adguardhome", func() {
			exec.RunIf("[ ! -e \"`which ~/adguardhome/bin/adguardhome`\" ]", []string{
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
		OnServers(func() []string {
			return infra.GetServers()
		})

	exec.Before("provision:adguardhome", "servers:setup")
}
