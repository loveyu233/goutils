package imagecode

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/loveyu233/goutils/imagecode/model"
	"image"
	"image/jpeg"
	"image/png"
	"math/big"
	"os"
	"strconv"
)

// 获取随机数
func getRandInt(max int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(max-1)))
	return int(num.Int64())
}

type CaptchaResource struct {
	Path      string // 资源存放路径
	BgPath    string // 背景图存放路径
	MaskPath  string // 遮盖图片存放路径
	NewBgPath string // 生成后新的图片存放位置
	IsSave    bool   // 是否存放新的图片
}

type CROption func(*CaptchaResource)

func WithPath(path string) CROption {
	return func(r *CaptchaResource) {
		r.Path = path
	}
}

func WithBgPath(path string) CROption {
	return func(r *CaptchaResource) {
		r.BgPath = path
	}
}

func WithMaskPath(path string) CROption {
	return func(r *CaptchaResource) {
		r.MaskPath = path
	}
}

func WithNewBgPath(path string) CROption {
	return func(r *CaptchaResource) {
		r.NewBgPath = path
	}
}

func WithIsSave(is bool) CROption {
	return func(r *CaptchaResource) {
		r.IsSave = is
	}
}

func NewCaptchaResource(opt ...CROption) *CaptchaResource {
	cr := new(CaptchaResource)
	for _, o := range opt {
		o(cr)
	}
	if cr.Path == "" {
		panic("path不能为空")
	}
	if cr.BgPath == "" {
		cr.BgPath = cr.Path + "/bg/"
	}
	if cr.MaskPath == "" {
		cr.MaskPath = cr.Path + "/mask.png"
	}
	if cr.NewBgPath == "" {
		cr.NewBgPath = cr.Path + "/newBg/"
	}
	return cr
}

// CreateCode 创建图片验证码
func (c CaptchaResource) CreateCode() *model.CaptchaModel {
	//生成随机数,用来随机选取图片
	nums := getRandInt(10)
	//用于生成的图片名称
	imageId := uuid.New().String()
	//获取图片
	f, _ := os.Open(c.BgPath + strconv.Itoa(nums) + ".png")
	//获取随机x坐标
	imageRandX := getRandInt(480 - 100)
	if imageRandX < 200 {
		imageRandX += 200
	}
	//获取随机y坐标
	imageRandY := getRandInt(240 - 100)
	if imageRandY < 100 {
		imageRandY += 100
	}
	//转化为image对象
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	//设置截取的最大坐标值和最小坐标值
	maxPotion := image.Point{
		X: imageRandX,
		Y: imageRandY,
	}
	minPotion := image.Point{
		X: imageRandX - 100,
		Y: imageRandY - 100,
	}
	subimg := image.Rectangle{
		Max: maxPotion,
		Min: minPotion,
	}
	//截取图像
	data := imaging.Crop(m, subimg)
	if c.IsSave {
		f, err = os.Create(c.NewBgPath + imageId + "screenshot.jpeg")
		defer f.Close()
		jpeg.Encode(f, data, nil)
	}
	//base64编码
	buffer := bytes.NewBuffer(nil)
	jpeg.Encode(buffer, data, nil)
	maskBase64 := base64.StdEncoding.EncodeToString(buffer.Bytes())
	//设置遮罩
	bkBase64 := createCodeImg(c.BgPath+strconv.Itoa(nums)+".png", minPotion, imageId, c.MaskPath, c.NewBgPath, c.IsSave)
	captchaModel := &model.CaptchaModel{
		SliderImg: maskBase64,
		BgImg:     bkBase64,
		X:         int32(imageRandX),
		Y:         int32(imageRandY),
	}
	return captchaModel
}

// 创建有遮盖后的图片
func createCodeImg(path string, minPotion image.Point, imageId string, mask string, newBg string, isSave bool) string {
	bg, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	maskFile, err := os.Open(mask)
	if err != nil {
		panic(err)
	}
	bgimg, err := png.Decode(bg)
	maskimg, err := png.Decode(maskFile)
	//参数：背景图，遮盖图，坐标，透明度
	data := imaging.Overlay(bgimg, maskimg, minPotion, 1.0)
	if isSave {
		f, _ := os.Create(newBg + imageId + ".jpeg")
		defer f.Close()
		jpeg.Encode(f, data, nil)
	}
	//base64编码
	buffer := bytes.NewBuffer(nil)
	jpeg.Encode(buffer, data, nil)
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}
