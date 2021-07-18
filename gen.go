package main

import (
	"crypto"
	"encoding/base64"
	"strings"

	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha512"
)

const (
	roundCount = 10
)

type Generator struct {
	Hash     string
	Password string
	Domain   string
	Length   int
}

func (g *Generator) getHash() crypto.Hash {
	name := strings.ToUpper(g.Hash)
	switch name {
	case "MD5":
		return crypto.MD5
	case "SHA1":
		return crypto.SHA1
	case "SHA512":
		return crypto.SHA512
	}
	return crypto.MD5
}

func (g *Generator) Generate() string {
	text := g.Password + ":" + g.Domain
	ch := g.getHash()
	generated := roundPassword(ch, text, g.Length, roundCount)
	return generated
}

func (g *Generator) MaxLength() int {
	ch := g.getHash()
	size := ch.Size()
	length := base64.StdEncoding.EncodedLen(size)
	return length
}

func roundPassword(ch crypto.Hash, input string, length int, round int) string {
	generated := input
	for {
		if round <= 0 {
			passwd := generated[:length]
			if validPassword(passwd) {
				return passwd
			}
		}

		round--
		generated = hashPassword(ch, generated)
	}
}

func hashPassword(ch crypto.Hash, input string) string {
	hh := ch.New()
	if _, err := hh.Write([]byte(input)); err != nil {
		panic(err)
	}
	digest := hh.Sum(nil)

	str := base64.StdEncoding.EncodeToString(digest)
	str = strings.ReplaceAll(str, "+", "9")
	str = strings.ReplaceAll(str, "/", "8")
	str = strings.ReplaceAll(str, "=", "A")
	return str
}

func validPassword(input string) bool {
	first := input[0]
	if first < 'a' || 'z' < first {
		return false
	}

	var upper bool
	var number bool
	for _, r := range input {
		if '0' <= r && r <= '9' {
			number = true
		} else if 'A' <= r && r <= 'Z' {
			upper = true
		}
	}
	return upper && number
}
