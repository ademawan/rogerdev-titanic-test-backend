package person

type Person interface {
	Create(int64, []interface{}) error
	GetAll() ([]interface{}, error)
}
