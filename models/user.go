package models

import (
	"GoDeveloperVK-testTask/utils"
)

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Role      string `json:"role" db:"role"`
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserTable() error {
	_, err := utils.DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(100),
            last_name VARCHAR(100),
            email VARCHAR(100),
            password VARCHAR(100),
            role VARCHAR(100),
            UNIQUE(email)
        )
    `)
	if err != nil {
		return err
	}

	return nil
}
