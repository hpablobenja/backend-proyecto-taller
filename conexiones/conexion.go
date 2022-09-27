package conexiones

import (
	"log"

	"crud-mysql/modelos"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/test?charset=utf8")
	//db, error := gorm.Open("mysql", "0ttq327hth9ufzp26nm8:pscale_pw_77kAJ2SbxBTlH2hrYBn2jQRZ3vbZ5rjU1eZr9bX9hwU@tcp(us-east.connect.psdb.cloud)/base1?tls=true")
	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrando....")

	db.AutoMigrate(&modelos.Persona{})
	db.AutoMigrate(&modelos.Noticia{})
}
