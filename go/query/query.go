package query

import (
	"math/rand"
	"time"
)

type Player struct {
	Cards []int32
}

func Shuffle(p1, p2 *Player) {
	rand.Seed(time.Now().Unix())
	numbers := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	p1.Cards = numbers[:5]
	p2.Cards = numbers[5:]
}

func (p *Player) Extract(idx int32) {
	p.Cards = p.Cards[:len(p.Cards)-1]
}

func (p *Player) IsEmpty() bool {
	return (len(p.Cards) == 0)
}

func (p *Player) FirstEl() int32 {
	return p.Cards[0]
}

func (p *Player) Clear() {
	p.Cards = []int32{}
}
