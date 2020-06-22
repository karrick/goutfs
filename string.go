package goutfs

import "golang.org/x/text/unicode/norm"

// String is a UTF-8 encoded string that allows per-character addressing,
// slicing, and truncation.
type String struct {
	buffer  []byte // buffer stores all bytes from the string
	offsets []int  // offsets[n] is the offset for the start of character N in sequence.
}

// NewString returns a new String by evaluating its input as UTF-8 sequence of
// bytes, and storing the offset to each addressable character.
//
//     func ExampleString() {
//         s := NewString("cafés")
//         fmt.Println(s.Len())
//         fmt.Println(string(s.Char(3)))
//         fmt.Println(string(s.Slice(0, 4)))
//         fmt.Println(string(s.Slice(4, -1)))
//         s.Trunc(3)
//         fmt.Println(string(s.Bytes()))
//         // Output:
//         // 5
//         // é
//         // café
//         // s
//         // caf
//     }
func NewString(s string) *String {
	buffer := make([]byte, 0, len(s)) // pre-allocate byte slice with cap as long as s, but with 0 len
	var offsets []int
	var offset int

	var ni norm.Iter
	ni.InitString(norm.NFKD, s)

	for !ni.Done() {
		// Initial testing revealed that norm.Iter reuses the same byte slice
		// each time, so will need to copy the bytes from the returned slice
		// each time through the loop into our own sequence.
		b := ni.Next()
		buffer = append(buffer, b...)
		offsets = append(offsets, offset)
		offset += len(b)
	}
	return &String{buffer: buffer, offsets: offsets}
}

// Bytes returns the entire slice of bytes that encode all characters in the
// String.
func (s *String) Bytes() []byte {
	return s.buffer
}

// Char returns the slice of bytes that encode the Ith character.
func (s *String) Char(i int) []byte {
	if i < 0 || i >= len(s.offsets) {
		return nil
	}
	if i < len(s.offsets)-1 {
		return s.buffer[s.offsets[i]:s.offsets[i+1]]
	}
	return s.buffer[s.offsets[i]:]
}

// Len returns the number of characters in the String.
func (s *String) Len() int {
	return len(s.offsets)
}

// Slice returns the slice of bytes that encode the Ith thru Jth-1 characters of
// the String. As two special cases, when i is -1, this returns nil; or when j
// is -1, this returns from the Ith character to the end of the string.
func (s *String) Slice(i, j int) []byte {
	if i < 0 || i >= len(s.offsets) {
		return nil
	}
	istart := s.offsets[i]

	if j == -1 || j >= len(s.offsets) {
		return s.buffer[istart:]
	}

	return s.buffer[istart:s.offsets[j]]
}

// Trunc truncates the String to max of i characters. As a special case the
// String is truncated to the empty string when i is less than or equal to 0. No
// operation is performed when i is greater than or equal to the number of
// characters in the String.
func (s *String) Trunc(i int) {
	if i > 0 {
		if i < len(s.offsets) {
			s.buffer = s.buffer[:s.offsets[i]]
			s.offsets = s.offsets[:i]
		}
	} else {
		s.buffer = nil
		s.offsets = nil
	}
}
