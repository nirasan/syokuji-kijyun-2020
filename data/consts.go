package data

import (
	"fmt"
	"strconv"
	"strings"
)

type Gender string

const (
	GenderMale   Gender = "男性"
	GenderFemale Gender = "女性"
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
	OptionEarlyPregnancy Option = "妊娠初期"
	OptionMidPregnancy   Option = "妊娠中期"
	OptionLatePregnancy  Option = "妊娠後期"
	OptionBreastfeeding  Option = "授乳中"
	OptionMenstruation   Option = "月経時"
)

type Age struct {
	Year  int
	Month int
}

func (a Age) String() string {
	return fmt.Sprintf("%d.%d", a.Year, a.Month)
}

type NilFloat struct {
	Float float64
	Valid bool
}

func (n NilFloat) String() string {
	if n.Valid {
		return fmt.Sprintf("%.2f", n.Float)
	} else {
		return "-"
	}
}

func NilFloatFromString(s string) NilFloat {
	s = strings.ReplaceAll(s, ",", "")
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return NilFloat{Valid: false}
	}
	return NilFloat{Float: v, Valid: true}
}

// Energy エネルギー P.91
type Energy struct {
	Gender        Gender
	From          Age
	To            Age
	ActivityLevel ActivityLevel
	Option        Option
	Value         NilFloat // kcal
}

// Protein タンパク質 P.133
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

func (n NilFloatRange) String() string {
	return fmt.Sprintf("%s:%s", n.Min, n.Max)
}

func ProteinKcalToGram(kcal float64) float64 {
	return kcal / 4
}

// Lipid 脂質 P.156
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

// SaturatedFattyAcids 飽和脂肪酸 P.157
type SaturatedFattyAcids struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	DG     NilFloat // 目標量の上限（%エネルギー）
}

// Omega3FattyAcids オメガ3系脂肪酸 P.158
type Omega3FattyAcids struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（g）
}

// Omega6FattyAcids オメガ6系脂肪酸 P.158
type Omega6FattyAcids struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（g）
}

// Carbohydrates 炭水化物 P.171
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

// DietaryFiber 食物繊維 P.172
type DietaryFiber struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	DG     NilFloat // 目標量の下限（g）
}

// VitaminA ビタミンA P.212
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

// VitaminD ビタミンD P.213
type VitaminD struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}

// VitaminE ビタミンE P.214
type VitaminE struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（mg）
	UL     NilFloat // 耐容上限量（mg）
}

// VitaminK ビタミンK P.215
type VitaminK struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（μg）
}

// VitaminB1 ビタミンB1 P.264
type VitaminB1 struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
}

// VitaminB2 ビタミンB2 P.265
type VitaminB2 struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
}

// Niacin ナイアシン P.266
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

// VitaminB6 ビタミンB6 P.267
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

// VitaminB12 ビタミンB12 P.268
type VitaminB12 struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（μg）
	RDA    NilFloat // 推奨量（μg）
	AI     NilFloat // 目安量（μg）
}

// FolicAcid 葉酸 P.269
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

// PantothenicAcid パントテン酸 P.270
type PantothenicAcid struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（mg）
}

// Biotin ビオチン P.271
type Biotin struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（μg）
}

// VitaminC ビタミンC P.272
type VitaminC struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
}

// Sodium ナトリウム P.313
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

// Potassium カリウム P.314
type Potassium struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（mg）
	DG     NilFloat // 目標量の下限（mg）
}

// Calcium カルシウム P.315
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

// Magnesium マグネシウム P.316
type Magnesium struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
}

// Phosphorus リン P.317
type Phosphorus struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（mg）
	UL     NilFloat // 耐容上限量（mg）
}

// Iron 鉄 P.373
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

// Zinc 亜鉛 P.374
type Zinc struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
	UL     NilFloat // 耐容上限量（mg）
}

// Copper 銅 P.375
type Copper struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（mg）
	RDA    NilFloat // 推奨量（mg）
	AI     NilFloat // 目安量（mg）
	UL     NilFloat // 耐容上限量（mg）
}

// Manganese マンガン P.376
type Manganese struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（mg）
	UL     NilFloat // 耐容上限量（mg）
}

// Iodine ヨウ素 P.377
type Iodine struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（μg）
	RDA    NilFloat // 推奨量（μg）
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}

// Selenium セレン P.378
type Selenium struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（μg）
	RDA    NilFloat // 推奨量（μg）
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}

// Chromium クロム P.379
type Chromium struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}

// Molybdenum モリブデン P.380
type Molybdenum struct {
	Gender Gender
	From   Age
	To     Age
	Option Option
	EAR    NilFloat // 推定平均必要量（μg）
	RDA    NilFloat // 推奨量（μg）
	AI     NilFloat // 目安量（μg）
	UL     NilFloat // 耐容上限量（μg）
}
