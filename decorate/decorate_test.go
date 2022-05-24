package decorate

import "testing"

func TestDecorate(t *testing.T) {
	w := NewBaseDecorator(&WecahtDecorator{})
	w.send("hello")

	e := NewBaseDecorator(&EmailDecorator{})

	e.send("hello")
}
