package pkg

// Carbohydrates は「日本人の食事摂取基準」（2020年版）の PDF 171 ページにある表「炭水化物の食事摂取基準」の情報を持つ
type Carbohydrates struct {
	gender, age int
	datum       *CarbohydratesDatum
}

// NewCarbohydrates は Carbohydrates を返す
func NewCarbohydrates(gender, age int) *Carbohydrates {
	c := &Carbohydrates{
		gender: gender,
		age:    age,
	}
	c.datum = c.GetDatum()
	return c
}

// NewCarbohydratesForPregnantWoman は妊婦向けに値を修正した Carbohydrates を返す
func NewCarbohydratesForPregnantWoman(gender, age int) *Carbohydrates {
	c := NewCarbohydrates(gender, age)
	if c.gender == GenderFemale {
		c.datum.DGMin = NewNullFloat64(50)
		c.datum.DGMax = NewNullFloat64(65)
	}
	return c
}

// NewCarbohydratesForLactatingWoman は授乳婦向けに値を修正した Carbohydrates を返す
func NewCarbohydratesForLactatingWoman(gender, age int) *Carbohydrates {
	c := NewCarbohydrates(gender, age)
	if c.gender == GenderFemale {
		c.datum.DGMin = NewNullFloat64(50)
		c.datum.DGMax = NewNullFloat64(65)
	}
	return c
}

// GetDatum は ProteinDatum を返す
func (c *Carbohydrates) GetDatum() *CarbohydratesDatum {
	for _, d := range c.Data() {
		if d.Gender == c.gender && d.Age == c.age {
			return &d
		}
	}
	return nil
}

// GetDG は目標量（%エネルギー）の上限と下限を返す
// DG とは tentative dietary goal for preventing life-style related diseases の略で、
// 生活習慣病の発症予防のために現在の日本人が当面の目標とすべき摂取量
func (c *Carbohydrates) GetDG() (float64, float64, bool) {
	d := c.datum
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
func (c *Carbohydrates) PercentEnergyToGram(kcal float64, ratio float64) float64 {
	return kcal * ratio / 100.0 / 4
}

// CarbohydratesDatum は「炭水化物の食事摂取基準」テーブルのデータを持つ
type CarbohydratesDatum struct {
	Gender int
	Age    int
	DGMax  NullFloat64 // 目標量 上限
	DGMin  NullFloat64 // 目標量 下限
}
