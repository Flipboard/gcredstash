package command

import (
	"github.com/Flipboard/gcredstash/src/gcredstash"
	"github.com/mitchellh/cli"
)

// Meta contain the meta-option that nearly all subcommand inherited.
type Meta struct {
	Ui      cli.Ui
	Table   string
	KmsKey  string
	Version string
	Driver  *gcredstash.Driver
}
