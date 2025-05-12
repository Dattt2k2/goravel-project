package main

import (
	"goravel/bootstrap"

	"github.com/goravel/framework/facades"

	_ "goravel/database/migrations" // Import để đăng ký tất cả migrations
)

func main() {
	// Khởi tạo ứng dụng
	bootstrap.Boot()

	// Thực hiện migrate
	if err := facades.Artisan().Call("migrate"); err != nil {
		facades.Log().Error("Migration failed: " + err.Error())
		return
	}

	facades.Log().Info("Migration completed successfully!")

	// Để chạy HTTP server (nếu cần)
	// facades.Route().Run(facades.Config().GetString("app.port", "8000"))
}
