package pattern

import (
	"fmt"
	"net/http"
)

//
// 桥接模式
//    是用于把抽象化与实现化解耦，使得二者可以独立变化。这种类型的设计模式属于结构型模式，
//    它通过提供抽象化和实现化之间的桥接结构，来实现二者的解耦。
//

//请求接口
type Request interface {
	HttpRequest() (*http.Request, error)
}

//客户端
type Client struct {
	Client *http.Client
}

func (c *Client) Query(req Request) (resp *http.Response, err error) {
	httpreq,_:=req.HttpRequest()
	resp, err = c.Client.Do(httpreq)
	return
}

type CdnRequest struct {
}

func (cdn *CdnRequest) HttpRequest() (*http.Request,  error) {
	return http.NewRequest("GET", "/cdn", nil)
}

type LiveRequest struct {
}

func (cdn *LiveRequest) HttpRequest() (*http.Request, error) {
	return http.NewRequest("GET", "/live", nil)
}

func TestBridge() {
	client := &Client{http.DefaultClient}

	cdnReq := &CdnRequest{}
	fmt.Println(client.Query(cdnReq))

	liveReq := &LiveRequest{}
	fmt.Println(client.Query(liveReq))
}
