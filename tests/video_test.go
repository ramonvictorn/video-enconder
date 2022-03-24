package tests

import (
	"testing"
	"time"
	"video-enconder/domain"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "another_id"
	video.FilePath = "video/path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
