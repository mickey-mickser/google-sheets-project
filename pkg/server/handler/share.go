package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/responses"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/mail"
	"strings"
)

// Share grants read access to a Google Sheet for a given email
// @Summary Share a Google Sheet (read access)
// @Description Grants "reader" permission to a user by email
// @Tags Sheets
// @Accept json
// @Produce json
// @Param request body models.ShareRequest true "Share payload"
// @Success 200 {object} map[string]interface{} "status + permission"
// @Failure 400 {object} responses.ResponseError "Invalid request body"
// @Failure 500 {object} responses.ResponseError "Internal server error"
// @Router /api/v1/sheets/share [post]
func (h *Handler) Share(c *gin.Context) {
	var req models.ShareRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequestBind(c, handlerNameShare, err)
		return
	}

	if req.Email == "" {
		responses.BadRequest(c, responses.NewMissingParamError("email"))
		return
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		responses.BadRequest(c, responses.NewInvalidBodyError(errors.New("invalid email format")))
		return
	}
	if req.SendEmail && strings.TrimSpace(req.EmailMessage) == "" {
		responses.BadRequest(c, responses.NewInvalidBodyError(errors.New("emailMessage is required when sendEmail=true")))
		return
	}

	h.Log.WithFields(logrus.Fields{
		"handler":   handlerNameShare,
		"email":     req.Email,
		"sendEmail": req.SendEmail,
	}).Info("sharing read access to sheet")

	resp, err := h.SheetUse.SharePermission(c.Request.Context(), req)
	if err != nil {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameShare,
			"error":   err.Error(),
		}).Error(errShareFailed)
		responses.InternalError(c)
		return
	}

	responses.OK(c, gin.H{"status": "ok", "data": resp})
}
