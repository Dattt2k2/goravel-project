package migratifunc (r *M20251205_CreateUsersTable) Up() error {
    if !facades.Schema().HasTable("users") {
        return facades.Schema().Create("users", func(table schema.Blueprint) {
            table.ID()
            table.String("name") // Bắt buộc có tên
            table.String("email").Unique() // Email phải duy nhất và bắt buộc
            table.String("password") // Mật khẩu bắt buộc
            table.String("phone") // Số điện thoại
            table.TinyInteger("type").Default(0) // 0: Recruiter, 1: Candidate
            table.Timestamps() // Thêm các trường created_at và updated_at
        }) (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251205_CreateUsersTable struct {

}

func (r *M20251205_CreateUsersTable) Signature() string {
	return "2025_12_05_000000_create_users_table"
}

func (r *M20251205_CreateUsersTable) Up() error {
    if !facades.Schema().HasTable("users") {
        return facades.Schema().Create("users", func(table schema.Blueprint) {
            table.ID()
            table.String("name") // Bắt buộc có tên
            table.Unique("email") // Email phải duy nhất và bắt buộc
            table.String("password") // Mật khẩu bắt buộc
			table.String("phone") // Số điện thoại
			table.Inte
            table.Timestamps() // Thêm các trường created_at và updated_at
        })
    }
    return nil
}

