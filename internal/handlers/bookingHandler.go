package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shiftinpro/internal/services"
)

type BookingHandler struct {
	bookingService services.BookingService
}

func NewBookingHandler(bookingService services.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}

func (h *BookingHandler) GetItemChecklists(c *gin.Context) {
	itemChecklists, err := h.bookingService.GetItemChecklists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to fetch Item Checklists", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": itemChecklists})
}
