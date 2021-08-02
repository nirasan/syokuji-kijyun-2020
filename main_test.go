package main

import "testing"

func TestGetEnergy(t *testing.T) {
	samples := []struct {
		Sex, Age, Work int
		Expect         float64
	}{
		{SexMale, Age0Month0to5, WorkI, 550},
		{SexFemale, Age6to7, WorkIII, 1650},
		{SexMale, Age18to29, WorkII, 2650},
	}
	for _, s := range samples {
		result := GetEnergy(s.Sex, s.Age, s.Work)
		if result != s.Expect {
			t.Errorf("invalid result. sample:%+v, result:%f", s, result)
		}
	}
}

func TestGetLipid(t *testing.T) {
	samples := []struct {
		Sex, Age, Work int
		Min, Max       float64
	}{
		{SexMale, Age0Month0to5, WorkI, 30.55, 30.55},
		{SexFemale, Age6to7, WorkIII, 36.66, 55},
		{SexMale, Age18to29, WorkII, 58.88, 88.33},
	}
	for _, s := range samples {
		a, b := GetLipid(s.Sex, s.Age, s.Work)
		t.Logf("min:%v, max:%v", a, b)
		if (a-s.Min) > 0.01 || (b-s.Max) > 0.01 {
			t.Errorf("invalid value, min:%v, max:%v, sample:%v", a, b, s)
		}
	}
}

func TestGetCarbohydrates(t *testing.T) {
	samples := []struct {
		Sex, Age, Work int
		Min, Max       float64
	}{
		{SexMale, Age0Month0to5, WorkI, 0, 0},
		{SexFemale, Age6to7, WorkIII, 206.25, 268.125},
		{SexMale, Age18to29, WorkII, 331.25, 430.625},
	}
	for _, s := range samples {
		a, b := GetCarbohydrates(s.Sex, s.Age, s.Work)
		t.Logf("min:%v, max:%v", a, b)
		if (a-s.Min) > 0.01 || (b-s.Max) > 0.01 {
			t.Errorf("invalid value, min:%v, max:%v, sample:%v", a, b, s)
		}
	}
}
