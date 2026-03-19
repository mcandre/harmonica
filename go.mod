module github.com/mcandre/harmonica

go 1.26.1

require (
	github.com/magefile/mage v1.16.1
	github.com/mcandre/mx v0.0.47
	github.com/saracen/fastzip v0.2.0
)

tool (
	github.com/alexkohler/nakedret/v2/cmd/nakedret
	github.com/kisielk/errcheck
	github.com/magefile/mage
	github.com/mcandre/tuco/cmd/tuco
	honnef.co/go/tools/cmd/staticcheck
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/alexkohler/nakedret/v2 v2.0.6 // indirect
	github.com/kisielk/errcheck v1.9.0 // indirect
	github.com/klauspost/compress v1.18.4 // indirect
	github.com/mcandre/tuco v0.0.22 // indirect
	github.com/saracen/zipextra v0.0.0-20250129175152-f1aa42d25216 // indirect
	golang.org/x/exp/typeparams v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
	golang.org/x/mod v0.33.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/tools v0.42.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)
