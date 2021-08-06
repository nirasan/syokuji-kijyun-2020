package pkg

// DietaryFiber は「日本人の食事摂取基準」（2020年版）の PDF 172 ページにある表「食物繊維の食事摂取基準」の情報を持つ
type DietaryFiber struct {
	gender, age int
	datum       *DietaryFiberDatum
}

// NewDietaryFiber は DietaryFiber を返す
func NewDietaryFiber(gender, age int) *DietaryFiber {
	d := &DietaryFiber{
		gender: gender,
		age:    age,
	}
	d.datum = d.GetDatum()
	return d
}

// NewDietaryFiberForPregnantWoman は妊婦向けに値を修正した DietaryFiber を返す
func NewDietaryFiberForPregnantWoman(gender, age int) *DietaryFiber {
	d := NewDietaryFiber(gender, age)
	if d.gender == GenderFemale {
		d.datum.DGMin = NewNullFloat64(18)
	}
	return d
}

// NewDietaryFiberForLactatingWoman は授乳婦向けに値を修正した DietaryFiber を返す
func NewDietaryFiberForLactatingWoman(gender, age int) *DietaryFiber {
	d := NewDietaryFiber(gender, age)
	if d.gender == GenderFemale {
		d.datum.DGMin = NewNullFloat64(18)
	}
	return d
}

// GetDatum は ProteinDatum を返す
func (d *DietaryFiber) GetDatum() *DietaryFiberDatum {
	for _, datum := range d.Data() {
		if datum.Gender == d.gender && datum.Age == d.age {
			return &datum
		}
	}
	return nil
}

// GetDG は目標量（g）の下限を返す
// DG とは tentative dietary goal for preventing life-style related diseases の略で、
// 生活習慣病の発症予防のために現在の日本人が当面の目標とすべき摂取量
func (d *DietaryFiber) GetDG() (float64, bool) {
	if d.datum == nil {
		return 0, false
	}
	return d.datum.DGMin.Flatten()
}

// DietaryFiberDatum は「食物繊維の食事摂取基準」テーブルのデータを持つ
type DietaryFiberDatum struct {
	Gender int
	Age    int
	DGMin  NullFloat64 // 目標量 下限
}
