//go:build !darwin && !nogui

package gui

import "context"

// followTheme: the Mac's window follows the saved theme (titlebar_darwin.go);
// Windows' title bar takes the page's colour (TintTitleBar), and Linux's is
// the page's or the window manager's.
func (h *host) followTheme(context.Context) {}
