package response

type Login struct {
	AccessToken string `json:"access_token"`
	Email       string `json:"email"`
	UserID      int    `json:"user_id"`
}
