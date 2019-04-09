package providers

import (
	"github.com/kidmam/icon-sdk-go/utils"
	"github.com/ybbus/jsonrpc"
)

var endpointURI string

type ParamsHeight struct {
	Height string `json:"height"`
}

type ParamsHash struct {
	Hash string `json:"hash"`
}

func Endpoint(URI string) {
	endpointURI = URI
}

func GeneralJSONRPCReq(out interface{}, method string, args ...interface{}) error {
	var err error
	apiURI := endpointURI + "/api/v3"
	rpcClient := jsonrpc.NewClient(apiURI)

	if args == nil {
		err = rpcClient.CallFor(&out, method)
	} else if len(utils.AssertionString(args[0])) == 66 {
		err = rpcClient.CallFor(&out, method, &ParamsHash{utils.AssertionString(args[0])})
	} else {
		err = rpcClient.CallFor(&out, method, &ParamsHeight{utils.Add0xPrefix(utils.AssertionString(args[0]))})
	}

	return err
}
