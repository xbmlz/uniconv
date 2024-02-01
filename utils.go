package uniconv

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// getDefaultOfficeHome returns A path that is the directory where lives the first detected office installation.
func getDefaultOfficeHome() (path string) {

	findOfficeHome := func(executablePath string, homePaths ...string) string {
		for _, homePath := range homePaths {
			if _, err := os.Stat(filepath.Join(homePath, executablePath)); err == nil {
				return homePath
			}
		}
		return ""
	}

	// env
	if path = os.Getenv("office.home"); path != "" {
		return
	} else if runtime.GOOS == "windows" {
		programFiles64 := os.Getenv("ProgramFiles")
		programFiles32 := os.Getenv("ProgramFiles(x86)")
		path = findOfficeHome(
			"program/soffice.exe",
			filepath.Join(programFiles64, "LibreOffice"),
			filepath.Join(programFiles64, "LibreOffice 5"),
			filepath.Join(programFiles32, "LibreOffice 5"),
			filepath.Join(programFiles32, "OpenOffice 4"),
			filepath.Join(programFiles64, "LibreOffice 4"),
			filepath.Join(programFiles32, "LibreOffice 4"),
			filepath.Join(programFiles64, "LibreOffice 3"),
			filepath.Join(programFiles32, "LibreOffice 3"),
			filepath.Join(programFiles32, "OpenOffice.org 3"),
		)
	} else if runtime.GOOS == "darwin" {
		homeDir := findOfficeHome(
			"MacOS/soffice",
			"/Applications/LibreOffice.app/Contents",
			"/Applications/OpenOffice.app/Contents",
			"/Applications/OpenOffice.org.app/Contents",
		)
		if homeDir == "" {
			homeDir =
				findOfficeHome(
					"program/soffice",
					"/Applications/LibreOffice.app/Contents",
					"/Applications/OpenOffice.app/Contents",
					"/Applications/OpenOffice.org.app/Contents",
				)
		}
		path = homeDir
	} else {
		path = findOfficeHome(
			"program/soffice.bin",
			"/usr/lib/libreoffice",
			"/usr/local/lib64/libreoffice",
			"/usr/local/lib/libreoffice",
			"/opt/libreoffice",
			"/usr/lib64/openoffice",
			"/usr/lib64/openoffice.org3",
			"/usr/lib64/openoffice.org",
			"/usr/lib/openoffice",
			"/usr/lib/openoffice.org3",
			"/usr/lib/openoffice.org",
			"/opt/openoffice4",
			"/opt/openoffice.org3",
		)
	}

	return
}

// fixPath returns an absolute path if the input path is not absolute.
func fixPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return absPath
}

// newUserInstallationDir generates a new unique path for a directory inside the working
func newUserInstallationDir() string {
	ts := time.Now().UnixNano()
	workingDir := fmt.Sprintf("%s/%s", os.TempDir(), "uniconv")
	return fmt.Sprintf("%s/%d", workingDir, ts)
}

// portIsAvailable checks if a port is available.
func isPortAvailable(port int) bool {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	defer listener.Close()
	return true
}
