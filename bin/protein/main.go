package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/nirasan/syokuji-kijyun-2020/pkg"
	"github.com/nirasan/syokuji-kijyun-2020/tools"
)

func main() {
	flag.Parse()
	pkgName := flag.Arg(0)
	outputFile := flag.Arg(1)

	data := make([]pkg.ProteinDatum, 0)

	lines := strings.Split(src, "\n")
	for i := 0; i <= 14; i++ {
		cols := strings.Split(lines[i], " ")
		j := len(cols) - 8
		p1 := newProtein(pkg.GenderMale, i, cols[j:j+4])
		p2 := newProtein(pkg.GenderFemale, i, cols[j+4:j+8])
		data = append(data, p1, p2)
	}

	g := tools.NewGenerator()
	g.MustWrite("// Code generated by \"bin/protein/main.go\"; DO NOT EDIT.\n")
	g.MustWrite("package %s\n", pkgName)
	g.MustWrite("func (p *Protein) Data() []ProteinDatum {\n")
	g.MustWrite("    return []ProteinDatum {\n")
	for _, d := range data {
		s := fmt.Sprintf("%#v,\n", d)
		s = strings.ReplaceAll(s, "pkg.ProteinDatum", "ProteinDatum")
		s = strings.ReplaceAll(s, "pkg.NullFloat64", "NullFloat64")
		g.MustWrite(s)
	}
	g.MustWrite("    }\n")
	g.MustWrite("}\n")

	g.Generate(outputFile)
}

func newProtein(gender, age int, cols []string) pkg.ProteinDatum {
	p := pkg.ProteinDatum{
		Gender: gender,
		Age:    age,
		EAR:    pkg.NullFloat64FromString(cols[0]),
		RDA:    pkg.NullFloat64FromString(cols[1]),
		AI:     pkg.NullFloat64FromString(cols[2]),
	}
	DGs := strings.Split(cols[3], "～")
	if len(DGs) == 2 {
		p.DGMin = pkg.NullFloat64FromString(DGs[0])
		p.DGMax = pkg.NullFloat64FromString(DGs[1])
	} else {
		p.DGMin = pkg.NewNullFloat64(nil)
		p.DGMax = pkg.NewNullFloat64(nil)
	}
	return p
}

// PDF 133 ページの「たんぱく質の食事摂取基準」のコピペ
const src = `0 ～ 5 （月） ─ ─ 10 ─ ─ ─ 10 ─
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