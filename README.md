# envfilesubst

envfilesubst is a variation of gettext's envsubst, with a different modus
operandi.

Firstly, instead of reading the current environment, it instead parses a file
in "envfile" format.

Secondly, it will read from standard input and replace all variable references
that can be found in the envfile. If variables are not explicitly mentioned in
the envfile, the references will be left intact (instead of replacing them with
an empty string).

## Installation

```
go install git.netflux.io/rob/envfilesubst@latest
```

## Usage

Given an envfile:

```
FOO=bar
X=1
```

Then:

```
echo "FOO is $FOO and X is ${X}. I don't know $BAR." | envfilesubst -f myenvfile
```

The output is:

```
FOO is bar and X is 1. I don't know $BAR.
```

## License

Copyright © 2022 Rob Watson.

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the “Software”), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
