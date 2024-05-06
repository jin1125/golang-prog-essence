package main

import "fmt"

func main() {
	a := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	n := 50
	a = a[:n+copy(a[n:], a[n+1:])] // a = a[:50+49]
	fmt.Println(a)

	s1 := "Hello"
	b := []byte(s1)
	b[0] = 'h'
	s1 = string(b)
	fmt.Println(s1)

	s2 := "こんにちわ世界"
	rs := []rune(s2)
	rs[4] = 'は'
	s2 = string(rs)
	fmt.Println(s2)

	var content = `
	複数行の
	文章からなる
	テキストです。
	`
	fmt.Println(content)
}
