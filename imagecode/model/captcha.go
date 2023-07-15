package model

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type CaptchaModel struct {
	SliderImg string `json:"sliderImg,omitempty"`
	BgImg     string `json:"bgImg,omitempty"`
	X         int32  `json:"x,omitempty"`
	Y         int32  `json:"y,omitempty"`
}

type CaptchaVo struct {
	SliderImg string `json:"sliderImg,omitempty"`
	BgImg     string `json:"bgImg,omitempty"`
	Y         int32  `json:"y,omitempty"`
}

// ToVo 转化为返回给前端的结构体和x轴坐标
func (m CaptchaModel) ToVo() (CaptchaVo, int32) {
	return CaptchaVo{
		SliderImg: m.SliderImg,
		BgImg:     m.BgImg,
		Y:         m.Y,
	}, m.X
}

// SaveRedis 吧x轴坐标保存到redis
func (m CaptchaModel) SaveRedis(c *redis.Client, key string, expiration time.Duration) error {
	return c.Set(context.TODO(), key, m.X, expiration).Err()
}
