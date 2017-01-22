/**
 * @author: puzhengchao <zhengchaopu@gmail.com>
 * @copyright (c) 2017 puzhengchao <zhengchaopu@gmail.com>
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	fmt.Println(strings.HasPrefix("abc", "a"))
}
