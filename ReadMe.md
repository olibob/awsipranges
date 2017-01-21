# awsipranges

awsipranges is a Go package to fetch the IP ranges of Amazon Web Services.

## What does it do?

It gets a JSON file made available by AWS containing all IP ranges of their services and converts it to a struct.

## Installation

Use `go get` to install the package.

```
go get github.com/olibob/awsipranges
```

## Package usage

Import the package: `import "github.com/olibob/awsipranges"`

```GO
package main

import (
	"fmt"
	"log"
  
  "github.com/olibob/awsipranges"
)

func main() {
	ipRanges, err := AwsIPRanges()
	if err != nil {
		log.Fatalf("Couldn't get IP ranges: %s", err)
	}

	ipArray := []string{}

	// Extract ip ranges
	for _, v := range ipRanges.Prefixes {
		if v.Region == "eu-west-1" {
			ipArray = append(ipArray, v.IPPrefix)
		}
	}

	for _, v := range ipArray {
		fmt.Println(v)
	}
}
```
