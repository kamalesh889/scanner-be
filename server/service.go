package server

import (
	"barcode/model"
	"barcode/utility"
	"strconv"
)

func (s *server) LoginService(req *LoginReq) (*LoginResp, error) {

	userReq := &model.User{
		Username: req.Username,
		Password: req.PassWord,
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

func (s *server) Createstore(req *Storedetails) (*StoreResp, error) {

	store := &model.Store{
		Name:     req.Storename,
		Location: req.Location,
		Schema:   req.Storename + "_schema",
	}

	store, err := s.db.CreateStore(store)
	if err != nil {
		return nil, err
	}

	resp := &StoreResp{
		StoreId:   store.Id,
		StoreName: store.Name,
	}

	return resp, nil

}

func (s *server) GetProductService(barcode string, code string) (*ProductResp, error) {

	storeCode, _ := strconv.Atoi(code)

	store, err := s.db.GetStore(uint64(storeCode))
	if err != nil {
		return nil, err
	}

	product, err := s.db.GetProduct(barcode, store.Schema)
	if err != nil {
		return nil, err
	}

	productPrice := product.Mrp - (product.Mrp * product.Discount)

	resp := &ProductResp{
		Name:    product.Name,
		Price:   productPrice,
		Barcode: barcode,
	}

	return resp, nil

}
