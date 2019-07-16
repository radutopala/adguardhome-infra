package caddy

import (
	"adguardhome/infra"
	"github.com/go-exec/exec"
)

func init() {
	exec.
		Task("provision:caddy", func() {
			exec.RunIf("[ ! -e \"`which ~/adguardhome/bin/caddy`\" ]", []string{
				"cd ~/adguardhome; curl --output caddy.tar.gz -s https://caddyserver.com/download/linux/amd64?license=personal 2>/dev/null",
				"tar xvzf caddy.tar.gz caddy",
				"mv caddy ~/adguardhome/bin/caddy",
				"rm -rf caddy.tar.gz",
			})

			context := infra.Config.Map(infra.Config.Get("caddy")).(map[string]interface{})

			exec.UploadTemplateFileSudo("./infra/config/caddy/Caddyfile", "~/adguardhome/config/Caddyfile", context)
			exec.UploadFileSudo("./infra/config/supervisor/caddy.conf", "/etc/supervisor/conf.d/caddy.conf")

			exec.RunIf(
				"[ ! \"`sudo supervisorctl status | grep 'no'`\" ]",
				"sudo supervisorctl update; sudo supervisorctl restart all; sudo supervisorctl status; else sudo service supervisor start; sudo supervisorctl status",
			)
		}).
		OnServers(func() []string {
			return infra.GetServers()
		})

	exec.Before("provision:caddy", "servers:setup")
}
