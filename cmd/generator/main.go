package main

import (
	"flag"
)

func main() {
	flag.Parse()
	service, output := flag.Arg(0), flag.Arg(1)
	switch service {
	case "energy":
		Energy(output)
	}
}
