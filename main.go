package main

func main() {

}

const (
	SexMale = iota
	SexFemale
)

const (
	Age0Month0to5 = iota
	Age0Month6to8
	Age0Month9to11
	Age1to2
	Age3to5
	Age6to7
	Age8to9
	Age10to11
	Age12to14
	Age15to17
	Age18to29
	Age30to49
	Age50to64
	Age65to74
	AgeOver75
)

const (
	WorkI = iota
	WorkII
	WorkIII
)

// GetEnergy は PDF の 91 ページ「推定エネルギー必要量（kcal/日）」を参考に
// 性別・年齢・運動量ごとの 1 日の必要エネルギー（kcal）を返します
func GetEnergy(sex, age, work int) float64 {
	list := []float64{
		550, 550, 550, 500, 500, 500,
		650, 650, 650, 600, 600, 600,
		700, 700, 700, 650, 650, 650,
		950, 950, 950, 900, 900, 900,
		1300, 1300, 1300, 1250, 1250, 1250,
		1350, 1550, 1750, 1250, 1450, 1650,
		1600, 1850, 2100, 1500, 1700, 1900,
		1950, 2250, 2500, 1850, 2100, 2350,
		2300, 2600, 2900, 2150, 2400, 2700,
		2500, 2800, 3150, 2050, 2300, 2550,
		2300, 2650, 3050, 1700, 2000, 2300,
		2300, 2700, 3050, 1750, 2050, 2350,
		2200, 2600, 2950, 1650, 1950, 2250,
		2050, 2400, 2750, 1550, 1850, 2100,
		1800, 2100, 2100, 1400, 1650, 1650,
	}
	index := sex*3 + age*6 + work
	return list[index]
}

// GetProtein は PDF の 133 ページ「たんぱく質の食事摂取基準」を参考に
// 性別・年齢・運動量ごとの 1 日のタンパク質の推奨量（g）を返します
func GetProtein(sex, age, work int) float64 {
	list := []float64{
		10, 10,
		15, 15,
		25, 25,
		20, 15,
		25, 20,
		30, 25,
		40, 30,
		45, 40,
		60, 45,
		65, 45,
		65, 40,
		65, 40,
		65, 40,
		60, 40,
		60, 40,
	}
	index := sex*1 + age*2
	return list[index]
}

// GetLipid は PDF の 156 ページ「脂質の食事摂取基準」を参考に
// 性別・年齢・運動量ごとの 1 日の脂質の摂取基準（g）の上限と下限を返します
func GetLipid(sex, age, work int) (float64, float64) {
	min, max := 20.0, 30.0
	switch age {
	case Age0Month0to5:
		min, max = 50, 50
	case Age0Month6to8, Age0Month9to11:
		min, max = 40, 40
	}
	energy := GetEnergy(sex, age, work)
	return energyToGramForLipid(energy, min), energyToGramForLipid(energy, max)
}

// energyToGramForLipid はエネルギー（kcal）から脂質（g）への変換
// 脂質 1g は 9kcal
func energyToGramForLipid(energy, rate float64) float64 {
	return energy * rate / 100 / 9
}

// GetSaturatedFattyAcids は PDF の 157 ページ「飽和脂肪酸の食事摂取基準」を参考に
// 性別・年齢・運動量ごとの 1 日の飽和脂肪酸の摂取基準（g）の上限を返します
func GetSaturatedFattyAcids(sex, age, work int) float64 {
	max := 0.0
	switch age {
	case Age0Month0to5, Age0Month6to8, Age0Month9to11, Age1to2:
		max = 0
	case Age3to5, Age6to7, Age8to9, Age10to11, Age12to14:
		max = 10
	case Age15to17:
		max = 8
	case Age18to29, Age30to49, Age50to64, Age65to74, AgeOver75:
		max = 7
	}
	energy := GetEnergy(sex, age, work)
	return energyToGramForLipid(energy, max)
}

// GetN3FattyAcids は PDF の 158 ページ「n─3 系脂肪酸の食事摂取基準」を参考に
// 性別・年齢・運動量ごとの 1 日のn─3 系脂肪酸の食事摂取基準の推奨量（g）を返します
func GetN3FattyAcids(sex, age, work int) float64 {
	list := []float64{
		0.9, 0.9,
		0.8, 0.8,
		0.7, 0.8,
		1.1, 1,
		1.5, 1.3,
		1.5, 1.3,
		1.6, 1.6,
		1.9, 1.6,
		2.1, 1.6,
		2, 1.6,
		2, 1.6,
		2.2, 1.9,
		2.2, 2,
		2.1, 1.8,
	}
	index := sex*1 + age*2
	return list[index]
}

// GetN6FattyAcids は PDF の 158 ページ「n─6 系脂肪酸の食事摂取基準」を参考に
// 性別・年齢・運動量ごとの 1 日のn─6 系脂肪酸の食事摂取基準の推奨量（g）を返します
func GetN6FattyAcids(sex, age, work int) float64 {
	list := []float64{
		4, 4,
		4, 4,
		4, 4,
		6, 6,
		8, 7,
		8, 7,
		10, 8,
		11, 9,
		13, 9,
		11, 8,
		10, 8,
		10, 8,
		9, 8,
		8, 7,
	}
	index := sex*1 + age*2
	return list[index]
}

// GetCarbohydrates は PDF の 171 ページ「炭水化物の食事摂取基準」を参考に
// 性別・年齢・運動量ごとの 1 日の炭水化物の摂取基準（g）の上限と下限を返します
func GetCarbohydrates(sex, age, work int) (float64, float64) {
	min, max := 50.0, 65.0
	switch age {
	case Age0Month0to5, Age0Month6to8, Age0Month9to11:
		min, max = 0, 0
	}
	energy := GetEnergy(sex, age, work)
	return energyToGramForCarbohydrates(energy, min), energyToGramForCarbohydrates(energy, max)
}

// energyToGramForCarbohydrates はエネルギー（kcal）から炭水化物（g）への変換
// 炭水化物 1g は 4kcal
func energyToGramForCarbohydrates(energy, rate float64) float64 {
	return energy * rate / 100 / 4
}

// GetDietaryFiber は PDF の 172 ページ「食物繊維の食事摂取基準」を参考に
// 性別・年齢・運動量ごとの 1 日の食物繊維の食事摂取基準の目標量（g）を返します
func GetDietaryFiber(sex, age, work int) float64 {
	list := []float64{
		0, 0,
		0, 0,
		0, 0,
		8, 8,
		10, 10,
		11, 11,
		13, 13,
		17, 17,
		19, 18,
		21, 18,
		21, 18,
		21, 18,
		20, 17,
		20, 17,
	}
	index := sex*1 + age*2
	return list[index]
}

// GetVitaminA は PDF の 212 ページ「ビタミン A の食事摂取基準」を参考に
// 性別・年齢・運動量ごとの 1 日のビタミンAの食事摂取基準の推奨量（μg）を返します
func GetVitaminA(sex, age, work int) float64 {
	list := []float64{
		300, 300,
		400, 400,
		400, 350,
		450, 500,
		400, 400,
		500, 500,
		600, 600,
		800, 700,
		900, 650,
		850, 650,
		900, 700,
		900, 700,
		850, 700,
		800, 650,
	}
	index := sex*1 + age*2
	return list[index]
}
