package main

import (
	"fmt"
	"os"

	"github.com/consensys/gurvy/internal/generator"
)

//go:generate go run main.go
func main() {

	bn256 := generator.NewCurveConfig("bn256",
		"21888242871839275222246405745257275088548364400416034343698204186575808495617",
		"21888242871839275222246405745257275088696311157297823662689037894645226208583",
		true,
		true,
	)

	bls377 := generator.NewCurveConfig("bls377",
		"8444461749428370424248824938781546531375899335154063827935233455917409239041",
		"258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177",
		true,
		true,
	)

	bls381 := generator.NewCurveConfig("bls381",
		"52435875175126190479447740508185965837690552500527637822603658699938581184513",
		"4002409555221667393417789825735904156556882819939007885332058136124031650490837864442687629129015664037894272559787",
		true,
		true,
	)

	bw761 := generator.NewCurveConfig("bw761",
		"258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177",
		"6891450384315732539396789682275657542479668912536150109513790160209623422243491736087683183289411687640864567753786613451161759120554247759349511699125301598951605099378508850372543631423596795951899700429969112842764913119068299",
		true,
		true,
	)
	bw761.CRange = []int{4, 8, 16}

	confs := []generator.CurveConfig{bn256, bls377, bls381, bw761}

	for i := 0; i < len(confs); i++ {

		assertNoError(generator.GenerateBaseFields(confs[i]))
		assertNoError(generator.GenerateMultiExpHelpers(confs[i]))
		assertNoError(generator.GenerateDoc(confs[i]))

		if confs[i].CurveName != "bw761" {

			// G1
			if confs[i].CurveName == "bn256" {
				confs[i].CofactorCleaning = false
			}
			assertNoError(generator.GeneratePoint(confs[i], "fp.Element", "g1"))

			// G2
			if confs[i].CurveName == "bn256" {
				confs[i].CofactorCleaning = true
			}
			assertNoError(generator.GeneratePoint(confs[i], "e2", "g2"))
			assertNoError(generator.GenerateFq12over6over2(confs[i]))
			assertNoError(generator.GeneratePairingTests(confs[i]))

		} else {
			// G1
			assertNoError(generator.GeneratePoint(confs[i], "fp.Element", "g1"))
			// G2
			assertNoError(generator.GeneratePoint(confs[i], "fp.Element", "g2"))
		}

	}

}

func assertNoError(err error) {
	if err != nil {
		fmt.Printf("\n%s\n", err.Error())
		os.Exit(-1)
	}
}
