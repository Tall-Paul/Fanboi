package main

import (
	"bufio"
	"fanboi/plugin"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "/var/local/emhttp/disks.ini"

type unraiddrivesPlugin struct {
	filePath   string
	driveTemps map[string]float32
}

func (pl unraiddrivesPlugin) SetValue(identifier string, value float32) {
	//noop
}

func (pl unraiddrivesPlugin) GetValue(identifier string) float32 {
	if identifier != "*" {
		return pl.driveTemps[identifier]
	} else {
		//get highest temp
		maxTemp := float32(0.00)
		for _, temp := range pl.driveTemps {
			if temp > maxTemp {
				maxTemp = temp
			}
		}
		return maxTemp
	}
}

func (pl unraiddrivesPlugin) StartHook() {
	//load file
	file, err := os.Open(pl.filePath)
	if err != nil {
		log.Fatal(err)
		fmt.Println("nuhuh")
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	curDrive := ""
	curLine := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curLine = scanner.Text()
		if strings.Contains(curLine, "[") {
			curDrive = strings.Replace(curLine, "[\"", "", 1)
			curDrive = strings.Replace(curDrive, "\"]", "", 1)
		}
		if strings.Contains(curLine, "temp") {
			buffer := strings.Replace(curLine, "temp=\"", "", 1)
			buffer = strings.Replace(buffer, "\"", "", 1)
			Temp, err := strconv.ParseFloat(buffer, 32)
			if err != nil {
				Temp = -1
			}
			pl.driveTemps[curDrive] = float32(Temp)
		}

	}

}

func (pl unraiddrivesPlugin) EndHook() {

}

func InitPlugin(pm *plugin.PluginManager) error {
	this := unraiddrivesPlugin{inputFile, make(map[string]float32)}
	pm.RegisterPlugin("unraiddrives", this)
	return nil
}
