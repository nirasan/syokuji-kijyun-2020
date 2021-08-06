package pkg

// VitaminA は「日本人の食事摂取基準」（2020年版）の PDF 212 ページにある表「ビタミン A の食事摂取基準」の情報を持つ
type VitaminA struct {
	gender, age int
	datum       *VitaminADatum
}

// NewVitaminA は VitaminA を返す
func NewVitaminA(gender, age int) *VitaminA {
	vit := &VitaminA{
		gender: gender,
		age:    age,
	}
	vit.datum = vit.GetDatum()
	return vit
}

// NewVitaminA は妊婦向けに値を修正した VitaminA を返す
func NewVitaminAForPregnantWoman(gender, age, term int) *VitaminA {
	vit := NewVitaminA(gender, age)
	if vit.gender == GenderFemale {
		switch term {
		case TermEarly:
			vit.datum.EAR.Value += 0
			vit.datum.RDA.Value += 0
		case TermMid:
			vit.datum.EAR.Value += 0
			vit.datum.RDA.Value += 0
		case TermLate:
			vit.datum.EAR.Value += 60
			vit.datum.RDA.Value += 80
		}
	}
	return vit
}

// NewVitaminA は授乳婦向けに値を修正した VitaminA を返す
func NewVitaminAForLactatingWoman(gender, age int) *VitaminA {
	vit := NewVitaminA(gender, age)
	if vit.gender == GenderFemale {
		vit.datum.EAR.Value += 300
		vit.datum.RDA.Value += 450
	}
	return vit
}

// GetDatum は VitaminADatum を返す
func (vit *VitaminA) GetDatum() *VitaminADatum {
	for _, d := range vit.Data() {
		if d.Gender == vit.gender && d.Age == vit.age {
			return &d
		}
	}
	return nil
}

// GetEAR は推定平均必要量（μg）を返す
// EAR とは estimated average requirement の略で、半数の者が必要量を満たす量
func (vit *VitaminA) GetEAR() (float64, bool) {
	d := vit.datum
	if d == nil {
		return 0, false
	}
	return d.EAR.Flatten()
}

// GetRDA は推奨量を（μg）返す
// RDA とは recommended dietary allowance の略で, ほとんどの者が充足している量
func (vit *VitaminA) GetRDA() (float64, bool) {
	d := vit.datum
	if d == nil {
		return 0, false
	}
	return d.RDA.Flatten()
}

// GetAI は目安量（μg）を返す
// AI とは adequate intake の略で、
// 一定の栄養状態を維持するのに十分な量であり、目安量以上を摂取している場合は不足のリスクはほとんどない
func (vit *VitaminA) GetAI() (float64, bool) {
	d := vit.datum
	if d == nil {
		return 0, false
	}
	return d.AI.Flatten()
}

// GetUL は耐容上限量（μg）を返す
// UL とは tolerable upper intake level の略で、
// 健康障害をもたらすリスクがないとみなされる習慣的な摂取量の上限
func (vit *VitaminA) GetUL() (float64, bool) {
	d := vit.datum
	if d == nil {
		return 0, false
	}
	return d.UL.Flatten()
}

// VitaminADatum は「ビタミン A の食事摂取基準」テーブルのデータをもつ
type VitaminADatum struct {
	Gender int
	Age    int
	EAR    NullFloat64 // 推定平均必要量
	RDA    NullFloat64 // 推奨量
	AI     NullFloat64 // 目安量
	UL     NullFloat64 // 耐容上限量
}
