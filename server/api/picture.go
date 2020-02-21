package api

import (
	"bufio"
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"math/rand"

	"github.com/goadesign/goa"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/lucasb-eyer/go-colorful"

	"github.com/fieldkit/cloud/server/api/app"
	"github.com/fieldkit/cloud/server/backend/repositories"
)

const (
	float64MaxUint64 = float64(math.MaxUint64)
)

type PictureController struct {
	*goa.Controller
	options *ControllerOptions
}

func NewPictureController(service *goa.Service, options *ControllerOptions) *PictureController {
	return &PictureController{
		options:    options,
		Controller: service.NewController("PictureController"),
	}
}

func UserDefaultPicture(id int64) ([]byte, error) {
	r := rand.New(rand.NewSource(id))

	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	gc := draw2dimg.NewGraphicContext(img)

	c := colorful.Hsv(r.Float64()*360., .8, .6)
	cRGBA := color.RGBAModel.Convert(c).(color.RGBA)
	gc.SetFillColor(cRGBA)
	gc.SetStrokeColor(cRGBA)
	gc.SetLineWidth(5)

	increments := float64(r.Intn(20) + 20)
	for i := 0.; i < increments; i += 1. {

		angle := math.Pi * 2. * i / increments
		radius := r.Float64() * 120.

		x := math.Cos(angle)*radius + 128.
		y := math.Sin(angle)*radius + 128.

		if i == .0 {
			gc.MoveTo(x, y)
		} else {
			gc.LineTo(x, y)
		}
	}

	gc.Close()
	gc.FillStroke()

	picture := bytes.NewBuffer([]byte{})
	if err := png.Encode(picture, img); err != nil {
		return nil, err
	}

	return picture.Bytes(), nil
}

func ProjectDefaultPicture(id int64) ([]byte, error) {
	r := rand.New(rand.NewSource(id))

	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	gc := draw2dimg.NewGraphicContext(img)

	c := colorful.Hsv(r.Float64()*360., .8, .6)
	cRGBA := color.RGBAModel.Convert(c).(color.RGBA)
	gc.SetFillColor(cRGBA)
	gc.SetStrokeColor(cRGBA)
	gc.SetLineWidth(5)

	for i := 0; i < 32; i++ {
		x := math.Max(math.Min(r.NormFloat64()*32+128, 255.-64-8-2.5), 8.+64+2.5)
		y := math.Max(math.Min(r.NormFloat64()*32+128, 255.-64-8-2.5), 8.+64+2.5)

		gc.MoveTo(x+r.Float64()*32-64, y+r.Float64()*32-64)
		gc.LineTo(x+r.Float64()*32+32, y+r.Float64()*32-64)
		gc.LineTo(x+r.Float64()*32+32, y+r.Float64()*32+32)
		gc.LineTo(x+r.Float64()*32-64, y+r.Float64()*32+32)
		gc.Close()
		gc.FillStroke()
	}

	picture := bytes.NewBuffer([]byte{})
	if err := png.Encode(picture, img); err != nil {
		return nil, err
	}

	return picture.Bytes(), nil
}

func ExpeditionDefaultPicture(id int64) ([]byte, error) {
	r := rand.New(rand.NewSource(id))

	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	gc := draw2dimg.NewGraphicContext(img)

	c := colorful.Hsv(r.Float64()*360., .8, .6)
	cRGBA := color.RGBAModel.Convert(c).(color.RGBA)
	gc.SetFillColor(cRGBA)
	gc.SetStrokeColor(cRGBA)
	gc.SetLineWidth(5)

	for i := 0; i < 32; i++ {
		x := math.Max(math.Min(r.NormFloat64()*32+128, 255.-64-8-2.5), 8.+64+2.5)
		y := math.Max(math.Min(r.NormFloat64()*32+128, 255.-64-8-2.5), 8.+64+2.5)

		gc.MoveTo(x+r.Float64()*32-16, y+r.Float64()*32-64)
		gc.LineTo(x+r.Float64()*32+32, y+r.Float64()*32-16)
		gc.LineTo(x+r.Float64()*32-16, y+r.Float64()*32+32)
		gc.LineTo(x+r.Float64()*32-64, y+r.Float64()*32-16)
		gc.Close()
		gc.FillStroke()
	}

	picture := bytes.NewBuffer([]byte{})
	if err := png.Encode(picture, img); err != nil {
		return nil, err
	}

	return picture.Bytes(), nil
}

