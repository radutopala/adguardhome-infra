package supervisor

import (
	"adguardhome/infra"
	"github.com/go-exec/exec"
)

func init() {
	exec.
		Task("provision:supervisor", func() {
			exec.RunIfNoBinary("supervisorctl", []string{
				"sudo apt install -y supervisor",
			})

			exec.Remote("sudo update-rc.d supervisor defaults")
		}).
		OnServers(func() []string {
			return infra.GetServers()
		})

	exec.Before("provision:supervisor", "servers:setup")
}
