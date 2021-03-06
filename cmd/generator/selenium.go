package main

import (
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/data"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

// Selenium は PDF 378 ページの「セレンの食事摂取基準」をパースして Go の構造体を作成します
func Selenium(output string) {
	lines := strings.Split(seleniumSrc, "\n")
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

	list := make([]data.Selenium, 0)
	for i := 0; i <= 13; i++ {
		line := strings.ReplaceAll(lines[i], ",", "")
		cols := strings.Split(line, " ")
		j := len(cols) - 8
		a := ages[i]
		list = append(list,
			newSelenium(data.GenderMale, a[0], a[1], data.OptionNone, cols[j:j+4]),
			newSelenium(data.GenderFemale, a[0], a[1], data.OptionNone, cols[j+4:j+8]),
		)
	}

	list2 := make([]data.Selenium, 0)
	for _, d := range list {
		if d.Gender != data.GenderFemale {
			continue
		}
		list2 = append(list2,
			newSeleniumWithOption(d, data.OptionEarlyPregnancy),
			newSeleniumWithOption(d, data.OptionMidPregnancy),
			newSeleniumWithOption(d, data.OptionLatePregnancy),
			newSeleniumWithOption(d, data.OptionBreastfeeding),
		)
	}

	list = append(list, list2...)

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"cmd/generator/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package data\n")
	g.MustWrite("func SeleniumList() []Selenium {\n")
	g.MustWrite("    return []Selenium {\n")
	for _, d := range list {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "data.", "")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(output)
}

func newSelenium(g data.Gender, from, to data.Age, o data.Option, cols []string) data.Selenium {
	return data.Selenium{
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

func newSeleniumWithOption(in data.Selenium, o data.Option) data.Selenium {
	d := data.Selenium{
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
		ear, rda = 5, 5
	case data.OptionBreastfeeding:
		ear, rda = 15, 20
	}
	if d.EAR.Valid {
		d.EAR.Float += ear
	}
	if d.RDA.Valid {
		d.RDA.Float += rda
	}
	return d
}

const seleniumSrc = `0 ～ 5 	（月） ─ ─ 15 ─ ─ ─ 15 ─
6 ～11（月） ─ ─ 15 ─ ─ ─ 15 ─
1 ～ 2 	（歳） 10 10 ─ 100 10 10 ─ 100
3 ～ 5 	（歳） 10 15 ─ 100 10 10 ─ 100
6 ～ 7 	（歳） 15 15 ─ 150 15 15 ─ 150
8 ～ 9 	（歳） 15 20 ─ 200 15 20 ─ 200
10～11（歳） 20 25 ─ 250 20 25 ─ 250
12～14（歳） 25 30 ─ 350 25 30 ─ 300
15～17（歳） 30 35 ─ 400 20 25 ─ 350
18～29（歳） 25 30 ─ 450 20 25 ─ 350
30～49（歳） 25 30 ─ 450 20 25 ─ 350
50～64（歳） 25 30 ─ 450 20 25 ─ 350
65～74（歳） 25 30 ─ 450 20 25 ─ 350
75 以上（歳） 25 30 ─ 400 20 25 ─ 350
妊婦（付加量） ＋5 ＋5 ─ ─
授乳婦（付加量） ＋15 ＋20 ─ ─`
