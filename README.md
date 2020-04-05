# Simple dns resolver implemented in go
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/securityguy/dns_resolver)

Forked from bogdanovich/dns_resolver to add IPv6 lookup

Based on based on miekg/dns.

## Features

- Uses provided dns servers array in random order
- Retries dns requests in case of i/o timeout
- IPv4 and IPv6

## Installing

### Using *go get*

    $ go get github.com/securityguy/dns_resolver

After this command *dns_resolver* is ready to use. Its source will be in:

    $GOPATH/src/github.com/securityguy/dns_resolver

## Example

``` go
package main

import (
	"log"
	"github.com/securityguy/dns_resolver"
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

    ip, err := resolver.LookupHost6("google.com")
    if err != nil {
        log.Fatal(err.Error())
    }
    log.Println(ip)
}

```
