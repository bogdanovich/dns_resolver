package dns_resolver

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/miekg/dns"
)

func TestNew(t *testing.T) {
	servers := []string{"8.8.8.8", "8.8.4.4"}
	expectedServers := []string{"8.8.8.8:53", "8.8.4.4:53"}
	resolver := New(servers)

	if !reflect.DeepEqual(resolver.Servers, expectedServers) {
		t.Error("resolver.Servers: ", resolver.Servers, "should be equal to", expectedServers)
	}
}

func TestNewFromResolvConf_ValidFile(t *testing.T) {
	path := "resolv.conf"
	resolver, err := NewFromResolvConf(path)
	config, _ := dns.ClientConfigFromFile(path)

	expectedServers := []string{}
	for _, server := range config.Servers {
		expectedServers = append(expectedServers, server+":53")
	}

	if !reflect.DeepEqual(resolver.Servers, expectedServers) {
		t.Error("resolver.Servers: ", resolver.Servers, "should be equal to", expectedServers)
	}

	if err != nil {
		t.Error("Should parse config file without errors. Error: ", err.Error())
	}
}

func TestNewFromResolvConf_InvalidFile(t *testing.T) {
	path := "resolv_err.conf"
	_, err := NewFromResolvConf(path)

	if err.Error() != "no such file or directory: resolv_err.conf" {
		t.Error("Should return error")
	}
}

func TestLookupHost_ValidServer(t *testing.T) {
	resolver := New([]string{"8.8.8.8", "8.8.4.4"})
	result, err := resolver.LookupHost("google-public-dns-a.google.com")
	if err != nil {
		fmt.Println(err.Error())
		t.Error("Should succeed dns lookup")
	}

	if result[0].String() != "8.8.8.8" {
		t.Error(fmt.Sprintf("google-public-dns-a.google.com should be resolved to 8.8.8.8. Got: %s", result[0].String()))
	}
}

func TestLookupTXT_ValidServer(t *testing.T) {
	resolver := New([]string{"8.8.8.8", "8.8.4.4"})
	result, err := resolver.LookupTXT("google-public-dns-a.google.com")
	if err != nil {
		fmt.Println(err.Error())
		t.Error("Should succeed dns txt lookup")
	}

	if result[0] != "http://xkcd.com/1361/" {
		t.Error(fmt.Sprintf("google-public-dns-a.google.com should have txt record http://xkcd.com/1361/, got: %s", result[0]))
	}
}
