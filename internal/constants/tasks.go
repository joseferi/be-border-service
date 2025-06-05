package constants

const (
	TaskSendEmail    = "email:send"
	TaskGeneratePDF  = "pdf:generate"
	TaskSendWhatsApp = "wa:send"
	TaskHealthCheck  = "health:check"
)

var AllowedTaskTypes = map[string]struct{}{
	TaskSendEmail:    {},
	TaskGeneratePDF:  {},
	TaskSendWhatsApp: {},
	TaskHealthCheck:  {},
}
