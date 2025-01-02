package database

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

func Connect()(*sqlx.DB,error){
	connection,connectionErr:=	sqlx.Connect("","")
	if connectionErr!=nil{
		return nil,errors.New("Error during connecting to pgsql ")
	}
	if connection.Ping()!=nil{
		return nil,errors.New("Error during pgsql ping")
	}
	return connection,nil
}
