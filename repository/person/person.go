package person

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type PersonRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *PersonRepository {
	return &PersonRepository{
		db: db,
	}
}

func (ur *PersonRepository) Create(time int64, persons []interface{}) error {

	jsonData, _ := json.Marshal(persons)

	_, err := ur.db.Exec("insert into person (date,data) values (?, ?)", time, jsonData)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
func (ur *PersonRepository) GetAll() ([]interface{}, error) {
	result := make([]interface{}, 0)
	type PersonData struct {
		ID   int           `json:"id"`
		Date string        `json:"date"`
		Data []interface{} `json:"data"`
	}
	rows, err := ur.db.Query("select * from person order by date desc")
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = PersonData{}
		var dateUnix int64
		var jsonData []uint8
		var source []byte
		var err = rows.Scan(&each.ID, &dateUnix, &jsonData)

		if err != nil {
			fmt.Println(err.Error())
			return result, err
		}
		source = []byte(jsonData)
		errj := json.Unmarshal(source, &each.Data)
		if err != nil {
			return result, errj
		}

		each.Date = ur.TimeToUser(dateUnix)
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return result, err
	}

	return result, nil
}
func (ur *PersonRepository) TimeToUser(timeInt int64) string {
	if timeInt <= 0 {
		return ""
	}
	i, err := strconv.ParseInt(strconv.Itoa(int(timeInt)), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	year := strconv.Itoa(tm.Year())
	month := strconv.Itoa(int(tm.Month()))
	h, m, s := tm.Clock()
	hour := strconv.Itoa(h)
	if len(hour) == 1 {
		hour = "0" + hour
	}
	minute := strconv.Itoa(m)
	if len(minute) == 1 {
		minute = "0" + minute
	}
	second := strconv.Itoa(s)

	switch month {
	case "1":
		month = "Januari"
	case "2":
		month = "Februari"
	case "3":
		month = "Maret"
	case "4":
		month = "April"
	case "5":
		month = "Mei"
	case "6":
		month = "Juni"
	case "7":
		month = "Juli"
	case "8":
		month = "Agustus"
	case "9":
		month = "September"
	case "10":
		month = "Oktober"
	case "11":
		month = "November"
	case "12":
		month = "Desember"
	}
	day := strconv.Itoa(tm.Day())
	createdOnString := day + " " + month + " " + year + " Pukul " + hour + ":" + minute + ":" + second + " WIB"
	return createdOnString
}
