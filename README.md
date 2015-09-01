# Simple dns resolver implemented in Go based on based on miekg/dns.

Uses provided dns servers in random order.


# Building

Building is done with the `go` tool. If you have setup your GOPATH
correctly, the following should work:

    go get github.com/bogdanovich/dns_resolver
    go build github.com/bogdanovich/dns_resolver

## Examples

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
