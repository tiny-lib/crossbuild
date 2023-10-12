package target

import "strings"

// the below code was from goreleaser
var (
	targetsCheckList = []string{
		"aixppc64",
		"android386",
		"androidamd64",
		"androidarm",
		"androidarm64",
		"darwinamd64",
		"darwinarm64",
		"dragonflyamd64",
		"freebsd386",
		"freebsdamd64",
		"freebsdarm",
		"freebsdarm64",
		"illumosamd64",
		"iosarm64",
		"jswasm",
		"wasip1wasm",
		"linux386",
		"linuxamd64",
		"linuxarm",
		"linuxarm64",
		"linuxppc64",
		"linuxppc64le",
		"linuxmips",
		"linuxmipsle",
		"linuxmips64",
		"linuxmips64le",
		"linuxs390x",
		"linuxriscv64",
		"linuxloong64",
		"netbsd386",
		"netbsdamd64",
		"netbsdarm",
		"openbsd386",
		"openbsdamd64",
		"openbsdarm",
		"openbsdarm64",
		"plan9386",
		"plan9amd64",
		"plan9arm",
		"solarisamd64",
		"windowsarm",
		"windowsarm64",
		"windows386",
		"windowsamd64",
	}

	validGoos = []string{
		"aix",
		"android",
		"darwin",
		"dragonfly",
		"freebsd",
		"illumos",
		"ios",
		"js",
		"linux",
		"netbsd",
		"openbsd",
		"plan9",
		"solaris",
		"windows",
		"wasip1",
	}

	validGoarch = []string{
		"386",
		"amd64",
		"arm",
		"arm64",
		"mips",
		"mips64",
		"mips64le",
		"mipsle",
		"ppc64",
		"ppc64le",
		"s390x",
		"wasm",
		"riscv64",
		"loong64",
	}

	validGoarm   = []string{"5", "6", "7"}
	validGomips  = []string{"hardfloat", "softfloat"}
	validGoamd64 = []string{"v1", "v2", "v3", "v4"}
)

// getBuildTarget returns the build target for the given target and build mode.
func getBuildExt(target, buildMode string) string {
	switch buildMode {
	case "c-shared":
		if strings.Contains(target, "darwin") {
			return ".dylib"
		}
		if strings.Contains(target, "windows") {
			return ".dll"
		}
		return ".so"
	case "c-archive":
		if strings.Contains(target, "windows") {
			return ".lib"
		}
		return ".a"
	}

	if target == "js_wasm" {
		return ".wasm"
	}

	if strings.Contains(target, "windows") {
		return ".exe"
	}

	return ""
}
