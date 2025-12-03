package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestPanicUsage(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "panic_test.go")
}

func TestPanicOtherPackage(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "other_package/panic_other_package.go")
}

func TestLogFatalUsage(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "log_fatal_test.go")
}

func TestLogFatalOtherPackage(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "other_package/log_fatal_other_package.go")
}

func TestOsExitUsage(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "os_exit_test.go")
}

func TestOsExitOtherPackage(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "other_package/os_exit_other_package.go")
}

func TestAllowedInMain(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "allowed_in_main.go")
}

func TestMixedCases(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "mixed_cases.go")
}
