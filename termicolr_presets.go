package termicolr

var ThemeDefault = Theme{
	// Base
	Background:    ColorFromANSI(0),  // black
	BackgroundAlt: ColorFromANSI(8),  // bright black (gray)
	Foreground:    ColorFromANSI(15), // bright white
	ForegroundAlt: ColorFromANSI(7),  // white

	// Emphasis
	Primary:   ColorFromANSI(4), // blue
	Secondary: ColorFromANSI(5), // magenta/purple
	Accent:    ColorFromANSI(6), // cyan (close enough for accent)
	Highlight: ColorFromANSI(3), // yellow
	Muted:     ColorFromANSI(8), // gray

	// Semantic
	Red:    ColorFromANSI(9),  // bright red
	Green:  ColorFromANSI(10), // bright green
	Yellow: ColorFromANSI(11), // bright yellow
	Blue:   ColorFromANSI(12), // bright blue
	Purple: ColorFromANSI(13), // bright magenta
	Orange: ColorFromANSI(3),  // reuse yellow as orange (ANSI has no native orange)
}

var ThemeCatppuccinMocha = Theme{
	Background:    ColorFromHex("#1e1e2e"), // base
	BackgroundAlt: ColorFromHex("#181825"), // mantle
	Foreground:    ColorFromHex("#cdd6f4"), // text
	ForegroundAlt: ColorFromHex("#bac2de"), // subtext1

	Primary:   ColorFromHex("#cba6f7"), // mauve
	Secondary: ColorFromHex("#f2cdcd"), // flamingo
	Accent:    ColorFromHex("#fab387"), // peach
	Highlight: ColorFromHex("#f5e0dc"), // rosewater
	Muted:     ColorFromHex("#6c7086"), // overlay0

	Red:    ColorFromHex("#f38ba8"),
	Green:  ColorFromHex("#a6e3a1"),
	Yellow: ColorFromHex("#f9e2af"),
	Blue:   ColorFromHex("#89b4fa"),
	Purple: ColorFromHex("#cba6f7"),
	Orange: ColorFromHex("#fab387"),
}

var ThemeCatppuccinLatte = Theme{
	Background:    ColorFromHex("#eff1f5"), // base
	BackgroundAlt: ColorFromHex("#e6e9ef"), // mantle
	Foreground:    ColorFromHex("#4c4f69"), // text
	ForegroundAlt: ColorFromHex("#5c5f77"), // subtext1

	Primary:   ColorFromHex("#8839ef"), // mauve
	Secondary: ColorFromHex("#ea76cb"), // flamingo
	Accent:    ColorFromHex("#fe640b"), // peach
	Highlight: ColorFromHex("#dc8a78"), // rosewater
	Muted:     ColorFromHex("#9ca0b0"), // overlay0

	Red:    ColorFromHex("#d20f39"),
	Green:  ColorFromHex("#40a02b"),
	Yellow: ColorFromHex("#df8e1d"),
	Blue:   ColorFromHex("#1e66f5"),
	Purple: ColorFromHex("#8839ef"),
	Orange: ColorFromHex("#fe640b"),
}

var ThemeGruvboxDark = Theme{
	Background:    ColorFromHex("#282828"),
	BackgroundAlt: ColorFromHex("#3c3836"),
	Foreground:    ColorFromHex("#ebdbb2"),
	ForegroundAlt: ColorFromHex("#d5c4a1"),

	Primary:   ColorFromHex("#b16286"), // purple
	Secondary: ColorFromHex("#d3869b"), // pinkish
	Accent:    ColorFromHex("#fe8019"), // orange
	Highlight: ColorFromHex("#fabd2f"), // yellow
	Muted:     ColorFromHex("#a89984"),

	Red:    ColorFromHex("#fb4934"),
	Green:  ColorFromHex("#b8bb26"),
	Yellow: ColorFromHex("#fabd2f"),
	Blue:   ColorFromHex("#83a598"),
	Purple: ColorFromHex("#d3869b"),
	Orange: ColorFromHex("#fe8019"),
}

var ThemeTokyoNight = Theme{
	Background:    ColorFromHex("#1a1b26"),
	BackgroundAlt: ColorFromHex("#24283b"),
	Foreground:    ColorFromHex("#c0caf5"),
	ForegroundAlt: ColorFromHex("#a9b1d6"),

	Primary:   ColorFromHex("#7aa2f7"), // blue
	Secondary: ColorFromHex("#bb9af7"), // purple
	Accent:    ColorFromHex("#7dcfff"), // cyan
	Highlight: ColorFromHex("#e0af68"), // yellow/orange
	Muted:     ColorFromHex("#565f89"),

	Red:    ColorFromHex("#f7768e"),
	Green:  ColorFromHex("#9ece6a"),
	Yellow: ColorFromHex("#e0af68"),
	Blue:   ColorFromHex("#7aa2f7"),
	Purple: ColorFromHex("#bb9af7"),
	Orange: ColorFromHex("#ff9e64"),
}

var ThemeKanagawa = Theme{
	Background:    ColorFromHex("#1f1f28"),
	BackgroundAlt: ColorFromHex("#2a2a37"),
	Foreground:    ColorFromHex("#dcd7ba"),
	ForegroundAlt: ColorFromHex("#c8c093"),

	Primary:   ColorFromHex("#7e9cd8"), // blue
	Secondary: ColorFromHex("#957fb8"), // purple
	Accent:    ColorFromHex("#ffa066"), // orange
	Highlight: ColorFromHex("#e6c384"), // yellow
	Muted:     ColorFromHex("#727169"),

	Red:    ColorFromHex("#c34043"),
	Green:  ColorFromHex("#98bb6c"),
	Yellow: ColorFromHex("#c0a36e"),
	Blue:   ColorFromHex("#7e9cd8"),
	Purple: ColorFromHex("#957fb8"),
	Orange: ColorFromHex("#ffa066"),
}

var ThemeDracula = Theme{
	Background:    ColorFromHex("#282a36"),
	BackgroundAlt: ColorFromHex("#44475a"),
	Foreground:    ColorFromHex("#f8f8f2"),
	ForegroundAlt: ColorFromHex("#e2e2dc"),

	Primary:   ColorFromHex("#bd93f9"), // purple
	Secondary: ColorFromHex("#ff79c6"), // pink
	Accent:    ColorFromHex("#50fa7b"), // green
	Highlight: ColorFromHex("#f1fa8c"), // yellow
	Muted:     ColorFromHex("#6272a4"),

	Red:    ColorFromHex("#ff5555"),
	Green:  ColorFromHex("#50fa7b"),
	Yellow: ColorFromHex("#f1fa8c"),
	Blue:   ColorFromHex("#8be9fd"),
	Purple: ColorFromHex("#bd93f9"),
	Orange: ColorFromHex("#ffb86c"),
}
