package main

func main() {

}

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	var data []byte

	return func(buf []byte, n int) int {
		b := make([]byte, 4)

		for len(data) < n {
			c := read4(b)
			if c == 0 {
				break
			}

			data = append(data, b[:c]...)
		}

		c := min(len(data), n)
		copy(buf, data[:c])
		data = data[c:]

		return c
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
