package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtein_GetEAR(t *testing.T) {
	p := NewProtein(GenderFemale, Age15To17Years)
	v, ok := p.GetEAR()
	assert.Equal(t, true, ok)
	assert.Equal(t, 45.0, v)
}

func TestProtein_GetRDA(t *testing.T) {
	p := NewProtein(GenderFemale, Age15To17Years)
	v, ok := p.GetRDA()
	assert.Equal(t, true, ok)
	assert.Equal(t, 55.0, v)
}

func TestProtein_GetAI(t *testing.T) {
	p := NewProtein(GenderFemale, Age0Years6To8Months)
	v, ok := p.GetAI()
	assert.Equal(t, true, ok)
	assert.Equal(t, 15.0, v)
}

func TestProtein_GetDG(t *testing.T) {
	p := NewProtein(GenderFemale, Age65To74Years)
	min, max, ok := p.GetDG()
	assert.Equal(t, true, ok)
	assert.Equal(t, 15.0, min)
	assert.Equal(t, 20.0, max)
}

func TestNewProteinForPregnantWoman(t *testing.T) {
	p := NewProteinForPregnantWoman(GenderFemale, Age15To17Years, TermLate)
	v, ok := p.GetEAR()
	assert.Equal(t, true, ok)
	assert.Equal(t, 65.0, v)
}

func TestNewProteinForLactatingWoman(t *testing.T) {
	p := NewProteinForLactatingWoman(GenderFemale, Age15To17Years)
	v, ok := p.GetEAR()
	assert.Equal(t, true, ok)
	assert.Equal(t, 60.0, v)
}

func TestProtein_PercentEnergyToGram(t *testing.T) {
	p := NewProtein(GenderFemale, Age15To17Years)
	min, max, ok := p.GetDG()
	assert.Equal(t, true, ok)
	assert.Equal(t, 13.0, min)
	assert.Equal(t, 20.0, max)

	e := NewEnergy(GenderFemale, Age15To17Years, ActivityLevelI)
	energy, ok := e.Get()
	assert.Equal(t, true, ok)
	assert.Equal(t, 2050.0, energy)

	gram := p.PercentEnergyToGram(energy, min)
	assert.Equal(t, 66.625, gram)
}
