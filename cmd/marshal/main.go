package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/nirasan/syokuji-kijyun-2020/data"
)

func main() {
	flag.Parse()
	dir := flag.Arg(0)

	targets := []struct {
		Filename string
		Data     interface{}
	}{
		{"energy.json", data.EnergyList()},
		{"protein.json", data.ProteinList()},
		{"lipid.json", data.LipidList()},
		{"saturated_fatty_acids.json", data.SaturatedFattyAcidsList()},
		{"omega3_fatty_acids.json", data.Omega3FattyAcidsList()},
		{"omega6_fatty_acids.json", data.Omega6FattyAcidsList()},
		{"carbohydrates.json", data.CarbohydratesList()},
		{"dietary_fiber.json", data.DietaryFiberList()},
		{"vitamin_a.json", data.VitaminAList()},
		{"vitamin_d.json", data.VitaminDList()},
		{"vitamin_e.json", data.VitaminEList()},
		{"vitamin_k.json", data.VitaminKList()},
		{"vitamin_b1.json", data.VitaminB1List()},
		{"vitamin_b2.json", data.VitaminB2List()},
		{"niacin.json", data.NiacinList()},
		{"vitamin_b6.json", data.VitaminB6List()},
		{"vitamin_b12.json", data.VitaminB12List()},
		{"folic_acid.json", data.FolicAcidList()},
		{"pantothenic_acid.json", data.PantothenicAcidList()},
		{"biotin.json", data.BiotinList()},
		{"vitamin_c.json", data.VitaminCList()},
		{"sodium.json", data.SodiumList()},
		{"potassium.json", data.PotassiumList()},
		{"calcium.json", data.CalciumList()},
		{"magnesium.json", data.MagnesiumList()},
		{"phosphorus.json", data.PhosphorusList()},
		{"iron.json", data.IronList()},
		{"zinc.json", data.ZincList()},
		{"copper.json", data.CopperList()},
		{"manganese.json", data.ManganeseList()},
		{"iodine.json", data.IodineList()},
		{"selenium.json", data.SeleniumList()},
		{"chromium.json", data.ChromiumList()},
		{"molybdenum.json", data.MolybdenumList()},
	}
	for _, t := range targets {
		bytes, err := json.Marshal(t.Data)
		if err != nil {
			panic(fmt.Sprintf("failed to marshal json. file:%s, error: %s", t.Filename, err.Error()))
		}
		err = ioutil.WriteFile(dir+"/"+t.Filename, bytes, 0666)
		if err != nil {
			panic(fmt.Sprintf("failed to write. file:%s, error: %s", t.Filename, err.Error()))
		}
	}
}
