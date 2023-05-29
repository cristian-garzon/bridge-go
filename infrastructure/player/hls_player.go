package player

import "fmt"

type HlsPlayer struct{}

func (player *HlsPlayer) Play() {
	fmt.Println("play with hls player")
}

func (player *HlsPlayer) Stop() {
	fmt.Println("stop with hls player")
}
