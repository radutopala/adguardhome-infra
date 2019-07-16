package generic

import (
	"adguardhome/infra"
	"github.com/go-exec/exec"
)

func init() {
	exec.
		Task("provision:generic", func() {
			exec.Remote("sudo apt-get update -y")
			exec.Remote("sudo service snapd stop; service lxcfs stop; sudo systemctl disable lxd; sudo systemctl disable snapd")

			var binaries = map[string]interface{}{
				"wget": "sudo apt-get install -y wget",
				"git":  "sudo apt-get install -y git",
				"nano": "sudo apt-get install -y nano",
				"htop": "sudo apt-get install -y htop",
				"add-apt-repository": []string{
					"sudo apt-get install -y python-software-properties",
					"sudo apt-get install -y software-properties-common",
				},
			}

			exec.RunIfNoBinaries(binaries)

			//disable resolv listening to 53
			exec.UploadFileSudo("./infra/config/systemd/resolved.conf", "/etc/systemd/resolved.conf")
			exec.Remote("rm -rf /etc/resolv.conf; cd /etc; ln -s /run/systemd/resolve/resolv.conf")
			exec.Remote("sudo service systemd-resolved restart")
		}).
		OnServers(func() []string {
			return infra.GetServers()
		})

	exec.Before("provision:generic", "servers:setup")
}
