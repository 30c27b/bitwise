package bitwise

import (
	"errors"
	"io"
)

type Reader struct {
	in   io.ByteReader
	rest byte
	size uint
}

func NewReader(in io.ByteReader) *Reader {
	return &Reader{in: in}
}

func (r *Reader) ReadBits(l uint) (n uint64, err error) {
	if l > 64 {
		return 0, errors.New("exceed the limit of bits read in single call (64)")
	}

	if l <= r.size {
		shift := r.size - l
		n = uint64(r.rest >> shift)
		r.rest &= 1<<shift - 1
		r.size -= l
	} else {
		n = uint64(r.rest)
		l -= r.size
		r.size = 0
		for l > 0 {
			b, err := r.in.ReadByte()
			if err != nil {
				return 0, err
			}
			if l >= 8 {
				n <<= 8
				n += uint64(b)
				l -= 8
			} else {
				shift := 8 - l
				n <<= l
				n += uint64(b >> byte(shift))
				r.rest = b & (1<<byte(shift) - 1)
				r.size = shift
				l = 0
			}
		}
	}

	return
}

func (r *Reader) Align() (n uint8, l uint) {
	l = r.size
	n = r.rest
	r.size = 0
	r.rest = 0
	return
}
