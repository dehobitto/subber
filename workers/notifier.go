package workers

import (
	"log"
	"subber/config"
)

type NotificationJob struct {
	Email   string
	Message string
}

type NotifierWorker struct {
	cfg *config.Config
}

func NewNotifierWorker(cfg *config.Config) *NotifierWorker {
	return &NotifierWorker{
		cfg: cfg,
	}
}

func (n *NotifierWorker) Start(jobs <-chan NotificationJob) {
	log.Println("Notifier Worker has started...")

	for job := range jobs {
		log.Printf("SENDING EMAIL %s | TEXT: %s\n", job.Email, job.Message)
	}
}
