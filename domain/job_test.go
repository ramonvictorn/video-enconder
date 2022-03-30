package domain_test

import (
	"testing"
	"time"
	"video-enconder/domain"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "CONVERTED", video)

	require.NotNil(t, job)
	require.Nil(t, err)
}
