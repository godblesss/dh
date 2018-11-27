package main
import (
	"DemoTest/dhkh"
	"fmt"
	"log"
	"math/big"
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
	fmt.Println("Client: Send Client public key: ")
	fmt.Println(new(big.Int).SetBytes(pub))
	_, err3 := udpconn.Write(pub)
	checkError(err3)
	buf := make([]byte, 256)
	//读取服务端发送的数据
	_, err4 := udpconn.Read(buf)
	serverPubKey := dhkx.NewPublicKey(buf)
	secret, _ := g.ComputeKey(serverPubKey, priv)
	checkError(err4)
	fmt.Println("Client: Receive Server public key: ")
	fmt.Println(new(big.Int).SetBytes(buf))
	fmt.Println("Client: secret key: ")
	fmt.Println(new(big.Int).SetBytes(secret.Bytes()))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}