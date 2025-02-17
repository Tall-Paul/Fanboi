package main

import (
	"fanboi/rules"
)

func main() {

	rulesFile := "./rules.fnb"
	ok, rm := rules.LoadRules(rulesFile)
	if ok {
		rm.RunRules()
	}

}
