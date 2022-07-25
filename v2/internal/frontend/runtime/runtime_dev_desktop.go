//go:build dev || hybrid || server || bindings || (!dev && !production && !bindings)
// +build dev hybrid server bindings !dev,!production,!bindings

package runtime

import _ "embed"

//go:embed runtime_dev_desktop.js
var RuntimeDesktopJS []byte
