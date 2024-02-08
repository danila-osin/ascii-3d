# ASCII-3D
## Description
ASCII 3d/2d render engine for console

## Showcase

https://github.com/danila-osin/ascii-3d/assets/55761790/62a0d56d-769c-4eeb-937a-bf6e9edd273c

## Usage
### go run
`go run cmd/main.go [params]`
```bash
params:
  -w int // Screen Width (default 50)
  
  -h int // Screen Height (default 50)
  
  -fr int // Frame Rate (default 20)
  
  -m string // App Mode [life, graph, controls, 3d] (default "unknown")
  
  -fa float // Font Aspect (default 0.4)
```
### make
`make run [params]`

```bash
params:
  w=int // Screen Width (default 50)
  
  h=int // Screen Height (default 50)
  
  fr=int // Frame Rate (default 20)
  
  m=string // App Mode [life, graph, controls, 3d] (default "unknown")
  
  fa=float // Font Aspect (default 0.4)
```


## TODO

### General
* ~~Framerate Setting~~
* ~~Screen Size Setting~~
* ~~Keyboard controls~~
  * ~~listening keyboard key presses~~
  * ~~controls description on screen~~
* Realtime screen size changing
* Configuration using TOML file
* Profiling(`pprof`)

### 2D
* Game of Life (`playable`)
  * ~~setting field size~~
  * setting initial state by keyboard
  * changing field size by keyboard
* Function Graph (`almost done`)
  * research rasterisation(`low priority`)
  * ~~scaling~~
  * ~~moving~~
  * ~~dynamic parameter~~ 
  * ~~keyboard(scaling, moving, changing parameter)~~
* ~~Screen Text~~(`done`)
  * ~~framing~~
* Text Generators(`low priority`)
  * scaling
  * moving
  * keyboard(scaling, moving)
  * text over any other modes

### 3D
* General
  * ~~color/light gradient~~
* Shapes
  * ~~setting shape~~
  * ~~setting camera point of view~~
  * ~~setting light source~~
  * ~~keyboard(camera moving, changing shape)~~
