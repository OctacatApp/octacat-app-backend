// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthMutation struct {
	Register *User        `json:"register"`
	Login    *JWTResponse `json:"login"`
}

type GetListParams struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type JWTResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type LoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Mutation struct {
}

type Query struct {
}

type RegisterParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"profileImage"`
	CreatedAt    string `json:"createdAt"`
	CreatedBy    string `json:"createdBy"`
	UpdatedAt    string `json:"updatedAt"`
	UpdatedBy    string `json:"updatedBy"`
	DeletedAt    string `json:"deletedAt"`
	DeletedBy    string `json:"deletedBy"`
	IsDeleted    bool   `json:"isDeleted"`
}

type UserPagination struct {
	Limit     int     `json:"limit"`
	Page      int     `json:"page"`
	TotalPage int     `json:"totalPage"`
	TotalData int     `json:"totalData"`
	Data      []*User `json:"data,omitempty"`
}

type UserQuery struct {
	GetList *UserPagination `json:"getList"`
	Me      *User           `json:"me"`
}
