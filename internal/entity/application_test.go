package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	a, err := NewApplication("Feature Flag", "feature_flag", "http://github.com/codevsk/feature_flag", "codevsk")

	assert.Nil(t, err)
	assert.NotNil(t, a)
	assert.NotEmpty(t, a.ID)
	assert.Equal(t, "Feature Flag", a.Name)
	assert.Equal(t, "feature_flag", a.Slug)
	assert.Equal(t, "http://github.com/codevsk/feature_flag", a.RepositoryUrl)
	assert.Equal(t, a.IsEnabled, true)
	assert.Equal(t, a.AutoRollback, false)
	assert.Equal(t, "codevsk", a.CreatedBy)
	assert.NotEmpty(t, a.CreatedAt)
}