package infra

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
	"github.com/tidwall/sjson"
	"io/ioutil"
	"log"
	"os"
)

type JsonConfig struct {
	result   gjson.Result
	filePath string
}

var Config JsonConfig

func (c *JsonConfig) Load(filePath string) {
	c.filePath = filePath

	json, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(color.RedString("Config file %s can't be read", filePath))
	}

	c.result = gjson.ParseBytes(json)
}

func (c *JsonConfig) Dump(filePath string) {
	pretty.DefaultOptions.Indent = "    "

	err := ioutil.WriteFile(filePath, pretty.Pretty([]byte(c.result.Raw)), os.FileMode(0644))
	if err != nil {
		log.Fatalln(color.RedString("Config file %s can't be written", filePath))
	}
}

func (c *JsonConfig) Set(path string, value interface{}) {
	c.result.Raw, _ = sjson.Set(c.result.Raw, path, value)

	go c.Dump(c.filePath)
}

func (c *JsonConfig) Get(path string, args ...interface{}) gjson.Result {
	return c.result.Get(fmt.Sprintf(path, args...))
}

func (c *JsonConfig) Exists(path string, args ...interface{}) bool {
	return c.result.Get(fmt.Sprintf(path, args...)).Exists()
}

func (c *JsonConfig) Delete(path string, args ...interface{}) {
	c.result.Raw, _ = sjson.Delete(c.result.Raw, fmt.Sprintf(path, args...))
}

// return $var.Result or Result
func (c *JsonConfig) Parse(result gjson.Result) gjson.Result {
	if len(result.String()) > 0 && result.String()[:1] == "$" {
		return c.Get(result.String()[1:])
	}
	return result
}

func (c *JsonConfig) Map(result gjson.Result) interface{} {
	if result.IsObject() {
		var m = map[string]interface{}{}
		result.ForEach(func(key, value gjson.Result) bool {
			m[key.String()] = c.Map(value)
			return true
		})
		return m
	}

	return c.Parse(result).Value()
}
