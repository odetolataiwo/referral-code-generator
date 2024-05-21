package src

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type ReferralCodeGenerator struct {
	Rng *rand.Rand
}

const (
	alphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
)

// NewReferralCodeGenerator creates a new instance of ReferralCodeGenerator with a seeded random number generator
func NewReferralCodeGenerator() *ReferralCodeGenerator {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rng := rand.New(source)
	return &ReferralCodeGenerator{Rng: rng}
}

// AlphaNumeric generates an alpha-numeric referral code
func (r *ReferralCodeGenerator) AlphaNumeric(wordCase string, wordLength int, numLength int) (string, error) {
	if wordCase == "" {
		return "", errors.New("word case type is required")
	}

	if wordLength < 1 {
		return "", errors.New("length of alphabets to be generated is required")
	}

	if numLength < 1 {
		return "", errors.New("number of random digits to be generated is required")
	}

	var picks []rune

	for i := 0; i < wordLength; i++ {
		index := r.Rng.Intn(len(alphabets))
		char := alphabets[index]
		if wordCase == "lowercase" {
			char = strings.ToLower(string(char))[0]
		}
		picks = append(picks, rune(char))

		for k := 0; k < numLength; k++ {
			numIndex := r.Rng.Intn(len(numbers))
			num := numbers[numIndex]
			picks = append(picks, rune(num))
		}
	}

	result := removeCommas(string(picks))
	return result, nil
}

// Alpha generates an alpha referral code
func (r *ReferralCodeGenerator) Alpha(wordCase string, wordLength int) (string, error) {
	if wordCase == "" {
		return "", errors.New("word case type is required")
	}

	if wordLength < 1 {
		return "", errors.New("length of alphabets to be generated is required")
	}

	var picks []rune

	for i := 0; i < wordLength; i++ {
		index := r.Rng.Intn(len(alphabets))
		char := alphabets[index]
		if wordCase == "lowercase" {
			char = strings.ToLower(string(char))[0]
		}
		picks = append(picks, rune(char))
	}

	result := removeCommas(string(picks))
	if wordCase == "uppercase" {
		return strings.ToUpper(result), nil
	} else {
		return strings.ToLower(result), nil
	}
}

// Custom generates a custom referral code
func (r *ReferralCodeGenerator) Custom(wordCase string, wordLength int, numLength int, username string) (string, error) {
	if len(username) < 1 {
		return "", errors.New("username is required")
	}

	if wordLength < 1 {
		return "", errors.New("word length to chunk must be a number")
	}

	if len(username) < wordLength {
		return "", errors.New("username's length should be greater than word length")
	}

	chunkedName := username[:wordLength]

	if numLength < 1 {
		return "", errors.New("random number length is required")
	}

	code := r.Rng.Intn(9*pow(10, numLength-1)) + pow(10, numLength-1)

	if wordCase == "uppercase" {
		return strings.ToUpper(chunkedName) + fmt.Sprintf("%d", code), nil
	} else if wordCase == "lowercase" {
		return strings.ToLower(chunkedName) + fmt.Sprintf("%d", code), nil
	} else {
		return chunkedName + fmt.Sprintf("%d", code), nil
	}
}
