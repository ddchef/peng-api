package services

import (
	"bytes"
	"context"
	"peng-api/global"
	"time"

	"github.com/dchest/captcha"
	"github.com/gofrs/uuid"
)

type captchaService struct{}

var CaptchaService = new(captchaService)

func (casbinService *captchaService) CreateCaptcha() (uuidStr string, err error) {
	captchaId := captcha.NewLen(global.App.Config.Captcha.Len)
	id, err := uuid.NewV4()
	uuidStr = id.String()
	err = global.App.Redis.SetNX(context.Background(), uuidStr, captchaId, time.Duration(global.App.Config.Captcha.CaptchaTtl)*time.Second).Err()
	return
}

func (casbinService *captchaService) CaptchaImage(id string) (img bytes.Buffer) {
	captchaId, err := global.App.Redis.Get(context.Background(), id).Result()
	if err != nil {
		return
	}
	captcha.WriteImage(&img, captchaId, global.App.Config.Captcha.Width, global.App.Config.Captcha.Height)
	return
}

func (casbinService *captchaService) Verify(id string, digits string) bool {
	captchaId, err := global.App.Redis.Get(context.Background(), id).Result()
	if err != nil {
		return false
	}
	if captcha.VerifyString(captchaId, digits) != true {
		global.App.Redis.Del(context.Background(), id)
		return false
	}
	return true
}
