package server

type LoginReq struct {
	Username string `json:"username"`
	PassWord string `json:"password"`
}

type LoginResp struct {
	Token    string `json:"token"`
	UserId   uint64 `json:"userid"`
	UserName string `json:"username"`
}
