package main
//import (
//	"encoding/binary"
//	"flag"
//	"fmt"
//	"net"
//	"os"
//	"time"
//)
//
////go run timeclient.go -host time.nist.gov
//func main() {
//	flag.Parse()
//	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1"+":"+"9005")
//	if err != nil {
//		fmt.Println("Can't resolve address: ", err)
//		os.Exit(1)
//	}
//
//	localaddr, err2 := net.ResolveUDPAddr("udp", "127.0.0.1"+":"+"10")
//	if err2 != nil {
//		fmt.Println("Can't resolve address: ", err2)
//		os.Exit(1)
//	}
//	conn, err := net.DialUDP("udp", localaddr, addr)
//	if err != nil {
//		fmt.Println("Can't dial: ", err)
//		os.Exit(1)
//	}
//	defer conn.Close()
//	_, err = conn.Write([]byte("hello, i'm client"))
//	if err != nil {
//		fmt.Println("failed:", err)
//		os.Exit(1)
//	}
//	data := make([]byte, 4)
//	_, err = conn.Read(data)
//	if err != nil {
//		fmt.Println("failed to read UDP msg because of ", err)
//		os.Exit(1)
//	}
//	t := binary.BigEndian.Uint32(data)
//	fmt.Println(time.Unix(int64(t), 0).String())
//	os.Exit(0)
//}

import (
	"dhkx-master"
	"fmt"
	"log"
	"net"
)

func main() {
	//获取udpaddr
	udpaddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	checkError(err);
	//连接，返回udpconn
	udpconn, err2 := net.DialUDP("udp", nil, udpaddr)
	checkError(err2)
	//写入数据
	//DH Key
	g,_ := dhkx.GetGroup(0)
	priv, _ := g.GeneratePrivateKey(nil)
	pub := priv.Bytes()
	fmt.Println("Client: Send Client public: ")
	fmt.Println(pub)
	_, err3 := udpconn.Write(pub)
	checkError(err3)
	buf := make([]byte, 256)
	//读取服务端发送的数据
	_, err4 := udpconn.Read(buf)
	serverPubKey := dhkx.NewPublicKey(buf)
	secret, _ := g.ComputeKey(serverPubKey, priv)
	checkError(err4)
	fmt.Println("Client: Receive Server public: ")
	fmt.Println(buf)
	fmt.Println("Client: secret: ")
	fmt.Println(secret.Bytes())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}