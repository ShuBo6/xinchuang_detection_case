package main

func main() {
	h := "00000000c80051000000000000000000000000990000000000000000010200000000010a0000000000000000000000000000000000000000000000000000000008000000382e312e332e363200400000000bfb5817584d0f38bb2a0e1f70f2d15d7e512fa885f0ddca1a7f3762b1d26f5cf669e2c9e59699c250813efd22ebc9c2f4daebc40ff9cb686dd5da3c1ed68b47"
	// 将h补充为\x00 这样的十六进制数据
	for i := 0; i < len(h); i++ {
		if i%2 == 0 {
			print("\\x")
		}
		print(string(h[i]))
	}
}