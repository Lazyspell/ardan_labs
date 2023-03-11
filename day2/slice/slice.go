package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int                // s is slice of int
	fmt.Println("len", len(s)) // len is "nil safe"

	if s == nil {
		fmt.Println("nil slice")
	}
	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4] //slicing operation, half-open range
	fmt.Printf("s3 = %#v\n", s3)
	result := concat([]string{"A", "B"}, []string{"C", "D", "E"})
	fmt.Println(result)

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs))
	vs = []float64{2, 1, 3, 4}
	fmt.Println(median(vs))
}

func concat(s1, s2 []string) []string {
	//Restiction: No "for" loop
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func median(values []float64) float64 {
	sort.Float64s(values)
	i := len(values) / 2
	if len(values)%2 == 1 {
		return values[i]
	}
	v := (values[i-1] + values[i]) / 2
	return v
}

func test() {
	fmt.Println("this is a test case for tagbar")
}
