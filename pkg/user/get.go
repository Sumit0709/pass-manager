package user

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/Sumit0709/pass-manager/pkg/auth"
)

// Returns user data matching to user
func Get(user User, secret string) ([]User, error) {

	// Authentication
	if err := auth.Authenticate(secret, filepath.Join(Dir, localDir, envFile)); err != nil {
		return nil, err
	}
	// Read data from file
	var users []User
	users, err := read(os.O_RDONLY)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return nil, errors.New("no user exists. need to add a user first")
		}
		return nil, err
	}

	// Gather matching values
	var result []User
	for _, val := range users {
		if user.match(val) {

			val.Password, err = decrypt(secret, val.Password)
			if err != nil {
				return nil, errors.New("error decrypting password")
			}
			result = append(result, val)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("no such user exists")
	}
	return result, nil
}
