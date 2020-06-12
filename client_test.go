package slack

import (
	"fmt"
	"github.com/tanaka-yui/go-slack-block-kit/blocks"
	"github.com/tanaka-yui/go-slack-block-kit/blocks/section"
	"github.com/tanaka-yui/go-slack-block-kit/utils"
	"testing"
)

const testUrl = "your slack incoming hock url"

func TestWebHock(t *testing.T) {
	slackClient := NewClient(testUrl)
	rootBlock := blocks.NewRootBlock(blocks.WithSize(5)).
		AddDivider().
		AddSection(section.NewBuilder(section.TextBlock{
			Type: section.MarkDown,
			Text: fmt.Sprintf("%s %s", ChannelKeyword, "mention sample"),
		}).Build()).
		AddSection(section.NewBuilder(section.TextBlock{
			Type: section.MarkDown,
			Text: fmt.Sprintf(fmt.Sprintf("%s : %s", Bold("bold text"), "1234567890")),
		}).Build()).
		AddDivider().
		Build()

	println(utils.ToJson(rootBlock))
	if err := slackClient.Send(rootBlock); err != nil {
		t.Fatal(err)
	}
}
