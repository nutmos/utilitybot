package fortune_teller

import "math/rand"

type Fortune string

const (
	FortuneGood Fortune = "Good"
	FortuneSuay Fortune = "Suay"
)

func TellMyFortune() Fortune {
	fortune := rand.Intn(2)
	switch fortune {
	case 0:
		return FortuneGood
	default:
		return FortuneSuay
	}
}
