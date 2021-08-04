package pkg

import "testing"

func TestEnergy_Get(t *testing.T) {
	samples := []struct {
		Gender, Age, ActivityLevel int
		Expect                     float64
		Valid                      bool
	}{
		{
			Gender:        GenderMale,
			Age:           Age18To29Years,
			ActivityLevel: ActivityLevelI,
			Expect:        2300,
			Valid:         true,
		},
		{
			Gender:        GenderFemale,
			Age:           Age3To5Years,
			ActivityLevel: ActivityLevelIII,
			Expect:        0,
			Valid:         false,
		},
		{
			Gender:        GenderFemale,
			Age:           Age3To5Years,
			ActivityLevel: ActivityLevelIII,
			Expect:        0,
			Valid:         false,
		},
	}
	for _, s := range samples {
		e := NewEnergy(s.Gender, s.Age, s.ActivityLevel)
		v, ok := e.Get()
		if ok != s.Valid || v != s.Expect {
			t.Errorf("invalid result. value:%v, ok:%v, sample:%v", v, ok, s)
		}
	}
}

func TestEnergy_GetForPregnantWoman(t *testing.T) {
	e := NewEnergy(GenderFemale, Age30To49Years, ActivityLevelII)
	v, ok := e.GetForPregnantWoman(TermEarly)
	if ok != true || v != 2100 {
		t.Errorf("invalid result. v:%v, ok:%v", v, ok)
	}
}
