package pkg

// VitaminD は「日本人の食事摂取基準」（2020年版）の PDF 213 ページにある表「ビタミン D の食事摂取基準」の情報を持つ
type VitaminD struct {
	gender, age int
	datum       *VitaminDDatum
}

// NewVitaminD は VitaminD を返す
func NewVitaminD(gender, age int) *VitaminD {
	vit := &VitaminD{
		gender: gender,
		age:    age,
	}
	vit.datum = vit.GetDatum()
	return vit
}

// NewVitaminD は妊婦向けに値を修正した VitaminD を返す
func NewVitaminDForPregnantWoman(gender, age int) *VitaminD {
	vit := NewVitaminD(gender, age)
	if vit.gender == GenderFemale {
		vit.datum.AI.Value = 8.5
	}
	return vit
}

// NewVitaminD は授乳婦向けに値を修正した VitaminD を返す
func NewVitaminDForLactatingWoman(gender, age int) *VitaminD {
	vit := NewVitaminD(gender, age)
	if vit.gender == GenderFemale {
		vit.datum.AI.Value = 8.5
	}
	return vit
}

// GetDatum は VitaminDDatum を返す
func (vit *VitaminD) GetDatum() *VitaminDDatum {
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
func (vit *VitaminD) GetAI() (float64, bool) {
	d := vit.datum
	if d == nil {
		return 0, false
	}
	return d.AI.Flatten()
}

// GetUL は耐容上限量（μg）を返す
// UL とは tolerable upper intake level の略で、
// 健康障害をもたらすリスクがないとみなされる習慣的な摂取量の上限
func (vit *VitaminD) GetUL() (float64, bool) {
	d := vit.datum
	if d == nil {
		return 0, false
	}
	return d.UL.Flatten()
}

// VitaminDDatum は「ビタミン D の食事摂取基準」テーブルのデータをもつ
type VitaminDDatum struct {
	Gender int
	Age    int
	AI     NullFloat64 // 目安量
	UL     NullFloat64 // 耐容上限量
}
