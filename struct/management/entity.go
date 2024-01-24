package management

import "fmt"

// User harus pake kapital supaya bisa dipake di package yang lain
// Semua fieldnya juga harus pake kapital
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// Method
// mirip function tapi sebuah method dimiliki oleh sebuah struct
func (user User) Display() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)

}

func (group Group) DisplayGroup() {
	fmt.Printf("Group Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count: %d", len(group.Users))
	fmt.Println("")
	fmt.Println("Users:")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
