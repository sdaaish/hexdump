package hexdump

import "testing"

func TestBasic(t *testing.T) {
	ans := 2
				if ans != 2 {
								t.Errorf("IntMin(2, -2) = %d; want -2", ans)
				}
}
