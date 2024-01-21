package screen

import "github.com/danila-osin/ascii-3d/pkg/geometry"

type Size struct {
	W, H int
}

type iteratorFunc func(cursor geometry.Vec2[int], value string)

type iteratorSetFunc func(cursor geometry.Vec2[int], value string) string
