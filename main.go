package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

func main() {
	l := []string{"0個", "1個", "2個", "3個"}

	a := New()
	fmt.Println("結果: ", a.timeRandChooseList(l))

	b := &Chooser{
		c: &Random{
			rng: rand.New(mt19937.New()),
		},
	}
	r := b.c.CryptoChooseIntnList(len(l))
	fmt.Println(l[r])
}

func New() *Chooser {
	a := &Chooser{
		c: &Random{
			rng: rand.New(rand.NewSource(time.Now().UnixNano())),
		},
	}
	return a
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

func (a *Chooser) timeRandChooseList(l []string) string {
	r := a.c.TimeChooseIntnList(len(l))
	return l[r]
}

func (a *Chooser) cryptoRandChooseList(l []string) string {
	r := a.c.CryptoChooseIntnList(len(l))
	return l[r]
}
