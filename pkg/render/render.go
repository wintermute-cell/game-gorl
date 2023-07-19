package render

import (
    "math"
    "boom-gorl/pkg/util"
    rl "github.com/gen2brain/raylib-go/raylib"
)

type RenderState struct {
    TargetTex        rl.RenderTexture2D
    RenderResolution rl.Vector2
    RenderScale      rl.Vector2
    MinScale         float32
}


func (rs *RenderState) BeginCustomRender() {
    if rs.TargetTex.ID == 0 {
        rs.TargetTex = rl.LoadRenderTexture(int32(rs.RenderResolution.X), int32(rs.RenderResolution.Y))
        rl.SetTextureFilter(rs.TargetTex.Texture, rl.FilterPoint)
    }
    scale := rl.Vector2{
        X: float32(rl.GetScreenWidth()) / rs.RenderResolution.X,
        Y: float32(rl.GetScreenHeight()) / rs.RenderResolution.Y,
    }
    rs.MinScale = float32(math.Min(float64(scale.X), float64(scale.Y)))

    rl.BeginTextureMode(rs.TargetTex)
    rs.RenderScale = scale

    mouse := rl.GetMousePosition()
    virtualMouse := rl.Vector2{}
    virtualMouse.X = (mouse.X - float32(rl.GetScreenWidth())-(rs.RenderResolution.X*rs.MinScale)*0.5) / rs.MinScale
    virtualMouse.Y = (mouse.Y - float32(rl.GetScreenHeight())-(rs.RenderResolution.Y*rs.MinScale)*0.5) / rs.MinScale
    virtualMouse = util.Vector2Clamp(virtualMouse, rl.Vector2{}, rs.RenderResolution)

    rl.SetMouseOffset(int(-(float32(rl.GetScreenWidth()) - rs.RenderResolution.X*rs.MinScale)*0.5),
        int(-(float32(rl.GetScreenHeight()) - rs.RenderResolution.Y*rs.MinScale)*0.5))
    rl.SetMouseScale(1/rs.MinScale, 1/rs.MinScale)
}

func (rs *RenderState) EndCustomRender() {
    rl.EndTextureMode()
    rl.DrawTexturePro(rs.TargetTex.Texture,
        rl.Rectangle{
            X: 0.0,
            Y: 0.0,
            Width: float32(rs.TargetTex.Texture.Width),
            Height: -float32(rs.TargetTex.Texture.Height),
        },
        rl.Rectangle{
            X: (float32(rl.GetScreenWidth()) - rs.RenderResolution.X*rs.MinScale) * 0.5,
            Y: (float32(rl.GetScreenHeight()) - rs.RenderResolution.Y*rs.MinScale) * 0.5,
            Width: rs.RenderResolution.X * rs.MinScale,
            Height: rs.RenderResolution.Y * rs.MinScale},
        rl.Vector2{X:0, Y:0}, 0, rl.White) 
}
