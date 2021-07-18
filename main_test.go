package main

import (
	"testing"
)

type testCase struct {
	expected string
	password string
	domain   string
	hash     string
	length   int
}

var (
	testCases = []testCase{
		{
			expected: "w9UbG0NEk7",
			password: "test",
			domain:   "example.com",
			hash:     "md5",
			length:   10,
		},
		{
			expected: "sJfoZg3nU8",
			password: "test",
			domain:   "example.com",
			hash:     "sha512",
			length:   10,
		},
		{
			expected: "aC81",
			password: "test",
			domain:   "example.com",
			hash:     "sha512",
			length:   4,
		},
		{
			expected: "vBKDNdjhhL6dBfgDSRxZxAAA",
			password: "test",
			domain:   "example.com",
			hash:     "md5",
			length:   24,
		},
		{
			expected: "sJfoZg3nU8y32EyHFRlSY08u",
			password: "test",
			domain:   "example.com",
			hash:     "sha512",
			length:   24,
		},
		{
			expected: "aRFG84Gim9",
			password: "test",
			domain:   "example.co.uk",
			hash:     "md5",
			length:   10,
		},
		{
			expected: "hSF8nTst4A",
			password: "test",
			domain:   "example.gov.ac",
			hash:     "md5",
			length:   10,
		},
		{
			expected: "ft8iv4t5sX",
			password: "Γαζέες καὶ μυρτιὲς δὲν θὰ βρῶ πιὰ στὸ χρυσαφὶ ξέφωτο",
			domain:   "example.com",
			hash:     "md5",
			length:   10,
		},
		{
			expected: "o1AWdbILuJ",
			password: "Benjamín pidió una bebida de kiwi y fresa",
			domain:   "example.com",
			hash:     "md5",
			length:   10,
		},
		{
			expected: "iUL7ndPlsD",
			password: "Árvíztűrő tükörfúrógép",
			domain:   "example.com",
			hash:     "md5",
			length:   10,
		},
		{
			expected: "fDOVXY6AhC",
			password: "わかよたれそつねならむ",
			domain:   "example.com",
			hash:     "md5",
			length:   10,
		},
		{
			expected: "i4LtmfRGl8",
			password: "ウヰノオクヤマ ケフコエテ",
			domain:   "example.com",
			hash:     "md5",
			length:   10,
		},
		{
			expected: "wD8T8KozGO",
			password: "מצא לו חברה איך הקליטה",
			domain:   "example.com",
			hash:     "md5",
			length:   10,
		},
		{
			expected: "jtUcAzTL4l",
			password: "В чащах юга жил бы цитрус? Да, но фальшивый экземпляр!",
			domain:   "example.com",
			hash:     "md5",
			length:   10,
		},
		{
			expected: "rnXePhv0JG",
			password: "จงฝ่าฟันพัฒนาวิชาการ",
			domain:   "example.com",
			hash:     "md5",
			length:   10,
		},
	}
)

func TestGenerator(t *testing.T) {
	for _, tc := range testCases {
		g := &Generator{
			Hash:     tc.hash,
			Password: tc.password,
			Domain:   tc.domain,
			Length:   tc.length,
		}
		p := g.Generate()
		if p != tc.expected {
			t.Errorf("%s != %s (%s %s %s %d)",
				p, tc.expected,
				tc.password, tc.domain, tc.hash, tc.length)
		}
	}
}
