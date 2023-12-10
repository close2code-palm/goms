package dtos

import "oauth2_api/domain"

type InsertedOne struct {
	InsertedId interface{} `json:"insertedId"`
}

type Token struct {
	AccessToken domain.TokenData
}
