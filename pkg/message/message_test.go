package message

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendMail(t *testing.T) {
	re := "313352050@qq.com"
	mes := "测试邮件"
	theme := "测试"
	err := SendMail(re, mes, theme)
	assert.Nil(t, err)
}
