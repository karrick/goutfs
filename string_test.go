package goutfs

import (
	"bytes"
	"fmt"
	"testing"
)

func ExampleString() {
	s := NewString("cafés")
	fmt.Println(s.Len())
	fmt.Println(string(s.Char(3)))
	fmt.Println(string(s.Slice(0, 4)))
	fmt.Println(string(s.Slice(4, -1)))
	s.Trunc(3)
	fmt.Println(string(s.Bytes()))
	// Output:
	// 5
	// é
	// café
	// s
	// caf
}

// Characters returns a slice of characters, each character being a slice of
// bytes of the respective encoded character.
func (s *String) characters() [][]byte {
	l := len(s.offsets)
	characters := make([][]byte, l)
	for i := 0; i < l; i++ {
		characters[i] = s.Char(i)
	}
	return characters
}

func TestString(t *testing.T) {
	t.Run("bytes", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			s := NewString("")
			if got, want := s.Bytes(), []byte(nil); !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("non-empty", func(t *testing.T) {
			s := NewString("cafés")
			if got, want := s.Bytes(), []byte("cafés"); !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})

	t.Run("char", func(t *testing.T) {
		t.Run("i less than 0", func(t *testing.T) {
			s := NewString("cafés")
			if got, want := s.Char(-1), []byte(nil); !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("i within range", func(t *testing.T) {
			s := NewString("cafés")
			if got, want := s.Char(0), []byte{'c'}; !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := s.Char(1), []byte{'a'}; !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := s.Char(2), []byte{'f'}; !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := s.Char(3), []byte("é"); !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := s.Char(4), []byte{'s'}; !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("i above range", func(t *testing.T) {
			s := NewString("cafés")
			if got, want := s.Char(5), []byte(nil); !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})

	t.Run("characters", func(t *testing.T) {
		t.Run("empty string", func(t *testing.T) {
			got, want := NewString("").characters(), [][]byte(nil)
			ensureSlicesOfByteSlicesMatch(t, got, want)
		})

		t.Run("a", func(t *testing.T) {
			got, want := NewString("a").characters(), [][]byte{[]byte{'a'}}
			ensureSlicesOfByteSlicesMatch(t, got, want)
		})

		t.Run("cafés", func(t *testing.T) {
			got := NewString("cafés").characters()
			want := [][]byte{[]byte{'c'}, []byte{'a'}, []byte{'f'}, []byte{101, 204, 129}, []byte{'s'}}
			ensureSlicesOfByteSlicesMatch(t, got, want)
		})
	})

	t.Run("len", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			s := NewString("")
			if got, want := s.Len(), 0; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("non-empty", func(t *testing.T) {
			s := NewString("cafés")
			if got, want := s.Len(), 5; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})

	t.Run("slice", func(t *testing.T) {
		t.Run("i negative", func(t *testing.T) {
			s := NewString("cafés")
			if got, want := s.Slice(-1, -1), []byte(nil); !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("i too large", func(t *testing.T) {
			s := NewString("cafés")

			if got, want := s.Slice(6, 13), []byte(nil); !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := s.Slice(6, -1), []byte(nil); !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("j too large", func(t *testing.T) {
			s := NewString("cafés")
			if got, want := string(s.Slice(0, 13)), "cafés"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("j is -1", func(t *testing.T) {
			s := NewString("cafés")

			if got, want := string(s.Slice(0, -1)), "cafés"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := string(s.Slice(1, -1)), "afés"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := string(s.Slice(2, -1)), "fés"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := string(s.Slice(3, -1)), "és"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := string(s.Slice(4, -1)), "s"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := string(s.Slice(5, -1)), ""; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("i and j within range", func(t *testing.T) {
			s := NewString("cafés")

			if got, want := string(s.Slice(0, 5)), "cafés"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := string(s.Slice(0, 4)), "café"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := string(s.Slice(0, 3)), "caf"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := string(s.Slice(0, 2)), "ca"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := string(s.Slice(0, 1)), "c"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}

			if got, want := string(s.Slice(0, 0)), ""; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})

	t.Run("trunc", func(t *testing.T) {
		t.Run("index zero", func(t *testing.T) {
			s := NewString("cafés")
			s.Trunc(0)
			if got, want := s.Len(), 0; got != want {
				t.Fatalf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := s.Bytes(), []byte(nil); !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("index one", func(t *testing.T) {
			s := NewString("cafés")
			s.Trunc(1)
			if got, want := s.Len(), 1; got != want {
				t.Fatalf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := s.Bytes(), []byte{'c'}; !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("index two", func(t *testing.T) {
			s := NewString("cafés")
			s.Trunc(2)
			if got, want := s.Len(), 2; got != want {
				t.Fatalf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := s.Bytes(), []byte{'c', 'a'}; !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("index before multi-code-point", func(t *testing.T) {
			s := NewString("cafés")
			s.Trunc(3)
			if got, want := s.Len(), 3; got != want {
				t.Fatalf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := s.Bytes(), []byte{'c', 'a', 'f'}; !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("index after multi-code-point", func(t *testing.T) {
			s := NewString("cafés")
			s.Trunc(4)
			if got, want := s.Len(), 4; got != want {
				t.Fatalf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := s.Bytes(), []byte{99, 97, 102, 101, 204, 129}; !bytes.Equal(got, want) {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("index equals length", func(t *testing.T) {
			s := NewString("cafés")
			s.Trunc(5)
			if got, want := s.Len(), 5; got != want {
				t.Fatalf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := string(s.Bytes()), "cafés"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})

		t.Run("index greater than length", func(t *testing.T) {
			s := NewString("cafés")
			s.Trunc(6)
			if got, want := s.Len(), 5; got != want {
				t.Fatalf("GOT: %v; WANT: %v", got, want)
			}
			if got, want := string(s.Bytes()), "cafés"; got != want {
				t.Errorf("GOT: %v; WANT: %v", got, want)
			}
		})
	})
}
