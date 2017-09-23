package main

import (
	"bufio"
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

	f, err := os.Open("templates/adventure.txt")
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
		time.Sleep(20 * time.Millisecond)
	}

	if cli {
		fmt.Fprintf(w, reset)
	} else {
		fmt.Fprintf(w, HTMLEND)
	}
}

func main() {
	http.HandleFunc("/adventure", render)
	http.ListenAndServe(":9000", nil)
}
