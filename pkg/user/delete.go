package user

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Sumit0709/pass-manager/pkg/auth"
)

func Delete(user User, secret string, force bool) ([]User, error) {

	var users []User

	// Authenticate user
	err := auth.Authenticate(secret, filepath.Join(Dir, localDir, envFile))
	if err != nil {
		return nil, err
	}

	// Filter users
	users, err = read(os.O_RDWR)
	if err != nil {
		return nil, err
	}
	filtered, matched := filter(user, users)
	if len(matched) == 0 {
		return nil, errors.New("no such user exists")
	}

	if len(matched) > 1 && !force {
		return users, nil
	}
	err = write(filtered)
	if err != nil {
		return matched, err
	}
	return matched, nil

}

// Filters all users matching to u in the slice
func filter(u User, users []User) ([]User, []User) {
	var filtered []User
	var matched []User

	for _, val := range users {
		if !u.match(val) {
			filtered = append(filtered, val)
		} else {
			matched = append(matched, val)
		}
	}
	return filtered, matched

}
