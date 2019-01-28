package proxy

type ReverseProxyBufferPool struct {
	bufferSize int
	ch         chan []byte
}

func NewReverseProxyBufferPool(size int) *ReverseProxyBufferPool {
	return &ReverseProxyBufferPool{
		bufferSize: size,
		ch:         make(chan []byte, 1024),
	}
}

func (r *ReverseProxyBufferPool) Get() []byte {
	select {
	case buf := <-r.ch:
		return buf
	default:
		return make([]byte, r.bufferSize)
	}
}

func (r *ReverseProxyBufferPool) Put(buf []byte) {
	if len(buf) != r.bufferSize {
		return
	}
	select {
	case r.ch <- buf:
	default:
		// no thing
	}
}
