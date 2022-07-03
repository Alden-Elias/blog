package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var (
	//验证码生成&存储
	captcha *base64Captcha.Captcha
)

func init() {
	// 设置自带的store
	store := base64Captcha.DefaultMemStore
	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "123456789qwertyuiplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString := captchaConfig
	driver := driverString.ConvertFonts()
	captcha = base64Captcha.NewCaptcha(driver, store)
}

func CaptGet() (id, b64s string, err error) {
	return captcha.Generate()
}

func CaptVerify(id, capt string) bool {
	return captcha.Store.Verify(id, capt, true)
}
