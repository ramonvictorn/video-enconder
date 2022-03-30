package repositories_test

import (
	"testing"
	"time"
	"video-enconder/application/repositories"
	"video-enconder/domain"
	"video-enconder/framework/database"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	// defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	videoFound, err := repo.Find(video.ID)

	require.NotEmpty(t, videoFound.ID)
	require.Nil(t, err)
	require.Equal(t, video.ID, videoFound.ID)
}
