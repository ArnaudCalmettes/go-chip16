package graphics

import (
	"image"
	"image/color"
)

const (
	ScreenW = 320
	ScreenH = 240
)

var defaultPalette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0x00}, // 0x0: Black (Transparent in FG)
	color.RGBA{0x00, 0x00, 0x00, 0xFF}, // 0x1: Black
	color.RGBA{0x88, 0x88, 0x88, 0xFF}, // 0x2: Gray
	color.RGBA{0xBF, 0x39, 0x32, 0xFF}, // 0x3: Red
	color.RGBA{0xDE, 0x7A, 0xAE, 0xFF}, // 0x4: Pink
	color.RGBA{0x4C, 0x3D, 0x21, 0xFF}, // 0x5: Dark Brown
	color.RGBA{0x90, 0x5F, 0x25, 0xFF}, // 0x6: Brown
	color.RGBA{0xE4, 0x94, 0x52, 0xFF}, // 0x7: Orange
	color.RGBA{0xEA, 0xD9, 0x79, 0xFF}, // 0x8: Yellow
	color.RGBA{0x53, 0x7A, 0x3B, 0xFF}, // 0x9: Green
	color.RGBA{0xAB, 0xD5, 0x4A, 0xFF}, // 0xA: Light Green
	color.RGBA{0x25, 0x2E, 0x38, 0xFF}, // 0xB: Dark Blue
	color.RGBA{0x00, 0x46, 0x7F, 0xFF}, // 0xC: Blue
	color.RGBA{0x68, 0xAB, 0xCC, 0xFF}, // 0xD: Light Blue
	color.RGBA{0xBC, 0xDE, 0xE4, 0xFF}, // 0xE: Sky Blue
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, // 0xF: White
}

// DefaultPalette returns a new palette initialized to chip16's default colors.
// This is kept a function so that a new color.Palette object gets created
// everytime, to avoid side effects during palette update.
func DefaultPalette() color.Palette {
	p := make(color.Palette, len(defaultPalette))
	copy(p, defaultPalette)
	return p
}

// State describes a state of the graphics system of the chip16.
type State struct {
	// Palette is the current color palette.
	Palette color.Palette

	// BG is the current background color (palette index).
	BG uint8

	// Screen is the current foreground image.
	FG *image.Paletted

	// SpriteW is the width of the sprite(s) to draw.
	SpriteW uint8

	// SpriteH is the height of the sprite(s) to draw.
	SpriteH uint8

	// HFlip tells whether the sprite(s) must be flipped horizontally.
	HFlip bool

	// VFlip tells whether the sprite(s) must be flipped vertically.
	VFlip bool
}

// NewState constructs and initialize a new graphics State
func NewState() *State {
	p := DefaultPalette()
	s := &State{
		Palette: p,
		FG:      image.NewPaletted(image.Rect(0, 0, ScreenW, ScreenH), p),
	}
	return s
}

var emptyFG = make([]uint8, ScreenW*ScreenH)

// Clear clears the screen and resets the background color.
func (s *State) Clear() {
	s.BG = 0

	// This is ~5x faster than recreating a paletted image altogether
	copy(s.FG.Pix, emptyFG)
}
