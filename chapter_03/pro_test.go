package chapter03_test

import (
	chapter03 "DDD/chapter_03"
	"testing"
)

func TestPoint(t *testing.T) {
	a := chapter03.NewPoint(1, 1)
	b := chapter03.NewPoint(1, 1)
	if a != b {
		t.Fatal("a and b were not equal")
	}
}