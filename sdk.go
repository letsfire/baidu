package baidu

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/letsfire/baidu/pkg/face"
	"github.com/letsfire/baidu/pkg/ocr"
	"github.com/letsfire/baidu/pkg/rni"
)

var client = resty.New()

const tokenURL = "https://aip.baidubce.com/oauth/2.0/token"

type SDK struct {
	apiKey       string        // 百度AI client_id
	secretKey    string        // 百度AI client_secret
	localCache   localCache    // Token本地缓存
	tokenStorage TokenStorage  // Token共享缓存
	safeDuration time.Duration // Token容错时长, 也是本地缓存有效时长
}

func (sdk *SDK) Rni() rni.SDK {
	return rni.NewSDK(sdk.token())
}

func (sdk *SDK) Ocr() ocr.SDK {
	return ocr.NewSDK(sdk.token())
}

func (sdk *SDK) Face() face.SDK {
	return face.NewSDK(sdk.token())
}

func (sdk *SDK) token() (v string, err error) {
	if v = sdk.localCache.get(); v != "" {
		return v, nil
	}
	defer func() {
		if err == nil && v != "" {
			exp := time.Now().Add(sdk.safeDuration)
			sdk.localCache = localCache{value: v, expiredAt: exp}
		}
	}()
	key := sdk.cacheKey()
	if v, err = sdk.tokenStorage.Fetch(key); err != nil {
		return "", err
	} else if v == "" {
		var exp time.Time
		if v, exp, err = sdk.requestToken(); err != nil {
			return "", err
		}
		_ = sdk.tokenStorage.Store(key, v, exp.Sub(time.Now()))
	}
	return
}

func (sdk *SDK) cacheKey() string {
	return fmt.Sprintf("%s@%s", sdk.apiKey, sdk.secretKey)
}

func (sdk *SDK) requestToken() (string, time.Time, error) {
	args := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     sdk.apiKey,
		"client_secret": sdk.secretKey,
	}
	res, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetQueryParams(args).Get(tokenURL)
	if err != nil {
		return "", time.Now(), err
	}
	var tokenRes = new(tokenResponse)
	err = json.Unmarshal(res.Body(), tokenRes)
	if err == nil {
		err = tokenRes.Error()
	}
	d := time.Duration(tokenRes.ExpiresIn) * time.Second
	return tokenRes.AccessToken, time.Now().Add(d - sdk.safeDuration), err
}

func New(apiKey, secretKey string, storage TokenStorage) *SDK {
	return &SDK{
		apiKey:       apiKey,
		secretKey:    secretKey,
		tokenStorage: storage,
		safeDuration: 5 * time.Minute,
	}
}
