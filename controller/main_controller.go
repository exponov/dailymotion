package controller

import (
	"github.com/astaxie/beego"
	"dailymotion/models"
	_"log"
)

type MainController struct {
	beego.Controller
}


// @router / [get]
func (m *MainController) Index() {
	m.TplName = "index.tpl"
	return
}

// @router /request [get]
func (m *MainController) RequestAd() {
	var rule models.ProxyRule
	var url string
	var err error
	
	ruleType := m.GetString("type")
	
	defer func () {
		m.ServeJSON(true)
	} ()
	
	if rule, err = models.NewRule(ruleType); err != nil {
		m.Data["json"] = map[string]string{"error": err.Error()}
		return
	}
	
	values := m.Input()
	if url, err = rule.Execute(&values); err != nil {
		m.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		m.Data["json"] = map[string]string{"url": url}	
	}
	
	return
}
