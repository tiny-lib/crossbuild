package crossbuild

import "context"

type LocalBuild struct {
	Goos       []string `yaml:"goos,omitempty" json:"goos,omitempty"`
	Goarch     []string `yaml:"goarch,omitempty" json:"goarch,omitempty"`
	Targets    []string `yaml:"targets,omitempty" json:"targets,omitempty"`
	GoProxy    string   `yaml:"go_proxy,omitempty" json:"go_proxy,omitempty"`
	Ldflags    string   `yaml:"ldflags,omitempty" json:"ldflags,omitempty"`
	CgoEnabled bool     `yaml:"cgo_enabled,omitempty" json:"cgo_enabled,omitempty"`

	OutPut         string `yaml:"output,omitempty" json:"output,omitempty"`
	CompressOutput bool   `yaml:"compress_output,omitempty" json:"compress_output,omitempty"`

	BeforeBuildFn    func(ctx context.Context) error
	AfterBuildFn     func(ctx context.Context) error
	BeforeCompressFn func(ctx context.Context) error
	AfterCompressFn  func(ctx context.Context) error
}
