package utils

import "awesomeProject/infrastructure/player"

func GetPlayerByHeader(playerHeader string) {
	var implementation player.Player

	switch playerHeader {
	case "mobile":
		implementation = &player.MssPlayer{}
	case "web":
		implementation = &player.HlsPlayer{}
	case "tv":
		implementation = &player.DashPlayer{}
	default:
		implementation = &player.HlsPlayer{}
	}

	implementation.Play()
}
