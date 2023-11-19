package controllers

type IAuthController interface {
	Get(ctx Context)
}

type authController struct {
}

func NewAuthController() IAuthController {
	return &authController{}
}

func (a *authController) Get(ctx Context) {

}
