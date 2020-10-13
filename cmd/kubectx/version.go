package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

// go build x -ldflags="-X main.version=123"
// expected default value from goreleaser -X main.version={{.Version}}
var (
	version = "unknown"
)

// VersionOp show app version.
type VersionOp struct{}

func (_ VersionOp) Run(stdout, _ io.Writer) error {
	return printVersion(stdout)
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s\n", version)
	return errors.Wrap(err, "write error")
}
