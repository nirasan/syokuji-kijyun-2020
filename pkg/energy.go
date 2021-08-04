package pkg

// Energy は「日本人の食事摂取基準」（2020年版）の PDF 91 ページにある表「推定エネルギー必要量」の情報を持つ
type Energy struct {
	gender, age, activityLevel int
}

// NewEnergy は Energy の構造体を返す
func NewEnergy(gender, age, activityLevel int) *Energy {
	return &Energy{
		gender:        gender,
		age:           age,
		activityLevel: activityLevel,
	}
}

// Get はエネルギーの必要量を返す
func (e *Energy) Get() (float64, bool) {
	for _, d := range e.Data() {
		if d.Gender == e.gender && d.Age == e.age && d.ActivityLevel == e.activityLevel {
			return d.Value, d.Valid
		}
	}
	return 0, false
}

// GetForPregnantWoman は妊婦のエネルギーの必要量を返す
func (e *Energy) GetForPregnantWoman(term int) (float64, bool) {
	v, ok := e.Get()
	if ok && e.gender == GenderFemale {
		switch term {
		case TermEarly:
			v += 50
		case TermMid:
			v += 250
		case TermLate:
			v += 450
		}
	}
	return v, ok
}

// GetForLactatingWoman は授乳婦のエネルギーの必要量を返す
func (e *Energy) GetForLactatingWoman() (float64, bool) {
	v, ok := e.Get()
	if ok && e.gender == GenderFemale {
		v += 350
	}
	return v, ok
}

// EnergyDatum はエネルギーの必要量テーブル情報を持つ
type EnergyDatum struct {
	Gender        int
	Age           int
	ActivityLevel int
	Value         float64
	Valid         bool
}
