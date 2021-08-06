package pkg

// Omega6FattyAcids は「日本人の食事摂取基準」（2020年版）の PDF 158 ページにある表「n─6 系脂肪酸の食事摂取基準」の情報を持つ
type Omega6FattyAcids struct {
	gender, age int
	datum       *Omega6FattyAcidsDatum
}

// NewOmega6FattyAcids は Omega6FattyAcids を返す
func NewOmega6FattyAcids(gender, age int) *Omega6FattyAcids {
	o := &Omega6FattyAcids{
		gender: gender,
		age:    age,
	}
	o.datum = o.GetDatum()
	return o
}

// NewOmega6FattyAcidsForPregnantWoman は妊婦向けに値を修正した Omega6FattyAcids を返す
func NewOmega6FattyAcidsForPregnantWoman(gender, age int) *Omega6FattyAcids {
	o := NewOmega6FattyAcids(gender, age)
	if o.gender == GenderFemale {
		o.datum.AI = NewNullFloat64(9)
	}
	return o
}

// NewOmega6FattyAcidsForLactatingWoman は授乳婦向けに値を修正した Omega6FattyAcids を返す
func NewOmega6FattyAcidsForLactatingWoman(gender, age int) *Omega6FattyAcids {
	o := NewOmega6FattyAcids(gender, age)
	if o.gender == GenderFemale {
		o.datum.AI = NewNullFloat64(10)
	}
	return o
}

// GetDatum は ProteinDatum を返す
func (o *Omega6FattyAcids) GetDatum() *Omega6FattyAcidsDatum {
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
func (o *Omega6FattyAcids) GetAI() (float64, bool) {
	if o.datum == nil {
		return 0, false
	}
	return o.datum.AI.Flatten()
}

// Omega6FattyAcidsDatum は「n─6 系脂肪酸の食事摂取基準」テーブルのデータを持つ
type Omega6FattyAcidsDatum struct {
	Gender int
	Age    int
	AI     NullFloat64 // 目安量
}
