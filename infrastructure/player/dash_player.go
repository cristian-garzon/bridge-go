package player

import "fmt"

type DashPlayer struct{}

func (player *DashPlayer) Play() {
	fmt.Println("play with dash player")
}

func (player *DashPlayer) Stop() {
	fmt.Println("play with dash player")
}
