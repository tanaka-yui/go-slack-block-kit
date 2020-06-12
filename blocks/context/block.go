package context

const blockType = "context"

type (
	Builder interface {
		WithBlockID(blockId string) Builder
		Build() Block
	}
	builder struct {
		blockType string
		elements  []interface{}
		blockID   *string
	}
	Block struct {
		Type     string        `json:"type"`
		Elements []interface{} `json:"elements"`
		BlockID  *string       `json:"block_id,omitempty"`
	}
)

func (b *builder) WithBlockID(blockId string) Builder {
	b.blockID = &blockId
	return b
}

func (b *builder) Build() Block {
	return Block{
		Type:     b.blockType,
		Elements: b.elements,
		BlockID:  b.blockID,
	}
}

func New() Builder {
	return &builder{
		blockType: blockType,
	}
}
