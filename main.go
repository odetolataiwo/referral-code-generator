package main

import (
	"fmt"

	cg "github.com/odetolataiwo/referral-code-generator/src"
)

func main() {
	r := cg.NewReferralCodeGenerator()

	alpha, err := r.Alpha("lowercase", 12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Alpha:", alpha)
	}

	alphaNumeric, err := r.AlphaNumeric("uppercase", 8, 7)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("AlphaNumeric:", alphaNumeric)
	}

	custom, err := r.Custom("lowercase", 6, 6, "rockstar")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Custom:", custom)
	}
}
