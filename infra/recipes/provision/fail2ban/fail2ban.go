package fail2ban

import (
	"adguardhome/infra"
	"github.com/go-exec/exec"
)

func init() {
	exec.
		Task("provision:fail2ban", func() {
			exec.RunIfNoBinary("fail2ban-client", []string{
				"sudo apt-get install -y --allow-unauthenticated fail2ban",
			})

			exec.UploadFileSudo("./infra/config/fail2ban/jail.local", "/etc/fail2ban/jail.local")

			exec.Remote("sudo service fail2ban restart")
		}).
		OnServers(func() []string {
			return infra.GetServers()
		})

	exec.Before("provision:fail2ban", "servers:setup")
}
