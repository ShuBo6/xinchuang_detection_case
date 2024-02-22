package main

import (
	"bytes"
	"encoding/hex"
	"net"
)

func main() {
	//00000000c800510000000000000000000000990000000000000000010200000000010a000000000000000000000000000000000000000000000000080000000382e312e3620040000000000bfb5817584d0f38bb2a0e1f70f2d15d7e512fa885f0ddca1a7f3762b1d26f5cf669e2c9e59699c250813efd22ebc9c2f4daebc40ff9cb686dd5da3c1ed68b47
	handshake := "00000000c80051000000000000000000000000990000000000000000010200000000010a0000000000000000000000000000000000000000000000000000000008000000382e312e332e363200400000000bfb5817584d0f38bb2a0e1f70f2d15d7e512fa885f0ddca1a7f3762b1d26f5cf669e2c9e59699c250813efd22ebc9c2f4daebc40ff9cb686dd5da3c1ed68b47"
	//00000000c80051000000000000000000000000990000000000000000010200000000010a0000000000000000000000000000000000000000000000000000000008000000382e312e332e363200400000000bfb5817584d0f38bb2a0e1f70f2d15d7e512fa885f0ddca1a7f3762b1d26f5cf669e2c9e59699c250813efd22ebc9c2f4daebc40ff9cb686dd5da3c1ed68b47
	data, err := hex.DecodeString(handshake)
	if err != nil {
		panic(err)
	}
	conn, err := net.Dial("tcp", "192.168.13.165:35236")
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
	if hex.EncodeToString(buffer.Bytes()[:6]) == "00000000e400" {
		println("是达梦")
	} else {
		println("不是达梦")
	}
}
