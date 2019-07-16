package ufw

import (
	"adguardhome/infra"
	"github.com/go-exec/exec"
)

func init() {
	exec.
		Task("provision:ufw", func() {
			exec.RunIfNoBinary("ufw", []string{
				"sudo apt-get install -y ufw",
			})

			exec.Remote("echo 'y' | sudo ufw reset")
			exec.Remote("sudo ufw default deny incoming")
			exec.Remote("sudo ufw allow ssh")
			exec.Remote("sudo ufw allow 80/tcp")
			exec.Remote("sudo ufw allow 443/tcp")
			exec.Remote("sudo ufw allow 53/tcp")
			exec.Remote("sudo ufw allow 53/udp")
			exec.Remote("sudo ufw allow 67/udp")
			exec.Remote("sudo ufw allow 68/tcp")
			exec.Remote("sudo ufw allow 68/tcp")

			exec.Remote("echo 'y' | sudo ufw enable")
			exec.Remote("sudo ufw status numbered")
		}).
		OnServers(func() []string {
			return infra.GetServers()
		})

	exec.Before("provision:ufw", "servers:setup")
}
