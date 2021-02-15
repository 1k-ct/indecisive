package random

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

func NewTime() *Chooser {
	a := &Chooser{
		c: &Random{
			rng: rand.New(rand.NewSource(time.Now().UnixNano())),
		},
	}
	return a
}
func NewCrypto() *Chooser {
	b := &Chooser{
		c: &Random{
			rng: rand.New(mt19937.New()),
		},
	}
	return b
}

type Choose interface {
	TimeChooseIntnList(n int) int
	CryptoChooseIntnList(n int) int
}
type Chooser struct {
	c Choose
}
type Random struct {
	rng *rand.Rand
}

// TimeChooseIntnList time時間を使ってrandom作る
func (d *Random) TimeChooseIntnList(n int) int {
	c := d.rng.Intn(n)
	return c
}

// CryptoChooseIntnList 暗号のとき使う
func (d *Random) CryptoChooseIntnList(n int) int {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	d.rng.Seed(seed.Int64())
	r := d.rng.Intn(n)
	return r
}

func (a *Chooser) TimeRandChooseList(l []string) string {
	r := a.c.TimeChooseIntnList(len(l))
	return l[r]
}

func (a *Chooser) CryptoRandChooseList(l []string) string {
	r := a.c.CryptoChooseIntnList(len(l))
	return l[r]
}
func example() {
	l := []string{"0個", "1個", "2個", "3個"}

	a := NewTime()
	fmt.Println("結果: ", a.TimeRandChooseList(l))

	b := NewCrypto()
	fmt.Println("結果: ", b.CryptoRandChooseList(l))
}
