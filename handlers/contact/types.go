package contact

import (
	"fmt"
)

// Contact contains the user sent data from website
type Contact struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// GetName returns the name
func (c *Contact) GetName() string {
	return c.Name
}

// GetEmail returns the email
func (c *Contact) GetEmail() string {
	return c.Email
}

// GetMessage returns the message
func (c *Contact) GetMessage() string {
	return c.Message
}

//ToString returns the entire struct as string
func (c *Contact) ToString() string {
	return fmt.Sprintf("Name: %s \n Email: %s \n Message: %s", c.Name, c.Email, c.Message)
}
