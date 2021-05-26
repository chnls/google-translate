/*
  @Date:   2021-05-25
  @File:   translator.go
  @Author: ls
*/
package translate

import (
	"bufio"
	"crypto/tls"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

var something0 interface{}
var pronounce  string
var something1 interface{}
var someBool   bool
var something2 interface{}
var data       []interface{}
var dataSlice = [][]interface{}{{&something0, &pronounce, &something1, &someBool, &something2, &data}}

var reslangTgt string
var someInt int
var reslangSrc string
var params  []interface{}
var resDataSlice = []interface{}{&dataSlice, &reslangTgt, &someInt, &reslangSrc, &params}

var other1 interface{}
var innerlangSrc string
var innerSlice = []interface{}{&other1, &resDataSlice, &innerlangSrc}

var wrb string
var rpcIds string
var jsonStr string
var doNotKnow0 interface{}
var doNotKnow1 interface{}
var doNotKnow2 interface{}
var other string
var responseSlice  = [][]interface{}{{&wrb, &rpcIds, &jsonStr, doNotKnow0, doNotKnow1, doNotKnow2, other}}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func parseRpc(text string, langSrc string, langTgt string) string {
	GOOGLE_TTS_RPC := []interface{}{"MkEWBc"}
	pi := []interface{}{text, langSrc, langTgt, true}
	p2 := []interface{}{nil}
	p1 := []interface{}{&pi, &p2}
	parameter := p1
	escapedParameter, _ := json.Marshal(parameter)
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(GOOGLE_TTS_RPC))
	rpc := [][][]interface{}{{{GOOGLE_TTS_RPC[randomIndex], string(escapedParameter), nil, "generic"}}}
	espacedRpc, _ := json.Marshal(rpc)
	//freq := fmt.Sprintf("f.req=%s&", url.QueryEscape(string(espacedRpc)))
	freq := "f.req=" + url.QueryEscape(string(espacedRpc)) + "&"
	return freq
}

func _translate(langTgt string, langSrc string, text string) ([]interface{}, string,  error) {
	src := LANGUAGES[langSrc]
	target := LANGUAGES[langTgt]
	if src != "" || target == "" {
		langSrc = "auto"
	}
	if len([]rune(text)) >= 5000 {
		return nil, "", errors.New("warning: can only detect less than 5000 characters")
	}
	if len([]rune(text)) == 0 {
		return nil, "", nil
	}

	//tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, Proxy: http.ProxyURL()}
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Timeout: 5 * time.Second, Transport: tr}
	reqUrl := "https://translate.google.cn/_/TranslateWebserverUi/data/batchexecute"
	freq := parseRpc(text, langSrc, langTgt)

	//post
	//response, _ := client.Post(reqUrl, "application/x-www-form-urlencoded", strings.NewReader(reqBody))
	request, err := http.NewRequest("POST", reqUrl, strings.NewReader(freq))
	if err != nil {
		return nil, "", err
	}

	//set header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	request.Header.Set("Referer", "http://translate.google.cn")

	response, err := client.Do(request)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}

	//http response
	//scanner := bufio.NewScanner(strings.NewReader(string(res)))
	//for scanner.Scan() {
	//	lineStr := scanner.Text()
	//	ok := strings.Contains(lineStr, "MkEWBc")
	//	if ok{
	//		lineStr += "]"
	//		err := json.Unmarshal([]byte(lineStr), &slice)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		fmt.Println(slice)
	//		//return lineStr, nil
	//	}
	//}
	//if err := scanner.Err(); err != nil {
	//	//return "", err
	//}
	reader := bufio.NewReader(strings.NewReader(string(res)))
	for {
		// read lines
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		ok := strings.Contains(str, "MkEWBc")
		if ok {
			str += "]"
			err := json.Unmarshal([]byte(str), &responseSlice)
			if err != nil {
				return nil, "", nil
			}
			err = json.Unmarshal([]byte(jsonStr), &innerSlice)
			if err != nil {
				return nil, "", nil
			}
		}
	}

	return data, pronounce, nil
	//return everyData, pronounce, nil
}

func rangeData(args ...[]interface{}) []string{
	var res []string
	if reflect.TypeOf(args).Kind() == reflect.Slice {
		s := reflect.ValueOf(args)
		ele := s.Index(0).Interface()
		ss := reflect.ValueOf(ele)
		for j := 0; j < ss.Len(); j++ {
			val := ss.Index(j).Interface()
			data := reflect.ValueOf(val)
			transData := data.Index(0).Interface().(string)
			res = append(res, transData)
		}
	}
	return res
}

func Translate(langTgt string, langSrc string, text string, pronounce bool) (string, string, error) {

	translate, s, err := _translate(langTgt, langSrc, text)
	if err != nil {
		return "", "", err
	}
	data := rangeData(translate)
	resStr := strings.Join(data, "")
	if !pronounce {
		return resStr, "", nil
	}else {
		return resStr, s, nil
	}

}
