package player

import "fmt"

type MssPlayer struct{}

func (player *MssPlayer) Play() {
	fmt.Println("play with mss player")
}

func (player *MssPlayer) Stop() {
	fmt.Println("stop with mss player")
}
