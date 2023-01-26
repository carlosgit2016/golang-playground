package entities

import (
	"github.com/icrowley/fake"
)

/*

Name: John Doe
Address: 123 Main St, Anytown USA 12345
Phone Number: 555-555-5555
Email: johndoe@example.com

Name: Jane Smith
Address: 456 Park Ave, Anothertown USA 67890
Phone Number: 555-555-5556
Email: janesmith@example.com

Name: Michael Johnson
Address: 789 Elm St, This Town USA 98765
Phone Number: 555-555-5557
Email: michaeljohnson@example.com

*/

type Person struct {
	Name    string
	Address string
	Phone   string
	Email   string
}

func NewPerson() *Person {

	return &Person{
		Name:    fake.FullName(),
		Address: fake.StreetAddress(),
		Phone:   fake.Phone(),
		Email:   fake.EmailAddress(),
	}

}
