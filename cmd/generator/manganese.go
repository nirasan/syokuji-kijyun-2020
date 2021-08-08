package main

import (
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/data"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

func Manganese(output string) {
	lines := strings.Split(manganeseSrc, "\n")
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

	list := make([]data.Manganese, 0)
	for i := 0; i <= 13; i++ {
		line := strings.ReplaceAll(lines[i], ",", "")
		cols := strings.Split(line, " ")
		j := len(cols) - 4
		a := ages[i]
		list = append(list,
			newManganese(data.GenderMale, a[0], a[1], data.OptionNone, cols[j:j+2]),
			newManganese(data.GenderFemale, a[0], a[1], data.OptionNone, cols[j+2:j+4]),
		)
	}

	list2 := make([]data.Manganese, 0)
	for _, d := range list {
		if d.Gender != data.GenderFemale {
			continue
		}
		list2 = append(list2,
			newManganeseWithOption(d, data.OptionEarlyPregnancy),
			newManganeseWithOption(d, data.OptionMidPregnancy),
			newManganeseWithOption(d, data.OptionLatePregnancy),
			newManganeseWithOption(d, data.OptionBreastfeeding),
		)
	}

	list = append(list, list2...)

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"cmd/generator/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package data\n")
	g.MustWrite("func ManganeseList() []Manganese {\n")
	g.MustWrite("    return []Manganese {\n")
	for _, d := range list {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "data.", "")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(output)
}

func newManganese(g data.Gender, from, to data.Age, o data.Option, cols []string) data.Manganese {
	return data.Manganese{
		Gender: g,
		From:   from,
		To:     to,
		Option: o,
		AI:     data.NilFloatFromString(cols[0]),
		UL:     data.NilFloatFromString(cols[1]),
	}
}

func newManganeseWithOption(in data.Manganese, o data.Option) data.Manganese {
	d := data.Manganese{
		Gender: in.Gender,
		From:   in.From,
		To:     in.To,
		Option: o,
		AI:     in.AI,
		UL:     in.UL,
	}
	var ai float64
	switch o {
	case data.OptionEarlyPregnancy, data.OptionMidPregnancy, data.OptionLatePregnancy:
		ai = 3.5
	case data.OptionBreastfeeding:
		ai = 3.5
	}
	if d.AI.Valid {
		d.AI.Float = ai
	}
	return d
}

// PDF 376 ページの「マンガンの食事摂取基準」
const manganeseSrc = `0 ～ 5 	（月） 0.01 ─ 0.01 ─
6 ～11（月） 0.5 ─ 0.5 ─
1 ～ 2 	（歳） 1.5 ─ 1.5 ─
3 ～ 5 	（歳） 1.5 ─ 1.5 ─
6 ～ 7 	（歳） 2.0 ─ 2.0 ─
8 ～ 9 	（歳） 2.5 ─ 2.5 ─
10～11（歳） 3.0 ─ 3.0 ─
12～14（歳） 4.0 ─ 4.0 ─
15～17（歳） 4.5 ─ 3.5 ─
18～29（歳） 4.0 11 3.5 11
30～49（歳） 4.0 11 3.5 11
50～64（歳） 4.0 11 3.5 11
65～74（歳） 4.0 11 3.5 11
75 以上（歳） 4.0 11 3.5 11
妊　婦 3.5 ─
授乳婦 3.5 ─`