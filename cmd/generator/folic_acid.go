package main

import (
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/data"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

func FolicAcid(output string) {
	lines := strings.Split(folicAcidSrc, "\n")
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

	list := make([]data.FolicAcid, 0)
	for i := 0; i <= 13; i++ {
		line := strings.ReplaceAll(lines[i], ",", "")
		cols := strings.Split(line, " ")
		j := len(cols) - 8
		a := ages[i]
		list = append(list,
			newFolicAcid(data.GenderMale, a[0], a[1], data.OptionNone, cols[j:j+4]),
			newFolicAcid(data.GenderFemale, a[0], a[1], data.OptionNone, cols[j+4:j+8]),
		)
	}

	list2 := make([]data.FolicAcid, 0)
	for _, d := range list {
		if d.Gender != data.GenderFemale {
			continue
		}
		if d.EAR.Valid && d.RDA.Valid {
			list2 = append(list2,
				newFolicAcidWithOption(d, data.OptionEarlyPregnancy),
				newFolicAcidWithOption(d, data.OptionMidPregnancy),
				newFolicAcidWithOption(d, data.OptionLatePregnancy),
				newFolicAcidWithOption(d, data.OptionBreastfeeding),
			)
		}
	}

	list = append(list, list2...)

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"cmd/generator/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package data\n")
	g.MustWrite("func FolicAcidList() []FolicAcid {\n")
	g.MustWrite("    return []FolicAcid {\n")
	for _, d := range list {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "data.", "")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(output)
}

func newFolicAcid(g data.Gender, from, to data.Age, o data.Option, cols []string) data.FolicAcid {
	return data.FolicAcid{
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

func newFolicAcidWithOption(in data.FolicAcid, o data.Option) data.FolicAcid {
	d := data.FolicAcid{
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
		ear, rda = 200, 240
	case data.OptionBreastfeeding:
		ear, rda = 80, 100
	}
	if d.EAR.Valid {
		d.EAR.Float += ear
	}
	if d.RDA.Valid {
		d.RDA.Float += rda
	}
	return d
}

// PDF 269 ページの「葉酸の食事摂取基準」のコピペ
const folicAcidSrc = `0 ～ 5 （月） ─ ─ 40 ─ ─ ─ 40 ─
6 ～11（月） ─ ─ 60 ─ ─ ─ 60 ─
1 ～ 2 （歳） 80 90 ─ 200 90 90 ─ 200
3 ～ 5 （歳） 90 110 ─ 300 90 110 ─ 300
6 ～ 7 （歳） 110 140 ─ 400 110 140 ─ 400
8 ～ 9 （歳） 130 160 ─ 500 130 160 ─ 500
10～11（歳） 160 190 ─ 700 160 190 ─ 700
12～14（歳） 200 240 ─ 900 200 240 ─ 900
15～17（歳） 220 240 ─ 900 200 240 ─ 900
18～29（歳） 200 240 ─ 900 200 240 ─ 900
30～49（歳） 200 240 ─ 1,000 200 240 ─ 1,000
50～64（歳） 200 240 ─ 1,000 200 240 ─ 1,000
65～74（歳） 200 240 ─ 900 200 240 ─ 900
75 以上（歳） 200 240 ─ 900 200 240 ─ 900
妊婦（付加量）3，4 ＋200 ＋240 ─ ─
授乳婦（付加量） ＋80 ＋100 ─ ─`
