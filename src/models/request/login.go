package request

type LoginRequest struct {
	UserName		string	`json:"user_name"`
	Password		string	`json:"user_password"`
	ProjectName     string  `json:"project_name"`
}
