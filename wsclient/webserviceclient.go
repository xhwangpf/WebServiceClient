package wsclient

import (
"bytes"
"fmt"
"io/ioutil"
"net/http"
)

func CallWebService(webserviceUrl string, methodName string, pars map[string]string, nameSpace string) string {

	postStr := CreateSOAPXml(nameSpace, methodName, pars)

	fmt.Println(postStr)

	output := PostWebService(webserviceUrl, postStr)
	fmt.Println("------", output)
	return output
}

func PostWebService(url string, value string) string {
	res, err := http.Post(url, "text/xml; charset=utf-8", bytes.NewBuffer([]byte(value)))
	//这里随便传递了点东西
	if err != nil {
		fmt.Println("post error", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	//取出主体的内容
	if err != nil {
		fmt.Println("read error", err)
	}
	res.Body.Close()
	// fmt.Printf("result----%s", data)
	return ByteToString(data)
}

func ByteToString(res []byte) string { return string(res) }

func CreateSOAPXml(nameSpace string, methodName string, valueStr map[string]string) string {

	/*拼装提交的xml数据*/
	soapBody := "<?xml version=\"1.0\" ?>"
	soapBody += "<S:Envelope xmlns:S=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:tde=\"" + nameSpace + "\">"
	soapBody += "<S:Header/>"
	soapBody += "<S:Body>"
	soapBody += "<tde:" + methodName + " >"

	//以下是具体参数
	for key, value := range valueStr {
		soapBody += "<tde:" + key + ">" + value + "</tde:" + key + ">"
	}

	soapBody += "</tde:" + methodName + ">"
	soapBody += "</S:Envelope></S:Body>"
	return soapBody
}
