package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/glacjay/goini"
)

type Config struct {
	Hash   string
	Length int
}

func loadConfig() *Config {
	dir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	path := filepath.Join(dir, ".config", "ngp", "config.ini")
	dict, err := ini.Load(path)
	if err != nil {
		return nil
	}

	section := dict[""]
	return &Config{
		Hash:   section["hash"],
		Length: parseInt(section["length"]),
	}
}

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
