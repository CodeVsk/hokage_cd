package git

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileContentFromGitHub(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	content, err := GetFileContentFromGitHub(ctx, repoOwner, repoName, filePath, fileName, repoToken)
	
	assert.Nil(t, err)
	assert.NotEmpty(t, content)
}