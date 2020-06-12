package section

const blockType = "section"

type (
	Builder interface {
		WithBlockID(blockId string) Builder
		WithFields(fields *[]interface{}) Builder
		WithAccessory(accessory *[]interface{}) Builder
		Build() Block
	}
	builder struct {
		blockType string
		text      TextBlock
		blockID   *string
		fields    *[]interface{}
		accessory *[]interface{}
	}
	Block struct {
		Type      string         `json:"type"`
		Text      TextBlock      `json:"text"`
		BlockID   *string        `json:"block_id,omitempty"`
		Fields    *[]interface{} `json:"fields,omitempty"`
		Accessory *[]interface{} `json:"accessory,omitempty"`
	}
)

type TextType = string

const (
	PlainText TextType = "plain_text"
	MarkDown  TextType = "mrkdwn"
)

type TextBlock struct {
	Type  TextType `json:"type"`
	Text  string   `json:"text"`
	Emoji bool     `json:"emoji,omitempty"`
}

func (b *builder) WithBlockID(blockId string) Builder {
	b.blockID = &blockId
	return b
}

func (b *builder) WithFields(fields *[]interface{}) Builder {
	b.fields = fields
	return b
}

func (b *builder) WithAccessory(accessory *[]interface{}) Builder {
	b.accessory = accessory
	return b
}

func (b *builder) Build() Block {
	return Block{
		Type:      b.blockType,
		Text:      b.text,
		BlockID:   b.blockID,
		Fields:    b.fields,
		Accessory: b.accessory,
	}
}

func NewBuilder(text TextBlock) Builder {
	return &builder{
		blockType: blockType,
		text:      text,
	}
}
