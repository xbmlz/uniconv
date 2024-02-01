package uniconv

import "testing"

func TestWithPath(t *testing.T) {
	c := NewConverter(WithPath("path"))

	if c.path != "path" {
		t.Errorf("WithPath() = %v, want %v", c.path, "path")
	}
}

func TestWithHost(t *testing.T) {
	c := NewConverter(WithHost("host"))

	if c.host != "host" {
		t.Errorf("WithHost() = %v, want %v", c.host, "host")
	}
}

func TestWithPort(t *testing.T) {
	c := NewConverter(WithPort(3000))

	if c.port != 3000 {
		t.Errorf("WithPort() = %v, want %v", c.port, 3000)
	}
}
