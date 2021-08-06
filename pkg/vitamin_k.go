package pkg

// VitaminK は「日本人の食事摂取基準」（2020年版）の PDF 215 ページにある表「ビタミン K の食事摂取基準」の情報を持つ
type VitaminK struct {
	gender, age int
	datum       *VitaminKDatum
}

// NewVitaminK は VitaminK を返す
func NewVitaminK(gender, age int) *VitaminK {
	vit := &VitaminK{
		gender: gender,
		age:    age,
	}
	vit.datum = vit.GetDatum()
	return vit
}

// NewVitaminK は妊婦向けに値を修正した VitaminK を返す
func NewVitaminKForPregnantWoman(gender, age int) *VitaminK {
	vit := NewVitaminK(gender, age)
	if vit.gender == GenderFemale {
		vit.datum.AI.Value = 150
	}
	return vit
}

// NewVitaminK は授乳婦向けに値を修正した VitaminK を返す
func NewVitaminKForLactatingWoman(gender, age int) *VitaminK {
	vit := NewVitaminK(gender, age)
	if vit.gender == GenderFemale {
		vit.datum.AI.Value = 150
	}
	return vit
}

// GetDatum は VitaminKDatum を返す
func (vit *VitaminK) GetDatum() *VitaminKDatum {
	for _, d := range vit.Data() {
		if d.Gender == vit.gender && d.Age == vit.age {
			return &d
		}
	}
	return nil
}

// GetAI は目安量（μg）を返す
// AI とは adequate intake の略で、
// 一定の栄養状態を維持するのに十分な量であり、目安量以上を摂取している場合は不足のリスクはほとんどない
func (vit *VitaminK) GetAI() (float64, bool) {
	d := vit.datum
	if d == nil {
		return 0, false
	}
	return d.AI.Flatten()
}

// VitaminKDatum は「ビタミン E の食事摂取基準」テーブルのデータをもつ
type VitaminKDatum struct {
	Gender int
	Age    int
	AI     NullFloat64 // 目安量
}
