package main

func main() {
	h := "520000000c00"
	// 将h补充为\x00 这样的十六进制数据
	for i := 0; i < len(h); i++ {
		if i%2 == 0 {
			print("\\x")
		}
		print(string(h[i]))
	}
}
