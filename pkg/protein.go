package pkg

// Protein は「日本人の食事摂取基準」（2020年版）の PDF 133 ページにある表「たんぱく質の食事摂取基準」の情報を持つ
type Protein struct {
	gender, age int
}

// NewProtein は Protein を返す
func NewProtein(gender, age int) *Protein {
	return &Protein{
		gender: gender,
		age:    age,
	}
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
	d := p.GetDatum()
	if d == nil {
		return 0, false
	}
	return d.EAR.Flatten()
}

func (p *Protein) GetEARForPregnantWoman(term int) (float64, bool) {
	v, ok := p.GetEAR()
	if p.gender == GenderFemale {
		switch term {
		case TermEarly:
			v += 0
		case TermMid:
			v += 5
		case TermLate:
			v += 20
		}
	}
	return v, ok
}

func (p *Protein) GetEARForLactatingWoman() (float64, bool) {
	v, ok := p.GetEAR()
	if ok && p.gender == GenderFemale {
		v += 15
	}
	return v, ok
}

// GetRDA は推奨量を（g）返す
// RDA とは recommended dietary allowance の略で, ほとんどの者が充足している量
func (p *Protein) GetRDA() (float64, bool) {
	d := p.GetDatum()
	if d == nil {
		return 0, false
	}
	return d.RDA.Flatten()
}

func (p *Protein) GetRDAForPregnantWoman(term int) (float64, bool) {
	v, ok := p.GetRDA()
	if p.gender == GenderFemale {
		switch term {
		case TermEarly:
			v += 0
		case TermMid:
			v += 5
		case TermLate:
			v += 25
		}
	}
	return v, ok
}

func (p *Protein) GetRDAForLactatingWoman() (float64, bool) {
	v, ok := p.GetRDA()
	if ok && p.gender == GenderFemale {
		v += 20
	}
	return v, ok
}

// GetAI は目安量（g）を返す
// AI とは adequate intake の略で、
//一定の栄養状態を維持するのに十分な量であり、目安量以上を摂取している場合は不足のリスクはほとんどない
func (p *Protein) GetAI() (float64, bool) {
	d := p.GetDatum()
	if d == nil {
		return 0, false
	}
	return d.AI.Flatten()
}

// GetDG は目標量（%）の上限と下限を返す
// DG とは tentative dietary goal for preventing life-style related diseases の略で、
// 生活習慣病の発症予防のために現在の日本人が当面の目標とすべき摂取量
// 目標量はエネルギーに対する割合であり 4kcal が 1g になる
func (p *Protein) GetDG() (float64, float64, bool) {
	d := p.GetDatum()
	if d == nil {
		return 0, 0, false
	}
	min, max := d.DGMin, d.DGMax
	if min.Valid && max.Valid {
		return min.Value, max.Value, false
	}
	return 0, 0, false
}

func (p *Protein) GetDGForPregnantWoman(term int) (float64, float64, bool) {
	switch term {
	case TermEarly, TermMid:
		return 13, 20, true
	case TermLate:
		return 15, 20, true
	}
	return 0, 0, false
}

func (p *Protein) GetDGForLactatingWoman() (float64, float64, bool) {
	return 15, 20, false
}

type ProteinDatum struct {
	Gender int
	Age    int
	EAR    NullFloat64 // 推定平均必要量
	RDA    NullFloat64 // 推奨量
	AI     NullFloat64 // 目安量
	DGMax  NullFloat64 // 目標量上限
	DGMin  NullFloat64 // 目標量下限
}
