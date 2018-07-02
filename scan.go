package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

const title = `
   ▄▄▄▄    ██▀███   ▄▄▄       ██▓ ███▄    █ 
  ▓█████▄ ▓██ ▒ ██▒▒████▄    ▓██▒ ██ ▀█   █ 
  ▒██▒ ▄██▓██ ░▄█ ▒▒██  ▀█▄  ▒██▒▓██  ▀█ ██▒
  ▒██░█▀  ▒██▀▀█▄  ░██▄▄▄▄██ ░██░▓██▒  ▐▌██▒
  ░▓█  ▀█▓░██▓ ▒██▒ ▓█   ▓██▒░██░▒██░   ▓██░
  ░▒▓███▀▒░ ▒▓ ░▒▓░ ▒▒   ▓▒█░░▓  ░ ▒░   ▒ ▒ 
  ▒░▒   ░   ░▒ ░ ▒░  ▒   ▒▒ ░ ▒ ░░ ░░   ░ ▒░
   ░    ░   ░░   ░   ░   ▒    ▒ ░   ░   ░ ░ 
   ░         ░           ░  ░ ░           ░ 
        ░                                   
    ██████  ▄████▄   ▄▄▄       ███▄    █    
  ▒██    ▒ ▒██▀ ▀█  ▒████▄     ██ ▀█   █    
  ░ ▓██▄   ▒▓█    ▄ ▒██  ▀█▄  ▓██  ▀█ ██▒   
    ▒   ██▒▒▓▓▄ ▄██▒░██▄▄▄▄██ ▓██▒  ▐▌██▒   
  ▒██████▒▒▒ ▓███▀ ░ ▓█   ▓██▒▒██░   ▓██░   
  ▒ ▒▓▒ ▒ ░░ ░▒ ▒  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒    
  ░ ░▒  ░ ░  ░  ▒     ▒   ▒▒ ░░ ░░   ░ ▒░   
  ░  ░  ░  ░          ░   ▒      ░   ░ ░    
        ░  ░ ░            ░  ░         ░    
           ░                                
                                made by wdhg

`

var wg = &sync.WaitGroup{}
var ms = flag.Int64("t", 200, "timeout in milliseconds")

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] host\n", os.Args[0])
	flag.PrintDefaults()
}

// basic connect to port
func scanPort(address string, port int, wait time.Duration) {
	c, err := net.DialTimeout("tcp", address+":"+strconv.Itoa(port), wait)
	if err == nil {
		fmt.Printf("%v/tcp		open\n", port)
		c.Close()
	}
	wg.Done()
}

func main() {
	fmt.Print(title)

	flag.Parse()
	if flag.NArg() != 1 {
		usage()
		return
	}
	host := flag.Arg(0)

	wg.Add(1000)
	fmt.Printf("PORT		STATUS\n")
	for i := 0; i < 1000; i++ {
		go scanPort(host, i, time.Duration(*ms)*time.Millisecond)
	}

	wg.Wait()
}
