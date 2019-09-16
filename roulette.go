package main

import(
	"math/rand"
)

func Roulette() (msg string) {
	results := make([]int, 9)
	emotes := [3]string{"<:astolfothink:551961210443268096> ", "<:NOTY:551975634054938634> ", "<:delete_this:551982979665362945> "}

	for index := range results {
		results[index] = rand.Intn(3)
	}

	return "results: \n" + emotes[results[0]] + emotes[results[1]] + emotes[results[2]] + "\n" + emotes[results[3]] + emotes[results[4]] + emotes[results[5]] + "\n" + emotes[results[6]] + emotes[results[7]] + emotes[results[8]]
}