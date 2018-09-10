package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPaasInternalCli(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PaaSInternalCli Suite")
}
