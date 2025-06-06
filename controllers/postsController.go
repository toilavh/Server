package controllers

import (
	"awesome_project/initializers"
	"awesome_project/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateRobot(c *gin.Context) {
	// Lấy dữ liệu từ body của request
	var body models.Robot
	c.Bind(&body)

	// Tạo một đối tượng Robot mới với dữ liệu từ body
	bot := models.Robot{
		Name:            body.Name,
		Description:     body.Description,
		Model:           body.Model,
		Version:         body.Version,
		ManufactureCost: body.ManufactureCost,
		ManufactureDate: body.ManufactureDate,
		Color:           body.Color,
		CameraCount:     body.CameraCount,
	}
	result := initializers.DB.Create(&bot)

	if result.Error != nil {
		// Nếu có lỗi trong quá trình tạo, trả về lỗi
		c.Status(400)
		return
	}
	// Trả về mã trạng thái 200 OK và đối tượng Robot đã tạo
	c.JSON(200, gin.H{
		"robot": bot,
	})
}

func GetAllRobots(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Bắt đầu với DB query
	db := initializers.DB

	// Lọc theo Model
	if models := c.Query("model"); models != "" {
		modelList := strings.Split(models, ",")
		db = db.Where("LOWER(model) IN ?", toLowerSlice(modelList))
	}

	// Lọc theo Color
	if colors := c.Query("color"); colors != "" {
		colorList := strings.Split(colors, ",")
		db = db.Where("LOWER(color) IN ?", toLowerSlice(colorList))
	}

	// Lọc theo khoảng ngày sản xuất
	start := c.Query("manufacture_date_start")
	end := c.Query("manufacture_date_end")
	if start != "" && end != "" {
		db = db.Where("manufacture_date BETWEEN ? AND ?", start, end)
	}

	//Sort theo ID
	sortBy := c.DefaultQuery("sort_by", "id") // mặc định là "id"
	order := c.DefaultQuery("order", "asc")   // mặc định là "asc"

	// ✅ Danh sách cột cho phép sắp xếp
	allowedSortColumns := map[string]bool{
		"id":               true,
		"name":             true,
		"model":            true,
		"version":          true,
		"manufacture_date": true,
		"manufacture_cost": true,
		"camera_count":     true,
	}

	if allowedSortColumns[sortBy] {
		if order != "asc" && order != "desc" {
			order = "asc"
		}
		db = db.Order(sortBy + " " + order)
	}

	// Truy vấn với phân trang
	var robots []models.Robot
	result := db.Limit(limit).Offset(offset).Find(&robots)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"robots": robots,
		"page":   page,
		"limit":  limit,
	})
}

func toLowerSlice(input []string) []string {
	// Chuyển đổi tất cả các phần tử trong slice thành chữ thường và loại bỏ khoảng trắng
	for i, v := range input {
		input[i] = strings.ToLower(strings.TrimSpace(v))
	}
	return input
}

func DeleteRobot(c *gin.Context) {
	// lấy id từ tham số URL
	id := c.Param("id")

	// Delete the robot
	initializers.DB.Delete(&models.Robot{}, id)
	// Nếu có lỗi trong quá trình xóa, trả về lỗi
	if initializers.DB.Error != nil {
		c.Status(500)
		return
	}
	// Nếu không có lỗi, trả về mã trạng thái 200 OK
	c.Status(200)
	// Hoặc có thể trả về một thông báo thành công
	c.JSON(200, gin.H{
		"message": "Robot deleted successfully",
	})
}
