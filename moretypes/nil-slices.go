/**
 * @author: puzhengchao <zhengchaopu@gmail.com>
 * @copyright (c) 2017 puzhengchao <zhengchaopu@gmail.com>
 */

package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
