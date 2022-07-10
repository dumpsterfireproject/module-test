package thing2_test

import (
	"github.com/dumpsterfireproject/module-test/thing2"
	"testing"
)

func TestSpeak(t *testing.T) {
	got, _ := thing2.Speak()
	if got != "foo" {
		t.Fail()
	}
}
