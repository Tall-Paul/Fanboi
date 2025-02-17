package main

import (
	"bufio"
	"fanboi/plugin"
	"log"
	"os"
	"strconv"
	"strings"
)

const ctrlfile = "/mnt/user/docker/laac/config.yaml"

type liquidctlPlugin struct {
	template map[int]string
	values   map[string]float32
}

func InitPlugin(pm *plugin.PluginManager) error {
	values := make(map[string]float32)
	template := make(map[int]string)
	this := liquidctlPlugin{template, values}
	pm.RegisterPlugin("liquidctl", this)
	return nil
}

func (tem liquidctlPlugin) SetValue(identifier string, value float32) {
	tem.values[identifier] = value
}

func (tem liquidctlPlugin) GetValue(identifier string) float32 {
	return tem.values[identifier]
}

func (tem liquidctlPlugin) StartHook() {
	file, err := os.Open(ctrlfile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	lineNo := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNo++
		tem.template[lineNo] = scanner.Text()
	}
}

func (tem liquidctlPlugin) EndHook() {

	for i := 1; i <= len(tem.template); i++ {
		for key, value := range tem.values {
			if strings.Contains(tem.template[i], key) {
				valtext := strconv.FormatFloat(float64(value), 'f', -1, 32)
				fandata := strings.Split(tem.template[i], ":")
				fandata[1] = " '" + valtext + "'"
				tem.template[i] = strings.Join(fandata, ":")
			}
		}
	}
	file, err := os.Open(ctrlfile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	write := bufio.NewWriter(file)
	for i := 1; i <= len(tem.template); i++ {
		write.WriteString(tem.template[i])
	}
	write.Flush()
}
