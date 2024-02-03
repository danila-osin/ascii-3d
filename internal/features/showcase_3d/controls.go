package showcase_3d

import (
	"fmt"
	"github.com/danila-osin/ascii-3d/internal/config"
	"github.com/danila-osin/ascii-3d/pkg/angle"
	"github.com/danila-osin/ascii-3d/pkg/controls"
	"github.com/danila-osin/ascii-3d/pkg/geometry"
	rot "github.com/danila-osin/ascii-3d/pkg/rotation"
	"os"
)

var cameraSens = 2.5

func setupControls(config config.Config, state *state) *controls.Controls {
	actions := []controls.Action{
		{
			Keys:        []string{"w"},
			Description: "Move Forward x+",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					v := state.cameraDir().MulN(0.1)
					state.cameraPos = state.cameraPos.Add(v)
				},
			},
		},
		{
			Keys:        []string{"s"},
			Description: "Move Backward x-",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					v := state.cameraDir().MulN(-0.1)
					state.cameraPos = state.cameraPos.Add(v)
				},
			},
		},
		{
			Keys:        []string{"d"},
			Description: "Move Right y+",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					camDir := state.cameraDir()
					camDir.Z = 0
					v := rot.RotateVec3(geometry.ZAxis, angle.Degree(90), camDir).MulN(0.1)
					state.cameraPos = state.cameraPos.Add(v)
				},
			},
		},
		{
			Keys:        []string{"a"},
			Description: "Move Left y-",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					camDir := state.cameraDir()
					camDir.Z = 0
					v := rot.RotateVec3(geometry.ZAxis, angle.Degree(-90), camDir).MulN(0.1)
					state.cameraPos = state.cameraPos.Add(v)
				},
			},
		},
		{
			Keys:        []string{"e"},
			Description: "Move Up z-",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraPos.Z -= 0.1
				},
			},
		},
		{
			Keys:        []string{"q"},
			Description: "Move Down z+",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraPos.Z += 0.1
				},
			},
		},
		{
			Keys:        []string{"i"},
			Description: "Rotate Up",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraRot.Yaw = state.cameraRot.Yaw.Add(angle.Degree(1 * cameraSens))
				},
			},
		},
		{
			Keys:        []string{"k"},
			Description: "Rotate Down",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraRot.Yaw = state.cameraRot.Yaw.Add(angle.Degree(-1 * cameraSens))
				},
			},
		},
		{
			Keys:        []string{"l"},
			Description: "Rotate Right",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraRot.Pitch = state.cameraRot.Pitch.Add(angle.Degree(1 * cameraSens))
				},
			},
		},
		{
			Keys:        []string{"j"},
			Description: "Rotate Left",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraRot.Pitch = state.cameraRot.Pitch.Add(angle.Degree(-1 * cameraSens))
				},
			},
		},
		{
			Keys:        []string{"u"},
			Description: "Roll Right",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraRot.Roll = state.cameraRot.Roll.Add(angle.Degree(1 * cameraSens))
				},
			},
		},
		{
			Keys:        []string{"o"},
			Description: "Roll Left",
			Handlers: []controls.ActionHandler{
				func(_ *controls.Controls) {
					state.cameraRot.Roll = state.cameraRot.Roll.Add(angle.Degree(-1 * cameraSens))
				},
			},
		},
		{
			Keys:        []string{"Q"},
			Description: "Exit",
			Handlers: []controls.ActionHandler{
				func(c *controls.Controls) {
					fmt.Println("Exit...")
					c.Close()
					os.Exit(1)
				},
			},
		},
	}

	return controls.New(config, actions)
}
