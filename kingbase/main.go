package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
)

func main() {
	handshake := "00000060000300006461746162617365005445535400636c69656e745f656e636f64696e6700555446380065787472615f666c6f61745f646967697473003200757365720053595354454d00646174657374796c650049534f2c204d44590000"
	data, err := hex.DecodeString(handshake)
	if err != nil {
		panic(err)
	}
	//conn, err := net.Dial("tcp", "192.168.13.165:4321") // 520000001700
	//conn, err := net.Dial("tcp", "192.168.13.165:54321") // 520000000c00
	//conn, err := net.Dial("tcp", "192.168.13.165:54311") // 520000000c00
	//conn, err := net.Dial("tcp", "192.168.13.165:54301") // 520000000c00 kingbase (Kingbase) V008R003C002B0320
	conn, err := net.Dial("tcp", "192.168.13.165:34321") // 520000001700 KINGBASE (KingbaseES) V008R006C005B0013

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	conn.Write(data)
	var buffer bytes.Buffer
	for {
		buf := make([]byte, 6)
		if buffer.Len() >= 6 {
			break
		}
		conn.Read(buf)
		if err != nil {
			panic(err)
		}
		buffer.Write(buf)
	}
	fmt.Println(hex.EncodeToString(buffer.Bytes()[:6]))
	if hex.EncodeToString(buffer.Bytes()[:4]) == "52000000" {
		//if hex.EncodeToString(buffer.Bytes()[:6]) == "52000000170000000a534352414d2d5348412d3235360000" {
		println("是金仓")
	} else {
		println("不是金仓")
	}
}
