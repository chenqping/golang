package main

import (
	"os"
	"strings"
	"net"
	"log"
	"io"
	"sync"
)

func main(){
	var wg sync.WaitGroup

	if len(os.Args) < 2{
		log.Fatal("server address not assigned!\n")
	}
	for _, serverItem := range os.Args[1:]{
		serverInfo := strings.Split(serverItem, "=")
		if len(serverInfo) != 2 || strings.ContainsRune(serverInfo[1], ':') != true{
			log.Println("skip the wrong parameter")
			continue
		}
		wg.Add(1)
		go func(serverName, serverAddress string){
			log.Println("connecting ", serverName, serverAddress)
			conn, err := net.Dial("tcp", serverAddress)
			if err != nil{
				log.Fatal(err)
			}
			defer conn.Close()
			mustCopy(os.Stdout, conn)
			wg.Done()
		}(serverInfo[0], serverInfo[1])
	}

	wg.Wait()
}


func mustCopy(dst io.Writer, src io.Reader){
	if _, err := io.Copy(dst, src); err != nil{
		log.Fatal(err)
	}
}
