package main

import "fmt"

func main() {

	m := map[string]func() int{}
	m["lzr"] = func() int {
		return 1
	}
	i := m["lzr"]
	fmt.Println(i())

	fmt.Println(sqlQuote(1))
}
func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x // (not shown)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
