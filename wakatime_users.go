package wakatime_users

import (
	"time"

	context "golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
)

// User
type User struct {
	ID        int64     `json:"id" datastore:"-"`
	APIKey    string    `json:"api_key" form:"api_key"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// key returns the a datastore key based of the user id
func (user *User) key(c context.Context) *datastore.Key {
	if user.ID == 0 {
		return datastore.NewIncompleteKey(c, "User", nil)
	}
	return datastore.NewKey(c, "User", "", user.ID, nil)
}

// Create creates a new user on google datastore
func (user *User) Create(c context.Context) error {
	key, err := datastore.Put(c, user.key(c), user)
	if err != nil {
		return err
	}
	user.ID = key.IntID()
	return nil
}

// GetUsers fetches all users
func GetUsers(c context.Context) ([]User, error) {
	query := datastore.NewQuery("User").Order("CreatedAt")
	var users []User
	keys, err := query.GetAll(c, &users)
	if err != nil {
		return users, err
	}
	usersLen := len(users)
	for i := 0; i < usersLen; i++ {
		users[i].ID = keys[i].IntID()
	}
	return users, nil
}

// GetUser gets a user by ID
func GetUser(c context.Context, id int64) (*User, error) {
	var user User
	user.ID = id
	key := user.key(c)
	err := datastore.Get(c, key, &user)
	if err != nil {
		return nil, err
	}
	user.ID = key.IntID()
	return &user, nil
}

// Delete deletes a user
func (user *User) Delete(c context.Context) error {
	err := datastore.Delete(c, user.key(c))
	if err != nil {
		return err
	}
	return nil
}
