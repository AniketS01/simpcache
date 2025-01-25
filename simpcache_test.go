package simpcache

import (
	"testing"
)

func TestSimpCache(t *testing.T) {
	tc := New[string, int](NoExp, 0)

	tc.add("john", 5, DefExp)

}
