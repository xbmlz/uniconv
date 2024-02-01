package uniconv

import "testing"

func TestNewProcessor(t *testing.T) {
	_ = NewProcessor()
}

func TestStart(t *testing.T) {
	p := NewProcessor()
	if err := p.Start(); err != nil {
		t.Errorf("Start() = %v, want %v", err, nil)
	}
}

func TestStop(t *testing.T) {
	p := NewProcessor()
	p.Start()
	if err := p.Stop(); err != nil {
		t.Errorf("Stop() = %v, want %v", err, nil)
	}
}

func TestIsRunning(t *testing.T) {
	p := NewProcessor()
	if p.IsRunning() {
		t.Errorf("IsRunning() = %v, want %v", true, false)
	}
}
