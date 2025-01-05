package validation

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type Register struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=20"`
	EmailCode string `json:"email_code" binding:"required,len=6"`
}

type ForgetPassword struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=20"`
	EmailCode string `json:"email_code" binding:"required,len=6"`
}

type SendCode struct {
	Email string `binding:"required,email"`
}
