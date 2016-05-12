package main

import (
	"flag"
	"fmt"
	"net/url"
)

// URLFlag defines a URL flag with specified name, default value, and usage string.
// The return value is the address of a urlFlag variable that stores the value of the flag.
// URLFlag panics if it fails to parse default value with url.Parse.
func URLFlag(name string, value string, usage string) *urlFlag {
	url, err := url.Parse(value)
	if err != nil {
		panic(fmt.Errorf("URLFlag: default value failed to url.Parse: %v", err))
	}
	p := &urlFlag{URL: *url}
	flag.CommandLine.Var(p, name, usage)
	return p
}

// urlFlag implements flag.Value for a url.URL flag.
type urlFlag struct{ url.URL }

func (u *urlFlag) Set(s string) error {
	url, err := url.Parse(s)
	if err != nil {
		return err
	}
	u.URL = *url
	return nil
}

func (u *urlFlag) String() string { return u.URL.String() }
