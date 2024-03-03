package models

type NtfyMessage struct {
	Actions  []string `json:"actions,omitempty"`
	Attach   string   `json:"attach,omitempty"`
	Click    string   `json:"click,omitempty"`
	Email    string   `json:"email,omitempty"`
	Priority string   `json:"priority,omitempty"`
	Tags     string   `json:"tags,omitempty"`
	Title    string   `json:"title,omitempty"`
	Message  string   `json:"message,omitempty"`
	Topic    string   `json:"topic,omitempty"`
	Filename string   `json:"filename,omitempty"`
}
