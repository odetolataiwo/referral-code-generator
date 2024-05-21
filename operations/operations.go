package operations

import (
	"math/rand"
	"time"
)

type Operations interface {
	AlphaNumeric(wordCase string, wordLength int, numLength int) (string, error)
	Alpha(wordCase string, wordLength int) (string, error)
	Custom(wordCase string, wordLength int, numLength int, username string) (string, error)
}

type ReferralCodeGenerator struct {
	Rng *rand.Rand
}

// NewReferralCodeGenerator creates a new instance of ReferralCodeGenerator with a seeded random number generator
func NewReferralCodeGenerator() *ReferralCodeGenerator {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rng := rand.New(source)
	return &ReferralCodeGenerator{Rng: rng}
}

func Init() Operations {
	return Operations(NewReferralCodeGenerator())
}
