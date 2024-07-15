package ezviz

const (
	gateway = "https://open.ys7.com"
)

type BaseRes struct {
	Code string
	Msg  string
}

func (c *BaseRes) IsOk() bool {
	return c.Code == "200"
}

type GetAccessTokenRes struct {
	BaseRes
	Data struct {
		AccessToken string `json:"accessToken"`
		ExpireTime  int64  `json:"expireTime"` //毫秒
	}
}

type GetPlayAddress struct {
	BaseRes
	Data struct {
		Id         string
		Url        string
		ExpireTime string
	}
}
