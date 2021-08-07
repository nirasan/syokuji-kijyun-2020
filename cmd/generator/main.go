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
	VitaminB1(dir + "/vitamin_b1_gen.go")
	VitaminB2(dir + "/vitamin_b2_gen.go")
	Niacin(dir + "/niacin_gen.go")
	VitaminB6(dir + "/vitamin_b6_gen.go")
	VitaminB12(dir + "/vitamin_b12_gen.go")
	FolicAcid(dir + "/folic_acid_gen.go")
	PantothenicAcid(dir + "/pantothenic_acid_gen.go")
	Biotin(dir + "/biotin_gen.go")
	VitaminC(dir + "/vitamin_c_gen.go")
}
