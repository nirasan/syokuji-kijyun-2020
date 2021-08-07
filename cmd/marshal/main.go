package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"github.com/nirasan/syokuji-kijyun-2020/data"
)

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	energyList := data.EnergyList()
	bytes, err := json.Marshal(energyList)
	if err != nil {
		panic("failed to marshal energyList. " + err.Error())
	}
	err = ioutil.WriteFile(dir+"/energy.json", bytes, 0666)
	if err != nil {
		panic("failed to write energy.json. " + err.Error())
	}
}