func StationDefaultPicture(id int64) ([]byte, error) {
	r := rand.New(rand.NewSource(id))
	x := 124
	y := 100

	img := image.NewRGBA(image.Rect(0, 0, x, y))
	gc := draw2dimg.NewGraphicContext(img)

	c := colorful.Hsv(r.Float64()*360., .8, .6)
	cRGBA := color.RGBAModel.Convert(c).(color.RGBA)
	gc.SetFillColor(cRGBA)
	gc.SetStrokeColor(cRGBA)
	gc.SetLineWidth(5)

	xd8 := float64(x / 8)
	yd8 := float64(y / 8)
	xd16 := float64(x / 16)
	yd16 := float64(y / 16)
	xd4 := float64(x / 4)
	yd4 := float64(y / 4)

	for i := 0; i < 32; i++ {
		x := math.Max(math.Min(r.NormFloat64()*xd8+float64(x/2.), float64(x)-float64(x/4)-8-2.5), 8.+float64(x/4)+2.5)
		y := math.Max(math.Min(r.NormFloat64()*yd8+float64(y/2.), float64(y)-float64(y/4)-8-2.5), 8.+float64(y/4)+2.5)

		gc.MoveTo(x+r.Float64()*xd8-xd16, y+r.Float64()*yd8-yd4)
		gc.LineTo(x+r.Float64()*xd8+xd8, y+r.Float64()*yd8-yd16)
		gc.LineTo(x+r.Float64()*xd8-xd16, y+r.Float64()*yd8+yd8)
		gc.LineTo(x+r.Float64()*xd8-xd4, y+r.Float64()*yd8-yd16)
		gc.Close()
		gc.FillStroke()
	}

	picture := bytes.NewBuffer([]byte{})
	if err := png.Encode(picture, img); err != nil {
		return nil, err
	}

	return picture.Bytes(), nil
}

func (c *PictureController) UserGetID(ctx *app.UserGetIDPictureContext) error {
	picture, err := UserDefaultPicture(int64(ctx.UserID))
	if err != nil {
		return err
	}

	mr := repositories.NewMediaRepository(c.options.Session)

	lm, err := mr.Load(ctx, "598c0282-5c8c-4bdd-9e82-950ab796c059")
	if err != nil {
		return err
	}

	if lm != nil {
		writer := bufio.NewWriter(ctx.ResponseData)

		_, err = io.Copy(writer, lm.Reader)
		if err != nil {
			return err
		}
		return nil
	}

	return ctx.OK(picture)
}

func (c *PictureController) UserSaveID(ctx *app.UserSaveIDPictureContext) error {
	mr := repositories.NewMediaRepository(c.options.Session)

	saved, err := mr.Save(ctx, ctx.RequestData)
	if err != nil {
		return err
	}

	// Need to save URL and ID with whatever this media should be attached to.

	_ = saved

	return ctx.OK(&app.MediaReferenceResponse{
		ID:  0,
		URL: "",
	})
}

func (c *PictureController) ProjectGetID(ctx *app.ProjectGetIDPictureContext) error {
	picture, err := ProjectDefaultPicture(int64(ctx.ProjectID))
	if err != nil {
		return err
	}

	return ctx.OK(picture)
}

func (c *PictureController) ExpeditionGetID(ctx *app.ExpeditionGetIDPictureContext) error {
	picture, err := ExpeditionDefaultPicture(int64(ctx.ExpeditionID))
	if err != nil {
		return err
	}

	return ctx.OK(picture)
}
