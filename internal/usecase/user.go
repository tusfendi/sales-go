package usecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/tusfendi/sales-go/config"
	"github.com/tusfendi/sales-go/internal/constants"
	"github.com/tusfendi/sales-go/internal/entity"
	"github.com/tusfendi/sales-go/internal/presenter"
	"github.com/tusfendi/sales-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(params presenter.Registration) (httpStatus int, err error)
	AuthUser(params presenter.Auth) (presenter.AuthResponse, int, error)
}

type userCtx struct {
	cfg      config.Config
	userRepo repository.UserRepository
}

func NewUserUsecase(cfg config.Config, userRepo repository.UserRepository) UserUsecase {
	return &userCtx{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (u *userCtx) CreateUser(params presenter.Registration) (int, error) {
	if errMessage, isError := params.Validate(); isError {
		return http.StatusBadRequest, errors.New(errMessage)
	}

	_, isExist, err := u.userRepo.GetByEmail(params.Email)

	if err != nil {
		return http.StatusInternalServerError, errors.New(constants.ErrDB)
	}

	if isExist {
		return http.StatusBadRequest, errors.New("Email sudah terdaftar")
	}

	passwordEncrypt, err := u.hash(params.Password)
	if err != nil {
		return http.StatusInternalServerError, errors.New(constants.ErrInternal)
	}

	_, err = u.userRepo.CreateUser(&entity.User{
		Email:    params.Email,
		Name:     params.Name,
		Password: passwordEncrypt,
	})

	if err != nil {
		return http.StatusInternalServerError, errors.New(constants.ErrDB)
	}

	return http.StatusCreated, nil
}

func (u *userCtx) AuthUser(params presenter.Auth) (presenter.AuthResponse, int, error) {
	response := presenter.AuthResponse{}

	if errMessage, isError := params.Validate(); isError {
		return response, http.StatusBadRequest, errors.New(errMessage)
	}

	user, isExist, err := u.userRepo.GetByEmail(params.Email)

	if err != nil {
		return response, http.StatusInternalServerError, errors.New(constants.ErrDB)
	}

	if !isExist || !u.isSame(params.Password, user.Password) {
		return response, http.StatusInternalServerError, errors.New("Email atau kata sandi salah")
	}

	token, err := u.generateToken(user)
	if err != nil {
		return response, http.StatusInternalServerError, errors.New(constants.ErrInternal)
	}
	response.Email = user.Email
	response.Token = token

	return response, http.StatusOK, nil
}

func (u *userCtx) hash(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashed), err
}

func (u *userCtx) isSame(str string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str)) == nil
}

func (u *userCtx) generateToken(params entity.User) (string, error) {
	now := time.Now().Unix()
	expired := time.Now().AddDate(0, 0, u.cfg.ExpiredToken).Unix()

	claims := &presenter.SSJWTClaim{
		Email: params.Email,
		ID:    params.ID,
		StandardClaims: &jwt.StandardClaims{
			Id:        fmt.Sprintf("%d", params.ID),
			IssuedAt:  now,
			NotBefore: now,
			ExpiresAt: expired,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(u.cfg.SecretKey))
	if err != nil {
		return "", errors.New("Terjadi kesalahan")
	}

	return tokenString, nil
}
