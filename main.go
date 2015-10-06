package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	formats map[string]string = map[string]string{
		"ansic":                   time.ANSIC,
		"unixdate":                time.UnixDate,
		"rubydate":                time.RubyDate,
		"rfc822":                  time.RFC822,
		"rfc822z":                 time.RFC822Z,
		"tfc850":                  time.RFC850,
		"rfc1123":                 time.RFC1123,
		"rfc1123z":                time.RFC1123Z,
		"rfc3339":                 time.RFC3339,
		"rfc3339Nano":             time.RFC3339Nano,
		"kitchen":                 time.Kitchen,
		"apahce httpd common log": "02/Jan/2006:15:04:05 -0700",
		"kibana4":                 "January 2th 2006, 15:04:05.999999999",
	}
	showFormats = flag.Bool("f", false, "Show available formats")
)

func main() {
	flag.Parse()
	now := time.Now()
	if *showFormats {
		keys := make([]string, 0, len(formats))
		for k := range formats {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Printf("%s: %s\n", k, now.Format(formats[k]))
		}
		os.Exit(-1)
	}
	if len(flag.Args()) == 0 {
		// Show unix time from current time
		fmt.Println(now.Unix())
		os.Exit(0)
	}
	utime := parseUnix(strings.Join(flag.Args(), " "))
	if utime != nil {
		// Show rfc3339 from unix time
		fmt.Println(utime.Format(time.RFC3339Nano))
		os.Exit(0)
	}
	if t := parseString(strings.Join(flag.Args(), " ")); t != nil {
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
	for _, format := range formats {
		if t, err := time.Parse(format, s); err == nil {
			return &t
		}
	}
	return nil
}
