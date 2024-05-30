package handler

type UserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}

type LoginRequest struct {
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RoleRequest struct {
	Role string `json:"role" form:"role"`
}

type UpdateProfilePictureRequest struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
}
