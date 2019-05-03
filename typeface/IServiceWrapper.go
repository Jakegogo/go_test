package main

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/log template

//go:generate gowrap gen -p go_test/typeface -i IService -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/log -o IServiceWrapper.go

import (
	"io"
	"log"
)

// IServiceWithLog implements IService that is instrumented with logging
type IServiceWithLog struct {
	_stdlog, _errlog *log.Logger
	_base            IService
}

// NewIServiceWithLog instruments an implementation of the IService with simple logging
func NewIServiceWithLog(base IService, stdout, stderr io.Writer) IServiceWithLog {
	return IServiceWithLog{
		_base:   base,
		_stdlog: log.New(stdout, "", log.LstdFlags),
		_errlog: log.New(stderr, "", log.LstdFlags),
	}
}

// Add implements IService
func (_d IServiceWithLog) Add(a int, b int) (i1 int, e1 error, i2 int) {
	_params := []interface{}{"IServiceWithLog: calling Add with params:", a, b}
	_d._stdlog.Println(_params...)
	defer func() {
		_results := []interface{}{"IServiceWithLog: Add returned results:", i1, e1, i2}
		_d._stdlog.Println(_results...)
	}()
	return _d._base.Add(a, b)
}
