package stringset_test

import (
	"fmt"
	"sort"

	"github.com/kentquirk/stringset/v2"
)

func ExampleStringSet_Add() {
	greetings := stringset.New().Add("hello")
	greetings.Add("aloha", "bonjour", "g'day")
	// the Strings() function will return results in random order
	output := greetings.Strings()
	sort.Strings(output)
	fmt.Println(output)
	// Output:
	// [aloha bonjour g'day hello]
}

func ExampleStringSet_Delete() {
	nums := stringset.New().Add("1", "2", "3", "4", "5", "6", "7")
	nums.Delete("2", "4", "6")
	output := nums.Strings()
	sort.Strings(output)
	fmt.Println(output)
	// Output:
	// [1 3 5 7]
}

func ExampleStringSet_Intersection() {
	ss1 := stringset.New().Add("a", "b", "c", "d")
	ss2 := stringset.New().Add("c", "d", "e", "f")

	inter := ss1.Intersection(ss2)
	output := inter.Strings()
	sort.Strings(output)
	fmt.Println(output)
	// Output:
	// [c d]
}
