// Package mail
// descr 邮件驱动
// author fm
// date 2022/11/18 14:17
package mail

type Driver interface {
	// Send 发送邮件
	Send(email Email, config map[string]string) bool
}
