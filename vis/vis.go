package vis

import (
    "image/gif"
    "image"
    "image/color"
    "os"
    "fmt"
    "log"
)

type GifWriter struct {
    dim image.Rectangle
    frames []*image.Paletted
    delays []int
    palette color.Palette
}

func (g *GifWriter) getLast() (*image.Paletted, error) {
    if len(g.frames) == 0 {
        return nil, fmt.Errorf("contains no frames")

    }
    return g.frames[len(g.frames)-1], nil
}

func (g *GifWriter) PushFrame(colorI, delayMS int) {
    g.frames = append(g.frames, image.NewPaletted(g.dim, g.palette))
    g.delays = append(g.delays, (delayMS+5)/10)
    for x := range g.dim.Dx() {
        for y := range g.dim.Dy() {
            f, _ := g.getLast()
            f.Set(x, y, g.palette[colorI])
        }
    }
}

func (g *GifWriter) SetPixel(x, y, cI int) {
    f, err := g.getLast()
    if err != nil {
        log.Panic(err)
    }
    f.Set(x,y,g.palette[cI])
}

func (g *GifWriter) Write(filename string) error {
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    err = gif.EncodeAll(f, &gif.GIF{
        Image: g.frames,
        Delay: g.delays,
    })
    return err
}

func NewGifWriter(x, y int, palette color.Palette) GifWriter {
    return GifWriter{
        dim: image.Rect(0, 0, x, y),
        frames: []*image.Paletted{},
        delays: []int{},
        palette: palette,
    } 
}

