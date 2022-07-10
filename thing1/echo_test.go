package thing1_test

import (
	"github.com/dumpsterfireproject/module-test/thing1"
	"testing"
)

func TestSpeak(t *testing.T) {
	got, _ := thing1.Speak()
	if got != "foo" {
		t.Fail()
	}
}
