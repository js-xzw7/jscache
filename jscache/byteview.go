package jscache

type ByteView struct {
	b []byte
}

func copyBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func (bv ByteView) Len() int {
	return len(bv.b)
}

func (bv ByteView) ByteSlice() []byte {
	return copyBytes(bv.b)
}

func (bv ByteView) String() string {
	return string(bv.b)
}
