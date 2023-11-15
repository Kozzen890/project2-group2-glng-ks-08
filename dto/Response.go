package dto

import "time"

type UpdateUserRes struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	Name  string    `json:"name"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UploadMediaRes struct {
	Message string `json:"message"`
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type UpdateMediaRes struct {
	Message string `json:"message"`
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserResponse struct {
	Id              uint   `json:"id"`
	Name        string `json:"name"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type PhotoUpdateRes struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

