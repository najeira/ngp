package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var length int
	var hash string
	flag.IntVar(&length, "n", 20, "password length")
	flag.StringVar(&hash, "h", "sha512", "hash algorithm")
	flag.Parse()

	password := flag.Arg(0)
	domain := flag.Arg(1)

	if cfg := loadConfig(); cfg != nil {
		if cfg.Length > 0 && !isFlagPassed("n") {
			length = cfg.Length
		}
		if cfg.Hash != "" && !isFlagPassed("h") {
			hash = cfg.Hash
		}
	}

	g := &Generator{
		Hash:     hash,
		Password: password,
		Domain:   domain,
		Length:   length,
	}

	if max := g.MaxLength(); length > max {
		fmt.Printf("%s length should less %d\n", hash, max)
		os.Exit(1)
		return
	}

	p := g.Generate()
	fmt.Println(p)
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
