package domain

type UserId int

type TokenData struct {
	Uid      UserId   `json:"userId,omitempty"` // null
	Username string   `json:"username"`
	Roles    []string `json:"roles,omitempty"`
	Expires  int      `json:"exp"`
}
