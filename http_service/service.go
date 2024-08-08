package httpservice

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	tpprotocolsdkgo "github.com/ThingsPanel/tp-protocol-sdk-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	aepapi "plugin_ctwing/apis/Aep_device_management"
	httpclient "plugin_ctwing/http_client"
	"plugin_ctwing/model"
)

var HttpClient *tpprotocolsdkgo.Client

func Init() {
	go start()
}

func start() {
	var handler tpprotocolsdkgo.Handler = tpprotocolsdkgo.Handler{
		OnDisconnectDevice: OnDisconnectDevice,
		OnGetForm:          OnGetForm,
		GetDeviceList:      OnGetDeviceList,
		OnNotifyEvent:      OnNotifyEvent,
	}
	addr := viper.GetString("http_server.address")
	log.Println("http服务启动：", addr)
	err := handler.ListenAndServe(addr)
	if err != nil {
		log.Println("ListenAndServe() failed, err: ", err)
		return
	}
}

// OnGetForm 获取协议插件的json表单
func OnGetForm(w http.ResponseWriter, r *http.Request) {
	logrus.Info("OnGetForm")
	r.ParseForm() //解析参数，默认是不会解析的
	logrus.Info("【收到api请求】path", r.URL.Path)
	logrus.Info("query", r.URL.Query())

	// device_type := r.URL.Query()["device_type"][0]
	form_type := r.URL.Query()["form_type"][0]
	// service_identifier := r.URL.Query()["protocol_type"][0]
	// 根据需要对服务标识符进行验证，可不验证
	// if service_identifier != "xxxx" {
	// 	RspError(w, fmt.Errorf("not support protocol type: %s", service_identifier))
	// 	return
	// }
	//CFG配置表单 VCR凭证表单 SVCR服务凭证表单
	switch form_type {
	case "VCR":
		RspSuccess(w, nil)
	case "SVCR":
		//服务凭证类型表单
		RspSuccess(w, readFormConfigByPath("./form_ctwing.json"))
	default:
		RspError(w, errors.New("not support form type: "+form_type))
	}
}

func OnDisconnectDevice(w http.ResponseWriter, r *http.Request) {
	logrus.Info("OnDisconnectDevice")
	r.ParseForm() //解析参数，默认是不会解析的
	logrus.Info("【收到api请求】path", r.URL.Path)
	logrus.Info("query", r.URL.Query())
	// 断开设备

	//RspSuccess(w, nil)
}

// ./form_config.json
func readFormConfigByPath(path string) interface{} {
	filePtr, err := os.Open(path)
	if err != nil {
		logrus.Warn("文件打开失败...", err.Error())
		return nil
	}
	defer filePtr.Close()
	var info interface{}
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		logrus.Warn("解码失败", err.Error())
		return info
	} else {
		logrus.Info("读取文件[form_config.json]成功...")
		return info
	}
}

type CwtingVoucher struct {
	ProductId string `json:"productId"`
	AppKey    string `json:"appKey"`
	AppSecret string `json:"AppSecret"`
	MasterKey string `json:"MasterKey"`
}

func OnGetDeviceList(w http.ResponseWriter, r *http.Request) {
	logrus.Info("OnGetDeviceList")
	//r.ParseForm() //解析参数，默认是不会解析的
	logrus.Info("【收到api请求】path", r.URL.Path)
	logrus.Info("query", r.FormValue("voucher"))
	var voucher CwtingVoucher
	err := json.Unmarshal([]byte(r.FormValue("voucher")), &voucher)
	if err != nil {
		return
	}
	page := r.FormValue("page")
	pageSize := r.FormValue("page_size")
	data := make(map[string]interface{})
	if page == "" || page == "0" {
		page = "1"
	}
	if pageSize == "" || pageSize == "0" {
		pageSize = "10"
	}
	resp, err := aepapi.QueryDeviceList(voucher.AppKey, voucher.AppSecret, voucher.MasterKey, voucher.ProductId, "", page, pageSize)
	//total, list, err := cache.GetDeviceList(r.Context(), voucher.ProductId, page, pageSize)
	if err != nil {
		RspError(w, err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		RspError(w, err)
		return
	}
	logrus.Debug("返回结果:", string(body))
	var deviceListResp model.CtwingDeviceListResp
	err = json.Unmarshal(body, &deviceListResp)
	if err != nil {
		RspError(w, err)
		return
	}
	var list []model.DeviceItem
	for _, v := range deviceListResp.Result.List {
		list = append(list, model.DeviceItem{
			DeviceNumber: fmt.Sprintf(viper.GetString("ctwing.device_number_key"), voucher.ProductId, v.DeviceId),
			DeviceName:   v.DeviceName,
			Description:  "",
		})
	}
	data["list"] = list
	data["total"] = deviceListResp.Result.Total

	// 获取设备列表
	RspSuccess(w, data)
}

// GetMD5Hash 计算给定字符串的MD5哈希值
func GetMD5Hash(text string) string {
	// 创建一个MD5哈希对象
	hasher := md5.New()

	// 将字符串写入哈希对象
	hasher.Write([]byte(text))

	// 计算哈希值并转换为16进制字符串
	return hex.EncodeToString(hasher.Sum(nil))
}

func OnNotifyEvent(w http.ResponseWriter, r *http.Request) {
	logrus.Info("OnNotifyEvent")
	r.ParseForm() //解析参数，默认是不会解析的
	logrus.Info("【收到api请求】path", r.URL.Path)
	logrus.Info("query", r.Body)
	// 读取body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Warn("读取body失败", err.Error())
		return
	}
	logrus.Info("body", string(body))
	type NotifyEvent struct {
		MessageType string `json:"message_type"`
		Message     string `json:"message"`
	}
	// 解析到NotifyEvent
	var notifyEvent NotifyEvent
	err = json.Unmarshal(body, &notifyEvent)
	if err != nil {
		logrus.Warn("解析body失败", err.Error())
		RspError(w, err)
		return
	}
	logrus.Info("notifyEvent", notifyEvent)
	if notifyEvent.MessageType == "1" {
		type NotifyEventData struct {
			ServiceAccessID string `json:"service_access_id"`
		}
		var notifyEventData NotifyEventData
		err = json.Unmarshal([]byte(notifyEvent.Message), &notifyEventData)
		if err != nil {
			logrus.Warn("解析message失败", err.Error())
			RspError(w, err)
			return
		}
		OnNotifyProperty(notifyEventData.ServiceAccessID)
		return
	} else {
		logrus.Warn("不支持的message_type", notifyEvent.MessageType)
	}
	RspSuccess(w, nil)
	// 处理事件通知
	//RspSuccess(w, nil)
}

// 配置变更发送属性
func OnNotifyProperty(serviceAccessID string) {
	rspData, err := httpclient.GetServiceAccessPoint(serviceAccessID)
	if err != nil {
		logrus.Warn("获取服务接入点失败", err.Error())
		return
	}
	if rspData.Code != 200 {
		logrus.Warn("获取服务接入点失败", rspData.Message)
		return
	}
	logrus.Info("获取服务接入点成功", rspData.Data)
	// 处理接入点
}
