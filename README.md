# goutfs

Go UTF-8 string.

Provides a String structure type that allows per-character addressing,
slicing, and truncation, all while ensuring characters that require
multiple code points are not split mid character.

## Example Use

```Go
func ExampleString() {
    s := goutfs.NewString("cafés")
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
```

## Definitions

For the purposes of this library, I have attempted to adopt the
universal and Go specific terminology for characters, code points,
runes, and bytes. There is a chance that I misread a resource and have
an error in my terminology, but a best effort has been attempted.

### Character

Each character occupies a single column in the output, and roughly
corresponds to what a human sees when they look at the printed text. A
human might see the latin letter e with an accent grave over it, for
example.

Characters are stored and transmitted using some encoding. In unicode
those encodings are called code points. Because of how combining
characters work in unicode, some characters could have multiple code
point representations. For instance, the lower case letter e with an
accent grave could be encoded as a single unicode code point, or
alternatively encoded by two code points: the first one being the
lower case latin letter e, the second as what is known as a combining
code point, in this case the combining code point for accent
grave. Both of these representations result in the same character
being displayed, but have two byte encodings. There are libraries to
normalize these encodings to one of various canonical
standards. However, I am not certain character normalization needs to
be addressed in this library.

### Code Point, a.k.a. Go rune

A code point is called a rune in Go parlance. A Go rune is stored as
an int32 value. Remember a rune is not necessarily a single
character. Some characters have multiple unicode encodings, each of
which could be single or multiple code points.

Another point--no pun intended--is there are look alike characters in
unicode. Not just different code point sequences that represent the
same character, but two different characters that happen to look
alike. For instance, the latin capitol K looks identical to the
unicode code point for the Kelvin symbol. This library need not worry
itself with look alike characters. In order to function correctly,
this library merely needs to know at one byte offset a particular
character ends and the next character begins.

### Strings

Go has no restrictions on the sequence of bytes stored in a
string. The only restriction Go puts on the bytes in a string are that
Go source code is defined as UTF-8, which means most string literal
values are valid UTF-8 encodings. This is not always the case,
however, as Go allows byte level escapes to be included in string
literals, which may or may not represent valid UTF-8 encoded data.

Iterating over a UTF-8 string will result in some runes that require
multiple bytes, and other runes that require a single byte.

### Starting vs Non-Starting (Combining) Rune

Unicode defines many code points that are called starting code
points. They may be displayed independently of any other code point,
and the may be modified indefinitely by appending non-starting code
points. These non-starting code points are more frequently called
combining code points in literature.

## References

1. https://blog.golang.org/strings
1. https://blog.golang.org/normalization
1. https://pkg.go.dev/golang.org/x/text/transform?tab=doc
