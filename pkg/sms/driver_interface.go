// Package sms
// descr
// author fm
// date 2022/11/17 15:12
package sms

type Driver interface {
	// Send 发送短信
	Send(phone string, message Message, config map[string]string) bool
}
