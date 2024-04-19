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

func RegisterConsumerWithArrears(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/arrears/:account", getConsumerWithArrears)
	}
}

func getConsumerWithArrears(c *gin.Context) {
	var consumerWithArrears models.ConsumerArrears

	account := c.Params.ByName("account")
	digit := c.Query("digit")

	if err := getArrears(account, digit, &consumerWithArrears); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   404,
			"status":  "Not found",
			"message": "Consumer not found!",
		})
		return
	}

	c.JSON(http.StatusOK, consumerWithArrears)
}

func getArrears(account string, digit string, consumerWithArrears *models.ConsumerArrears) error {
	stmt, err := conf.GetDBConnection().Prepare(constants.GetConsumerArrears)

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
		&consumerWithArrears.Account,
		&consumerWithArrears.Arrears,
	)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
