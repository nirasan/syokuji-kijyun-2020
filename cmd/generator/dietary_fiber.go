package main

import (
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/data"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

// DietaryFiber は PDF 172 ページの「食物繊維の食事摂取基準」をパースして Go の構造体を作成します
func DietaryFiber(output string) {
	lines := strings.Split(dietaryFiberSrc, "\n")
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

	list := make([]data.DietaryFiber, 0)
	for i := 0; i <= 13; i++ {
		line := strings.ReplaceAll(lines[i], " 以上", "")
		cols := strings.Split(line, " ")
		j := len(cols) - 2
		a := ages[i]
		list = append(list,
			newDietaryFiber(data.GenderMale, a[0], a[1], data.OptionNone, cols[j]),
			newDietaryFiber(data.GenderFemale, a[0], a[1], data.OptionNone, cols[j+1]),
		)
	}

	list2 := make([]data.DietaryFiber, 0)
	for _, d := range list {
		if d.Gender != data.GenderFemale {
			continue
		}
		if d.DG.Valid {
			list2 = append(list2,
				newDietaryFiberWithOption(d, data.OptionEarlyPregnancy),
				newDietaryFiberWithOption(d, data.OptionMidPregnancy),
				newDietaryFiberWithOption(d, data.OptionLatePregnancy),
				newDietaryFiberWithOption(d, data.OptionBreastfeeding),
			)
		}
	}

	list = append(list, list2...)

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"cmd/generator/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package data\n")
	g.MustWrite("func DietaryFiberList() []DietaryFiber {\n")
	g.MustWrite("    return []DietaryFiber {\n")
	for _, d := range list {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "data.", "")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(output)
}

func newDietaryFiber(g data.Gender, from, to data.Age, o data.Option, col string) data.DietaryFiber {
	return data.DietaryFiber{
		Gender: g,
		From:   from,
		To:     to,
		Option: o,
		DG:     data.NilFloatFromString(col),
	}
}

func newDietaryFiberWithOption(in data.DietaryFiber, o data.Option) data.DietaryFiber {
	d := data.DietaryFiber{
		Gender: in.Gender,
		From:   in.From,
		To:     in.To,
		Option: o,
		DG:     in.DG,
	}
	switch o {
	case data.OptionEarlyPregnancy, data.OptionMidPregnancy, data.OptionLatePregnancy, data.OptionBreastfeeding:
		d.DG.Float = 18
	}
	return d
}

// PDF 172 ページの「食物繊維の食事摂取基準」
const dietaryFiberSrc = `0 ～ 5 （月） ─ ─
6 ～11（月） ─ ─
1 ～ 2 （歳） ─ ─
3 ～ 5 （歳） 8 以上 8 以上
6 ～ 7 （歳） 10 以上 10 以上
8 ～ 9 （歳） 11 以上 11 以上
10～11（歳） 13 以上 13 以上
12～14（歳） 17 以上 17 以上
15～17（歳） 19 以上 18 以上
18～29（歳） 21 以上 18 以上
30～49（歳） 21 以上 18 以上
50～64（歳） 21 以上 18 以上
65～74（歳） 20 以上 17 以上
75 以上（歳） 20 以上 17 以上
妊　婦 18 以上
授乳婦 18 以上`
