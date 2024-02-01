package uniconv

import "testing"

func TestNewConverter(t *testing.T) {
	c := NewConverter()

	if c.host == "" || c.port == 0 {
		t.Errorf("NewConverter() = %v, want %v", c, "localhost:2002")
	}
}

func TestConvert(t *testing.T) {
	c := NewConverter()

	err := c.Convert("./testdata/test.docx", "./testdata/test.pdf")

	if err != nil {
		t.Errorf("Convert() = %v, want %v", err, nil)
	}
}
