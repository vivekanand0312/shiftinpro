package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shiftinpro/internal/services"
)

type ReqAddress struct {
	Pincode *float64 `json:"pincode,omitempty"`
	State   *string  `json:"state,omitempty"`
}

type AddressHandler struct {
	addressService services.AddressService
}

func NewAddressHandler(addressService services.AddressService) *AddressHandler {
	return &AddressHandler{addressService: addressService}
}

func (h *AddressHandler) GetAddress(c *gin.Context) {
	var input ReqAddress

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid input format", "error": err.Error()})
		return
	}

	addresses, err := h.addressService.FetchAddress(input.Pincode, input.State)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to fetch addresses", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": addresses})
}
