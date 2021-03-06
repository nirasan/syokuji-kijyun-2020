package main

import (
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/data"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

// VitaminB2 は 265 ページの「ビタミン B2 の食事摂取基準」をパースして Go の構造体を作成します
func VitaminB2(output string) {
	lines := strings.Split(vitaminB2Src, "\n")
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

	list := make([]data.VitaminB2, 0)
	for i := 0; i <= 13; i++ {
		line := strings.ReplaceAll(lines[i], ",", "")
		cols := strings.Split(line, " ")
		j := len(cols) - 6
		a := ages[i]
		list = append(list,
			newVitaminB2(data.GenderMale, a[0], a[1], data.OptionNone, cols[j:j+3]),
			newVitaminB2(data.GenderFemale, a[0], a[1], data.OptionNone, cols[j+3:j+6]),
		)
	}

	list2 := make([]data.VitaminB2, 0)
	for _, d := range list {
		if d.Gender != data.GenderFemale {
			continue
		}
		if d.EAR.Valid && d.RDA.Valid {
			list2 = append(list2,
				newVitaminB2WithOption(d, data.OptionEarlyPregnancy),
				newVitaminB2WithOption(d, data.OptionMidPregnancy),
				newVitaminB2WithOption(d, data.OptionLatePregnancy),
				newVitaminB2WithOption(d, data.OptionBreastfeeding),
			)
		}
	}

	list = append(list, list2...)

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"cmd/generator/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package data\n")
	g.MustWrite("func VitaminB2List() []VitaminB2 {\n")
	g.MustWrite("    return []VitaminB2 {\n")
	for _, d := range list {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "data.", "")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(output)
}

func newVitaminB2(g data.Gender, from, to data.Age, o data.Option, cols []string) data.VitaminB2 {
	return data.VitaminB2{
		Gender: g,
		From:   from,
		To:     to,
		Option: o,
		EAR:    data.NilFloatFromString(cols[0]),
		RDA:    data.NilFloatFromString(cols[1]),
		AI:     data.NilFloatFromString(cols[2]),
	}
}

func newVitaminB2WithOption(in data.VitaminB2, o data.Option) data.VitaminB2 {
	d := data.VitaminB2{
		Gender: in.Gender,
		From:   in.From,
		To:     in.To,
		Option: o,
		EAR:    in.EAR,
		RDA:    in.RDA,
		AI:     in.AI,
	}
	var ear, rda float64
	switch o {
	case data.OptionEarlyPregnancy, data.OptionMidPregnancy, data.OptionLatePregnancy:
		ear, rda = 0.2, 0.3
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

const vitaminB2Src = `0 ～ 5 （月） ─ ─ 0.3 ─ ─ 0.3
6 ～11（月） ─ ─ 0.4 ─ ─ 0.4
1 ～ 2 （歳） 0.5 0.6 ─ 0.5 0.5 ─
3 ～ 5 （歳） 0.7 0.8 ─ 0.6 0.8 ─
6 ～ 7 （歳） 0.8 0.9 ─ 0.7 0.9 ─
8 ～ 9 （歳） 0.9 1.1 ─ 0.9 1.0 ─
10～11（歳） 1.1 1.4 ─ 1.0 1.3 ─
12～14（歳） 1.3 1.6 ─ 1.2 1.4 ─
15～17（歳） 1.4 1.7 ─ 1.2 1.4 ─
18～29（歳） 1.3 1.6 ─ 1.0 1.2 ─
30～49（歳） 1.3 1.6 ─ 1.0 1.2 ─
50～64（歳） 1.2 1.5 ─ 1.0 1.2 ─
65～74（歳） 1.2 1.5 ─ 1.0 1.2 ─
75 以上（歳） 1.1 1.3 ─ 0.9 1.0 ─
妊婦（付加量） ＋0.2 ＋0.3 ─
授乳婦（付加量） ＋0.5 ＋0.6 ─`
