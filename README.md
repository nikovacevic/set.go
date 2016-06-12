# Set
*Go interface and implementations of a set data structure*

## Install
```bash
go get github.com/nikovacevic/set.go
```
## Import
```go
import "github.com/nikovacevic/set.go"
```
## Interface
```go
Add(e interface{}) bool

Remove(e interface{}) bool

Size() int

Clear() bool

Contains(e interface{}) bool

Equals(s Set) bool

IsSubsetOf(s Set) bool

Union(s Set) Set

Intersection(s Set) Set

Difference(s Set) Set
```
## Implementations
### MapSet
Leverages built in `map` type with `interface{}` keys and `struct{}` zero-width values.
### SliceSet
Coming soon!

## Acknowlegments
This package draws ideas from the following community writing and examples:
- http://dave.cheney.net/2014/03/25/the-empty-struct
- https://github.com/fatih/set
- https://github.com/deckarep/golang-set

**NOTE** Currently, the implemented MapSet is not threadsafe. Once again, see other community imlementations for examples.

# License

The MIT License (MIT)
Copyright (c) 2016 Niko Kovacevic
