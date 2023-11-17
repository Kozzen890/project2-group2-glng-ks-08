package dto

import "time"

// User Response
type UpdateUserRes struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	Name  string    `json:"name"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserRes struct {
	Id        uint      `json:"id"`
	Name			string		`json:"name"`
	Email			string		`json:"email"`
	Age				int				`json:"age"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
	Photos 		[]GetUserPhotos
	Comments	[]GetCommentsUser
	Media			[]GetUserMedia
}

type GetCommentsUser struct {
	Id        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	PhotoId   uint      `json:"photo_id"`
	Message		string 		`json:"message"`
}

type GetUserMedia struct {
	Id        uint       `json:"id"`
	Name      string		`json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type UserResponse struct {
	Id              uint   `json:"id"`
	Name        string `json:"name"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type GetUserPhotos struct {
	Id        uint       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
}

// Photo Response
type GetPhotosWithUser struct {
	Id        uint       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserId    uint       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      GetUsersPhoto
}

type GetUsersPhoto struct {
	Email    string `json:"email"`
	Name string `json:"name"`
}

type PhotoUpdateRes struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Comments Response
type GetCommentsByUser struct {
	Id        uint      `json:"id"`
	Message     string    `json:"message"`
	PhotoId  uint    `json:"photo_id"`
}

type GetCommentsRes struct {
	Id        uint      `json:"id"`
	Message     string    `json:"message"`
	PhotoId  uint    `json:"photo_id"`
	UserId       int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User  GetUserComments
	Photo GetPhotoComments
}

type GetPhotoComments struct {
	Id  uint       `json:"id"`
	Title    string `json:"title"`
	Caption string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserId    uint       `json:"user_id"`
}

type GetUserComments struct {
	Id  uint       `json:"id"`
	Email    string `json:"email"`
	Name string `json:"name"`
}

type UploadCommentResponse struct {
	Status string `json:"status"`
	Id        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   uint      `json:"photo_id"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Social Media Response
type GetSocialMediaRes struct {
	Id        uint      `json:"id"`
	Name string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId uint `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User GetUserSocialMediaRes
}

type GetUserSocialMediaRes struct {
	Id        uint      `json:"id"`
	Username string			`json:"username"`
	ProfileImageUrl string	`json:"profile_image_url"`
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



