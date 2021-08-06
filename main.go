package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"golang.org/x/term"
)

func main() {
	var length int
	var hash string
	flag.IntVar(&length, "n", 20, "password length")
	flag.StringVar(&hash, "h", "sha512", "hash algorithm")
	flag.Parse()

	//password := flag.Arg(0)
	//domain := flag.Arg(1)

	if cfg := loadConfig(); cfg != nil {
		if cfg.Length > 0 && !isFlagPassed("n") {
			length = cfg.Length
		}
		if cfg.Hash != "" && !isFlagPassed("h") {
			hash = cfg.Hash
		}
	}

	password := promptPassword()
	domain := promptDomain()

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
	//fmt.Println(p)
	_ = clipboard.WriteAll(p)
	fmt.Println("The password was copied to the clipboard")
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

func promptPassword() string {
	fmt.Print("Master passphrase: ")
	defer fmt.Print("\n")

	fd := os.Stdin.Fd()
	b, _ := term.ReadPassword(int(fd))
	return string(b)
}

func promptDomain() string {
	fmt.Print("Domain: ")
	defer fmt.Print("\n")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	return strings.TrimRight(s, "\r\n")
}
