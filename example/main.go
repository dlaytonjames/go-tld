package main

import (
	"fmt"

	"github.com/jpillora/go-tld"
)

func main() {
	u, _ := tld.Parse("http://a.very.complex-domain.co.uk:8080/foo/bar")

	fmt.Printf("[ %s ] [ %s ] [ %s ] [ %s ] [ %s ]", u.Subdomain, u.Domain, u.TLD, u.Port, u.Path)
}
