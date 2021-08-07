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
	OptionEarlyPregnancy Option = "EarlyPregnancy"
	OptionMidPregnancy   Option = "MidPregnancy"
	OptionLatePregnancy  Option = "LatePregnancy"
	OptionBreastfeeding  Option = "Breastfeeding"
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
