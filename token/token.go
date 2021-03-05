package token

import (

	"context"
	"easycrm/models"
	"errors"
	"github.com/dsurush/jwt/pkg/jwt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type TokenSvc struct {
	secret []byte
	pool   *pgxpool.Pool
}

func NewTokenSvc(secret []byte, pool *pgxpool.Pool) *TokenSvc {
	return &TokenSvc{secret: secret, pool: pool}
}

type Payload struct {
	Id   uuid.UUID  `json:"id"`
	Exp   int64  `json:"exp"`
	Login string `json:"login"`
	Role  string `json:"role"`
}

type RequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	Token string `json:"token"`
//	Role string `json:"role"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}

//var ErrInvalidLogin = errors.New("invalid login or password")
var ErrInvalidPasswordOrLogin = errors.New("invalid password")

func (receiver *TokenSvc) FindUserForPassCheck(login string) (User models.Admin, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	err = receiver.pool.QueryRow(context.Background(), `Select ID, first_name, last_name, password_hash from users where username = ($1)`, login).Scan(
		&User.ID,
		&User.FirstName,
		&User.LastName,
		&User.Password)
	if err != nil {
		log.Printf("Can't scan %e", err)
	//	fmt.Printf("Can't scan %e", err)
		return
	}
	return
}
func (receiver *TokenSvc) Generate(context context.Context, request *RequestDTO) (response ResponseDTO, err error) {
	//user, err := models.FindUserByLogin(request.Username)
	user, err := receiver.FindUserForPassCheck(request.Username)
	if err != nil {
		err = ErrInvalidPasswordOrLogin
		return
	}
	//fmt.Println(user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		err = ErrInvalidPasswordOrLogin
		return
	}
	response.Name = user.FirstName
	response.Surname = user.LastName
	response.Token, err = jwt.Encode(Payload{
		Id:  user.ID,
		Exp: time.Now().Add(time.Hour * 10).Unix(),
		///Exp:   time.Now().Add(time.Second * 10).Unix(),
		//Login: user.Login,
		//Role:  user.Role,
	}, receiver.secret)
	return
}