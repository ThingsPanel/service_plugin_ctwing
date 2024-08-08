package aepapi

import (
	"net/http"
	aepsdkcore "plugin_ctwing/apis/core"
)

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func SaasCreateRule(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/api/v2/rule/sass/createRule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20200111000503"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数ruleId: 类型String, 参数可以为空
//
//	描述:
//
// 参数productId: 类型String, 参数不可以为空
//
//	描述:
//
// 参数pageNow: 类型long, 参数可以为空
//
//	描述:
//
// 参数pageSize: 类型long, 参数可以为空
//
//	描述:
func SaasQueryRule(appKey string, appSecret string, productId string, ruleId string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_rule_engine/api/v2/rule/sass/queryRule"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["ruleId"] = ruleId
	param["productId"] = productId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20200111000633"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func SaasUpdateRule(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/api/v2/rule/sass/updateRule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20200111000540"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func SaasDeleteRuleEngine(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/api/v2/rule/sass/deleteRule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20200111000611"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func CreateRule(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/createRule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210327062633"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func UpdateRule(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/updateRule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210327062642"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func DeleteRule(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/deleteRule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210327062626"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数ruleId: 类型String, 参数不可以为空
//
//	描述:
//
// 参数productId: 类型String, 参数可以为空
//
//	描述:
//
// 参数pageNow: 类型long, 参数可以为空
//
//	描述:
//
// 参数pageSize: 类型long, 参数可以为空
//
//	描述:
func GetRules(appKey string, appSecret string, ruleId string, productId string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/getRules"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["ruleId"] = ruleId
	param["productId"] = productId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20210327062616"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func GetRuleRunStatus(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/getRuleRunningStatus"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210327062610"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func UpdateRuleRunStatus(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/modifyRuleRunningStatus"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210327062603"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func CreateForward(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/addForward"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210327062556"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func UpdateForward(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/updateForward"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210327062549"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func DeleteForward(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/deleteForward"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210327062539"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数ruleId: 类型String, 参数不可以为空
//
//	描述:
//
// 参数productId: 类型String, 参数可以为空
//
//	描述:
//
// 参数pageNow: 类型long, 参数可以为空
//
//	描述:
//
// 参数pageSize: 类型long, 参数可以为空
//
//	描述:
func GetForwards(appKey string, appSecret string, ruleId string, productId string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/getForwards"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["ruleId"] = ruleId
	param["productId"] = productId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20210327062531"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

// 参数ruleId: 类型String, 参数不可以为空
//
//	描述:
//
// 参数pageNow: 类型long, 参数可以为空
//
//	描述:
//
// 参数pageSize: 类型long, 参数可以为空
//
//	描述:
func GetWarns(appKey string, appSecret string, ruleId string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/getWarns"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["ruleId"] = ruleId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20210423162903"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func DeleteWarn(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/deleteWarn"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210423162859"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func UpdateWarn(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/updateWarn"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210423162906"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func CreateWarn(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/addWarn"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210423162909"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func CreateAction(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/addAction"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210423162837"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func UpdateAction(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/updateAction"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210423162842"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func DeleteAction(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/deleteAct"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20210423162848"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数ruleId: 类型String, 参数不可以为空
//
//	描述:
//
// 参数pageNow: 类型long, 参数可以为空
//
//	描述:
//
// 参数pageSize: 类型long, 参数可以为空
//
//	描述:
func GetActions(appKey string, appSecret string, ruleId string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_rule_engine/v3/rule/getActions"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["ruleId"] = ruleId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20211028100156"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
