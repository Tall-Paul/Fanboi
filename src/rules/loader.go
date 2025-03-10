package rules

import (
	"bufio"
	"fanboi/plugin"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadRules(ruleFile string) (bool, RuleManager) {
	pm := plugin.NewPluginManager()
	file, err := os.Open(ruleFile)
	if err != nil {
		log.Fatal(err)
		return false, RuleManager{nil, nil}
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)

		}
	}()
	scanner := bufio.NewScanner(file)
	out := make(map[int]Rule)
	line := 0
	curLine := ""
	ruleInputText := ""
	ruleOutputText := ""

	for scanner.Scan() {
		line++
		curLine = scanner.Text()
		ruleParts := strings.Split(curLine, "THEN")
		if len(ruleParts) == 1 {
			ruleInputText = ""
			ruleOutputText = ruleParts[0]
		} else {
			ruleInputText = ruleParts[0]
			ruleOutputText = ruleParts[1]
		}
		ruleInputObj := RuleInput{true, nil, "", "", 0}
		if ruleInputText != "" {
			inputparts := strings.Split(strings.TrimSpace(ruleInputText), " ")
			if len(inputparts) == 5 {
				pl := pm.GetPlugin(inputparts[1])
				if pl == nil {
					pl, _ = pm.LoadPlugin(inputparts[1])
				}
				if pl != nil {
					val, err := strconv.ParseFloat(inputparts[4], 32)
					if err == nil {
						ruleInputObj = RuleInput{false, pl, inputparts[2], inputparts[3], float32(val)}
					}
				}
			}
		}
		if ruleOutputText != "" {
			outputparts := strings.Split(strings.TrimSpace(ruleOutputText), " ")
			if len(outputparts) == 3 {
				pl := pm.GetPlugin(outputparts[0])
				if pl == nil {
					pl, _ = pm.LoadPlugin(outputparts[0])
				}
				if pl != nil {
					val, err := strconv.ParseFloat(outputparts[2], 32)
					if err == nil {
						ruleOutputObj := RuleOutput{pl, outputparts[1], float32(val)}
						rule := Rule{line, ruleInputObj, ruleOutputObj}
						out[line] = rule
					}
				}
			}
		}

	}
	rm := RuleManager{pm, out}
	return true, rm
}
