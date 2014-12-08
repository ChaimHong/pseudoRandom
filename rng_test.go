package pseudoRandom

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type player struct {
	seed int64
	num  int
}

var playerA = &player{
	seed: 1,
	num:  0,
}

func TestRnd(t *testing.T) {
	for i := 0; i < 10; i++ {
		globalRng := NewRng(playerA.seed, playerA.num)
		var rs rand.Source = globalRng
		r := rand.New(rs)
		fmt.Println("res", r.Intn(10000))
		playerA.num++
	}
}

func TestNormal(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				time.Sleep(3e7)
				rand.Int()
			}
		}()
	}

	for i := 0; i < 5; i++ {
		time.Sleep(3e8)
		fmt.Println(rand.Intn(10000))
	}

	time.Sleep(10e9)
}
