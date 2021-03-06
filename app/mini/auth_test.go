package mini_test

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"

	"github.com/godcong/wego/app/mini"
)

func TestAuth_Session(t *testing.T) {
	auth := mini.NewAuth(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))
	resp := auth.Session("0022IX8c1OPfgv0tOQ6c1tGZ8c12IX8E")
	t.Log(resp.String())
}

func TestAuth_UserInfo(t *testing.T) {
	auth := mini.NewAuth(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))
	resp := auth.UserInfo("002JXxze2ilgfB0zNmAe2Amsze2JXxzJ", "rCmWuMckRqkw33i+s+NCh32iPdO+yiPS/FWJInan6XUdnXROIC8vXm7clc5NlRMFjI1hPo59eWWeLeLyfZs5lzuzOHASH2VVnwwetAjwbt9KC9v8zWGAZfvlweQWlBtKpSNS0H9dc1bhXafuA763mRq0v01Uq/LAktVAcyd1l/2JCKPhosRSov9F8FTCTt4YL1S4NeYGcjPDb+Mgb9LeRleseMZuziZbKvs66XnPw2ARtrGsiU3uyB4/WZGKERMJll3eRmgYe98F+q4ey0VAz3+Ah5x5NHDfrmxFgm4t3U78VF9q7IB706ULUgMozXJlU5cjsuaVNROXpBmWT/3fHpL3XIWl6U/m7V9o8RiLmmxSSChGCpq2zMjPqj741Z1gKe0wuQ7RpKAWrd1Ui2tG23r6TCigYCE7cb4BEI/KRJkWP0LbfTG8S/9tvuX+xuSgd78qc5nXGqEpMz+FR+b0yC2UcBBup3HO9WZ/3Ut8BjA=", "rVJM6LaFd8PboQCHvwDelQ==")
	t.Log(string(resp))
}
