# Server
# Robot Management API

## 🧱 Cấu trúc Source Code

```bash
.
├── main.go                 # Điểm khởi động của chương trình
├── .env                   # Biến môi trường (DB_URL, API_KEY, PORT...)
├── initializers/          # Thư mục khởi tạo ban đầu (DB, env)
│   ├── database.go        # Kết nối database
│   └── loadEnv.go         # Load file .env
├── models/                # Khai báo cấu trúc dữ liệu (struct)
│   └── robot.go           # Định nghĩa struct Robot
├── controllers/           # Xử lý logic API (CRUD, filter, sort...)
│   └── robotController.go # Các hàm xử lý request liên quan đến robot
├── middlewares/           # Middleware kiểm tra API key
│   └── middleware.go
├── routes/                # Khai báo các route API
│   └── router.go
└── go.mod / go.sum        # Quản lý module Go

| Công nghệ      | Lý do chọn                                                                 |
| -------------- | -------------------------------------------------------------------------- |
| **Golang**     | Ngôn ngữ nhẹ, tốc độ cao, dễ mở rộng, phổ biến trong viết API backend      |
| **PostgreSQL** | Hệ quản trị cơ sở dữ liệu mạnh mẽ, hỗ trợ tốt cho lọc, phân trang, sắp xếp |
| **Gin**        | Framework HTTP nhanh và đơn giản nhất cho Go                               |
| **GORM**       | ORM giúp làm việc với database dễ dàng, rõ ràng                            |
| **godotenv**   | Load file `.env` để dễ quản lý cấu hình môi trường                         |
| **Postman**    | Test API thủ công một cách trực quan                                       |

