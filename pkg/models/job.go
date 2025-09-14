package models

import (
	"empregufu/pkg/config"
	"time"
	"gorm.io/gorm"
)

var db *gorm.DB

type Job struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Title string `gorm:"not null" json:"title" required:"true"`
	Description string `gorm:"not null" json:"description"`
	Company string `gorm:"not null" json:"company"`
	Location string `gorm:"not null" json:"location" required:"true"`
	Salary float64 `gorm:"not null" json:"salary"`
	PostedAt time.Time `gorm:"not null" json:"posted_at"`
}

func Init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Job{})
}

func CreateJob(job *Job) (*Job, error) {
	err := db.Create(job).Error
	if err != nil {
		return nil, err
	}
	return job, nil	
}
func GetJobs() ([]Job, error) {
	var jobs []Job
	err := db.Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return jobs, nil	
}
func GetJob(id uint) (*Job, error) {
	var job Job
	err := db.First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}
func UpdateJob(job *Job) (*Job, error) {
	err := db.Save(job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}
func DeleteJob(id uint) error {
	err := db.Delete(&Job{}, id).Error
	if err != nil {
		return err
	}
	return nil
}


