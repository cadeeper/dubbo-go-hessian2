package hessian

type BytePool struct {
	c chan []byte
}

func NewPool(max int) *BytePool {
	return &BytePool{
		c: make(chan []byte, max),
	}
}

func (b *BytePool) Get(size int) []byte {
	var c []byte
	select {
	case c = <-b.c:
	default:
		return make([]byte, size)
	}
	if cap(c) < size {
		return make([]byte, size)
	}
	return c[:size]
}

func (b *BytePool) Put(data []byte) {
	select {
	case b.c <- data:
	default:

	}
}
