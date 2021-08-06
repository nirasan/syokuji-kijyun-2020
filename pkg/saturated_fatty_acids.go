package pkg

// SaturatedFattyAcids は「日本人の食事摂取基準」（2020年版）の PDF 157 ページにある表「飽和脂肪酸の食事摂取基準」の情報を持つ
type SaturatedFattyAcids struct {
	gender, age int
	datum       *SaturatedFattyAcidsDatum
}

// NewSaturatedFattyAcids は SaturatedFattyAcids を返す
func NewSaturatedFattyAcids(gender, age int) *SaturatedFattyAcids {
	s := &SaturatedFattyAcids{
		gender: gender,
		age:    age,
	}
	s.datum = s.GetDatum()
	return s
}

// NewSaturatedFattyAcidsForPregnantWoman は妊婦向けに値を修正した SaturatedFattyAcids を返す
func NewSaturatedFattyAcidsForPregnantWoman(gender, age int) *SaturatedFattyAcids {
	s := NewSaturatedFattyAcids(gender, age)
	if s.gender == GenderFemale {
		s.datum.DGMin = NewNullFloat64(0)
		s.datum.DGMax = NewNullFloat64(7)
	}
	return s
}

// NewSaturatedFattyAcidsForLactatingWoman は授乳婦向けに値を修正した SaturatedFattyAcids を返す
func NewSaturatedFattyAcidsForLactatingWoman(gender, age int) *SaturatedFattyAcids {
	s := NewSaturatedFattyAcids(gender, age)
	if s.gender == GenderFemale {
		s.datum.DGMin = NewNullFloat64(0)
		s.datum.DGMax = NewNullFloat64(7)
	}
	return s
}

// GetDatum は ProteinDatum を返す
func (s *SaturatedFattyAcids) GetDatum() *SaturatedFattyAcidsDatum {
	for _, d := range s.Data() {
		if d.Gender == s.gender && d.Age == s.age {
			return &d
		}
	}
	return nil
}

// GetDG は目標量（%エネルギー）の上限と下限を返す
// DG とは tentative dietary goal for preventing life-style related diseases の略で、
// 生活習慣病の発症予防のために現在の日本人が当面の目標とすべき摂取量
func (s *SaturatedFattyAcids) GetDG() (float64, float64, bool) {
	d := s.datum
	if d == nil {
		return 0, 0, false
	}
	min, max := d.DGMin, d.DGMax
	if min.Valid && max.Valid {
		return min.Value, max.Value, true
	}
	return 0, 0, false
}

// PercentEnergyToGram は目安量と目標量の単位である（%エネルギー）を（g）に変換します
func (s *SaturatedFattyAcids) PercentEnergyToGram(kcal float64, ratio float64) float64 {
	return kcal * ratio / 100.0 / 9
}

// SaturatedFattyAcidsDatum は「飽和脂肪酸の食事摂取基準」テーブルのデータを持つ
type SaturatedFattyAcidsDatum struct {
	Gender int
	Age    int
	DGMax  NullFloat64 // 目標量 上限
	DGMin  NullFloat64 // 目標量 下限
}
