package image

const blockType = "image"

type (
	Builder interface {
		WithTitle(title string) Builder
		WithBlockID(blockId string) Builder
		Build() Block
	}
	builder struct {
		blockType string
		imageUrl  string
		altText   string
		title     *string
		blockID   *string
	}
	Block struct {
		Type     string  `json:"type"`
		ImageUrl string  `json:"image_url"`
		AltText  string  `json:"alt_text"`
		Title    *string `json:"title,omitempty"`
		BlockID  *string `json:"block_id,omitempty"`
	}
)

func (b *builder) WithTitle(title string) Builder {
	b.title = &title
	return b
}

func (b *builder) WithBlockID(blockId string) Builder {
	b.blockID = &blockId
	return b
}

func (b *builder) Build() Block {
	return Block{
		Type:     b.blockType,
		ImageUrl: b.imageUrl,
		AltText:  b.altText,
		Title:    b.title,
		BlockID:  b.blockID,
	}
}

func NewBuilder(imageUrl string, altText string) Builder {
	return &builder{
		blockType: blockType,
		imageUrl:  imageUrl,
		altText:   altText,
	}
}
