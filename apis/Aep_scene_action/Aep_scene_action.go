package aepapi

import (
	"net/http"
	aepsdkcore "plugin_ctwing/apis/core"
)

// 参数productId: 类型long, 参数可以为空
//
//	描述:可通过产品ID搜索
//
// 参数pageNow: 类型long, 参数可以为空
//
//	描述:当前页
//
// 参数pageSize: 类型long, 参数可以为空
//
//	描述:每页记录数
func QueryActionList(appKey string, appSecret string, productId string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_scene_action/getActionList"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20240126035641"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

// 参数pageNow: 类型long, 参数可以为空
//
//	描述:
//
// 参数pageSize: 类型long, 参数可以为空
//
//	描述:
func QuerySceneList(appKey string, appSecret string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_scene_action/getSceneList"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20240126035701"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func ActionExecute(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_scene_action/action/execute"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20240126035654"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func SceneExecute(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_scene_action/scene/execute"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20240126035634"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}
