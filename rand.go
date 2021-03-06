package rand4go

import (
	"math/rand"
	"sync"
	"time"
)

var shared *Rand
var once sync.Once

func sharedRand() *Rand {
	once.Do(func() {
		shared = NewRand()
	})
	return shared
}

type Rand struct {
	r *rand.Rand
}

func NewRand() *Rand {
	var r = &Rand{}
	r.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return r
}

func (this *Rand) Int() int {
	return this.r.Int()
}

func (this *Rand) Intn(n int) int {
	return this.r.Intn(n)
}

func (this *Rand) IntRange(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return this.Intn(max-min) + min
}

func (this *Rand) Int32() int32 {
	return this.r.Int31()
}

func (this *Rand) Int32n(n int32) int32 {
	return this.r.Int31n(n)
}

func (this *Rand) Int32Range(min, max int32) int32 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return this.Int32n(max-min) + min
}

func (this *Rand) Int64() int64 {
	return this.r.Int63()
}

func (this *Rand) Int64n(n int64) int64 {
	return this.r.Int63n(n)
}

func (this *Rand) Int64Range(min, max int64) int64 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return this.Int64n(max-min) + min
}

func (this *Rand) Uint32() uint32 {
	return this.r.Uint32()
}

func (this *Rand) Uint64() uint64 {
	return this.r.Uint64()
}

func (this *Rand) Float32() float32 {
	return this.r.Float32()
}

func (this *Rand) Float64() float64 {
	return this.r.Float64()
}
