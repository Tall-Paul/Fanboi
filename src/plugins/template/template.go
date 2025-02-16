package main

import (
	"fanboi/plugin"
	"fmt"
	"strconv"
	"strings"
)

/*
This struct represents the plugin and implements the plugin interface
You can add any data fields you want here.  For a plugin that sets fan speeds via a file (like liquidctl) you probably want to stash the values
each run of fanboi, then write the file at the end, so we have a template as a string and the values as a map
*/
type templatePlugin struct {
	template string
	values   map[string]float32
}

/*
This gets called at program startup, when the plugins are loaded.  We initialise the struct and pass it to the pluginmanager
*/
func InitPlugin(pm *plugin.PluginManager) error {
	values := make(map[string]float32)
	this := templatePlugin{"fan1: {fan1_speed}, fan2: {fan2_speed}", values}
	pm.RegisterPlugin("template", this)
	return nil
}

/*
This function is part of the plugin interface and must be implemented in your struct
Since we don't want to write a value every time, just a big file at the end here we just store the value
*/
func (tem templatePlugin) SetValue(identifier string, value float32) {
	tem.values[identifier] = value
}

/*
This function is part of the plugin interface and must be implemented in your struct
This isn't very useful for a template plugin, you'd use this more for temperature sensing plugins
*/
func (tem templatePlugin) GetValue(identifier string) float32 {
	return tem.values[identifier]
}

/*
This function is part of the plugin interface and must be implemented in your struct
This runs once at the start of every 'fun' of fanboi.  Here you would for example, read the file that you're going to write values into
*/
func (tem templatePlugin) StartHook() {

}

/*
This function is part of the plugin interface and must be implemented in your struct
this runs at the end of a run of fanboi
*/
func (tem templatePlugin) EndHook() {
	out := tem.template
	for key, value := range tem.values {
		textValue := strconv.FormatFloat(float64(value), 'f', -1, 32)
		out = strings.Replace(out, "{"+key+"}", textValue, -1)
	}
	fmt.Println(out)
}
