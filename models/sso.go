package models

type SSOUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"picture"`
}

type PictureData struct {
	Data struct {
		Height       int    `json:"height"`
		IsSilhouette bool   `json:"is_silhouette"`
		Url          string `json:"url"`
		Width        int    `json:"width"`
	} `json:"data"`
}

type FacebookUser struct {
	SSOUser
	Picture PictureData `json:"picture"`
}

type GoogleUser struct {
	SSOUser
	Picture string `json:"picture"`
}
