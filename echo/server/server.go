package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/tls"
)

type echoServer struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool
}

func (es *echoServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("echo server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
	return gnet.None
}

func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
	defer func() {
		c.Close()
	}()
	buf, _ := c.Next(-1)
	c.Write(buf)
	return gnet.None
}

func main() {
	var port int
	var multicore bool
	cerRSA, err := tls.LoadX509KeyPair("../../certs/server.rsa4096.crt", "../../certs/server.rsa4096.key")
	if err != nil {
		log.Println(err)
		return
	}
	cerDSA, err := tls.LoadX509KeyPair("../../certs/server.ecdsa.secp521r1.crt", "../../certs/server.ecdsa.secp521r1.key")
	if err != nil {
		log.Println(err)
		return
	}
	// server only uses TLS 1.2 and TLS 1.3
	config := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{cerRSA, cerDSA},
	}
	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9003, "--port 9003")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()
	echo := &echoServer{addr: fmt.Sprintf("tcp://:%d", port), multicore: multicore}
	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(multicore), gnet.WithTLS(config)))
}
