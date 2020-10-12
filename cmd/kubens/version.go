package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

// go build x -ldflags="-X main.Version=123"
var (
	Version = "unknown"
)

// VersionOp show app version.
type VersionOp struct{}

func (_ VersionOp) Run(stdout, _ io.Writer) error {
	return printVersion(stdout)
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s\n", Version)
	return errors.Wrap(err, "write error")
}
