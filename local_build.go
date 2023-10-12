package crossbuild

import "context"

type LocalBuild struct {
	Goos       []string `yaml:"goos,omitempty" json:"goos,omitempty"`
	Goarch     []string `yaml:"goarch,omitempty" json:"goarch,omitempty"`
	Targets    []string `yaml:"targets,omitempty" json:"targets,omitempty"`
	GoProxy    string   `yaml:"go_proxy,omitempty" json:"go_proxy,omitempty"`
	Ldflags    string   `yaml:"ldflags,omitempty" json:"ldflags,omitempty"`
	BuildMode  string   `yaml:"build_mode,omitempty" json:"build_mode,omitempty"`
	CgoEnabled bool     `yaml:"cgo_enabled,omitempty" json:"cgo_enabled,omitempty"`

	OutPut         string `yaml:"output,omitempty" json:"output,omitempty"`
	CompressOutput bool   `yaml:"compress_output,omitempty" json:"compress_output,omitempty"`

	BeforeBuildFn    func(ctx context.Context) error
	AfterBuildFn     func(ctx context.Context) error
	BeforeCompressFn func(ctx context.Context) error
	AfterCompressFn  func(ctx context.Context) error
}

func (l LocalBuild) BeforeBuild(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (l LocalBuild) AfterBuild(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (l LocalBuild) BeginBuild(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (l LocalBuild) BeforeCompress(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (l LocalBuild) AfterCompress(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
