package main

import (
	"DemoTest/dhkh"
	"math/big"

	"fmt"
	"net"
	"log"
)

func chkError(err error) {
	if err != nil {
		log.Fatal(err);
	}
}

func clientHandle(conn *net.UDPConn) {
	defer conn.Close()
	buf := make([]byte, 256)
	_, udpaddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		return
	}
	fmt.Println("Server: Receive Client public key: ")
	fmt.Println(new(big.Int).SetBytes(buf))
	g, _ := dhkx.GetGroup(0)
	priv, err := g.GeneratePrivateKey(nil)
	if (err != nil) {
		fmt.Println("Server: GeneratePrivate Error")
		return
	}
	clientPubKey := dhkx.NewPublicKey(buf)
	if (clientPubKey == nil) {
		fmt.Println("Server: clientpublickey is nil: ")
		return
	}
	k, err2 := g.ComputeKey(clientPubKey, priv)
	if (err2 != nil) {
		fmt.Println("Server: ComputeKey Error: " + err2.Error())
		return
	}

	fmt.Println("Server: Send Server public key: ")
	fmt.Println(new(big.Int).SetBytes(priv.Bytes()))
	fmt.Println("Server: secret key: ")
	fmt.Println(new(big.Int).SetBytes(k.Bytes()))

	conn.WriteToUDP(priv.Bytes(), udpaddr)
}

func main() {
	udpaddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	chkError(err)
	//监听端口
	udpconn, err2 := net.ListenUDP("udp", udpaddr)
	chkError(err2)
	//udp没有对客户端连接的Accept函数
	for {
		clientHandle(udpconn)
	}
}