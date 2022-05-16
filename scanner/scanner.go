package scanner

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/hashicorp/go-envparse"
)

var envvarRegex = regexp.MustCompile(`\$\{?([A-Z][A-Z0-9_]*)\}?`)

type Scanner struct {
	envfile, r io.Reader
	w          io.Writer
}

// New returns a new Scanner with the provided arguments.
func New(w io.Writer, r io.Reader, envfile io.Reader) *Scanner {
	return &Scanner{
		w:       w,
		r:       r,
		envfile: envfile,
	}
}

const nl = "\n"

func (s *Scanner) Scan() error {
	vars, err := envparse.Parse(s.envfile)
	if err != nil {
		return fmt.Errorf("error parsing envfile: %v", err)
	}

	scanner := bufio.NewScanner(s.r)
	for scanner.Scan() {
		text := scanner.Text()
		matchIndices := envvarRegex.FindAllStringSubmatchIndex(text, -1)

		var sb strings.Builder
		var c int
		for _, idx := range matchIndices {
			m1, m2, n1, n2 := idx[0], idx[1], idx[2], idx[3]
			writeString(&sb, text[c:m1])
			c = m2

			name := text[n1:n2]
			if val, ok := vars[name]; ok {
				writeString(&sb, val)
			} else {
				writeString(&sb, text[m1:m2])
			}
		}

		writeString(&sb, text[c:])
		writeString(&sb, nl)

		if _, err := s.w.Write([]byte(sb.String())); err != nil {
			return fmt.Errorf("error writing to output: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	return nil
}

// writeString writes a string to a StringBuilder, discarding the result and
// (non-existent) error.
func writeString(sb *strings.Builder, s string) {
	_, _ = sb.WriteString(s)
}
