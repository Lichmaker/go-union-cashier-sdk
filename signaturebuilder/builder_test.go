package signaturebuilder_test

import (
	"testing"

	"github.com/lichmaker/go-union-cashier-sdk/signaturebuilder"
)

func TestSign(t *testing.T) {
	priKeyPath := "./testfile/private.key"
	testCase := []struct {
		body     string
		wantSign string
	}{
		{
			body:     `app_id=5dc387b32c07871d371334e9c45120ba&biz_content=%7B%22sndDt%22%3A%2220220527150857%22%2C%22busiMerNo%22%3A%22100000000000002%22%2C%22msgBody%22%3A%7B%22merOrdrNo%22%3A%2210000000000000220220527150857079%22%2C%22trxTtlAmt%22%3A%221%22%2C%22bizFunc%22%3A%22%22%2C%22bizType%22%3A%2210101%22%2C%22subject%22%3A%22%E6%B5%8B%E8%AF%95%22%2C%22ordrDesc%22%3A%22%E6%B5%8B%E8%AF%95123%22%2C%22oprId%22%3A%22%22%2C%22timeoutExpress%22%3A%22%22%2C%22trxChnlType%22%3A%2202%22%2C%22userNo%22%3A%22123456%22%2C%22userType%22%3A%22USER_ID%22%2C%22remark1%22%3A%22123%22%7D%2C%22notifyUrl%22%3A%22https%3A%2F%2Fbaidu.com%22%7D&method=gnete.upbc.cashier.trade.prePay&sign_alg=1&timestamp=2022-05-27+15%3A22%3A07&v=1.0.1`,
			wantSign: "e6728f990ca17381107a642d215960207ba380dd7755b496608c85e2dc194dc21804b8e616afbdde7dd48fd6902a5bf6fbaefdfbb5b7f166c9a7ee20a5f8c6bfd34981ad2b7081b485424151c3921339cfc45a3902d07ac3bc2b08c7a6514434dedc7824e157c243148fd7e470585f638c6399724e19e9b72110d1404135a9dd39ebf8a4f2139dc90b480d2d7e795e80433f3c6d1d68de84f75b0a9460db753baad9c86740e518d0f027d4c0266e133f37ac26ad5a7144c3df61438f38290d4680c1c75d94153f3159c261cea9b56c394593c97d08c82461d8429f7b78eaf44164e6940902cd22935a316ea8e4a4fda14e16a0c18adae912d251eae7b5cd71b2",
		},
	}
	for _, tc := range testCase {
		buildSign, err := signaturebuilder.Sign(tc.body, priKeyPath)
		if err != nil {
			t.Error(err)
		}
		if buildSign != tc.wantSign {
			t.Errorf("build sign failed. want : %s , build : %s", tc.wantSign, buildSign)
		}
	}
}

func TestVerify(t *testing.T) {
	pubKeyPath := "./testfile/public.crt"
	testCase := []struct {
		body string
	}{
		{
			body: `{"code":"10010","msg":"验签不通过","sign":"4de03d5b7708684779997a6665c15b1f42da65d959228f605b7bbfbdc81981ceda05138f1f01325d2612233f160ccac6c924019eb00ff5b8799e97adb1fc27721ace934ae8361ec64648c22f821104342617321c5cb9fb5673e4950362c79869a532c1c4b86d0620685b9db9de0b406c43e1877a8e920226182e6e71f6e4f81ca370a21df5cfaba0b055c621f6e997f3a9382bd704c917b255c7428621a698931c702445625a8ca803c1cf275de7845b2de7664a36cf7d7eb95283de016a6bdd075301261b77ac4f3a6efb5816a6f05538316a4b76b3a0d0c1f495a8717affe3dd6f10cfd596659cba1b0205040c0742392d44e78570751797051a089ba17189"}`,
		},
	}
	for _, tc := range testCase {
		bodyByte := []byte(tc.body)
		err := signaturebuilder.Verify(bodyByte, pubKeyPath)
		if err != nil {
			t.Error(err)
		}
	}
}
