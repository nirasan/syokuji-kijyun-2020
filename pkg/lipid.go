package pkg

// Lipid は「日本人の食事摂取基準」（2020年版）の PDF 156 ページにある表「脂質の食事摂取基準」の情報を持つ
type Lipid struct {
	gender, age int
	datum       *LipidDatum
}

// NewLipid は Lipid を返す
func NewLipid(gender, age int) *Lipid {
	l := &Lipid{
		gender: gender,
		age:    age,
	}
	l.datum = l.GetDatum()
	return l
}

// NewLipidForPregnantWoman は妊婦向けに値を修正した Lipid を返す
func NewLipidForPregnantWoman(gender, age int) *Lipid {
	l := NewLipid(gender, age)
	if l.gender == GenderFemale {
		l.datum.DGMin = NewNullFloat64(20)
		l.datum.DGMax = NewNullFloat64(30)
	}
	return l
}

// NewLipidForLactatingWoman は授乳婦向けに値を修正した Lipid を返す
func NewLipidForLactatingWoman(gender, age int) *Lipid {
	l := NewLipid(gender, age)
	if l.gender == GenderFemale {
		l.datum.DGMin = NewNullFloat64(20)
		l.datum.DGMax = NewNullFloat64(30)
	}
	return l
}

// GetDatum は ProteinDatum を返す
func (l *Lipid) GetDatum() *LipidDatum {
	for _, d := range l.Data() {
		if d.Gender == l.gender && d.Age == l.age {
			return &d
		}
	}
	return nil
}

// GetAI は目安量（%エネルギー）を返す
// AI とは adequate intake の略で、
// 一定の栄養状態を維持するのに十分な量であり、目安量以上を摂取している場合は不足のリスクはほとんどない
func (l *Lipid) GetAI() (float64, bool) {
	if l.datum == nil {
		return 0, false
	}
	return l.datum.AI.Flatten()
}

// GetDG は目標量（%エネルギー）の上限と下限を返す
// DG とは tentative dietary goal for preventing life-style related diseases の略で、
// 生活習慣病の発症予防のために現在の日本人が当面の目標とすべき摂取量
func (l *Lipid) GetDG() (float64, float64, bool) {
	d := l.datum
	if d == nil {
		return 0, 0, false
	}
	min, max := d.DGMin, d.DGMax
	if min.Valid && max.Valid {
		return min.Value, max.Value, false
	}
	return 0, 0, false
}

// PercentEnergyToGram は目安量と目標量の単位である（%エネルギー）を（g）に変換します
func (l *Lipid) PercentEnergyToGram(kcal float64, ratio float64) float64 {
	return kcal * ratio / 100.0 / 9
}

// LipidDatum は「脂質の食事摂取基準」テーブルのデータを持つ
type LipidDatum struct {
	Gender int
	Age    int
	AI     NullFloat64 // 目安量
	DGMax  NullFloat64 // 目標量 上限
	DGMin  NullFloat64 // 目標量 下限
}
