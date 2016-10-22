package service

import (
	"ap/models"
)

type AccountService struct {
}

var Account *AccountService

func (service *AccountService) Find(accountId int64) (models.Account, bool) {

	_, info := models.FindByAccount(accountId)
	return info, true
}
