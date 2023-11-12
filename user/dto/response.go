package dto

import "github.com/tae2089/gin-boilerplate/common/domain"

type ResponseLogin struct {
	domain.JwtToken
}
