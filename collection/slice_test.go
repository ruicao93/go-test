package collection

import "testing"

func TestSlice(t *testing.T) {
	sa := []int{1, 2, 3, 4, 5}
	sb := sa[0:1]
	sb[0] = 999
	t.Log(sa)
	t.Log(sb)
	t.Log(cap(sa))
	t.Log(cap(sb))

	sa = append(sa, 1)
	sb[0] = 1000
	t.Log(sa)
	t.Log(sb)
	t.Log(cap(sa))
	t.Log(cap(sb))
	t.Log(sb[3])
}
