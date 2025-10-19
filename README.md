# TERMICOLR

> Lightweight terminal theming package for Go CLI apps

TERMICOLR (pronounced "termee color") provides a simple, expressive way to manage colors, styles, and themes in Go CLI environments. It helps to build consistent, colorful, and human-friendly command-line interfaces with minimal dependencies.

### ✨ Features

- TrueColor (RGB / HEX) and ANSI-256 color support  
- Chainable, composable style definitions  
- Cross-platform with `TTY` and `NO_COLOR` detection  
- Reusable theme presets (Catppuccin, Gruvbox, Dracula, and more)

### 🚀 Installation

```sh
go get github.com/patppuccin/termicolr
```

### 📖 Usage

#### Basic usage

```go
package main

import (
	"fmt"
	"github.com/patppuccin/termicolr"
)

func main() {
	theme := termicolr.CatppuccinMocha
	style := termicolr.NewStyle().FG(theme.Primary).Bold().Underline()
	fmt.Println(style.Sprint("Hello, Termicolr!"))
}
```

#### Custom themes

```go
package main

import (
	"fmt"
	"github.com/patppuccin/termicolr"
)

func main() {
    myTheme := termicolr.Theme{
        Background:    termicolr.ColorFromHex("#1e1e2e"),
        BackgroundAlt: termicolr.ColorFromHex("#181825"),
        Foreground:    termicolr.ColorFromHex("#cdd6f4"),
        ForegroundAlt: termicolr.ColorFromHex("#bac2de"),
        Primary:       termicolr.ColorFromHex("#cba6f7"),
        Secondary:     termicolr.ColorFromHex("#f2cdcd"),
        Accent:        termicolr.ColorFromHex("#fab387"),
        Highlight:     termicolr.ColorFromHex("#f5e0dc"),
        Muted:         termicolr.ColorFromHex("#6c7086"),
        Red:           termicolr.ColorFromHex("#f38ba8"),
        Green:         termicolr.ColorFromHex("#a6e3a1"),
        Yellow:        termicolr.ColorFromHex("#f9e2af"),
        Blue:          termicolr.ColorFromHex("#89b4fa"),
        Purple:        termicolr.ColorFromHex("#cba6f7"),
        Orange:        termicolr.ColorFromHex("#fab387"),
    }
    style := termicolr.NewStyle().FG(myTheme.Primary).Bold().Underline()
    fmt.Println(style.Sprint("Hello, Termicolr!"))
}
```

### 🧱 Projects using Termicolr

- [Asky](https://github.com/patppuccin/asky) – Simple interactive prompt library for Go

### 🪪 License

[MIT License](/LICENSE) - © 2025 Patrick Ambrose