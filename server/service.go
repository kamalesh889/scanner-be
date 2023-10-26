package server

import (
	"barcode/model"
	"barcode/utility"
)

func (s *server) LoginService(req *LoginReq) (*LoginResp, error) {

	userReq := &model.User{
		UserName: req.Username,
		PassWord: req.PassWord,
	}

	user, err := s.db.GetUser(userReq)
	if err != nil {
		return nil, err
	}

	tokenReq := &utility.TokenReq{
		Id: user.Id,
	}

	//need to check , should work with same instance as i am passing pointer

	token, err := tokenReq.CreateJwtToken()
	if err != nil {
		return nil, err
	}

	resp := &LoginResp{
		Token:    token.Token,
		UserId:   user.Id,
		UserName: req.Username,
	}

	return resp, nil

}
