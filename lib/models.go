package lib

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
)

type App interface {
	Run() bool
	GetDetail(propName string) string
}

func RunSingle(app App) {
	app.Run()
}

type AppInfo struct {
	Id     int
	Name   string
	Source string
}

type Script struct {
	Id     int
	Name   string
	Source string
}

func (appInfo AppInfo) Run() bool {
	log.Println("Running app :", appInfo.Id, appInfo.Name)
	return true
}

func (appInfo AppInfo) GetDetail(propName string) string {
	v := reflect.ValueOf(appInfo)
	field := v.FieldByName(propName)
	if !field.IsValid() {
		log.Println("no such field", propName)
		return "default"
	}
	return fmt.Sprintf("%v", field.Interface())
}

func GetPropertyValue(obj interface{}, propertyName string) string {
	v := reflect.ValueOf(obj)
	field := v.FieldByName(propertyName)
	if !field.IsValid() {
		log.Println("no such field", propertyName)
		return "default"
	}
	return fmt.Sprintf("%v", field.Interface())
}

func (script Script) Run() bool {
	log.Println("Running app :", script.Source, script.Name)

	cmd := exec.Command("code", script.Source)
	err := cmd.Run()
	if err != nil {
		return false
	} else {
		return true
	}
}

func (appInfo Script) GetDetail(propName string) string {
	v := reflect.ValueOf(appInfo)
	field := v.FieldByName(propName)
	if !field.IsValid() {
		log.Println("no such field", propName)
		return "default"
	}
	return fmt.Sprintf("%v", field.Interface())
}

func ShortenSentenceToFixChar(sentence string, charLen int) string {
	length := len(sentence)
	if length > charLen {
		return "..." + sentence[length-30:]
	}
	return sentence
}
