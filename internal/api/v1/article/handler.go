package article

import (
	"net/http"
	"oncomapi/internal/api/v1/article/dto"
	"oncomapi/pkg/response"
	"oncomapi/pkg/validation"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// GetAll godoc
// @Summary Get all articles
// @Tags Article
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.ArticleResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /articles [get]
func (h *Handler) GetAll(c *gin.Context) {
	articles, err := h.service.GetAll()
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SendSuccess(c, DataList(articles))
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid ID")
	}

	article, err := h.service.GetByID(uint(id))
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	/*if err != nil {
		logger.ErrorSentry(err, "Gagal ambil article ID %v", id)
		if appErr, ok := err.(*apperror.AppError); ok {
			response.SendError(c, appErr.Code, appErr.Message)
			return
		}
		response.SendError(c, http.StatusInternalServerError, "Internal server error")
		return
	}*/
	response.SendSuccess(c, DetailResponse(&article))
}

func (h *Handler) Create(c *gin.Context) {
	input := dto.CreateRequest{}
	if !validation.Validate(c, &input) {
		return
	}

	result, err := h.service.Create(&input)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SendSuccess(c, DetailResponse(result), "Article created")
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	input := dto.UpdateRequest{}
	if !validation.Validate(c, &input) {
		return
	}

	input.ID = uint(id)
	result, err := h.service.Update(&input)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SendSuccess(c, DetailResponse(result), "Article updated")
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		response.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SendMessage(c, "Article deleted")
}

func (h *Handler) GetUserArticles(c *gin.Context) {
	userID := c.GetInt("userID")

	articles, err := h.service.FindByUserID(userID)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, "Failed to read articles")
		return
	}

	response.SendSuccess(c, articles)
}
