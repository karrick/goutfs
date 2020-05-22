package goutfs

import "testing"

func ensureByteSlicesMatch(tb testing.TB, got, want []byte) {
	tb.Helper()

	la, lb := len(got), len(want)

	max := la
	if max < lb {
		max = lb
	}

	for i := 0; i < max; i++ {
		if i < la && i < lb {
			if g, w := got[i], want[i]; g != w {
				tb.Errorf("%d: GOT: %q; WANT: %q", i, got, want)
			}
		} else if i < la {
			tb.Errorf("%d: GOT: extra byte: %q", i, got[i])
		} else /* i < lb */ {
			tb.Errorf("%d: WANT extra byte: %q", i, want[i])
		}
	}
}

func ensureSlicesOfByteSlicesMatch(tb testing.TB, got, want [][]byte) {
	tb.Helper()

	la, lb := len(got), len(want)

	max := la
	if max < lb {
		max = lb
	}

	for i := 0; i < max; i++ {
		if i < la && i < lb {
			ensureByteSlicesMatch(tb, got[i], want[i])
		} else if i < la {
			tb.Errorf("%d: GOT: extra slice: %v", i, got[i])
		} else /* i < lb */ {
			tb.Errorf("%d: WANT: extra slice: %v", i, want[i])
		}
	}
	if tb.Failed() {
		tb.Logf("GOT: %v; WANT: %v", got, want)
	}
}
