package crossbuild

import "context"

type Builder interface {
	// BeforeBuild before build hook
	BeforeBuild(ctx context.Context) error
	// AfterBuild after build hook
	AfterBuild(ctx context.Context) error
	// BeginBuild start to build
	BeginBuild(ctx context.Context) error
	// BeforeCompress before compress hook
	BeforeCompress(ctx context.Context) error
	// AfterCompress after compress hook
	AfterCompress(ctx context.Context) error
}
