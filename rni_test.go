package baidu

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ai-sdk/baidu/pkg/rni"
)

func TestSDK_Rni_H5SASSToken(t *testing.T) {
	token, err := sdk.Rni().H5SASSToken("10653")
	assert.Nil(t, err)
	url := rni.H5SASSRedirect(token, "https://www.taobao.com", "https://www.baidu.com")
	fmt.Println(url)
	fmt.Println(token)
	res, err := sdk.Rni().H5SASSResult("XLO819EBhUwj6u0f3YPVldnH")
	assert.Nil(t, err)
	fmt.Println(res)
}
