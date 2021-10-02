package tempflag

import (
	"flag"
	"fmt"

	"github.com/JalanJiang/the-go-programming-language-works/ch7/7.4/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
