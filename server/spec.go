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

type Storedetails struct {
	Storename string `json:"storename"`
	Location  string `json:"location"`
}

type StoreResp struct {
	StoreId   uint64 `json:"storeid"`
	StoreName string `json:"storename"`
}

type ProductResp struct {
	Name    string  `json:"name"`
	Mrp     float64 `json:"mrp"`
	Price   float64 `json:"price"`
	Barcode string  `json:"barcode"`
}
