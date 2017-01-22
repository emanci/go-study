/**
 * @author: puzhengchao <zhengchaopu@gmail.com>
 * @copyright (c) 2017 puzhengchao <zhengchaopu@gmail.com>
 */

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	fmt.Println("all:", m)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	fmt.Println("all:", m)

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 100

	v2, ok2 := m["Answer"]
	fmt.Println("The value:", v2, "Present?", ok2)
}
