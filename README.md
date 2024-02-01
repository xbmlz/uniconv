# Uniconv

[![Build Status](https://github.com/xbmlz/uniconv/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/features/actions)
[![Coverage Status](https://coveralls.io/repos/github/xbmlz/uniconv/badge.svg?branch=main)](https://coveralls.io/github/xbmlz/uniconv?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/xbmlz/uniconv)](https://goreportcard.com/report/github.com/xbmlz/uniconv)
[![Go Doc](https://godoc.org/github.com/xbmlz/uniconv?status.svg)](https://godoc.org/github.com/xbmlz/uniconv)
[![Code Size](https://img.shields.io/github/languages/code-size/xbmlz/uniconv.svg?style=flat-square)](https://github.com/xbmlz/uniconv)
[![Release](https://img.shields.io/github/release/xbmlz/uniconv.svg?style=flat-square)](https://github.com/xbmlz/uniconv/releases)

Using LibreOffice as a server for converting documents.

## Requirements

- Go 1.18 or later

- [Apache OpenOffice](https://www.openoffice.org/) or [LibreOffice](https://www.libreoffice.org/); the latest stable version is usually recommended.

## Installation

```bash
go get -u github.com/xbmlz/uniconv
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/xbmlz/uniconv"
)

func main() {
    // create a new processor, the default port is 2002
    p := uniconv.NewProcessor()
    
    // Start an office process and connect to the started instance (on port 2002).
    p.Start()
    defer p.Stop()
    

    // Convert
    c := uniconv.NewConverter()
    c.Convert("input.docx", "output.pdf")
}
```