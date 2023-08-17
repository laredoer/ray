# ray

The fantastic ORM library for Golang, aims to be developer friendly.

## Quick Start

```golang

type User struct {
  ID   int64  `ray:"id"`
  Name string `ray:"name"`
  Age  int    `ray:"age"`
}

func (User) TableName() string {
  return "users"
}


func GetUsers() {
  users, err := From[User](
    Select("id", "name"),
    Where("id", "=", 1),
    And("name", "=", "test"),
    Or("age", "=", 18),
    Order("id desc"),
  ).All()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(users)
}
```
