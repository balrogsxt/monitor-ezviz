package ezviz

import (
	"errors"
	"github.com/imroc/req/v3"
)

type Ezviz struct {
	appKey    string
	appSecret string
}

func NewClient(appKey, appSecret string) *Ezviz {
	return &Ezviz{
		appKey:    appKey,
		appSecret: appSecret,
	}
}

func (c *Ezviz) GetAccessToken() (string, error) {
	res, err := c.buildRequest().SetFormDataAnyType(map[string]interface{}{
		"appKey":    c.appKey,
		"appSecret": c.appSecret,
	}).Post("/api/lapp/token/get")
	if err != nil {
		return "", err
	}
	var r GetAccessTokenRes
	if err := res.Unmarshal(&r); err != nil {
		return "", err
	}
	if !r.IsOk() {
		return "", errors.New(r.Msg)
	}
	return r.Data.AccessToken, nil
}

// GetPlayAddress 获取播放地址
func (c *Ezviz) GetPlayAddress(deviceSerial string) (string, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return "", err
	}
	res, err := c.buildRequest().SetFormDataAnyType(map[string]interface{}{
		"accessToken":  accessToken,
		"deviceSerial": deviceSerial,
		"channelNo":    1,
		"protocol":     4, //流播放协议，1-ezopen、2-hls、3-rtmp、4-flv，默认为1
		"quality":      1, //视频清晰度，1-高清（主码流）、2-流畅（子码流）
		"expireTime":   86400 * 365,
	}).Post("/api/lapp/v2/live/address/get")
	if err != nil {
		return "", err
	}
	var r GetPlayAddress
	if err := res.Unmarshal(&r); err != nil {
		return "", err
	}
	if !r.IsOk() {
		return "", errors.New(r.Msg)
	}
	return r.Data.Url, nil
}
func (c *Ezviz) buildRequest() *req.Request {
	return req.C().SetBaseURL(gateway).R()
}
