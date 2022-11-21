// Package jwt
// descr JSON WEB TOKEN
// author fm
// date 2022/11/18 17:28
package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	golangJwt "github.com/golang-jwt/jwt"
	"gohub-lesson/pkg/app"
	"gohub-lesson/pkg/config"
	"gohub-lesson/pkg/logger"
)

var (
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrHeaderEmpty            = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        = errors.New("请求头中 Authorization 格式有误")
)

// JWT 定义一个jwt对象
type JWT struct {

	// 秘钥，用以加密 JWT，读取配置信息 app.key
	SignKey []byte

	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration
}

// JWTCustomClaims 自定义载荷
type JWTCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	golangJwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

func (class *JWT) ParseToken(ctx *gin.Context) (claims *JWTCustomClaims, err error) {

	// 1. 从 Header 里获取 token
	tokenString, err := class.getTokenFromHeader(ctx)

	if err != nil {
		return
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := class.parseTokenString(tokenString)

	// 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		validationErr, ok := err.(*golangJwt.ValidationError)
		if ok {
			if validationErr.Errors == golangJwt.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == golangJwt.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	// 4. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken 更新 Token，用以提供 refresh token 接口
func (class *JWT) RefreshToken(ctx *gin.Context) (string, error) {

	// 1. 从 Header 里获取 token
	tokenString, parseErr := class.getTokenFromHeader(ctx)

	if parseErr != nil {
		return "", parseErr
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := class.parseTokenString(tokenString)

	// 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		validationErr, ok := err.(*golangJwt.ValidationError)
		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != golangJwt.ValidationErrorExpired {
			return "", err
		}
	}

	// 4. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*JWTCustomClaims)

	// 5. 检查是否过了『最大允许刷新的时间』
	x := app.TimeNowInTimezone().Add(-class.MaxRefresh).Unix()

	if claims.IssuedAt > x {
		// 修改过期时间
		claims.StandardClaims.ExpiresAt = class.expireAtTime()
		return class.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

// GenerateToken 生成  Token，在登录成功时调用
func (class *JWT) GenerateToken(userId string, userName string) (token string) {

	// 1. 构造用户 claims 信息(负荷)
	expireAtTime := class.expireAtTime()

	claims := JWTCustomClaims{
		userId,
		userName,
		expireAtTime,
		golangJwt.StandardClaims{
			NotBefore: app.TimeNowInTimezone().Unix(), // 签名生效时间
			IssuedAt:  app.TimeNowInTimezone().Unix(), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: expireAtTime,                   // 签名过期时间
			Issuer:    config.GetString("app.name"),   // 签名颁发者
		},
	}

	// 2. 根据 claims 生成token对象
	token, err := class.createToken(claims)

	if err != nil {
		logger.LogIf(err)
		return
	}

	return
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (class *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// 使用HS256算法进行token生成
	token := golangJwt.NewWithClaims(golangJwt.SigningMethodHS256, claims)
	return token.SignedString(class.SignKey)
}

// expireAtTime 过期时间
func (class *JWT) expireAtTime() int64 {

	timenow := app.TimeNowInTimezone()

	var expireTime int64

	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute

	return timenow.Add(expire).Unix()
}

// parseTokenString 使用 golangJwt.ParseWithClaims 解析 Token
func (class *JWT) parseTokenString(token string) (*golangJwt.Token, error) {
	return golangJwt.ParseWithClaims(
		token,
		&JWTCustomClaims{},
		func(token *golangJwt.Token) (any, error) {
			return class.SignKey, nil
		},
	)
}

// getTokenFromHeader 使用 golangJwt.ParseWithClaims 解析 Token
// Authorization:Bearer xxxxx
func (class *JWT) getTokenFromHeader(ctx *gin.Context) (token string, err error) {

	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)

	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}

	return parts[1], nil
}
