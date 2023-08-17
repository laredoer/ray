package ray

import (
	"fmt"
	"log"
	"testing"
)

type User struct {
	ID   int64  `ray:"id"`
	Name string `ray:"name"`
	Age  int    `ray:"age"`
}

func (User) TableName() string {
	return "users"
}

func TestFrom(t *testing.T) {
	users, err := From[User](
		Select("id", "name"),
		Where("id", "=", 1),
		And("name", "=", "test"),
		Or("age", "=", 18),
		Order("id desc"),
		GroupBy("id"),
	).All()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}
