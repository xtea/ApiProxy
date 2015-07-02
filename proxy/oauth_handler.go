package proxy

import (
	"fmt"
	"github.com/EE-Tools/goauth/common"
	"github.com/EE-Tools/goauth/models/auth/checker"
	_ "log"
	"net/http"
)

// oauth check handler
type OauthCheckHandler struct {
}

func (this *OauthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	accessToken := r.Form.Get("access_token")
	if accessToken == "" {
		msg := "param [access_token] is empty"
		WriteHttpErrorMessage(HTTP_CODE_FORBIDDEN, msg, w)
		return false
	}
	tokenChecker := new(checker.TokenChecker)
	_, errCode := tokenChecker.CheckAccessToken(map[string]string{"access_token": accessToken}, true)
	if errCode != common.API_EC_SUCCESS {
		aec := common.NewApiErrorCode(errCode)
		msg := fmt.Sprintf("%d:%s", aec.ErrNo, aec.ErrMsg)
		WriteHttpErrorMessage(HTTP_CODE_FORBIDDEN, msg, w)
		return false
	}
	return true
}
