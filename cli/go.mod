module github.com/way-platform/vin-go/cli

go 1.24.4

require (
	github.com/spf13/cobra v1.9.1
	github.com/way-platform/vin-go v0.0.0
	google.golang.org/protobuf v1.36.10
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
)

replace github.com/way-platform/vin-go => ../
