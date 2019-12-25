package factoryFont

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/font"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontFamily"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontStyle"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontVariant"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontWeight"
	"image/color"
)

// en: Create a new font struct
//     size: size of the font
//     sizeUnit: size unit of the font
//     fontColor: a color.RGBA struct used with text
//     family: font name. Must be a string or a constant fontFamily.FontFamily
//             Example: 'Verdana' or fontFamily.KHelvetica
//     style: must be fontStyle.KItalic or fontStyle.KOblique
//     variant: fontVariant.KNormal or fontVariant.KSmallCaps
//     weight: must be fontWeight.KNormal, fontWeight.KBold, fontWeight.KBolder,
//     fontWeight.KLighter or fontWeight.K100 to fontWeight.K900
//     density: Please, see a density methods
//     iDensity: Please, see a density methods
//
// pt_br: Cria um novo struct de fonte
//     size: tamanho da fonte
//     sizeUnit: unidade do tamanho da fonte, exemplo: 'px'
//     fontColor: um struct color.RGBA usado com texto
//     family: nome da fonte escolhida. Pode ser string ou uma constante
//     fontFamily.FontFamily.
//             Exemplo: 'Verdana' ou fontFamily.KHelvetica
//     style: deve ser fontStyle.KItalic ou fontStyle.KOblique
//     variant: fontVariant.KSmallCaps
//     weight: deve ser fontWeight.KNormal, fontWeight.KBold, fontWeight.KBolder,
//     fontWeight.KLighter ou fontWeight.K100 a fontWeight.K900
//     density: Por favor, veja o elemento densidade
//     iDensity: Por favor, veja o elemento densidade
func NewFontWithStyleVariantWeight(size int, sizeUnit string, fontColor color.RGBA, family fontFamily.FontFamily, style fontStyle.FontStyle, variant fontVariant.FontVariant, weight fontWeight.FontWeight, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) font.Font {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(size)
	size = densityCalc.Int()

	f := font.Font{
		Size:     size,
		SizeUnit: sizeUnit,
		Color:    fontColor,
		Family:   family,
		Style:    style,
		Variant:  variant,
		Weight:   weight,
	}
	return f
}