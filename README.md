# Simple dns resolver implemented in go
[![Build Status](https://travis-ci.org/bogdanovich/dns_resolver.svg?branch=master)](https://travis-ci.org/bogdanovich/dns_resolver)
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/bogdanovich/dns_resolver)

Based on based on miekg/dns.

## Features

- Uses provided dns servers array in random order
- Retries dns requests in case of i/o timeout

## Installing

### Using *go get*

    $ go get github.com/bogdanovich/dns_resolver

After this command *dns_resolver* is ready to use. Its source will be in:

    $GOPATH/src/github.com/bogdanovich/dns_resolver

## Example

``` go
package main

import (
	"log"
	"github.com/bogdanovich/dns_resolver"
)

func main() {
	resolver := dns_resolver.New([]string{"8.8.8.8", "8.8.4.4"})
	// OR
	// resolver := dns_resolver.NewFromResolvConf("resolv.conf")

	// In case of i/o timeout
	resolver.RetryTimes = 5

	ip, err := resolver.LookupHost("google.com")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(ip)
	// Output [216.58.192.46]
}

```
