# AdGuardHome Infra

We all care about privacy, right?

If you do, you should start proxying your network's DNS requests through [AdGuardHome](https://github.com/AdguardTeam/AdGuardHome) and filter out the unwanted like ad trackers, malware, phishing and adult websites.

## Setup

An AdGuardHome instance will be installed on your own server, to fully comply with your own rules.

### A. Prerequisites

#### Server

A fresh Server with:
* Ubuntu 18.04
* a public IP
* at least 1 CPU and 1 GB of RAM
* ssh access to the server using your local private key

Few recommended hosting providers: [Hetzner](https://www.hetzner.com/), [DigitalOcean](https://www.digitalocean.com).

#### Clone this repo locally
`git clone git@github.com:radutopala/adguardhome-infra.git` 

#### Go installed

Install Go language from https://golang.org/dl/.

### B. Update configuration

Check `./config.json` and update:
* `server.ip` to your server's public IP
* `adguardhome.auth.user` and `adguardhome.auth.password` to some random long values, for basic auth on the AdGuardHome's Dashboard
* `caddy.domain` to a domain where you want to see the AdGuard's Dashboard

### C. Domain A entry
Add an A entry into your DNS manager for `caddy.domain` to your `server.ip`.

### D. Build the infra
```
./main.go build:infra
```

### E. Dashboard
You should now be able to see the AdGuardHome's Dashboard at `caddy.domain`.

### F. Local router's DNS setup
In the Dashboard open the Setup Guide (at `/#guide`). 

Recommended is to setup the router's DNS, using the `server.ip`. You should be using only this `server.ip`, to proxy DNS requests through it and filter out.

---

## AdGuardHome Update

You can easily update your AdGuardHome instance to the latest version by running: 
```
./main.go provision:adguardhome --update
```

---

## Mac on Dev
In case you want to test locally first on your Mac:
```
local/bin/mac/adguardhome -c config/config.yml -w local
```

--- 

## Credits

This repo uses https://github.com/go-exec/exec, to build the infra.

Special thanks to [AdGuardHome](https://github.com/AdguardTeam/AdGuardHome)'s team for their very useful privacy tool.
