package main

import (
	"encoding/json"
	"github.com/r3inbowari/zserver"
	"io/ioutil"
	"net/http"
	"pixiapi/src"
)

func spriteallgetting(writer http.ResponseWriter, request *http.Request) {
	var m Messages
	splist, err := src.GetALLSpritefromDB()

	if err != nil {
		m.Data = "null"
		m.Msg = err.Error()
		ms, _ := json.Marshal(m)
		writer.Write(ms)
		return
	}
	m.Data = splist
	m.Msg = "success get all spritedata"
	ms, _ := json.Marshal(m)
	writer.Write(ms)

}
func postsprite(writer http.ResponseWriter, request *http.Request) {
	var m Messages
	var sp src.SpriteEnt
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		m.Data = "null"
		m.Msg = err.Error()
		ms, _ := json.Marshal(m)
		writer.Write(ms)
		return
	}
	err = json.Unmarshal(body, &sp)
	if err != nil {
		m.Data = "null"
		m.Msg = err.Error()
		ms, _ := json.Marshal(m)
		writer.Write(ms)
		return
	}
	_, err = src.InsertSpritetoDB(sp)
	if err != nil {
		m.Data = "null"
		m.Msg = err.Error()
		ms, _ := json.Marshal(m)
		writer.Write(ms)
		return
	}
	m.Data = sp
	m.Msg = "success post spritedata"
	ms, _ := json.Marshal(m)
	writer.Write(ms)
}
func spritebynamegetting(writer http.ResponseWriter, request *http.Request) {
	var sp src.SpriteEnt
	var m Messages
	var result = make(map[string]string)
	keys := request.URL.Query()
	for k, v := range keys {
		result[k] = v[0]
	}

	name := result["name"]
	sp, err := src.GetSpritebyName(name)
	if err != nil {
		m.Data = "null"
		m.Msg = err.Error()
		ms, _ := json.Marshal(m)
		writer.Write(ms)
		return
	}
	m.Data = sp
	m.Msg = "success get spritedata"
	ms, _ := json.Marshal(m)
	writer.Write(ms)

}
func updatasprite(writer http.ResponseWriter, request *http.Request) {
	var sp src.SpriteEnt
	var m Messages
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		m.Data = "null"
		m.Msg = err.Error()
		ms, _ := json.Marshal(m)
		writer.Write(ms)
		return
	}
	err = json.Unmarshal(body, &sp)
	if err != nil {
		m.Data = "null"
		m.Msg = err.Error()
		ms, _ := json.Marshal(m)
		writer.Write(ms)
		return
	}
	_, err = src.UpdataSprite(sp)
	if err != nil {
		m.Data = "null"
		m.Msg = err.Error()
		ms, _ := json.Marshal(m)
		writer.Write(ms)
		return
	}
	m.Data = sp
	m.Msg = "success updata spritedata"
	ms, _ := json.Marshal(m)
	writer.Write(ms)
}

type Messages struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func main() {

	d := zserver.DefaultServer(zserver.Options{})

	d.Map("/spriteGetALL", spriteallgetting, "GET")
	d.Map("/spritePost", postsprite, "POST")
	d.Map("/spriteGetByname", spritebynamegetting, "GET")
	d.Map("/spriteUpdata", updatasprite, "POST")
	d.Start()

}
