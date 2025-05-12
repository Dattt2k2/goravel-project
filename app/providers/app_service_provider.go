package providers

import (
	"github.com/goravel/framework/contracts/foundation"
)

type AppServiceProvider struct {
}

// Register đăng ký các ràng buộc cần thiết vào container.
func (app *AppServiceProvider) Register(app foundation.Application) {
	//
}

// Boot khởi động bất kỳ dịch vụ nào.
func (app *AppServiceProvider) Boot(app foundation.Application) {
	//
}
