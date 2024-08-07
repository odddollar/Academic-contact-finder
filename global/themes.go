package global

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var TextColour color.NRGBA = color.NRGBA{86, 86, 86, 255}
var ButtonColour color.NRGBA = color.NRGBA{138, 94, 169, 255}
var ButtonTextColour color.NRGBA = color.NRGBA{255, 255, 255, 255}

// Theme struct that implements fyne's Theme interface
type MainTheme struct{}

// Return custom colours, falling back to defaults otherwise
func (t MainTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

// Return default icons
func (t MainTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Return default fonts
func (t MainTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Return default sizes
func (t MainTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

// Main button theme to give purple background with white text
type ButtonTheme struct{}

// Return custom colours, falling back to defaults otherwise
func (t ButtonTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameForeground:
		return ButtonTextColour
	case theme.ColorNameButton:
		return ButtonColour
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

// Return default icons
func (t ButtonTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Return default fonts
func (t ButtonTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Return default sizes
func (t ButtonTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}