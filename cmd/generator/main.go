package main

import (
	"flag"
)

func main() {
	flag.Parse()
	dir := flag.Arg(0)

	Energy(dir + "/energy_gen.go")
	Protein(dir + "/protein_gen.go")
	Lipid(dir + "/lipid_gen.go")
	SaturatedFattyAcids(dir + "/saturated_fatty_acids_gen.go")
	Omega3FattyAcids(dir + "/omega3_fatty_acids_gen.go")
	Omega6FattyAcids(dir + "/omega6_fatty_acids_gen.go")
	Carbohydrates(dir + "/carbohydrates_gen.go")
	DietaryFiber(dir + "/dietary_fiber_gen.go")
	VitaminA(dir + "/vitamin_a_gen.go")
	VitaminD(dir + "/vitamin_d_gen.go")
	VitaminE(dir + "/vitamin_e_gen.go")
	VitaminK(dir + "/vitamin_k_gen.go")
}
