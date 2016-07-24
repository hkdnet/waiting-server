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
	fmt.Printf("%s START\n", opts.name)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("sleeping...")
		var sec int
		if opts.sec > 0 {
			sec = opts.sec
		} else {
			sec = rand.Intn(5) + 1
		}
		fmt.Printf("%d sec sleeping...\n", sec)
		time.Sleep(time.Duration(sec) * time.Second)
		fmt.Fprintf(w, opts.name)
	})
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
