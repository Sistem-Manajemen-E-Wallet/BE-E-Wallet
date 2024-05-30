package handler

type UserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	PhoneNumber string `gorm:"unique" json:"phone_number" form:"phone_number"`
	Pin         string `json:"pin" form:"pin"`
	PinConfirm  string `json:"confirm_pin" form:"confirm_pin"`
}

type UserUpdateRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	PhoneNumber string `gorm:"unique" json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
}

type LoginRequest struct {
	PhoneNumber string `gorm:"unique" json:"phone_number" form:"phone_number"`
	Pin         string `json:"pin" form:"pin"`
}

type UpdateProfilePictureRequest struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
}
