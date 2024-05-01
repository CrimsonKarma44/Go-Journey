// package main

// import (
// 	"fmt"
// 	"os"
// )

// type Page struct {
// 	Title string
// 	Body []byte
// }

// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	return os.WriteFile(filename, p.Body, 0600)
// }

// // func loadPage(title string) *Page {
// // 	filename := title + ".txt"
// // 	body, _ := os.ReadFile(filename)
// // 	return &Page{Title: title, Body: body}
// // }

// func loadPage(title string) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }

// func main() {
// 	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
// 	p1.save()
// 	p2, _ := loadPage("TestPage")
// 	fmt.Println(string(p2.Body))
// }

//go:build ignore

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func main() {
// 	fmt.Println("starting******")

// 	http.HandleFunc("/api-test", handler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payment struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Status   string  `json:"status"`
}

var payments []Payment

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.GET("/api/payments", getPayments)
	router.POST("/api/payments", createPayment)

	// Start server
	router.Run(":8080")
}

func getPayments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": payments,
	})
}

func createPayment(c *gin.Context) {
	var payment Payment

	// Bind JSON body to Payment struct
	if err := c.BindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add payment to payments slice
	payments = append(payments, payment)

	c.JSON(http.StatusCreated, gin.H{"message": "Payment created successfully", "data": payment})
}
