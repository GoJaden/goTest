package main

import (
	"encoding/json"
	"fmt"
)

type ProductsResult struct {
	Status int         `json:"status"`
	Data   ProductData `json:"data"`
	Msg    string      `json:"msg"`
}
type ProductData struct {
	ProductList []*ProSku `json:"list"`
}

type ProSku struct {
	ProductBaseInfo ProductBaseInfo `json:"productBaseInfo"`
	SkuList         []*SkuBaseInfo  `json:"skuList"`
}
type SkuBaseInfo struct {
	Id                 uint64 `json:"id"`
	MasterName         string `json:"masterName"`
	SlaveName          string `json:"slaveName"`
	Image              string `json:"image"`
	Price              uint32 `json:"price"`
	Score              uint32 `json:"score"`
	Inventory          uint32 `json:"inventory"`
	SalesCount         uint32 `json:"salesCount"`
	MarketPrice        uint32 `json:"marketPrice"`
	SlavePrice         uint32 `json:"slavePrice"`
	AvailableInventory uint32 `json:"availableInventory"`
}

type ProductBaseInfo struct {
	Id              uint64 `json:"id"`
	IsMember        uint32 `json:"isMember"`
	MasterName      string `json:"masterName"`
	SlaveName       string `json:"slaveName"`
	SupplierId      uint64 `json:"supplierId"`
	DefaultSkuId    uint64 `json:"defaultSkuId"`
	AttrSet         uint32 `json:"attrSet"`
	SalesCount      uint32 `json:"salesCount"`
	ProductStatus   uint32 `json:"productStatus"`
	TotailInventory uint32 `json:"totailInventory"`
}


func main() {
	var bb uint8 = 258
	fmt.Println(bb)

	productsResult := &ProductsResult{}
	data := "{\"status\":6602002,\"data\":\"\",\"msg\":\"商品不存在\"}"
	err := json.Unmarshal([]byte(data), productsResult)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(productsResult)
}
