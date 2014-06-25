package binding

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

// Umarshals JSON-encoded data and returns the result as
// map[string][]string
// to be feeded into binding.bindForm as "formData".
//
// This is basically an emulation of net/http/Request.ParseForm for JSON.
//
// Does neither use "json" struct tags nor applies reflection on JSON values.
//
// Unlike with "form-urlencoded" data, where keys can be repeated,
// only one value for a key can exist. Therefore, although it could be a
// JSON-encoded list, only the first string contains anything.
func FlatDecode(r io.Reader) (map[string][]string, error) {
	// converts the return type of flatDecode
	p, err := flatDecode(json.NewDecoder(r))
	retval := make(map[string][]string, len(p))
	for key := range p {
		retval[key] = []string{string(p[key])}
	}
	return retval, err
}

// Unmarshals JSON-encoded data and stores the result in
// map[string]interface{}.
//
// Explores nested JSON objects and "flattens" them. Let 'r' be a JSON name
// containing a JSON object 'so' and 'sk' a JSON name in 'so', then a
// "r.sk" ← so[sk] results.
//
// JSON values are not unmarshaled with the exception of nested JSON objects
// and opportunistic un-quoting because binding.bindForm is responsible for
// (reflection-less) processing of the values.
func flatDecode(dec *json.Decoder) (map[string]json.RawMessage, error) {
	// Our struct is their 'object', a key is called "name" in JSON.
	var flatObject map[string]json.RawMessage
	if err := dec.Decode(&flatObject); err != nil {
		return flatObject, err
	}
	namesToDelete := make([]string, 0, len(flatObject))
	for name := range flatObject {
		switch flatObject[name][0] {
		case '{': // explore and flatten sub-objects
			namesToDelete = append(namesToDelete, name)
			subObject, err := flatDecode(json.NewDecoder(bytes.NewReader(flatObject[name])))
			if err != nil {
				return flatObject, err
			}
			for sk := range subObject {
				flatObject[name+"."+sk] = subObject[sk]
			}
		case '"':
			// Please note: e.g. "a"" results in 'invalid character' error.
			// Unquoting can fail even for valid JSON data. Such a value
			// is expected to not pass validation stage.
			if u, ok := unquoteBytes(flatObject[name]); ok {
				flatObject[name] = u
			}
		}
	}
	for _, name := range namesToDelete {
		delete(flatObject, name)
	}
	return flatObject, nil
}

/*
The following lines contain "unexported functions" written by the authors of Go.

Regenerate using::

  mv json.go json.go~
  grep -B 1000 -F "****" json.go~ > json.go
  grep -Pzo '(?s)(//\N*\n)*^func getu4.*?^}' /usr/lib/go/src/pkg/encoding/json/decode.go >> json.go
  grep -Pzo '(?s)(//\N*\n)*^func unquoteBytes.*?^}' /usr/lib/go/src/pkg/encoding/json/decode.go >> json.go
  goimports -w=true json.go

Copyright 2010 The Go Authors. All rights reserved.
Use of the following source code is governed by a BSD-style
license that can be found in Go's LICENSE file.
****/

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
// or it returns -1.
func getu4(s []byte) rune {
	if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
		return -1
	}
	r, err := strconv.ParseUint(string(s[2:6]), 16, 64)
	if err != nil {
		return -1
	}
	return rune(r)
}

func unquoteBytes(s []byte) (t []byte, ok bool) {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return
	}
	s = s[1 : len(s)-1]

	// Check for unusual characters. If there are none,
	// then no unquoting is needed, so return a slice of the
	// original bytes.
	r := 0
	for r < len(s) {
		c := s[r]
		if c == '\\' || c == '"' || c < ' ' {
			break
		}
		if c < utf8.RuneSelf {
			r++
			continue
		}
		rr, size := utf8.DecodeRune(s[r:])
		if rr == utf8.RuneError && size == 1 {
			break
		}
		r += size
	}
	if r == len(s) {
		return s, true
	}

	b := make([]byte, len(s)+2*utf8.UTFMax)
	w := copy(b, s[0:r])
	for r < len(s) {
		// Out of room?  Can only happen if s is full of
		// malformed UTF-8 and we're replacing each
		// byte with RuneError.
		if w >= len(b)-2*utf8.UTFMax {
			nb := make([]byte, (len(b)+utf8.UTFMax)*2)
			copy(nb, b[0:w])
			b = nb
		}
		switch c := s[r]; {
		case c == '\\':
			r++
			if r >= len(s) {
				return
			}
			switch s[r] {
			default:
				return
			case '"', '\\', '/', '\'':
				b[w] = s[r]
				r++
				w++
			case 'b':
				b[w] = '\b'
				r++
				w++
			case 'f':
				b[w] = '\f'
				r++
				w++
			case 'n':
				b[w] = '\n'
				r++
				w++
			case 'r':
				b[w] = '\r'
				r++
				w++
			case 't':
				b[w] = '\t'
				r++
				w++
			case 'u':
				r--
				rr := getu4(s[r:])
				if rr < 0 {
					return
				}
				r += 6
				if utf16.IsSurrogate(rr) {
					rr1 := getu4(s[r:])
					if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
						// A valid pair; consume.
						r += 6
						w += utf8.EncodeRune(b[w:], dec)
						break
					}
					// Invalid surrogate; fall back to replacement rune.
					rr = unicode.ReplacementChar
				}
				w += utf8.EncodeRune(b[w:], rr)
			}

		// Quote, control characters are invalid.
		case c == '"', c < ' ':
			return

		// ASCII
		case c < utf8.RuneSelf:
			b[w] = c
			r++
			w++

		// Coerce to well-formed UTF-8.
		default:
			rr, size := utf8.DecodeRune(s[r:])
			r += size
			w += utf8.EncodeRune(b[w:], rr)
		}
	}
	return b[0:w], true
}
