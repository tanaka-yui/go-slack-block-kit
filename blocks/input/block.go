package input

import "github.com/tanaka-yui/go-slack-block-kit/blocks/section"

const blockType = "input"

type (
	Builder interface {
		WithBlockID(blockId string) Builder
		WithHint(hint section.TextBlock) Builder
		WithOptional(b bool) Builder
		Build() Block
	}
	builder struct {
		blockType string
		label     string
		element   interface{}
		blockID   *string
		hint      *section.TextBlock
		optional  *bool
	}
	Block struct {
		Type     string             `json:"type"`
		Label    string             `json:"label"`
		Element  interface{}        `json:"element"`
		BlockID  *string            `json:"block_id,omitempty"`
		Hint     *section.TextBlock `json:"hint,omitempty"`
		Optional *bool              `json:"optional,omitempty"`
	}
)

func (b *builder) WithBlockID(blockId string) Builder {
	b.blockID = &blockId
	return b
}

func (b *builder) WithHint(hint section.TextBlock) Builder {
	b.hint = &hint
	return b
}

func (b *builder) WithOptional(optional bool) Builder {
	b.optional = &optional
	return b
}

func (b *builder) Build() Block {
	return Block{
		Type:     b.blockType,
		Label:    b.blockType,
		Element:  b.element,
		BlockID:  b.blockID,
		Hint:     b.hint,
		Optional: b.optional,
	}
}

func NewBuilder(label string, element interface{}) Builder {
	return &builder{
		blockType: blockType,
		label:     label,
		element:   element,
	}
}
