package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	defaultPort = 8080
	defaultSec  = -1
	defaultName = "NONAME"
)

type serverOption struct {
	port int
	sec  int // if non-positive, wait random sec(1-5)
	name string
}

func main() {
	run(newOption())
}

func run(opts serverOption) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if opts.sec > 0 {
			time.Sleep(time.Duration(opts.sec) * time.Second)
		} else {
			sec := rand.Intn(5) + 1
			time.Sleep(time.Duration(sec) * time.Second)
		}
		fmt.Fprintf(w, opts.name)
	})
	fmt.Printf("%s: SERVER START", opts.name)
	http.ListenAndServe(fmt.Sprintf(":%d", opts.port), nil)
}

func newOption() serverOption {
	return serverOption{port: port(), sec: sec(), name: name()}
}

func port() int {
	str := os.Getenv("PORT")
	port, err := strconv.Atoi(str)
	if err != nil {
		return defaultPort
	}
	return port
}

func sec() int {
	str := os.Getenv("WAITING_SEC")
	sec, err := strconv.Atoi(str)
	if err != nil {
		return defaultSec
	}
	return sec
}
func name() string {
	name := os.Getenv("NAME")
	if name == "" {
		return defaultName
	}
	return name
}
