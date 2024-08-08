package aepapi

import (
	"net/http"
	aepsdkcore "plugin_ctwing/apis/core"
)

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func DeleteArchivesInfo(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/device_archives/deleteArchivesInfo"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20231117042743"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func GetArchivesAttribute(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/device_archives/getArchivesAttr"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20231117042748"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func UpdateArchivesInfo(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/device_archives/updateArchivesInfo"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20231117042738"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func AddArchivesInfo(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/device_archives/addArchivesInfo"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20231215034317"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// 参数searchValueType: 类型long, 参数可以为空
//
//	描述:1：按设备id查询，2：按设备类型查询
//
// 参数searchValue: 类型String, 参数可以为空
//
//	描述:
//
// 参数pageNow: 类型long, 参数不可以为空
//
//	描述:
//
// 参数pageSize: 类型long, 参数不可以为空
//
//	描述:
func GetArchivesInfo(appKey string, appSecret string, pageNow string, pageSize string, searchValueType string, searchValue string) (*http.Response, error) {
	path := "/device_archives/getArchivesInfo"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["searchValueType"] = searchValueType
	param["searchValue"] = searchValue
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20231215034340"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

// 参数body: 类型json, 参数不可以为空
//
//	描述:body,具体参考平台api说明
func GetDeviceType(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/device_archives/getDeviceType"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20231215034248"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}
