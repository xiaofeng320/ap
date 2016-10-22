package models

import "github.com/astaxie/beego/orm"

type Account struct {
	AccountId int `orm:"pk;auto"`
	UserId    string
	Type      int
}

func FindByAccount(accoungId int64) (bool, Account) {
	o := orm.NewOrm()
	var account Account
	err := o.QueryTable(account).Filter("AccountId", accoungId).One(&account)
	return err != orm.ErrNoRows, account
}
