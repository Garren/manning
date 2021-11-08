package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBddWithGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BddWithGinkgo Suite")
}
