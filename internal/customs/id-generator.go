package customs

import (
	"sync/atomic"
	"time"
)

type Gigantic struct {
	count uint64
}

func InitializeGigantic() *Gigantic {
	return &Gigantic{
		count: 0,
	}
}

func (g *Gigantic) NextID() uint64 {
	count := atomic.AddUint64(&g.count, 1)
	cur := count ^ uint64(time.Now().UnixNano())
	return g.Scramble(cur)
}

func (g Gigantic) Scramble(n uint64) uint64 {
	n ^= n >> 33
	n *= 0xff51afd7ed558ccd
	n ^= n >> 33
	n *= 0xc4ceb9fe1a85ec53
	n ^= n >> 33
	return n
}
