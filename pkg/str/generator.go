package str

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const DefaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*?"

type Generator struct {
	Length int
	Chars  string
}

func (s *Generator) SetLength(length int) *Generator {
	s.Length = length
	return s
}

func (s *Generator) SetChars(chars string) *Generator {
	s.Chars = chars
	return s
}

func (s *Generator) Generate() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	shuffledCharset := make([]byte, len(s.Chars))
	perm := seededRand.Perm(len(s.Chars))
	for i, v := range perm {
		shuffledCharset[v] = s.Chars[i]
	}

	s.Chars = string(shuffledCharset)

	if len(s.Chars) == 0 {
		return ""
	}

	b := make([]byte, s.Length)

	for i := range b {
		newChar := s.Chars[seededRand.Intn(len(s.Chars))]

		if i > 0 && b[i-1] == newChar {
			for newChar == b[i-1] {
				newChar = s.Chars[seededRand.Intn(len(s.Chars))]
			}
		}

		b[i] = newChar
	}

	return string(b)
}

func New() *Generator {
	return &Generator{
		Length: 24,
		Chars:  DefaultCharset,
	}
}
func Default() *Generator {
	generator := New()

	flag.IntVar(&generator.Length, "length", 24, "Length of the generated the password")
	flag.StringVar(&generator.Chars, "chars", DefaultCharset, "Base characters used to generate the password")

	flag.Usage = func() {
		fmt.Println("String Generator Usage:")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Validate charset length
	if len(generator.Chars) < generator.Length {
		fmt.Printf(
			"Error: Cannot generate string of length %d - charset only contains %d characters\n",
			generator.Length,
			len(generator.Chars),
		)
		fmt.Println("Either:")
		fmt.Println("- Reduce the length with -length=<smaller number>")
		fmt.Println("- Provide a larger charset with -chars=<your characters>")
		os.Exit(1)
	}

	return generator
}
