package main

import (
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/data"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

// Omega3FattyAcids は PDF 158 ページの「n─3 系脂肪酸の食事摂取基準」をパースして Go の構造体を作成します
func Omega3FattyAcids(output string) {
	lines := strings.Split(omega3FattyAcidsSrc, "\n")
	ages := [][]data.Age{
		{{Month: 0}, {Month: 5}},
		{{Month: 6}, {Month: 11}},
		{{Year: 1}, {Year: 2}},
		{{Year: 3}, {Year: 5}},
		{{Year: 6}, {Year: 7}},
		{{Year: 8}, {Year: 9}},
		{{Year: 10}, {Year: 11}},
		{{Year: 12}, {Year: 14}},
		{{Year: 15}, {Year: 17}},
		{{Year: 18}, {Year: 29}},
		{{Year: 30}, {Year: 49}},
		{{Year: 50}, {Year: 64}},
		{{Year: 65}, {Year: 74}},
		{{Year: 75}, {Year: 100}},
	}

	list := make([]data.Omega3FattyAcids, 0)
	for i := 0; i <= 13; i++ {
		cols := strings.Split(lines[i], " ")
		j := len(cols) - 2
		a := ages[i]
		list = append(list,
			newOmega3FattyAcids(data.GenderMale, a[0], a[1], data.OptionNone, cols[j]),
			newOmega3FattyAcids(data.GenderFemale, a[0], a[1], data.OptionNone, cols[j+1]),
		)
	}

	list2 := make([]data.Omega3FattyAcids, 0)
	for _, d := range list {
		if d.Gender != data.GenderFemale {
			continue
		}
		list2 = append(list2,
			newOmega3FattyAcidsWithOption(d, data.OptionEarlyPregnancy),
			newOmega3FattyAcidsWithOption(d, data.OptionMidPregnancy),
			newOmega3FattyAcidsWithOption(d, data.OptionLatePregnancy),
			newOmega3FattyAcidsWithOption(d, data.OptionBreastfeeding),
		)
	}

	list = append(list, list2...)

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"cmd/generator/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package data\n")
	g.MustWrite("func Omega3FattyAcidsList() []Omega3FattyAcids {\n")
	g.MustWrite("    return []Omega3FattyAcids {\n")
	for _, d := range list {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "data.", "")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(output)
}

func newOmega3FattyAcids(g data.Gender, from, to data.Age, o data.Option, col string) data.Omega3FattyAcids {
	return data.Omega3FattyAcids{
		Gender: g,
		From:   from,
		To:     to,
		Option: o,
		AI:     data.NilFloatFromString(col),
	}
}

func newOmega3FattyAcidsWithOption(in data.Omega3FattyAcids, o data.Option) data.Omega3FattyAcids {
	d := data.Omega3FattyAcids{
		Gender: in.Gender,
		From:   in.From,
		To:     in.To,
		Option: o,
		AI:     in.AI,
	}
	switch o {
	case data.OptionEarlyPregnancy, data.OptionMidPregnancy, data.OptionLatePregnancy:
		d.AI.Float = 1.6
	case data.OptionBreastfeeding:
		d.AI.Float = 1.8
	}
	return d
}

const omega3FattyAcidsSrc = `0 ～ 5 （月） 0.9 0.9
6 ～11（月） 0.8 0.8
1 ～ 2 （歳） 0.7 0.8
3 ～ 5 （歳） 1.1 1.0
6 ～ 7 （歳） 1.5 1.3
8 ～ 9 （歳） 1.5 1.3
10～11（歳） 1.6 1.6
12～14（歳） 1.9 1.6
15～17（歳） 2.1 1.6
18～29（歳） 2.0 1.6
30～49（歳） 2.0 1.6
50～64（歳） 2.2 1.9
65～74（歳） 2.2 2.0
75 以上（歳） 2.1 1.8
妊　婦 1.6
授乳婦 1.8`
