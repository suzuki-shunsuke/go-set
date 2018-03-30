# go-set

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-set)
[![Build Status](https://travis-ci.org/suzuki-shunsuke/go-set.svg?branch=master)](https://travis-ci.org/suzuki-shunsuke/go-set)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-set/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-set)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-set)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-set)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-set.svg)](https://github.com/suzuki-shunsuke/go-set)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/go-set.svg)](https://github.com/suzuki-shunsuke/go-set/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-set/master/LICENSE)

set data structure for golang.

## Example

```golang
import (
	"fmt"
	"github.com/suzuki-shunsuke/go-set"
)

func main() {
	s := set.NewStrSet("a", "b")
	fmt.Println(s.Len()) // 2
	fmt.Println(s.Has("a")) // true
	fmt.Println(s.HasAll("a", "b")) // true
	fmt.Println(s.HasAny("a", "c")) // true
	s.Add("c")
	s.Adds("d", "e")
	s.Remove("a")
	s.Removes("b", "c")
	fmt.Println(s.ToList()) // []string{"d", "e"}
	// Iteration
	for k, _ := range s.ToMap(false) {
		fmt.Println(k)
	}
	s.Clear()
	fmt.Println(s.Len()) // 0
}
```

## Other libraries

* https://github.com/avelino/awesome-go#data-structures
* https://github.com/deckarep/golang-set
* https://github.com/Workiva/go-datastructures#set
* https://github.com/yourbasic/bit
* https://github.com/xtgo/set
* https://github.com/goware/set
* https://github.com/zoumo/goset

## Change Log

See [CHANGELOG](CHANGELOG.md).

## License

[MIT](LICENSE)
