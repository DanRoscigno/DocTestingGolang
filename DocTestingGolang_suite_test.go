package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDocTestingGolang(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DocTestingGolang Suite")
}
