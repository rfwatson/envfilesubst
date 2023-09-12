# envfilesubst

envfilesubst is a variation of gettext's
[envsubst](https://www.gnu.org/software/gettext/manual/html_node/envsubst-Invocation.html),
with a different modus operandi.

Firstly, instead of reading the current environment it reads from file in
traditional "envfile" format.

Secondly, it will read input from stdin replacing all variable references that
can be matched with the envfile. If variables are not explicitly mentioned in
the envfile, the references will be left untouched (instead of replacing them
with an empty string).

## Git

The main git repo is: https://git.netflux.io/rob/envfilesubst

It is also mirrored on GitHub: https://github.com/rfwatson/envfilesubst

## Installation

```
go install git.netflux.io/rob/envfilesubst@latest
```

## Usage

Given an envfile:

```
# myenvfile
FOO=bar
X=1
```

Then:

```
echo "FOO is $FOO and X is ${X}. I don't know $BAZ." | envfilesubst -f myenvfile
```

The output is:

```
FOO is bar and X is 1. I don't know $BAZ.
```
