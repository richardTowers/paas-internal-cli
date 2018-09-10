package main_test

import (
	"github.com/richardTowers/paas-internal-cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Main", func() {
	var origStdout *os.File
	var origStderr *os.File
	var stdout *os.File
	var stderr *os.File

	BeforeEach(func () {
		_, stdout, _ = os.Pipe()
		_, stderr, _ = os.Pipe()
		origStdout = os.Stdout
		os.Stdout = stdout
		origStderr = os.Stderr
		os.Stderr = stderr
	})

	AfterEach(func () {
		os.Stdout = origStdout
		stdout.Close()
		os.Stderr = origStderr
		stderr.Close()
	})

	It("should not error when given valid options", func () {
		Expect(main.Main([]string{"--verbose"})).To(Succeed())
		Expect(main.Main([]string{"some arbitrary arguments for now"})).To(Succeed())
	})

	It("should error when given invalid options", func () {
		Expect(main.Main([]string{"--excessive-verbiage-mode"})).NotTo(Succeed())
		Expect(main.Main([]string{"-u"})).NotTo(Succeed())
	})
})
