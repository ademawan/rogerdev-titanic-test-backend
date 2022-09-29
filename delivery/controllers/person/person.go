package person

import (
	"fmt"
	"net/http"
	"rogerdev-titanic-test-backend/delivery/controllers/common"
	"rogerdev-titanic-test-backend/repository/person"
	"time"

	// utils "todo-list-app/utils/aws_S3"

	// "github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
)

type PersonController struct {
	repo person.Person
	// conn *session.Session
}

func New(repository person.Person /*, S3 *session.Session*/) *PersonController {
	return &PersonController{
		repo: repository,
		// conn: S3,
	}
}

func (pc *PersonController) Post() echo.HandlerFunc {
	return func(c echo.Context) error {

		persons := []CreatePersonRequestFormat{}

		c.Bind(&persons)
		if len(persons) < 1 {
			errMessage := fmt.Sprintf("data json is required")
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, errMessage, nil))

		}
		for i, val := range persons {
			err := c.Validate(&val)

			if err != nil {
				// fmt.Println(err.Error())
				errMessage := fmt.Sprintf("There is some problem from input data ke-%v", i+1)
				return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, errMessage, nil))
			}
		}
		data := make(map[int][]string)
		for _, val := range persons {
			if _, checked := data[val.Age]; !checked {
				data[val.Age] = []string{val.Name}
			} else {
				data[val.Age] = append(data[val.Age], val.Name)
			}
		}
		higerTotal := 0
		dataPerson := make([]interface{}, 0)
		for _, val := range data {
			if len(val) > higerTotal {
				higerTotal = len(val)
				dataPerson = make([]interface{}, 0)
				dataPerson = append(dataPerson, val)
			} else if len(val) == higerTotal {
				dataPerson = append(dataPerson, val)
			}
		}

		timeLocation, _ := time.LoadLocation("Asia/Jakarta")
		timeNow := time.Now().In(timeLocation).Unix()

		err_repo := pc.repo.Create(timeNow, dataPerson)
		errMessageFromRepo := "Success filter person"
		if err_repo != nil {
			errMessageFromRepo += " but data can't saved history to database"
		} else {
			errMessageFromRepo += " ,history saved"
		}

		// response := PersonCreateResponse{}
		// response.Name = res.Name
		// response.Age = res.Age
		// response.Roles = res.Roles
		// response.Image = res.Image

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusOK, errMessageFromRepo, dataPerson))

	}
}
func (pc *PersonController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		data, err_repo := pc.repo.GetAll()
		errMessageFromRepo := "Success get history"
		if err_repo != nil {
			errMessageFromRepo += " Can't connect to database"
		}

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusOK, errMessageFromRepo, data))

	}
}
