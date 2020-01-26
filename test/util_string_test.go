package test

import (
	"github.com/RiverDanceGit/yeepayGo/util"
	"testing"
)

func TestUuid(t *testing.T) {
	t.Log(util.Uuid())
}

func TestStrrpos(t *testing.T) {
	foo := "0123456789a123456789b123456789c"
	t.Log(util.Strrpos(foo, "7", -5))
	t.Log(util.Strrpos(foo, "7", 20))
	t.Log(util.Strrpos(foo, "7", 28))
}

func TestStrrchr(t *testing.T) {
	foo := "0123456789a123456789b123456789c"
	t.Log(util.Strrchr(foo, "7"))
}
