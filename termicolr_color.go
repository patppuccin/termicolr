package termicolr

import (
	"strconv"
	"strings"
)

type color string

func (c color) toSGR(bg bool) (string, bool) {
	if c == "" {
		return "", false
	}
	s := string(c)

	if len(s) > 4 && s[:4] == "rgb:" {
		s := s[4:] // after "rgb:"
		// find two commas in one pass
		c1, c2 := -1, -1
		for i := 0; i < len(s); i++ {
			if s[i] == ',' {
				if c1 == -1 {
					c1 = i
				} else {
					c2 = i
					break
				}
			}
		}
		if c1 == -1 || c2 == -1 {
			return "", false
		}
		r := s[:c1]
		g := s[c1+1 : c2]
		b := s[c2+1:]
		if bg {
			return "48;2;" + r + ";" + g + ";" + b, true
		}
		return "38;2;" + r + ";" + g + ";" + b, true
	}

	if len(s) > 5 && s[:5] == "ansi:" {
		n := strings.TrimPrefix(s, "ansi:")
		if n == "" {
			return "", false
		}
		if bg {
			return "48;5;" + n, true
		}
		return "38;5;" + n, true
	}
	return "", false
}

// --- Color Conversion Helpers ----------------------------
func ColorFromHex(hx string) color {
	hx = strings.TrimPrefix(strings.TrimSpace(hx), "#")
	if len(hx) != 6 {
		return ""
	}
	var result strings.Builder
	result.WriteString("rgb:")
	for i := 0; i < 6; i += 2 {
		val, err := strconv.ParseUint(hx[i:i+2], 16, 8)
		if err != nil {
			return ""
		}
		if i > 0 {
			result.WriteString(",")
		}
		result.WriteString(strconv.Itoa(int(val)))
	}
	return color(result.String())
}
func ColorFromRGB(r, g, b int) color {
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return ""
	}
	return color("rgb:" +
		strconv.Itoa(r) + "," +
		strconv.Itoa(g) + "," +
		strconv.Itoa(b))
}
func ColorFromANSI(i int) color {
	if i < 0 || i > 255 {
		return ""
	}
	return color("ansi:" + strconv.Itoa(i))
}
