package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nyushi/unixtime/parse"
)

var (
	showFormats = flag.Bool("f", false, "Show available formats")
)

func main() {
	flag.Parse()
	now := time.Now()
	if *showFormats {
		for _, v := range parse.Formats() {
			fmt.Println(v)
		}
		os.Exit(-1)
	}
	if len(flag.Args()) == 0 {
		// Show unix time from current time
		fmt.Println(now.Unix())
		os.Exit(0)
	}
	utime := parse.FromUnix(strings.Join(flag.Args(), " "))
	if utime != nil {
		// Show rfc3339 from unix time
		fmt.Println(utime.Format(time.RFC3339Nano))
		os.Exit(0)
	}
	if t := parse.FromDateString(strings.Join(flag.Args(), " ")); t != nil {
		// Show unix time from string
		fmt.Println(t.Unix())
		os.Exit(0)
	}
}


