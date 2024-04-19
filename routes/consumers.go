package routes

import (
	"database/sql"
	conf "github.com/ericsison/bill-inquiry/config"
	"github.com/ericsison/bill-inquiry/constants"
	"github.com/ericsison/bill-inquiry/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterConsumerRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// group routes into v1
		v1.GET("/consumers/:account", getBill)
	}
}

/**
* There are three (3) possible scenarios where billing might return default values
* (1) - The account number has arrears (balance from preview month/s)
* (2) - The account number has advance payment
* (3) - The account number has fully paid
 */
func getBill(ctx *gin.Context) {
	var bill models.Bill

	account := ctx.Params.ByName("account")
	digit := ctx.Query("digit")

	if err := getConsumerInfo(account, digit, &bill); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   404,
			"status":  "Not found",
			"message": "Consumer not found!",
		})
		return
	}

	getBalance(account, digit, &bill)

	getOtherCharges(account, digit, &bill)

	ctx.JSON(http.StatusOK, bill)
}

func getConsumerInfo(account string, digit string, bill *models.Bill) error {
	stmt, err := conf.GetDBConnection().Prepare(constants.GetCustomerInfoQuery)

	if err != nil {
		log.Print(err)
	}

	row := stmt.QueryRow(sql.Named("account", account), sql.Named("digit", digit))

	defer func(s *sql.Stmt) {
		if err := s.Close(); err != nil {
			log.Print(s)
		}
	}(stmt)

	err = row.Scan(
		&bill.Info.Account,
		&bill.Info.Name,
		&bill.Info.Address,
		&bill.Info.MeterNo,
		&bill.Info.CheckDigit,
		&bill.Info.Description,
		&bill.Info.Status,
		&bill.Info.AverageUsage,
	)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func getBalance(account string, digit string, bill *models.Bill) {
	stmt, err := conf.GetDBConnection().Prepare(constants.GetBalanceQuery)

	if err != nil {
		log.Print(err)
	}

	row := stmt.QueryRow(sql.Named("account", account), sql.Named("digit", digit))

	err = row.Scan(
		&bill.Billing.WaterBill,
		&bill.Billing.DueDate,
		&bill.Billing.SeniorDiscount,
		&bill.Billing.AdvancePayment,
		&bill.Billing.Penalty,
		&bill.Billing.Total,
		&bill.Billing.WaterUsage,
	)

	//fmt.Println(&bill.Billing.WaterUsage)

	if err != nil {
		log.Print(err)
	}
}

func getOtherCharges(account string, digit string, bill *models.Bill) {
	var other models.OtherCharges

	stmt, err := conf.GetDBConnection().Prepare(constants.GetOtherChargesQuery)

	if err != nil {
		log.Print(err)
	}

	rows, err := stmt.Query(sql.Named("account", account), sql.Named("digit", digit))

	for rows.Next() {
		err := rows.Scan(&other.Remarks, &other.Due, &other.Balance)

		if err != nil {
			log.Print(err)
		}

		bill.Billing.Other = append(bill.Billing.Other, other)
	}
}
