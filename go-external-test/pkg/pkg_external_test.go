package pkg_test

import (
	"testing"

	"github.com/sagikazarmark/please-issues/go-external-test/pkg"
)

func TestHello(t *testing.T) {
	hello := pkg.Hello()

	if hello != pkg.CompareWith {
		t.Error("hello mismatch")
	}
}
