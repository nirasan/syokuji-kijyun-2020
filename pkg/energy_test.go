package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	for i, s := range samples {
		e := NewEnergy(s.Gender, s.Age, s.ActivityLevel)
		t.Logf("[%d]%+v::: %+v", i, s, *e.datum)
		v, ok := e.Get()
		assert.Equal(t, s.Valid, ok)
		assert.Equal(t, s.Expect, v)
	}
}

func TestNewEnergyForPregnantWoman(t *testing.T) {
	e := NewEnergyForPregnantWoman(GenderFemale, Age30To49Years, ActivityLevelII, TermEarly)
	v, ok := e.Get()
	assert.Equal(t, true, ok)
	assert.Equal(t, 2100.0, v)
}
