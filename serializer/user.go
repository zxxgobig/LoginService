package serializer

import "LoginService/model"

type User struct {
	ID uint `form:"id" json:"id"`
	UserName string `form:"name" json:"user_name"`
	NickName string `form:"name" json:"name"`
	Status string `json:"status"`
	Avatar string `json:"avatar"`
	CreatedAt int64 `json:created_at`

}

type UserResponse struct {
	Response
	Data User `json:"data"`
}
func BuildUserResponse(user model.User) UserResponse{
	return UserResponse{
		Data: User{
			ID:user.ID,
			UserName:user.UserName,
			NickName:user.NickName,
			Status:user.Status,
			Avatar:user.Avatar,
			CreatedAt: user.CreatedAt.Unix(),
		},
	}
}