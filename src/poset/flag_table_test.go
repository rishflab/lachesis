package poset

import (
	"fmt"
	"testing"
)

func TestMergeFlagTable(t *testing.T) {

	ft1 := make(FlagTable)
	ft2 := make(FlagTable)

	ft1["a"] = true
	ft1["b"] = false
	ft1["c"] = false

	ft2["b"] = true
	ft2["c"] = true

	merged := ft1.mergeFlagTable(ft2)

	fmt.Println(merged)

}
