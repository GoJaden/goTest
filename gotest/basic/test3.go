package main

import (
	httputil "common-lib/http-util"
	"fmt"
)

func main() {
	s := fmt.Sprintf("http://business-examine.sdyxmall/api/order/list?userId=%v&pageSize=%v&pageNum=%v&orderStatus=%v",
		691719, 10, 1, 0)
	fmt.Println(s)
	resp, _, _ := httputil.NewHttpClient().DoGet(fmt.Sprintf("http://business-examine.sdyxmall/api/order/list?userId=%v&pageSize=%v&pageNum=%v&orderStatus=%v",
		691719, 10, 1, 0), "").Response()
	fmt.Println(resp)
}
