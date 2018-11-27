package main

import (
	"dhkx-master"
	//"encoding/binary"
	//"flag"
	"fmt"
	"net"
	//"os"
	//"time"
	"log"
)
//
//func main() {
//	flag.Parse()
//	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1"+":"+"9005")
//	if err != nil {
//		fmt.Println("Can't resolve address: ", err)
//		os.Exit(1)
//	} else {
//		fmt.Println(addr)
//	}
//	conn, err := net.ListenUDP("udp", addr)
//	if err != nil {
//		fmt.Println("Error listening:", err)
//		os.Exit(1)
//	}
//	defer conn.Close()
//	for {
//		handleClient(conn)
//	}
//}
//func handleClient(conn *net.UDPConn) {
//	data := make([]byte, 1024)
//	n, remoteAddr, err := conn.ReadFromUDP(data)
//	if err != nil {
//		fmt.Println("failed to read UDP msg because of ", err.Error())
//		return
//	} else {
//		fmt.Println("already handleClient")
//	}
//	daytime := time.Now().Unix()
//	fmt.Println(n, remoteAddr)
//	b := make([]byte, 4)
//	binary.BigEndian.PutUint32(b, uint32(daytime))
//	conn.WriteToUDP(b, remoteAddr)
//}

func chkError(err error) {
	if err != nil {
		log.Fatal(err);
	}
}

func clientHandle(conn *net.UDPConn) {
	defer conn.Close()
	buf := make([]byte, 256)
	//读取数据
	//注意这里返回三个参数
	//第二个是udpaddr
	//下面向客户端写入数据时会用到
	_, udpaddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		return
	}
	fmt.Println("Server: Receive Client public: ")
	fmt.Println(buf)
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

	fmt.Println("Server: Send Server public: ")
	fmt.Println(priv.Bytes())
	fmt.Println("Server: secret: ")
	fmt.Println(k.Bytes())

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