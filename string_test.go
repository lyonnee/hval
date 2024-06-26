package hvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringContains(t *testing.T) {
	var s1 string = "hello,world"
	substr := "llo"
	err := Validate[string](s1, ContainsStr(substr))
	assert.NoError(t, err)

	var s2 string = "iamlyon.ne"
	err = Validate[string](s2, ContainsStr("nee", "不包含指定字符串"))
	assert.Error(t, err)
}

func TestIsEmail(t *testing.T) {
	var e1 string = "lyon.nee@outlook.com"
	err := Validate[string](e1, IsEmail("无效的邮箱地址"))
	assert.NoError(t, err)

	var e2 string = "lyon.neeoutlook.com"
	err = Validate[string](e2, IsEmail("无效的邮箱地址"))
	assert.Error(t, err)
}

func TestIsUrl(t *testing.T) {
	err := Validate[string]("testURL", IsUrl("无效的url"))
	assert.Error(t, err)

	err = Validate[string]("lyon.nee/", IsUrl())
	assert.Error(t, err)

	err = Validate[string]("http://github.com", IsUrl())
	assert.NoError(t, err)

	err = Validate[string]("https://github.com/lyonnee/hvalid", IsUrl())
	assert.NoError(t, err)
}

func TestIsIP(t *testing.T) {
	err := Validate[string]("192.168.0.1", IsIPv4())
	assert.NoError(t, err)

	err = Validate[string]("0.0.0.0", IsIPv4())
	assert.NoError(t, err)

	err = Validate[string]("266.0.0.-1", IsIPv4("无效的IPv4地址"))
	assert.Error(t, err)

	err = Validate[string]("ac.0.a.-1", IsIPv4("无效的IPv4地址"))
	assert.Error(t, err)

	err = Validate[string]("FC00:0000:130F:0000:0000:09C0:876A:130B", IsIPv6())
	assert.NoError(t, err)

	err = Validate[string]("XXXX:0000:130F:0000:0000:09C0:876A:130B", IsIPv6())
	assert.Error(t, err)

	err = Validate[string]("abcf:0000:130F:efgd:0000:09C0:876A:130B", IsIPv6())
	assert.Error(t, err)
}
