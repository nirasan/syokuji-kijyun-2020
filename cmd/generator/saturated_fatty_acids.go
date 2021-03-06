package main

import (
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/data"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

// SaturatedFattyAcids は PDF 157 ページの「飽和脂肪酸の食事摂取基準」をパースして Go の構造体を作成します
func SaturatedFattyAcids(output string) {
	lines := strings.Split(saturatedFattyAcidsSrc, "\n")
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

	list := make([]data.SaturatedFattyAcids, 0)
	for i := 0; i <= 13; i++ {
		line := strings.ReplaceAll(lines[i], " 以下", "")
		cols := strings.Split(line, " ")
		j := len(cols) - 2
		a := ages[i]
		list = append(list,
			newSaturatedFattyAcids(data.GenderMale, a[0], a[1], data.OptionNone, cols[j]),
			newSaturatedFattyAcids(data.GenderFemale, a[0], a[1], data.OptionNone, cols[j+1]),
		)
	}

	list2 := make([]data.SaturatedFattyAcids, 0)
	for _, d := range list {
		if d.Gender != data.GenderFemale {
			continue
		}
		if d.DG.Valid {
			list2 = append(list2,
				newSaturatedFattyAcidsWithOption(d, data.OptionEarlyPregnancy),
				newSaturatedFattyAcidsWithOption(d, data.OptionMidPregnancy),
				newSaturatedFattyAcidsWithOption(d, data.OptionLatePregnancy),
				newSaturatedFattyAcidsWithOption(d, data.OptionBreastfeeding),
			)
		}
	}

	list = append(list, list2...)

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"cmd/generator/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package data\n")
	g.MustWrite("func SaturatedFattyAcidsList() []SaturatedFattyAcids {\n")
	g.MustWrite("    return []SaturatedFattyAcids {\n")
	for _, d := range list {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "data.", "")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(output)
}

func newSaturatedFattyAcids(g data.Gender, from, to data.Age, o data.Option, col string) data.SaturatedFattyAcids {
	return data.SaturatedFattyAcids{
		Gender: g,
		From:   from,
		To:     to,
		Option: o,
		DG:     data.NilFloatFromString(col),
	}
}

func newSaturatedFattyAcidsWithOption(in data.SaturatedFattyAcids, o data.Option) data.SaturatedFattyAcids {
	d := data.SaturatedFattyAcids{
		Gender: in.Gender,
		From:   in.From,
		To:     in.To,
		Option: o,
		DG:     in.DG,
	}
	switch o {
	case data.OptionEarlyPregnancy, data.OptionMidPregnancy, data.OptionLatePregnancy, data.OptionBreastfeeding:
		d.DG.Float = 7
	}
	return d
}

const saturatedFattyAcidsSrc = `0 ～ 5 （月） ─ ─
6 ～11（月） ─ ─
1 ～ 2 （歳） ─ ─
3 ～ 5 （歳） 10 以下 10 以下
6 ～ 7 （歳） 10 以下 10 以下
8 ～ 9 （歳） 10 以下 10 以下
10～11（歳） 10 以下 10 以下
12～14（歳） 10 以下 10 以下
15～17（歳） 8 以下 8 以下
18～29（歳） 7 以下 7 以下
30～49（歳） 7 以下 7 以下
50～64（歳） 7 以下 7 以下
65～74（歳） 7 以下 7 以下
75 以上（歳） 7 以下 7 以下
妊　婦 7 以下
授乳婦 7 以下`
