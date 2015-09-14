package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		// Show unix time from current time
		fmt.Println(time.Now().Unix())
		os.Exit(0)
	}
	utime := parseUnix(strings.Join(os.Args[1:], " "))
	if utime != nil {
		// Show rfc3339 from unix time
		fmt.Println(utime.Format(time.RFC3339Nano))
		os.Exit(0)
	}
	if t := parseString(strings.Join(os.Args[1:], " ")); t != nil {
		// Show unix time from string
		fmt.Println(t.Unix())
		os.Exit(0)
	}
}

func parseUnix(s string) *time.Time {
	var (
		r       *regexp.Regexp
		matched []string
	)

	// parse unix time with nanoseconds
	r = regexp.MustCompile(`^(\d+)\.(\d+)$`)
	matched = r.FindStringSubmatch(s)
	if len(matched) > 0 {
		sec, _ := strconv.ParseInt(matched[1], 10, 64)
		nanosec, _ := strconv.ParseInt(matched[2], 10, 64)
		t := time.Unix(sec, nanosec)
		return &t
	}

	// parse unix time
	r = regexp.MustCompile(`^(\d+)$`)
	matched = r.FindStringSubmatch(s)
	if len(matched) > 0 {
		sec, _ := strconv.ParseInt(matched[1], 10, 64)
		t := time.Unix(sec, 0)
		return &t
	}

	return nil
}

func parseString(s string) *time.Time {
	for _, format := range []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		"02/Jan/2006:15:04:05 -0700",           // apache httpd common log
		"January 2th 2006, 15:04:05.999999999", // kibana4
	} {
		if t, err := time.Parse(format, s); err == nil {
			return &t
		}
	}
	return nil
}
