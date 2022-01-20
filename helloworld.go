package main

import "fmt"

//ex1:
// func WordCount(s string) map[string]int {
// 	a := make([]string, 10)
// 	m := make(map[string]int)
// 	a = strings.Fields(s)
// 	for i := range a {
// 		x, ok := m[a[i]]
// 		if ok == true {
// 			m[a[i]] = x + 1
// 		} else {
// 			m[a[i]] = 1
// 		}
// 	}
// 	return m
// }

// func main() {
// 	fmt.Println(WordCount("I am learning Go!"))
// }

//ex2
// func fibonacci() func(i int) int {
// 	return func(i int) int {
// 		if i == 0 || i == 1 {
// 			return 1
// 		}
// 		// return int(func(i-2)+func(i-1))

// 		return fibonacci()(i-2) + fibonacci()(i-1)

// 	}
// }

// func main() {
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f(i))
// 	}
// }

//ex3
// type IPAddr [4]byte

// // TODO: Add a "String() string" method to IPAddr.
// func (i IPAddr) String() string {
// 	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
// }

// func main() {
// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	for name, ip := range hosts {
// 		fmt.Printf("%v: %v\n", name, ip)
// 	}
// }

//ex4
// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	y := 0.0
// 	for math.Abs(z-y) > 0.00000001 {
// 		y = z
// 		z -= (z*z - x) / (2 * z)
// 	}
// 	return z
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// }

//ex5
type MyReader struct {
	s string
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	copied := copy(b, []byte(mr.s))
	return copied, nil
}

func main() {
	b := make([]byte, 8)
	fmt.Println(MyReader{"12345678abcd1234"}.Read(b))
	fmt.Printf("%v", b)
}
