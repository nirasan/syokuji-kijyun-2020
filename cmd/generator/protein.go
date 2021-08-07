package main

import (
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/data"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

func Protein(output string) {
	lines := strings.Split(proteinSrc, "\n")
	ages := [][]data.Age{
		{{Month: 0}, {Month: 5}},
		{{Month: 6}, {Month: 8}},
		{{Month: 9}, {Month: 11}},
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

	list := make([]data.Protein, 0)
	for i := 0; i <= 14; i++ {
		cols := strings.Split(lines[i], " ")
		j := len(cols) - 8
		a := ages[i]
		list = append(list,
			newProtein(data.GenderMale, a[0], a[1], data.OptionNone, cols[j:j+4]),
			newProtein(data.GenderFemale, a[0], a[1], data.OptionNone, cols[j+4:j+8]),
		)
	}

	list2 := make([]data.Protein, 0)
	for _, d := range list {
		if d.Gender != data.GenderFemale {
			continue
		}
		if d.EAR.Valid && d.RDA.Valid && d.DG.Min.Valid && d.DG.Max.Valid {
			list2 = append(list2,
				newProteinWithOption(d, data.OptionEarlyPregnancy),
				newProteinWithOption(d, data.OptionMidPregnancy),
				newProteinWithOption(d, data.OptionLatePregnancy),
				newProteinWithOption(d, data.OptionBreastfeeding),
			)
		}
	}

	list = append(list, list2...)

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"cmd/generator/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package data\n")
	g.MustWrite("func ProteinList() []Protein {\n")
	g.MustWrite("    return []Protein {\n")
	for _, d := range list {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "data.", "")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(output)
}

func newProtein(g data.Gender, from, to data.Age, o data.Option, cols []string) data.Protein {
	return data.Protein{
		Gender: g,
		From:   from,
		To:     to,
		Option: o,
		EAR:    data.NilFloatFromString(cols[0]),
		RDA:    data.NilFloatFromString(cols[1]),
		AI:     data.NilFloatFromString(cols[2]),
		DG:     newProteinDG(cols[3]),
	}
}

func newProteinDG(s string) data.NilFloatRange {
	pair := strings.Split(s, "～")
	if len(pair) == 2 {
		return data.NilFloatRange{
			Min: data.NilFloatFromString(pair[0]),
			Max: data.NilFloatFromString(pair[1]),
		}
	}
	return data.NilFloatRange{
		Min: data.NilFloat{},
		Max: data.NilFloat{},
	}
}

func newProteinWithOption(in data.Protein, o data.Option) data.Protein {
	d := data.Protein{
		Gender: in.Gender,
		From:   in.From,
		To:     in.To,
		Option: o,
		EAR:    in.EAR,
		RDA:    in.RDA,
		AI:     in.AI,
		DG:     in.DG,
	}
	var ear, rda, dgmin, dgmax float64
	switch o {
	case data.OptionEarlyPregnancy:
		ear, rda, dgmin, dgmax = 0, 0, 13, 20
	case data.OptionMidPregnancy:
		ear, rda, dgmin, dgmax = 5, 5, 13, 20
	case data.OptionLatePregnancy:
		ear, rda, dgmin, dgmax = 20, 25, 15, 20
	case data.OptionBreastfeeding:
		ear, rda, dgmin, dgmax = 15, 20, 15, 20
	}
	if d.EAR.Valid {
		d.EAR.Float += ear
	}
	if d.RDA.Valid {
		d.RDA.Float += rda
	}
	if d.DG.Min.Valid {
		d.DG.Min.Float = dgmin
	}
	if d.DG.Max.Valid {
		d.DG.Max.Float = dgmax
	}
	return d
}

// PDF 133 ページの「たんぱく質の食事摂取基準」のコピペ
const proteinSrc = `0 ～ 5 （月） ─ ─ 10 ─ ─ ─ 10 ─
6 ～ 8 （月） ─ ─ 15 ─ ─ ─ 15 ─
9 ～11（月） ─ ─ 25 ─ ─ ─ 25 ─
1 ～ 2 （歳） 15 20 ─ 13～20 15 20 ─ 13～20
3 ～ 5 （歳） 20 25 ─ 13～20 20 25 ─ 13～20
6 ～ 7 （歳） 25 30 ─ 13～20 25 30 ─ 13～20
8 ～ 9 （歳） 30 40 ─ 13～20 30 40 ─ 13～20
10～11（歳） 40 45 ─ 13～20 40 50 ─ 13～20
12～14（歳） 50 60 ─ 13～20 45 55 ─ 13～20
15～17（歳） 50 65 ─ 13～20 45 55 ─ 13～20
18～29（歳） 50 65 ─ 13～20 40 50 ─ 13～20
30～49（歳） 50 65 ─ 13～20 40 50 ─ 13～20
50～64（歳） 50 65 ─ 14～20 40 50 ─ 14～20
65～74（歳）2 50 60 ─ 15～20 40 50 ─ 15～20
75 以上（歳）2 50 60 ─ 15～20 40 50 ─ 15～20
妊婦（付加量）　
初期
中期
後期
＋0
＋5
＋20
＋0
＋5
＋25
─ ─ 3
─ 3
─ 4
授乳婦（付加量） ＋15 ＋20 ─ ─ 4`
