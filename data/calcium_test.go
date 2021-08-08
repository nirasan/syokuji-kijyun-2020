package data

import (
	"testing"
)

func TestCalciumList(t *testing.T) {
	list := CalciumList()
	t.Logf("---- 男性")
	for _, d := range list {
		if d.Gender == GenderMale {
			t.Log(d)
		}
	}
	t.Logf("---- 女性")
	for _, d := range list {
		if d.Gender == GenderFemale && d.Option == OptionNone {
			t.Log(d)
		}
	}
	t.Logf("---- 女性（妊娠初期）")
	for _, d := range list {
		if d.Gender == GenderFemale && d.Option == OptionEarlyPregnancy {
			t.Log(d)
		}
	}
	t.Logf("---- 女性（妊娠中期）")
	for _, d := range list {
		if d.Gender == GenderFemale && d.Option == OptionMidPregnancy {
			t.Log(d)
		}
	}
	t.Logf("---- 女性（妊娠後期）")
	for _, d := range list {
		if d.Gender == GenderFemale && d.Option == OptionLatePregnancy {
			t.Log(d)
		}
	}
	t.Logf("---- 女性（授乳中）")
	for _, d := range list {
		if d.Gender == GenderFemale && d.Option == OptionBreastfeeding {
			t.Log(d)
		}
	}
}
