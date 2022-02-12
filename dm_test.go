package dm

import (
	"fmt"
	dmSchema "github.com/nfjBill/gorm-driver-dm/schema"
	"gorm.io/gorm"
	"testing"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "dm://sysdba:aaaaaaaaa@192.168.31.192:5236?autoCommit=true"
	db, err = gorm.Open(Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Printf("Error: failed to connect dm server: %v\n", err)
		return
	}

	TB(db)
}

type Model struct {
	ID        uint `gorm:"primarykey,autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Model
	Key      string `gorm:"index:idx_key,unique,comment:备注"`
	Name     string
	Age      int
	Content  dmSchema.Clob `gorm:"size:1024000"`
	Birthday time.Time
}

func TestAutoMigrate(t *testing.T) {
	if Table().HasTable(&User{}) {
		err := Table().DropTable(&User{})

		if err != nil {
			fmt.Printf("Error: failed to DropTable: %v\n", err)
			return
		}
	}

	err := Table().AutoMigrate(&User{})

	if err != nil {
		fmt.Printf("Error: failed to AutoMigrate: %v\n", err)
		return
	}
}

func TestCreate(t *testing.T) {
	err := Table(&User{Key: "1", Name: "Jinzhu", Age: 18, Content: "asdfdasfasdfasdfj手机卡是点击", Birthday: time.Now()}).Create()
	_ = Table(&User{Key: "2", Name: "Jinzhu", Age: 19, Content: "bbb", Birthday: time.Now()}).Create()
	_ = Table(&User{Key: "3", Name: "Jinzhu2", Age: 20, Content: "ccc", Birthday: time.Now()}).Create()

	if err != nil {
		fmt.Printf("Error: failed to Create: %v\n", err)
		return
	}
}

func TestGet(t *testing.T) {
	var data User
	err := Table(&User{Name: "Jinzhu"}).Get(&data)

	if err != nil {
		fmt.Printf("Error: failed to Get: %v\n", err)
		return
	}
}

func TestWhere(t *testing.T) {
	var data []User
	err := Table(&User{Name: "Jinzhu"}).GetWhere(&data)

	if err != nil {
		fmt.Printf("Error: failed to Where: %v\n", err)
		return
	}
}

func TestAll(t *testing.T) {
	var data []User
	err := Table().GetAll(&data)

	if err != nil {
		fmt.Printf("Error: failed to All: %v\n", err)
		return
	}
}

// err
func TestClausesAssignmentColumns(t *testing.T) {
	err := Table(&User{Key: "3", Content: "DDDD"}).ClausesAssignmentColumns("KEY", []string{"DELETED_AT", "CONTENT"})

	if err != nil {
		fmt.Printf("Error: failed to ClausesAssignmentColumns: %v\n", err)
		return
	}
}
