# Bogosearch

## Overview

The `bogosearch` package implements a bogo approach to searching for a target value within an integer array.

## Performance

![Alt text](search_growth.png?raw=true "Title")

## Usage

Below is a simple example demonstrating how to use the Bogosearch:

```go
package main

import (
    "fmt"
    "github.com/rshcoding/bogosearch"
)

func main() {
    arr := []int{5, 3, 7, 1, 4, 9, 8}
    target := 4
    index, found := bogosearch.Search(arr, target)
    if found {
        fmt.Printf("bogo found target %d at index %d.\n", target, index)
    } else {
        fmt.Println("bogo was unable to find target")
    }
}
```