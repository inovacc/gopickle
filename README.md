# gopickle [![Test](https://github.com/inovacc/gopickle/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/inovacc/gopickle/actions/workflows/test.yml)

`gopickle` is a Go package that provides functionality similar to Python's `pickle` module. It allows you to serialize and deserialize Go data structures to and from files using the `gob` encoding.

## Installation

To install the package, use the following command:

```sh
go get github.com/inovacc/gopickle/
```

## Usage
Here is an example of how to use the gopickle package:

```go
package main

import (
	"fmt"
	"log"
	"github.com/inovacc/gopickle"
)

func main() {
	data := struct {
		Name string
		Age  int
	}{
		Name: "John Doe",
		Age:  30,
	}

	// Serialize data to a file
	if err := gopickle.Dump("data.pgo", data); err != nil {
		log.Fatalf("Failed to dump data: %v", err)
	}

	var loadedData struct {
		Name string
		Age  int
	}

	// Deserialize data from a file
	if err := gopickle.Load("data.pgo", &loadedData); err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	fmt.Printf("Loaded data: %+v\n", loadedData)
}
```

## Note
This package is a closed implementation of the pickle module that only exists in the Python world. It provides similar functionality for Go developers.

## License
This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.
```
