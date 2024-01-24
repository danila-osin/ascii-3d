package screen

import "github.com/danila-osin/ascii-3d/pkg/geometry"

type Size struct {
	W, H int
}

// BRenderFn Calls before Render()
type BRenderFn func()

// ARenderFn Calls after Render()
type ARenderFn func()

type iteratorFn func(cursor geometry.Vec2[int], value string)

type iteratorSetFn func(cursor geometry.Vec2[int], value string) string
