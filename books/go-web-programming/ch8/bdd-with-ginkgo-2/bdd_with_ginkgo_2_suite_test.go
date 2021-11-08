package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBddWithGinkgo2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BddWithGinkgo2 Suite")
}
