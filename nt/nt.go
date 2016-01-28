package nt

// #cgo CPPFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -L${SRCDIR}/libs -lntcore
// #include "networktables/NetworkTable.h"
import "C"

import (
	"unsafe"
)
