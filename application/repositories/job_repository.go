package repositories

import (
	"fmt"
	"video-enconder/domain"

	"gorm.io/gorm"
)

type JobRepository interface {
	Insert(job *domain.Job) *domain.Job
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDb struct {
	Db *gorm.DB
}

func (repo JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (repo JobRepositoryDb) Find(id string) (*domain.Job, error) {
	var job domain.Job
	repo.Db.Preload("Jobs").First(&job, "id = ?", id)

	if job.ID == "" {
		return nil, fmt.Errorf("JOB NOT EXIST")
	}

	return &job, nil
}

func (repo JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Save(&job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}
