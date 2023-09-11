package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	ATokenExpiredDuration = 2 * time.Hour
	RTokenExpiredDuration = 30 * 24 * time.Hour
	TokenIssuer           = "admin"
)

var (
	mySecret          = []byte("my Secret Decode")
	ErrorInvalidToken = errors.New("verify Token Failed")
)

type PayLoad struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func getJWTTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}

func keyFunc(token *jwt.Token) (any, error) {
	return mySecret, nil
}

// GenToken 颁发token access token 和 refresh token
func GenToken(userID uint64, userName string, secret []byte) (atoken, rtoken string, err error) {
	// 构建 凭证 基础信息
	rc := jwt.RegisteredClaims{
		Issuer:    TokenIssuer,                       // 颁发人
		ExpiresAt: getJWTTime(ATokenExpiredDuration), // 到期时间
	}
	// 绑定载荷信息
	at := PayLoad{userID, userName, rc}
	// 使用SHA256对载荷非对称加密，进行签名和加盐
	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, at).SignedString(secret)

	// refresh token 长token用来刷新，所以不需要载荷。
	rt := rc
	rt.ExpiresAt = getJWTTime(RTokenExpiredDuration)
	rtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rt).SignedString(secret)

	return atoken, rtoken, err
}

// VerifyToken 验证Token
func VerifyToken(tokenId string) (PayLoad, error) {
	var payLoad PayLoad
	token, err := jwt.ParseWithClaims(tokenId, &payLoad, keyFunc)
	if err != nil {
		return payLoad, err
	}
	// 解析成功后为True
	if !token.Valid {
		err = ErrorInvalidToken
		return payLoad, err
	}
	return payLoad, nil
}

// RefreshToken 通过refresh token 刷新 短token(atoken)
func RefreshToken(atoken, rtoken string, secret []byte) (newAtoken, newRtoken string, err error) {
	// rtoken 无效退出
	if _, err = jwt.Parse(rtoken, keyFunc); err != nil {
		return
	}
	// 从旧的access token 中解析出 payload 数据信息
	var claim PayLoad
	// 校验不通过，并且该错误是因为Token过期引起的，那么进行续签。
	_, err = jwt.ParseWithClaims(atoken, &claim, keyFunc)
	if err != nil && err.Error() == "token has invalid claims: token is expired" {
		return GenToken(claim.UserID, claim.Username, secret)
	}
	return
}
