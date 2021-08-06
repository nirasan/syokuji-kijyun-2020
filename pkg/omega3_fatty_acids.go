package pkg

// Omega3FattyAcids は「日本人の食事摂取基準」（2020年版）の PDF 158 ページにある表「n─3 系脂肪酸の食事摂取基準」の情報を持つ
type Omega3FattyAcids struct {
	gender, age int
	datum       *Omega3FattyAcidsDatum
}

// NewOmega3FattyAcids は Omega3FattyAcids を返す
func NewOmega3FattyAcids(gender, age int) *Omega3FattyAcids {
	o := &Omega3FattyAcids{
		gender: gender,
		age:    age,
	}
	o.datum = o.GetDatum()
	return o
}

// NewOmega3FattyAcidsForPregnantWoman は妊婦向けに値を修正した Omega3FattyAcids を返す
func NewOmega3FattyAcidsForPregnantWoman(gender, age int) *Omega3FattyAcids {
	o := NewOmega3FattyAcids(gender, age)
	if o.gender == GenderFemale {
		o.datum.AI = NewNullFloat64(1.6)
	}
	return o
}

// NewOmega3FattyAcidsForLactatingWoman は授乳婦向けに値を修正した Omega3FattyAcids を返す
func NewOmega3FattyAcidsForLactatingWoman(gender, age int) *Omega3FattyAcids {
	o := NewOmega3FattyAcids(gender, age)
	if o.gender == GenderFemale {
		o.datum.AI = NewNullFloat64(1.8)
	}
	return o
}

// GetDatum は ProteinDatum を返す
func (o *Omega3FattyAcids) GetDatum() *Omega3FattyAcidsDatum {
	for _, d := range o.Data() {
		if d.Gender == o.gender && d.Age == o.age {
			return &d
		}
	}
	return nil
}

// GetAI は目安量（g）を返す
// AI とは adequate intake の略で、
// 一定の栄養状態を維持するのに十分な量であり、目安量以上を摂取している場合は不足のリスクはほとんどない
func (o *Omega3FattyAcids) GetAI() (float64, bool) {
	if o.datum == nil {
		return 0, false
	}
	return o.datum.AI.Flatten()
}

// Omega3FattyAcidsDatum は「n─3 系脂肪酸の食事摂取基準」テーブルのデータを持つ
type Omega3FattyAcidsDatum struct {
	Gender int
	Age    int
	AI     NullFloat64 // 目安量
}
