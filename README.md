# go-slack-block-kit

- this library is [slack Block Kit](https://api.slack.com/block-kit) wrapper.

- use pure golang library 

## Usage

```go
package main

import "fmt"

import "github.com/tanaka-yui/go-slack-block-kit"
import "github.com/tanaka-yui/go-slack-block-kit/blocks"
import "github.com/tanaka-yui/go-slack-block-kit/blocks/section"

const testUrl = "your slack incoming hock url"

func main() {
    slackClient := slack.NewClient(testUrl)

    rootBlock := blocks.NewRootBlock(blocks.WithSize(5)).
        AddDivider().
        AddSection(section.NewBuilder(section.TextBlock{
            Type: section.MarkDown,
            Text: fmt.Sprintf("%s %s", slack.ChannelKeyword, "@channel mention sample"),
        }).Build()).
        AddSection(section.NewBuilder(section.TextBlock{
            Type: section.MarkDown,
            Text: fmt.Sprintf(fmt.Sprintf("%s : %s", slack.Bold("bold text"), "1234567890")),
        }).Build()).
        AddDivider().
        Build()

    if err := slackClient.Send(rootBlock); err != nil {
        println(err.Error())
    }
}
```

## License
Licensed under the Apache License, Version 2.0: http://www.apache.org/licenses/LICENSE-2.0