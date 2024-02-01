package uniconv

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os/exec"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type Processor interface {
	Start() error
	Stop() error
	IsRunning() bool
}

type Process struct {
	Office
	isRunning atomic.Bool
	locker    sync.RWMutex
	cmd       *exec.Cmd
	userDir   string
}

type ProcessOption func(*Process)

func NewProcessor(options ...ProcessOption) Processor {
	o := &Process{}
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

	o.isRunning.Store(false)
	return o
}

func (p *Process) Start() error {
	if p.isRunning.Load() {
		return errors.New("LibreOffice is already running")
	}

	if !isPortAvailable(p.port) {
		return fmt.Errorf("port %d is not available", p.port)
	}

	userProfileDirPath := newUserInstallationDir()

	args := []string{
		"--headless",
		"--invisible",
		"--nocrashreport",
		"--nodefault",
		"--nologo",
		"--nofirststartwizard",
		"--norestore",
		fmt.Sprintf("--accept=socket,host=%s,port=%d,tcpNoDelay=1;urp;StarOffice.ComponentContext", p.host, p.port),
	}

	// run as daemon
	cmd := exec.Command(p.path+"/program/soffice", args...)

	err := cmd.Start()
	if err != nil {
		return errors.New("LibreOffice failed to start")
	}

	waitChan := make(chan error, 1)
	go func() {
		// By waiting the process, we avoid the creation of a zombie process
		// and make sure we catch an early exit if any.
		waitChan <- cmd.Wait()
	}()

	connChan := make(chan error, 1)
	go func() {
		// check tcp
		for {
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", p.host, p.port))
			if err != nil {
				continue
			}
			connChan <- nil

			_ = conn.Close()

			break
		}
	}()

	var success bool
	defer func() {
		if success {
			p.locker.Lock()
			defer p.locker.Unlock()

			p.cmd = cmd
			p.userDir = userProfileDirPath
			p.isRunning.Store(true)
			return
		}

		// Let's make sure the process is killed.
		err := cmd.Process.Signal(syscall.SIGKILL)
		if err != nil {
			log.Fatalf(fmt.Sprintf("kill LibreOffice process: %v", err))
		}

	}()

	for {
		select {
		case err = <-connChan:
			if err != nil {
				return fmt.Errorf("LibreOffice socket not available: %w", err)
			}
			log.Printf("LibreOffice process started")
			success = true
			return nil
		case err = <-waitChan:
			return fmt.Errorf("LibreOffice process exited: %w", err)
		}

	}
}

func (o *Process) Stop() error {
	if !o.isRunning.Load() {
		// No big deal? Like calling cancel twice.
		return nil
	}
	o.locker.Lock()
	defer o.locker.Unlock()

	err := o.cmd.Process.Signal(syscall.SIGKILL)
	if err != nil {
		log.Fatalf(fmt.Sprintf("kill LibreOffice process: %v", err))
	}

	o.userDir = ""
	o.cmd = nil
	o.isRunning.Store(false)
	log.Println("LibreOffice process stopped")
	return nil
}

func (o *Process) IsRunning() bool {
	if !o.isRunning.Load() {
		// Non-started browser but not restarting?
		return false
	}
	o.locker.RLock()
	defer o.locker.RUnlock()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%d", o.port), time.Duration(10)*time.Second)
	if err == nil {
		err = conn.Close()
		if err != nil {
			log.Printf("close connection after health checking LibreOffice: %v", err)
		}

		return true
	}

	return false
}
