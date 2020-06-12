package file

const blockType = "file"

type (
	Builder interface {
		WithBlockID(blockId string) Builder
		Build() Block
	}
	builder struct {
		blockType  string
		externalID string
		source     string
		blockID    *string
	}
	Block struct {
		Type       string  `json:"type"`
		ExternalID string  `json:"external_id"`
		Source     string  `json:"source"`
		BlockID    *string `json:"block_id,omitempty"`
	}
)

func (b *builder) WithBlockID(blockId string) Builder {
	b.blockID = &blockId
	return b
}

func (b *builder) Build() Block {
	return Block{
		Type:       b.blockType,
		ExternalID: b.externalID,
		Source:     b.source,
		BlockID:    b.blockID,
	}
}

func NewBuilder(externalId string, source string) Builder {
	return &builder{
		blockType:  blockType,
		externalID: externalId,
		source:     source,
	}
}
