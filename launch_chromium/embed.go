package main

import _ "embed"

//go:embed chromium/build/chromium.zip
var embeded []byte
var build_time string
