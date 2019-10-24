package defs

import "io"

// CallData data struct used in api-request-response
type CallData struct {
	// CallID 识别该消息的唯一编码
	CallID string
	// Method 请求使用的方法
	Method string
	// EndPoint 请求该Rest资源的定位符url
	EndPoint string
	// Type 该请求属于Rest API中的哪个级别
	Type string
	// Body POST请求的消息体
	Body string
	// Data 请求返回的数据
	Data io.Writer
	// PlatForm 平台：binance,bithumb,okex
	PlatForm string
}
