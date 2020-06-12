package divider

const blockType = "divider"

type Block struct {
	Type string `json:"type"`
}

func New() *Block {
	return &Block{
		Type: blockType,
	}
}
