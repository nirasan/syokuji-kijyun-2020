package pkg

// VitaminE は「日本人の食事摂取基準」（2020年版）の PDF 214 ページにある表「ビタミン E の食事摂取基準」の情報を持つ
type VitaminE struct {
	gender, age int
	datum       *VitaminEDatum
}

// NewVitaminE は VitaminE を返す
func NewVitaminE(gender, age int) *VitaminE {
	vit := &VitaminE{
		gender: gender,
		age:    age,
	}
	vit.datum = vit.GetDatum()
	return vit
}

// NewVitaminE は妊婦向けに値を修正した VitaminE を返す
func NewVitaminEForPregnantWoman(gender, age int) *VitaminE {
	vit := NewVitaminE(gender, age)
	if vit.gender == GenderFemale {
		vit.datum.AI.Value = 6.5
	}
	return vit
}

// NewVitaminE は授乳婦向けに値を修正した VitaminE を返す
func NewVitaminEForLactatingWoman(gender, age int) *VitaminE {
	vit := NewVitaminE(gender, age)
	if vit.gender == GenderFemale {
		vit.datum.AI.Value = 7
	}
	return vit
}

// GetDatum は VitaminEDatum を返す
func (vit *VitaminE) GetDatum() *VitaminEDatum {
	for _, d := range vit.Data() {
		if d.Gender == vit.gender && d.Age == vit.age {
			return &d
		}
	}
	return nil
}

// GetAI は目安量（mg）を返す
// AI とは adequate intake の略で、
// 一定の栄養状態を維持するのに十分な量であり、目安量以上を摂取している場合は不足のリスクはほとんどない
func (vit *VitaminE) GetAI() (float64, bool) {
	d := vit.datum
	if d == nil {
		return 0, false
	}
	return d.AI.Flatten()
}

// GetUL は耐容上限量（mg）を返す
// UL とは tolerable upper intake level の略で、
// 健康障害をもたらすリスクがないとみなされる習慣的な摂取量の上限
func (vit *VitaminE) GetUL() (float64, bool) {
	d := vit.datum
	if d == nil {
		return 0, false
	}
	return d.UL.Flatten()
}

// VitaminEDatum は「ビタミン E の食事摂取基準」テーブルのデータをもつ
type VitaminEDatum struct {
	Gender int
	Age    int
	AI     NullFloat64 // 目安量
	UL     NullFloat64 // 耐容上限量
}
