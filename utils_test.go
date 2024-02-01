package uniconv

import "testing"

func TestGetDefaultOfficeHome(t *testing.T) {
	path := getDefaultOfficeHome()
	if path == "" {
		t.Errorf("getDefaultOfficeHome() = %v, want %v", path, "path")
	}
}

func TestFixPath(t *testing.T) {
	path := fixPath("path")
	if path == "" {
		t.Errorf("fixPath() = %v, want %v", path, "path")
	}
}

func TestNewUserInstallationDir(t *testing.T) {
	path := newUserInstallationDir()
	if path == "" {
		t.Errorf("newUserInstallationDir() = %v, want %v", path, "path")
	}
}
