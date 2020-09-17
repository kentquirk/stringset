package stringset_test

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kentquirk/stringset/v2"
)

func Example() {
	s := "this is a test it is only a test"
	nodups := stringset.New().Add(strings.Split(s, " ")...).Strings()
	sort.Strings(nodups)
	fmt.Println(nodups)
	// Output:
	// [a is it only test this]
}
