package apiserver

type userSaveRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type userLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
