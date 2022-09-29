package utils

import (
	"database/sql"
	"fmt"
	"rogerdev-titanic-test-backend/configs"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(config *configs.AppConfig) (*sql.DB, error) {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {

		fmt.Println(err.Error(), connectionString)
		return nil, err
	}

	InitMigrate(db)
	return db, nil
}

func InitMigrate(db *sql.DB) {
	// db.Migrator().DropTable(&entities.User{})
	// db.AutoMigrate(&entities.User{})
	rows, err := db.Query("CREATE TABLE IF NOT EXISTS `user` (`uid` varchar(32) NOT NULL,`name` varchar(50) NOT NULL,`email` varchar(50) NOT NULL,`password` varchar(100) NOT NULL,`address` varchar(50),`gender` varchar(10) NOT NULL,created_at INT,updated_at INT,deleted_at INT,UNIQUE KEY unique_uid (uid),UNIQUE KEY unique_email (email)) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rows2, err2 := db.Query("CREATE TABLE IF NOT EXISTS `person` ( `id` int NOT NULL AUTO_INCREMENT,`date` int NOT NULL,`data` json DEFAULT NULL,PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
	if err != nil {
		fmt.Println(err2.Error())
		return
	}

	defer rows.Close()
	defer rows2.Close()

}
