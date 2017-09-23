package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mgutz/ansi"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	HTMLSTART = `<html><head>
    <style>body{background:black;color:green}</style>
    </head><body><pre>`
	HTMLEND = `</pre></body></html>`
)

var (
	delay   int
	tplPath string
	port    string
)

func isCliClient(ua string) bool {
	cliUA := [3]string{"curl", "Wget", "HTTPie"}

	for _, elem := range cliUA {
		if strings.HasPrefix(ua, elem) {
			return true
		}
	}

	return false
}

func render(w http.ResponseWriter, r *http.Request) {

	green := ansi.ColorCode("green")
	reset := ansi.ColorCode("reset")

	f, err := os.Open(tplPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ua := r.Header["User-Agent"][0]
	fmt.Println(ua)

	cli := isCliClient(ua)
	if cli {
		fmt.Fprintf(w, green)
	} else {
		fmt.Fprintf(w, HTMLSTART)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Fprintln(w, scanner.Text())
		if fls, ok := w.(http.Flusher); ok {
			fls.Flush()
		} else {
			fmt.Println("Damn, no flush")
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	if cli {
		fmt.Fprintf(w, reset)
	} else {
		fmt.Fprintf(w, HTMLEND)
	}
}

func main() {
	flag.StringVar(&port, "p", "9000", "Port to listen on")
	flag.StringVar(&tplPath, "t", "templates/adventure.txt", "Path to template file")
	flag.IntVar(&delay, "d", 20, "Delay between lines")
	flag.Parse()

	http.HandleFunc("/adventure", render)
	http.ListenAndServe(":"+port, nil)
}
