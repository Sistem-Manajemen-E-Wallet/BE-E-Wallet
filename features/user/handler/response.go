package handler

type UserResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	Role           string `json:"role"`
	ProfilePicture string `json:"profile_picture,omitempty"`
}
