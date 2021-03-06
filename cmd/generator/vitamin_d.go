package main

import (
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/data"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

// VitaminD は 213 ページの「ビタミン D の食事摂取基準」をパースして Go の構造体を作成します
func VitaminD(output string) {
	lines := strings.Split(vitaminDSrc, "\n")
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

	list := make([]data.VitaminD, 0)
	for i := 0; i <= 13; i++ {
		line := strings.ReplaceAll(lines[i], ",", "")
		cols := strings.Split(line, " ")
		j := len(cols) - 4
		a := ages[i]
		list = append(list,
			newVitaminD(data.GenderMale, a[0], a[1], data.OptionNone, cols[j:j+2]),
			newVitaminD(data.GenderFemale, a[0], a[1], data.OptionNone, cols[j+2:j+4]),
		)
	}

	list2 := make([]data.VitaminD, 0)
	for _, d := range list {
		if d.Gender != data.GenderFemale {
			continue
		}
		list2 = append(list2,
			newVitaminDWithOption(d, data.OptionEarlyPregnancy),
			newVitaminDWithOption(d, data.OptionMidPregnancy),
			newVitaminDWithOption(d, data.OptionLatePregnancy),
			newVitaminDWithOption(d, data.OptionBreastfeeding),
		)
	}

	list = append(list, list2...)

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"cmd/generator/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package data\n")
	g.MustWrite("func VitaminDList() []VitaminD {\n")
	g.MustWrite("    return []VitaminD {\n")
	for _, d := range list {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "data.", "")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(output)
}

func newVitaminD(g data.Gender, from, to data.Age, o data.Option, cols []string) data.VitaminD {
	return data.VitaminD{
		Gender: g,
		From:   from,
		To:     to,
		Option: o,
		AI:     data.NilFloatFromString(cols[0]),
		UL:     data.NilFloatFromString(cols[1]),
	}
}

func newVitaminDWithOption(in data.VitaminD, o data.Option) data.VitaminD {
	d := data.VitaminD{
		Gender: in.Gender,
		From:   in.From,
		To:     in.To,
		Option: o,
		AI:     in.AI,
		UL:     in.UL,
	}
	var ai float64
	switch o {
	case data.OptionEarlyPregnancy, data.OptionMidPregnancy, data.OptionLatePregnancy, data.OptionBreastfeeding:
		ai = 8.5
	}
	if d.AI.Valid {
		d.AI.Float = ai
	}
	return d
}

const vitaminDSrc = `0 ～ 5 （月） 5.0 25 5.0 25
6 ～11（月） 5.0 25 5.0 25
1 ～ 2 （歳） 3.0 20 3.5 20
3 ～ 5 （歳） 3.5 30 4.0 30
6 ～ 7 （歳） 4.5 30 5.0 30
8 ～ 9 （歳） 5.0 40 6.0 40
10～11（歳） 6.5 60 8.0 60
12～14（歳） 8.0 80 9.5 80
15～17（歳） 9.0 90 8.5 90
18～29（歳） 8.5 100 8.5 100
30～49（歳） 8.5 100 8.5 100
50～64（歳） 8.5 100 8.5 100
65～74（歳） 8.5 100 8.5 100
75 以上（歳） 8.5 100 8.5 100
妊　婦 8.5 ─
授乳婦 8.5 ─`
