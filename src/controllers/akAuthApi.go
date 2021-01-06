package controllers

type AuthUserRequest struct {
	UserName	string
	UserPwd		string
	XToken 		string
}

type AuthUserResponse struct {
	Status		bool
	Msg 		string
}
