package str

import (
	"strconv"
	"strings"
	"testing"
)

func TestStr(t *testing.T) {
	s := "aa"
	t.Logf("%v", strings.Split(s, ":"))
}

func TestStrConvert(t *testing.T) {
	str := "112#14#12"
	res := []byte{}

	i := len(str) - 1
	for i >= 0 {
		if str[i] == '#' {
			n, _ := strconv.Atoi(str[i-2 : i])
			res = append([]byte{byte(n) - 1 + 'a'}, res...)
			i -= 3
		} else {
			n := str[i] - '0'
			res = append([]byte{n - 1 + 'a'}, res...)
			i -= 1
		}
	}

	length := len(res)
	res2 := make([]byte, len(res))
	for i, ch := range res {
		res2[length-i-1] = ch
	}

	t.Logf("%s", string(res))
	t.Logf("%s", string(res2))
}

func TestTrim(t *testing.T) {
	str := `
  up
`
	t.Logf("Trimmed string: %s", strings.TrimSpace(str))
}
