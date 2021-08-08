package main

import (
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/data"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

func Copper(output string) {
	lines := strings.Split(copperSrc, "\n")
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

	list := make([]data.Copper, 0)
	for i := 0; i <= 13; i++ {
		line := strings.ReplaceAll(lines[i], ",", "")
		cols := strings.Split(line, " ")
		j := len(cols) - 8
		a := ages[i]
		list = append(list,
			newCopper(data.GenderMale, a[0], a[1], data.OptionNone, cols[j:j+4]),
			newCopper(data.GenderFemale, a[0], a[1], data.OptionNone, cols[j+4:j+8]),
		)
	}

	list2 := make([]data.Copper, 0)
	for _, d := range list {
		if d.Gender != data.GenderFemale {
			continue
		}
		list2 = append(list2,
			newCopperWithOption(d, data.OptionEarlyPregnancy),
			newCopperWithOption(d, data.OptionMidPregnancy),
			newCopperWithOption(d, data.OptionLatePregnancy),
			newCopperWithOption(d, data.OptionBreastfeeding),
		)
	}

	list = append(list, list2...)

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"cmd/generator/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package data\n")
	g.MustWrite("func CopperList() []Copper {\n")
	g.MustWrite("    return []Copper {\n")
	for _, d := range list {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "data.", "")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(output)
}

func newCopper(g data.Gender, from, to data.Age, o data.Option, cols []string) data.Copper {
	return data.Copper{
		Gender: g,
		From:   from,
		To:     to,
		Option: o,
		EAR:    data.NilFloatFromString(cols[0]),
		RDA:    data.NilFloatFromString(cols[1]),
		AI:     data.NilFloatFromString(cols[2]),
		UL:     data.NilFloatFromString(cols[3]),
	}
}

func newCopperWithOption(in data.Copper, o data.Option) data.Copper {
	d := data.Copper{
		Gender: in.Gender,
		From:   in.From,
		To:     in.To,
		Option: o,
		EAR:    in.EAR,
		RDA:    in.RDA,
		AI:     in.AI,
		UL:     in.UL,
	}
	var ear, rda float64
	switch o {
	case data.OptionEarlyPregnancy, data.OptionMidPregnancy, data.OptionLatePregnancy:
		ear, rda = 0.1, 0.1
	case data.OptionBreastfeeding:
		ear, rda = 0.5, 0.6
	}
	if d.EAR.Valid {
		d.EAR.Float += ear
	}
	if d.RDA.Valid {
		d.RDA.Float += rda
	}
	return d
}

// PDF 374 ページの「銅の食事摂取基準」
const copperSrc = `0 ～ 5 	（月） － － 0.3 － － － 0.3 －
6 ～11（月） － － 0.3 － － － 0.3 －
1 ～ 2 	（歳） 0.3 0.3 － － 0.2 0.3 － －
3 ～ 5 	（歳） 0.3 0.4 － － 0.3 0.3 － －
6 ～ 7 	（歳） 0.4 0.4 － － 0.4 0.4 － －
8 ～ 9 	（歳） 0.4 0.5 － － 0.4 0.5 － －
10～11（歳） 0.5 0.6 － － 0.5 0.6 － －
12～14（歳） 0.7 0.8 － － 0.6 0.8 － －
15～17（歳） 0.8 0.9 － － 0.6 0.7 － －
18～29（歳） 0.7 0.9 － 7 0.6 0.7 － 7
30～49（歳） 0.7 0.9 － 7 0.6 0.7 － 7
50～64（歳） 0.7 0.9 － 7 0.6 0.7 － 7
65～74（歳） 0.7 0.9 － 7 0.6 0.7 － 7
75 以上（歳） 0.7 0.8 － 7 0.6 0.7 － 7
妊婦（付加量） ＋0.1 ＋0.1 ─ ─
授乳婦（付加量） ＋0.5 ＋0.6 ─ ─`
