package aepapi

import (
    "net/http"
    aepsdkcore "plugin_ctwing/apis/core"
)

// 参数MasterKey: 类型String, 参数不可以为空
//
//	描述:
//
// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func CancelAllCommand(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_device_command_cancel/cancelAllCommand"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20230419143717"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}
