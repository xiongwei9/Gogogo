package registerCenter

import "time"

// 服务描述信息
type ServiceDescInfo struct {
	ServiceName  string        // 服务名称
	Host         string        // ip地址
	Port         int           // 端口
	IntervalTime time.Duration // 心跳间隔 秒

}

// 服务注册和下线的接口
type RegisterI interface {

	// 服务注册
	Register(serviceInfo ServiceDescInfo) error

	// 服务下线
	UnRegister(serviceInfo ServiceDescInfo) error
}
