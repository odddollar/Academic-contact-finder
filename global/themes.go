package global

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Global colours
var (
	Grey        = color.NRGBA{86, 86, 86, 255}
	Purple      = color.NRGBA{138, 94, 169, 255}
	LightPurple = color.NRGBA{198, 154, 229, 255}
	White       = color.NRGBA{255, 255, 255, 255}
)

// Theme struct that implements fyne's Theme interface
type MainTheme struct{}

// Return custom colours, falling back to defaults otherwise
func (t MainTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		return Purple
	case theme.ColorNameWarning: // Used for setting things like backgrounds in certain widgets
		return LightPurple
	default:
		return theme.DefaultTheme().Color(name, theme.VariantLight)
	}
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
	switch name {
	case theme.SizeNameSubHeadingText:
		return 19
	case theme.SizeNameText:
		return 15
	case theme.SizeNameCaptionText:
		return 12
	default:
		return theme.DefaultTheme().Size(name)
	}
}
