package server

import (
	"testing"

	"github.com/mattn/lile-example/gene9go"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestTranslate(t *testing.T) {
	ctx := context.Background()
	req := &gene9go.Request{}

	res, err := cli.Translate(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
