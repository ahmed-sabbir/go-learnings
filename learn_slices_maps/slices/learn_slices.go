package main

import "fmt"

func stringSliceOperations() {
	uniHello := "🖐  🌍"
	bytes := []byte(uniHello)
	fmt.Println(bytes)
	runes := []rune(uniHello)
	fmt.Println(runes)
	runes[1] = 'a'
	fmt.Println(runes)
	fmt.Println(uniHello)

}

func main() {

	stringSliceOperations()

	s3 := []string{"a", "b", "c"}
	for k, v := range s3 {
		fmt.Println(k, v)
	}

	s4 := s3[0:2]
	fmt.Println("s4:", s4)
	s3[0] = "d"
	fmt.Println("s4", s4)

	var s5 []string
	fmt.Println("s5:", s5)
	s5 = s3
	s5[1] = "camel"
	fmt.Println("s3:", s3)
	modSlice(s3)
	fmt.Println("s3 after modification:", s3)

	fmt.Println("------")

	s := make([]string, 0)
	fmt.Println("length of s:", len(s))
	s = append(s, "hello")
	fmt.Println("length of s:", len(s))
	fmt.Println("what is s:", s)
	fmt.Println("contens of s[0]:", s[0])
	s[0] = "Goodbye"
	fmt.Println("contents of s[0]:", s[0])

	fmt.Println("--------")

	s2 := make([]string, 2)
	fmt.Println("contents of s2:", s2)
	fmt.Println("contents of s2[0]:", s2[0])
	s2 = append(s2, "hello")
	fmt.Println("contents of s2:", s2)
	fmt.Println("contents of s2[0]:", s2[0])
	fmt.Println("contents of s2[2]:", s2[2])
	fmt.Println("length of s2:", len(s2))
}

func modSlice(s []string) {
	s[0] = "hello"
}
