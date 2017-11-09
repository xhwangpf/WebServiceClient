package wsclient

import "fmt"

func CallWebServiceTest(){

	pars := make(map[string]string)

	pars["param1"] = "0"
	pars["param2"] = "1"

	//test webservice uri
	webservicePath := "http://192.168.7.77:8080/sysware/services/TemplateWS"

	nameSpace := "http://tdeide.webservice.integration.datacenter.xxxx.com"

	methodName := "theWebServiceMethodName"

	resultInfo := CallWebService(webservicePath, methodName, pars, nameSpace)

	fmt.Println(resultInfo)
}
