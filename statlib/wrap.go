package statlib

const minWidth = 20
const maxWidth = 200

func clampWidth(width *int) {
	if *width < minWidth {
		*width = minWidth
	} else if *width > maxWidth {
		*width = maxWidth
	}
}
