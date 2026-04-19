package tests

import (
	ssov1 "github.com/Vladislav747/protos/gen/go/sso"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sso/tests/suite"
	"testing"
)

const (
	emptyAppID = 0
	appID      = 1
	// такой же секрет, как и в 1_init_apps.up.sql
	appSecret = "test-secret"

	passDefaultLen = 10
)

func TestAuthRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()

	pass := randomFakePassword()

	respReq, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: pass,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, respReq.GetUserId())

	//loginTime := time.Now()

	respLogin, err := st.AuthClient.Login(ctx, &ssov1.LoginRequest{
		Email:    email,
		Password: pass,
		AppId:    appID,
	})
	require.NoError(t, err)

	token := respLogin.GetToken()

	assert.NotEmpty(t, token)

	//tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(appSecret), nil
	//})
	//
	//require.NoError(t, err)
	//
	//claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	//assert.True(t, ok)
	//
	//assert.Equal(t, respReq.GetUserId(), int64(claims["uid"].(float64)))
	//assert.Equal(t, email, claims["uid"].(string))
	//assert.Equal(t, appID, int(claims["app_id"].(float64)))
	//
	//const deltaSeconds = 1
	//
	//assert.InDelta(t, loginTime.Add(st.Cfg.TokenTTL).Unix(), claims["exp"].(float64), deltaSeconds)
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, passDefaultLen)
}
