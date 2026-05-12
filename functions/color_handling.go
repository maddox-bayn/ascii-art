package functions

const (
	Black   = "\033[0;30m"
	Red     = "\033[0;31m"
	Green   = "\033[0;32m"
	Yellow  = "\033[0;33m"
	Blue    = "\033[0;34m"
	Magenta = "\033[0;35m"
	Cyan    = "\033[0;36m"
	White   = "\033[0;37m"
	Orange  = "\033[0;38;5;208m"
	Reset   = "\033[0m"
)

var ColorMap = map[string]string{
	"black":   Black,
	"red":     Red,
	"green":   Green,
	"yellow":  Yellow,
	"blue":    Blue,
	"magenta": Magenta,
	"cyan":    Cyan,
	"orange":  Orange,
	"white":   White,
}

func GetColorMusk(word, subStr, color string) []bool {
	needColor := make([]bool, len(word))
	if subStr == "" && color != "" {
		// If no substring is specified, color the whole word
		for i := range needColor {
			needColor[i] = true
		}
		return needColor
	}

	// Instead of just one strings.Index:
	if subStr != "" {
		for i := 0; i <= len(word)-len(subStr); i++ {
			if word[i:i+len(subStr)] == subStr {
				for j := 0; j < len(subStr); j++ {
					needColor[i+j] = true
				}
				// Optional: i += len(subStr) - 1 to skip overlapping matches
			}
		}
	}
	return needColor
}
