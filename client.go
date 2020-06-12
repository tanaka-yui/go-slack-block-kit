package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tanaka-yui/go-slack-block-kit/blocks"
	"github.com/tanaka-yui/go-slack-block-kit/utils"
	"net/http"
	"strconv"
)

type (
	Client interface {
		Send(block blocks.Builder) error
	}
	clientImpl struct {
		webHookUrl string
		httpClient *http.Client
	}
)

func (c *clientImpl) Send(block blocks.Block) error {
	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(&block); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.webHookUrl, &requestBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Content-Length", strconv.Itoa(requestBody.Len()))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		e := errors.New(fmt.Sprintf("slack notification request error. status code: %d, body: %v", res.StatusCode, utils.ToString(res.Body)))
		return e
	}
	defer func() {
		_ = res.Body.Close()
	}()

	return nil
}

func NewClient(url string) *clientImpl {
	return &clientImpl{
		webHookUrl: url,
		httpClient: http.DefaultClient,
	}
}
