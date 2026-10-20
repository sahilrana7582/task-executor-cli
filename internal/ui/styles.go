package ui

import "github.com/charmbracelet/lipgloss"

var bannerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#fc036f"))

const bannerArt = `
      _____                    _____
     /\    \                  /\    \
    /::\    \                /::\    \
    \:::\    \              /::::\    \
     \:::\    \            /::::::\    \
      \:::\    \          /:::/\:::\    \
       \:::\    \        /:::/__\:::\    \
       /::::\    \      /::::\   \:::\    \
      /::::::\    \    /::::::\   \:::\    \
     /:::/\:::\    \  /:::/\:::\   \:::\    \
    /:::/  \:::\____\/:::/__\:::\   \:::\____\
   /:::/    \::/    /\:::\   \:::\   \::/    /
  /:::/    / \/____/  \:::\   \:::\   \/____/
 /:::/    /            \:::\   \:::\    \
/:::/    /              \:::\   \:::\____\
\::/    /                \:::\   \::/    /
 \/____/                  \:::\   \/____/
                           \:::\    \
                            \:::\____\
                             \::/    /
                              \/____/

TaskWorker - minimal task executor framework`

// Banner returns the colored ASCII banner.
func Banner() string {
	return bannerStyle.Render(bannerArt)
}
