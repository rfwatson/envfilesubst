package scanner_test

import (
	"bytes"
	"strings"
	"testing"

	"git.netflux.io/rob/envfilesubst/scanner"
	"github.com/stretchr/testify/assert"
)

func TestScanner(t *testing.T) {
	envfile := `FOO=bar
BAR=baz
BAZ=123
FOO_BAR=true
QUUX1=2
`

	testCases := []struct {
		name       string
		envfile    string
		input      string
		wantOutput string
		wantError  string
	}{
		{
			name:       "single variable",
			envfile:    envfile,
			input:      "$FOO",
			wantOutput: "bar\n",
		},
		{
			name:       "input with prefix text",
			envfile:    envfile,
			input:      "baz $FOO",
			wantOutput: "baz bar\n",
		},
		{
			name:       "input with suffix text",
			envfile:    envfile,
			input:      "$FOO baz",
			wantOutput: "bar baz\n",
		},
		{
			name:       "input with prefix and suffix text",
			envfile:    envfile,
			input:      "baz $FOO qux",
			wantOutput: "baz bar qux\n",
		},
		{
			name:       "input with prefix and suffix text and whitespace",
			envfile:    envfile,
			input:      "\tbaz $FOO qux  ",
			wantOutput: "\tbaz bar qux  \n",
		},
		{
			name:       "single variable with curly brackets",
			envfile:    envfile,
			input:      "${FOO}",
			wantOutput: "bar\n",
		},
		{
			name:       "multiple variables on a single line",
			envfile:    envfile,
			input:      "qux ${FOO} quxx $BAR $BAZ",
			wantOutput: "qux bar quxx baz 123\n",
		},
		{
			name:       "non-existent variable",
			envfile:    envfile,
			input:      "$NOPE",
			wantOutput: "$NOPE\n",
		},
		{
			name:       "non-existent variable with curly brackets",
			envfile:    envfile,
			input:      "${NOPE}",
			wantOutput: "${NOPE}\n",
		},
		{
			name:       "multiple variables including non-existent",
			envfile:    envfile,
			input:      "$FOO $BAR $NOPE $BAZ",
			wantOutput: "bar baz $NOPE 123\n",
		},
		{
			name:       "variable name with an underscore",
			envfile:    envfile,
			input:      "$FOO_BAR is true",
			wantOutput: "true is true\n",
		},
		{
			name:       "variable name with a number",
			envfile:    envfile,
			input:      "$QUUX1 + ${QUUX1} = 4",
			wantOutput: "2 + 2 = 4\n",
		},
		{
			name:    "multiline input ending with newline",
			envfile: envfile,
			input: `---
metadata:
	name: "$FOO"
	labels:
		bar: "$BAR"
		baz: "$BAZ"
`,
			wantOutput: `---
metadata:
	name: "bar"
	labels:
		bar: "baz"
		baz: "123"
`,
		},
		{
			name:    "multiline input not ending with newline",
			envfile: envfile,
			input: `---
metadata:
	name: "$FOO"
	labels:
		bar: "$BAR"
		baz: "$BAZ"`,
			wantOutput: `---
metadata:
	name: "bar"
	labels:
		bar: "baz"
		baz: "123"
`,
		},
		{
			name:       "empty string",
			envfile:    envfile,
			input:      "",
			wantOutput: "",
		},
		{
			name:       "multiline with only newlines",
			envfile:    envfile,
			input:      "\n\n\n\n\n",
			wantOutput: "\n\n\n\n\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			envfile := strings.NewReader(strings.TrimSpace(tc.envfile))
			input := strings.NewReader(tc.input)
			var output bytes.Buffer

			scanner := scanner.New(&output, input, envfile)

			err := scanner.Scan()

			if tc.wantError == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantOutput, output.String())
			} else {
				assert.EqualError(t, err, tc.wantError)
			}
		})
	}
}
