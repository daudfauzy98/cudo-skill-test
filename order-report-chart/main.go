package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type OrderReport struct {
	OrderDate  string      `json:"orderDate"`
	Categories []*Category `json:"categories"`
}

type Category struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var ORDER_REPORT []*OrderReport
var Productcategory = []string{
	"Beverages",
	"Condiments",
	"Confections",
	"Dairy Products",
	"Grains/Cereals",
	"Meat/Poultry",
	"Produce",
	"Seafood",
}

func main() {
	initOrderReport()

	routes := gin.Default()
	routes.LoadHTMLFiles("views/index.html")

	routes.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	routes.GET("/api/order", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"status": "success",
			"data":   ORDER_REPORT,
			"lables": Productcategory,
		})
	})
	routes.Run("localhost:8000")
}

func initOrderReport() {
	b, err := ioutil.ReadFile("order-report.json")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(b, &ORDER_REPORT)
	if err != nil {
		log.Fatalln(err)
	}

	for i, val := range ORDER_REPORT {
		notExists := notExistCategories(val.Categories, Productcategory)
		for _, val := range notExists {
			ORDER_REPORT[i].Categories = append(ORDER_REPORT[i].Categories, &Category{
				Name:     val,
				Quantity: 0,
			})
		}
	}
}

func notExistCategories(arr []*Category, categories []string) []string {
	var missing []string
	for _, category := range categories {
		if !isExist(arr, category) {
			missing = append(missing, category)
		}
	}
	return missing
}

func isExist(arr []*Category, s string) bool {
	for _, val := range arr {
		if val.Name == s {
			return true
		}
	}
	return false
}
