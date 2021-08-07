package data

import (
	"strconv"
	"strings"
)

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
)

type ActivityLevel string

const (
	ActivityLevel1 ActivityLevel = "Level1"
	ActivityLevel2 ActivityLevel = "Level2"
	ActivityLevel3 ActivityLevel = "Level3"
)

type Option string

const (
	OptionNone           Option = ""
	OptionEarlyPregnancy Option = "EarlyPregnancy"
	OptionMidPregnancy   Option = "MidPregnancy"
	OptionLatePregnancy  Option = "LatePregnancy"
	OptionBreastfeeding  Option = "Breastfeeding"
)

type Age struct {
	Year  int
	Month int
}

type NilFloat struct {
	Float float64
	Valid bool
}

func NilFloatFromString(s string) NilFloat {
	s = strings.ReplaceAll(s, ",", "")
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return NilFloat{Valid: false}
	}
	return NilFloat{Float: v, Valid: true}
}

type Energy struct {
	Gender        Gender
	From          Age
	To            Age
	ActivityLevel ActivityLevel
	Option        Option
	Value         NilFloat
}
