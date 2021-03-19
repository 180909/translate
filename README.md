## Installation
To install translate package  you need to install Go and set your Go workspace first.
                             
1. The first need [Go](https://golang.org/) installed (**version 1.12+ is required**), then you can use the below Go command to install translate.

```sh
$ go get -u github.com/180909/translate
```

2. Import it in your code
```go
import "github.com/180909/translate"
```

## Demo

```sh
$ cat example.go
```

```go
package main

import (
	"fmt"
	"github.com/180909/translate"
)

func main() {
	str := translate.Translate("你好", 0)
	fmt.Print(str)
}
```

```
# run example.go
$ go run example.go
```

## Translate
|types|function|
|-----|--------|
|0|Chinese to English|
|1|Chinese to Japanese|
|2|Chinese to Korean|
|3|Chinese to French|
|4|Chinese to Russian|
|5|Chinese to Spanish|
|6|English to Chinese|
|7|Japanese to Chinese|
|8|Korean to Chinese|
|9|French to Chinese|
|10|Russian to Chinese|
|11|Spanish to Chinese|