package over

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var buf = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 64)
		return &b
	},
}

func GetGoroutineID() uint64 {
	bp := buf.Get().(*[]byte)
	defer buf.Put(bp)
	b := *bp
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	idx := bytes.IndexByte(b, ' ')
	if idx < 0 {
		panic(fmt.Sprintf("no space found in %q", b))
	}
	b = b[:idx]
	gID, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse goroutine ID out of %q: %v", b, err))
	}
	return gID
}
