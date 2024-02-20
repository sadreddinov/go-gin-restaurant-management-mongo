package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/go-gin-restaurant-management-mongo/pkg/handler"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", handler.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", handler.GetInvoice())
	incomingRoutes.POST("/invoices", handler.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", handler.UpdateInvoice())
}
