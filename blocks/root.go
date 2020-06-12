package blocks

import (
	"github.com/tanaka-yui/go-slack-block-kit/blocks/actions"
	"github.com/tanaka-yui/go-slack-block-kit/blocks/divider"
	"github.com/tanaka-yui/go-slack-block-kit/blocks/file"
	"github.com/tanaka-yui/go-slack-block-kit/blocks/input"
	"github.com/tanaka-yui/go-slack-block-kit/blocks/section"
)

type options struct {
	size uint
}

type Option func(*options)

func WithSize(size uint) Option {
	return func(o *options) {
		o.size = size
	}
}

type (
	Builder interface {
		AddActions(block actions.Block) Builder
		AddSection(block section.Block) Builder
		AddInput(block input.Block) Builder
		AddFile(block file.Block) Builder
		AddDivider() Builder
		Build() Block
	}
	builder struct {
		rootBlocks []interface{}
	}
	Block struct {
		RootBlocks []interface{} `json:"blocks,omitempty"`
	}
)

func (r *builder) AddActions(block actions.Block) Builder {
	r.rootBlocks = append(r.rootBlocks, &block)
	return r
}

func (r *builder) AddSection(block section.Block) Builder {
	r.rootBlocks = append(r.rootBlocks, &block)
	return r
}

func (r *builder) AddInput(block input.Block) Builder {
	r.rootBlocks = append(r.rootBlocks, block)
	return r
}

func (r *builder) AddFile(block file.Block) Builder {
	r.rootBlocks = append(r.rootBlocks, block)
	return r
}

func (r *builder) AddDivider() Builder {
	r.rootBlocks = append(r.rootBlocks, divider.New())
	return r
}

func (r *builder) Build() Block {
	return Block{
		RootBlocks: r.rootBlocks,
	}
}

func NewRootBlock(opts ...Option) *builder {
	dopt := &options{}
	for i := range opts {
		opts[i](dopt)
	}
	return &builder{
		rootBlocks: make([]interface{}, 0, dopt.size),
	}
}
