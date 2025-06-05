package constants

type QueuePriority string

const (
	Critical QueuePriority = "critical"
	Default  QueuePriority = "default"
	Low      QueuePriority = "low"
)
