package pkg

// Energy は「日本人の食事摂取基準」（2020年版）の PDF 91 ページにある表「推定エネルギー必要量」の情報を持つ
type Energy struct {
	gender, age, activityLevel int
	datum                      *EnergyDatum
}

// NewEnergy は Energy を返す
func NewEnergy(gender, age, activityLevel int) *Energy {
	e := &Energy{
		gender:        gender,
		age:           age,
		activityLevel: activityLevel,
	}
	e.datum = e.GetDatum()
	return e
}

// NewEnergyForPregnantWoman は妊婦向けに付加量を加えられた Energy を返す
func NewEnergyForPregnantWoman(gender, age, activityLevel, term int) *Energy {
	e := NewEnergy(gender, age, activityLevel)
	if gender == GenderFemale {
		switch term {
		case TermEarly:
			e.datum.Value.Value += 50
		case TermMid:
			e.datum.Value.Value += 250
		case TermLate:
			e.datum.Value.Value += 450
		}
	}
	return e
}

// NewEnergyForLactatingWoman は授乳婦向けに付加量を加えられた Energy を返す
func NewEnergyForLactatingWoman(gender, age, activityLevel int) *Energy {
	e := NewEnergy(gender, age, activityLevel)
	if gender == GenderFemale {
		e.datum.Value.Value += 350
	}
	return e
}

// GetDatum は EnergyDatum を返す
func (e *Energy) GetDatum() *EnergyDatum {
	for _, d := range e.Data() {
		if d.Gender == e.gender && d.Age == e.age && d.ActivityLevel == e.activityLevel {
			return &d
		}
	}
	return nil
}

// Get はエネルギーの必要量（kcal）を返す
func (e *Energy) Get() (float64, bool) {
	if e.datum != nil {
		return e.datum.Value.Flatten()
	}
	return 0, false
}

// EnergyDatum はエネルギーの必要量テーブル情報を持つ
type EnergyDatum struct {
	Gender        int
	Age           int
	ActivityLevel int
	Value         NullFloat64
}
