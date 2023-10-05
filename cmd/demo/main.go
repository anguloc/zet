package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.bitget.com/v1/finance/trend/product/list"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	//req.Header.Set("Cookie", "cf_chl_2=e0348379c623e39; cf_clearance=_PX2oM8ikiRzHxjvdXwA.AwijXomPo8oViMMEwoXk3E-1695459500-0-1-1e08f744.4e3d7682.4909a8a4-150.0.0; __cf_bm=JP3D1ifHWq3n9aib6M6du8YgrKTmH4fEyySlguh8Rxo-1695459503-0-Ab/BGH02CTGLxmftRYNZ/MtpPfi49F9+NYzhwF25E5X0ZOWpNz9UUKS0gX4S2oz3biTOionPsJ9YIzKnJiKwA1E=; _cfuvid=O8yulUvZrfUWBlyvo_3tFPCDYCb8DIdovN.lTG8T.jo-1695459503118-0-604800000")

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		println(123)
		panic(err)
	}
	defer rsp.Body.Close()
	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		println(456)
		panic(err)
	}
	fmt.Println(string(data))
	println("httpcode:", rsp.StatusCode)
}
