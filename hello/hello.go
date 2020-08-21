package main

import "fmt"

const engHelloPrefix = "Hello "
const frenchHelloPrefix = "Bonjour "
const chineseHelloPrefix = "你好 "
const french = "French"
const china = "China"

//Hello World字符串
func Hello(name string, language string) string {
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case china:
		prefix = chineseHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "China"))
}
