package uniconv

import (
	_ "embed"
	"errors"
	"os/exec"
	"strconv"
)

//go:embed uniconv
var script []byte

func NewConverter(options ...Option) *Office {
	o := &Office{}
	for _, option := range options {
		option(o)
	}
	if o.path == "" {
		o.path = getDefaultOfficeHome()
	}
	if o.host == "" {
		o.host = defaultHost
	}
	if o.port == 0 {
		o.port = defaultPort
	}
	return o
}

func (c *Office) Convert(input, output string) error {

	in := fixPath(input)
	out := fixPath(output)

	args := []string{
		"-c",
		string(script),
		in,
		"-O",
		out,
		"-H",
		c.host,
		"-P",
		strconv.Itoa(c.port),
	}

	cmd := exec.Command(c.path+"/program/python", args...)

	o, _ := cmd.CombinedOutput()

	if !cmd.ProcessState.Success() {
		return errors.New(string(o))
	}
	return nil
}
