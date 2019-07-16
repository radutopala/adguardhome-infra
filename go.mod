module adguardhome

require (
	github.com/fatih/color v1.7.0
	github.com/go-exec/exec v0.0.0-20190715174909-f3ac22ac3ec0
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/tidwall/gjson v1.1.5
	github.com/tidwall/match v1.0.1 // indirect
	github.com/tidwall/pretty v0.0.0-20180105212114-65a9db5fad51
	github.com/tidwall/sjson v1.0.3
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/sys v0.0.0-20190712062909-fae7ac547cb7 // indirect
)

replace github.com/go-exec/exec v0.0.0-20190715174909-f3ac22ac3ec0 => ../exec
