package models

import (
		"fmt"
		"log"
		"net/url"
		)


type ProxyRule interface {
	Execute  (u *url.Values) (url string, err error)
}

var ruleMap = map[string]func() ProxyRule{
					"geolocation": func() ProxyRule {return  &geolocationRule{rule:rule{name:"geolocation"}} },
		}

func NewRule(typ string ) (r ProxyRule, err error) {

  var create func()(ProxyRule)
  var ok bool
  	
  if len(typ) == 0 {	
  	err = fmt.Errorf("Type is blank")
  	return
  }
  
  if create, ok = ruleMap[typ]; !ok {
  	err = fmt.Errorf("%s could not be found", typ)
  	return
  }
  r = create()
  return
}

type rule struct {
	name 	string
}

func (r *geolocationRule) Execute (input *url.Values) (url string, err error) {
	if  err = r.ingest(input); err != nil {
		err = fmt.Errorf("Error while ingesting url for rule:'%s', error:'%s'", r.name, err.Error() )		
		return
	}
	
	if url, err =  r.getUrl(); err != nil {
		err = fmt.Errorf("Error while attempting to get url for rule:'%s', error:'%s'", r.name, err.Error())		
		return
	}

	return
}

type geolocationRule struct {
	rule
	ip 		string
}

func (r *geolocationRule) ingest (input *url.Values) (err error) {
	var ip string

	if ip = input.Get("ip"); len(ip) == 0 {
		err = fmt.Errorf("param Ip is missing")
		return
	}
	//TODO add more validation for IP address
	r.ip = ip
	return
}

func (r *geolocationRule) getUrl() (url string, err error) {
	log.Println("looking for location of ip address %s ",r.ip)
	url = "http://www.dailymotion.com/video/x4w613j_wasteland-3-gameplay-trailer-a-frosty-reception-2016_videogames"
	return
} 






