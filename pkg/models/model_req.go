package models

type BatchUpdateRequest struct {
	Updates []UpdateRequest `json:"updates" binding:"required,dive"`
}

type UpdateRequest struct {
	Range  string          `json:"range" binding:"required"`
	Value  string          `json:"value" binding:"required"`
	Values [][]interface{} `json:"values,omitempty"`
}
type Properties struct {
	Title string `json:"title" binding:"required"`
}

type CreateRequest struct {
	Properties Properties `json:"properties" binding:"required"`
}

type DeleteRequest struct {
	Range string `json:"range"`
}

type ShareRequest struct {
	Email        string `json:"email"`
	SendEmail    bool   `json:"sendEmail,omitempty"`
	EmailMessage string `json:"emailMessage,omitempty"`
}
