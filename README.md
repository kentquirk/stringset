# stringset
Implementation of a Set of strings in Go. Supports all the standard set operations.

Package stringset provides the capabilities of a standard set with strings, including
the standard set operations (add, delete, check membership, intersection, union) and a
specialized ability to have a "negative" set for intersections.

Negative sets are useful for inverting the sense of a set when combining them
using Intersection. Negative sets do not change the behavior of any other
set operations, including Union and Difference. Negative operations were created
for managing a set of tags, where it is useful to be able to search for items
that contain some tags but do not contain others.

The API is designed to be chainable so that common operations become one-liners,
and it is designed to be easy to use with slices of strings.

This package does not return errors. Set operations should be fast and
chainable; adding items more than once, or attempting to remove things
that don't exist are not errors.

**NOTE**
This repository contains two versions of this code, conforming to Go's versioning rules.
Version 1, the original, is in the root of the repository, and conforms to the specification noted above (with negative sets). Documentation for that version is
[here](https://pkg.go.dev/github.com/kentquirk/stringset).

Version 2 removes the negative sets (they're not generally useful), cleans up the tests, and is a proper Go module. [Documentation](https://pkg.go.dev/github.com/kentquirk/stringset/v2).

# Installation

```bash
$ go get github.com/kentquirk/stringset/v2
```

Example usage:

```go
s := "this is a test it is only a test"
nodups := stringset.New().Add(strings.Split(s, " ")...).Strings()
sort.Strings(nodups)
fmt.Println(nodups)
```
Result:
```text
[a is it only test this]
```
