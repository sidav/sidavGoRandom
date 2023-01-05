package prng

type PRNG interface {
	SetSeed(int)
	Rand(int) int
	RandInRange(int, int) int
	RollDice(int, int, int) int
}
