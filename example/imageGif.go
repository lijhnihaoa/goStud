package example

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //完整的X振荡器旋转数量
		res     = 0.001 //角度分辨率
		size    = 100   //图像画布覆盖[-size .. +size]
		nfranme = 64    //动画框架数量
		delay   = 8     //10ms单位帧之间的延迟
	)
	freq := rand.Float64() * 3.0 //Y振荡器的相对频率
	anim := gif.GIF{LoopCount: nfranme}
	phase := 0.0 //相位差
	for i := 0; i < nfranme; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func StudImageGif() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func Call(out io.Writer) {
	lissajous(out)
}
