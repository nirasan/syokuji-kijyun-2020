package data

import (
	"strconv"
	"strings"
)

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
)

type ActivityLevel string

const (
	ActivityLevel1 ActivityLevel = "Level1"
	ActivityLevel2 ActivityLevel = "Level2"
	ActivityLevel3 ActivityLevel = "Level3"
)

type Option string

const (
	OptionNone           Option = ""
	OptionEarlyPregnancy Option = "EarlyPregnancy" // 妊娠初期
	OptionMidPregnancy   Option = "MidPregnancy"   // 妊娠中期
	OptionLatePregnancy  Option = "LatePregnancy"  // 妊娠後期
	OptionBreastfeeding  Option = "Breastfeeding"  // 授乳中
	OptionMenstruation   Option = "Menstruation"   // 月経時
)

type Age struct {
	Year  int
	Month int
}

type NilFloat struct {
	Float float64
	Valid bool
}

func NilFloatFromString(s string) NilFloat {
	s = strings.ReplaceAll(s, ",", "")
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return NilFloat{Valid: false}
	}
	return NilFloat{Float: v, Valid: true}
}

type Energy struct {
	Gender        Gender
	From          Age
	To            Age
	ActivityLevel ActivityLevel
	Option        Option
	Value         NilFloat // kcal
}

type Protein struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat      // 推定平均必要量（g）
	RDA    NilFloat      // 推奨量（g）
	AI     NilFloat      // 目安量（g）
	DG     NilFloatRange // 目標量（%エネルギー）
}

type NilFloatRange struct {
	Min NilFloat // 下限
	Max NilFloat // 上限
}

func ProteinKcalToGram(kcal float64) float64 {
	return kcal / 4
}

type Lipid struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat      // 目安量（%エネルギー）
	DG     NilFloatRange // 目標量（%エネルギー）
}

func LipidKcalToGram(kcal float64) float64 {
	return kcal / 9
}

type SaturatedFattyAcids struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	DG     NilFloat // 目標量の上限（%エネルギー）
}

type Omega3FattyAcids struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（g）
}

type Omega6FattyAcids struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（g）
}

type Carbohydrates struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	DG     NilFloatRange // 目標量（%エネルギー）
}

func CarbohydratesKcalToGram(kcal float64) float64 {
	return kcal / 4
}

type DietaryFiber struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	DG     NilFloat // 目標量の下限（g）
}

type VitaminA struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（μg）
	RDA    NilFloat // 推奨量（μg）
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}

type VitaminD struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}

type VitaminE struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（mg）
	UL     NilFloat // 耐容上限量（mg）
}

type VitaminK struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（μg）
}

type VitaminB1 struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
}

type VitaminB2 struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
}

type Niacin struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
	UL     NilFloat // 耐容上限量（mg）
}

type VitaminB6 struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
	UL     NilFloat // 耐容上限量（mg）
}

type VitaminB12 struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（μg）
	RDA    NilFloat // 推奨量（μg）
	AI     NilFloat // 目安量（μg）
}

type FolicAcid struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（μg）
	RDA    NilFloat // 推奨量（μg）
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}

type PantothenicAcid struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（mg）
}

type Biotin struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（μg）
}

type VitaminC struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
}

type Sodium struct {
	Gender  Gender
	From    Age
	To      Age
	Option  Option
	EAR     NilFloat // 推定平均必要量（mg）
	EARSalt NilFloat // 推定平均必要量の食塩相当量（g）
	AI      NilFloat // 目安量（mg）
	AISalt  NilFloat // 目安量の食塩相当量（g）
	DGSalt  NilFloat // 目標量の食塩相当量（g）
}

type Potassium struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（mg）
	DG     NilFloat // 目標量の下限（mg）
}

type Calcium struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（μg）
	RDA    NilFloat // 推奨量（μg）
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}

type Magnesium struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
}

type Phosphorus struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（mg）
	UL     NilFloat // 耐容上限量（mg）
}

type Iron struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
	UL     NilFloat // 耐容上限量（mg）
}

type Zinc struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（μg）
	RDA    NilFloat // 推奨量（μg）
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}

type Copper struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（μg）
	RDA    NilFloat // 推奨量（μg）
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}
