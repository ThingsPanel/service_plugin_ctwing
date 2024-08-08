package model

type DeviceItem struct {
	DeviceNumber string `json:"device_number"`
	DeviceName   string `json:"device_name"`
	Description  string `json:"description"`
}

// 事件/命令
type EventInfo struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

type CtwingDeviceListResp struct {
	Code   int              `json:"code"`
	Msg    string           `json:"msg"`
	Result CtwingDeviceList `json:"result"`
}

type CtwingDeviceList struct {
	PageNum int                `json:"pageNum"`
	Total   int                `json:"total"`
	List    []CtwingDeviceItem `json:"list"`
}

type CtwingDeviceItem struct {
	DeviceId   string `json:"deviceId"`
	DeviceName string `json:"deviceName"`
}
