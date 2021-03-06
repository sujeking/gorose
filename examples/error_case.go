package main

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	"github.com/gohouse/gorose/v2/examples/dbobj"
)

func getdb() gorose.IOrm {
	engin := dbobj.GetSqliteEngin()
	return engin.NewOrm()
}
func main() {
	var res interface{}
	var err error
	db := getdb()
	res, err = db.Table("users").Insert(gorose.Data{"uid2": 222})
	if err != nil {
		fmt.Println("0 - err:", err.Error())
	}
	fmt.Println("0 - res:", res)
	res, err = db.Reset().Table("users").Where("uid2", ">", 1).First()
	if err != nil {
		fmt.Println("1 - err:", err.Error())
	}
	fmt.Println("1 - res:", res)

	res, err = db.Reset().Table("users").Where("age", ">", 0).First()
	if err != nil {
		fmt.Println("2 - err:", err.Error())
	}
	fmt.Println("2 - res:", res)
}
