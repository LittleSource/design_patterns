package decorate

import "fmt"

type Component interface {
	send(msg string)
}

type Concrete struct {
}

func (c *Concrete) send(msg string) {
	fmt.Println("基本的发送消息：" + msg)
}

type BaseDecorator struct {
	component Component
}

func NewBaseDecorator(component Component) *BaseDecorator {
	return &BaseDecorator{component: component}
}

func (b *BaseDecorator) send(msg string) {
	b.component.send(msg)
}

// 下面模拟两个具体的发送消息对象

type WecahtDecorator struct {
}

func (w *WecahtDecorator) send(msg string) {
	fmt.Println("微信发送消息：" + msg)
}

type EmailDecorator struct{}

func (e *EmailDecorator) send(msg string) {
	fmt.Println("邮件发送消息：" + msg)
}
