// github issue: https://github.com/cloudcmds/tamarin/issues/6
// expected value: 11
// expected type: int

s := "\ntest\t\"str\\"

raw := `
test	"str\`

assert(s == raw)

len(s)
