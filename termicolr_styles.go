package termicolr

import (
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

var (
	noTTY = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))
	noColor = os.Getenv("NO_COLOR") != ""
)

// Style represents text formatting attributes and colors.
type Style struct {
	fg, bg                   color
	dim, bold, italic        bool
	underline, strikethrough bool
}

// NewStyle returns a new, empty Style instance.
func NewStyle() *Style { return &Style{} }

// Chained configuration methods.
func (s *Style) FG(c color) *Style     { s.fg = c; return s }
func (s *Style) BG(c color) *Style     { s.bg = c; return s }
func (s *Style) Dim() *Style           { s.dim = true; return s }
func (s *Style) Bold() *Style          { s.bold = true; return s }
func (s *Style) Italic() *Style        { s.italic = true; return s }
func (s *Style) Underline() *Style     { s.underline = true; return s }
func (s *Style) Strikethrough() *Style { s.strikethrough = true; return s }

func (s *Style) isEmpty() bool {
	return !s.bold && !s.dim && !s.italic &&
		!s.underline && !s.strikethrough &&
		s.fg == "" && s.bg == ""
}

// Sprint applies the style to the given text and returns the styled string.
func (s *Style) Sprint(text string) string {
	if text == "" || noTTY || noColor || s.isEmpty() {
		return text
	}

	var b strings.Builder
	b.WriteString("\x1b[")

	first := true
	write := func(code string) {
		if code == "" {
			return
		}
		if !first {
			b.WriteByte(';')
		}
		first = false
		b.WriteString(code)
	}

	if s.bold {
		write("1")
	}
	if s.dim {
		write("2")
	}
	if s.italic {
		write("3")
	}
	if s.underline {
		write("4")
	}
	if s.strikethrough {
		write("9")
	}

	if code, ok := s.fg.toSGR(false); ok {
		write(code)
	}
	if code, ok := s.bg.toSGR(true); ok {
		write(code)
	}

	if first {
		return text
	}

	b.WriteString("m")
	b.WriteString(text)
	b.WriteString("\x1b[0m")
	return b.String()
}
