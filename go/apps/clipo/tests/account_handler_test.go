package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M)  {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
}

func shutdown() {

}
