package iconsdk

import (
	"encoding/json"
	"strconv"

	"github.com/kidmam/icon-sdk-go/providers"
	"github.com/kidmam/icon-sdk-go/utils"
)

func Setendpoint(URI string) {
	providers.Endpoint(URI)
}

func GetLastBlockHeight() (int64, error) {
	var body map[string]interface{}
	err := providers.GeneralJSONRPCReq(&body, "icx_getLastBlock")
	height := int64(body["height"].(float64))

	if err == nil {
		return height, err
	} else {
		return -1, err
	}
}

func GetBlock(value interface{}) (string, error) {
	var result []byte
	var err error
	var body map[string]interface{}

	if utils.IsPredefinedBlockValue(value) {
		// by value
		err = providers.GeneralJSONRPCReq(&body, "icx_getLastBlock")
		if err == nil {
			result, _ = json.Marshal(body)
		}

	} else if utils.IsHexBlockHash(value) {
		// by hash
		params := utils.AssertionString(value)
		err = providers.GeneralJSONRPCReq(&body, "icx_getBlockByHash", params)
		if err == nil {
			result, _ = json.Marshal(body)
		}

	} else if utils.IsBlockHeight(value) {
		// by height
		params := strconv.FormatInt(utils.AssertionInteger(value), 16)
		err = providers.GeneralJSONRPCReq(&body, "icx_getBlockByHeight", params)
		if err == nil {
			result, _ = json.Marshal(body)
		}
	}

	// ...

	return string(result), err
}
