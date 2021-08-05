package pkg

// Protein は「日本人の食事摂取基準」（2020年版）の PDF 133 ページにある表「たんぱく質の食事摂取基準」の情報を持つ
type Protein struct {
	gender, age int
	datum       *ProteinDatum
}

// NewProtein は Protein を返す
func NewProtein(gender, age int) *Protein {
	p := &Protein{
		gender: gender,
		age:    age,
	}
	p.datum = p.GetDatum()
	return p
}

// NewProtein は Protein を返す
func NewProteinForPregnantWoman(gender, age, term int) *Protein {
	p := NewProtein(gender, age)
	if p.gender == GenderFemale {
		switch term {
		case TermEarly:
			p.datum.EAR.Value += 0
			p.datum.RDA.Value += 0
			p.datum.DGMin.Value = 13
			p.datum.DGMax.Value = 20
		case TermMid:
			p.datum.EAR.Value += 5
			p.datum.RDA.Value += 5
			p.datum.DGMin.Value = 13
			p.datum.DGMax.Value = 20
		case TermLate:
			p.datum.EAR.Value += 20
			p.datum.RDA.Value += 25
			p.datum.DGMin.Value = 15
			p.datum.DGMax.Value = 20
		}
	}
	return p
}

// NewProtein は Protein を返す
func NewProteinForLactatingWoman(gender, age int) *Protein {
	p := NewProtein(gender, age)
	if p.gender == GenderFemale {
		p.datum.EAR.Value += 15
		p.datum.RDA.Value += 20
		p.datum.DGMin.Value = 15
		p.datum.DGMax.Value = 20
	}
	return p
}

// GetDatum は ProteinDatum を返す
func (p *Protein) GetDatum() *ProteinDatum {
	for _, d := range p.Data() {
		if d.Gender == p.gender && d.Age == p.age {
			return &d
		}
	}
	return nil
}

// GetEAR は推定平均必要量（g）を返す
// EAR とは estimated average requirement の略で、半数の者が必要量を満たす量
func (p *Protein) GetEAR() (float64, bool) {
	d := p.datum
	if d == nil {
		return 0, false
	}
	return d.EAR.Flatten()
}

// GetRDA は推奨量を（g）返す
// RDA とは recommended dietary allowance の略で, ほとんどの者が充足している量
func (p *Protein) GetRDA() (float64, bool) {
	d := p.datum
	if d == nil {
		return 0, false
	}
	return d.RDA.Flatten()
}

// GetAI は目安量（g）を返す
// AI とは adequate intake の略で、
// 一定の栄養状態を維持するのに十分な量であり、目安量以上を摂取している場合は不足のリスクはほとんどない
func (p *Protein) GetAI() (float64, bool) {
	d := p.datum
	if d == nil {
		return 0, false
	}
	return d.AI.Flatten()
}

// GetDG は目標量（%エネルギー）の上限と下限を返す
// DG とは tentative dietary goal for preventing life-style related diseases の略で、
// 生活習慣病の発症予防のために現在の日本人が当面の目標とすべき摂取量
// 目標量はエネルギーに対する割合でありタンパク質の場合 4kcal が 1g になる
func (p *Protein) GetDG() (float64, float64, bool) {
	d := p.datum
	if d == nil {
		return 0, 0, false
	}
	min, max := d.DGMin, d.DGMax
	if min.Valid && max.Valid {
		return min.Value, max.Value, true
	}
	return 0, 0, false
}

// PercentEnergyToGram は目標量の単位である（%エネルギー）を（g）に変換します
func (p *Protein) PercentEnergyToGram(kcal float64, ratio float64) float64 {
	return kcal * ratio / 100.0 / 4
}

// ProteinDatum は「たんぱく質の食事摂取基準」テーブルのデータをもつ
type ProteinDatum struct {
	Gender int
	Age    int
	EAR    NullFloat64 // 推定平均必要量
	RDA    NullFloat64 // 推奨量
	AI     NullFloat64 // 目安量
	DGMin  NullFloat64 // 目標量 下限
	DGMax  NullFloat64 // 目標量 上限
}
