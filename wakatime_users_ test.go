package wakatime_users

import (
	"log"
	"os"
	"testing"

	context "golang.org/x/net/context"

	"errors"

	"google.golang.org/appengine/aetest"
)

var (
	ctx     context.Context
	testErr error
	done    func()
)

// TestMain initializes tests
func TestMain(m *testing.M) {
	ctx, done, testErr = aetest.NewContext()
	if testErr != nil {
		log.Fatalf("Failed to initialize app engine test context. %v", testErr)
	}
	defer done()
	os.Exit(m.Run())
}

// TestCreate tests create
func TestCreate(t *testing.T) {
	user := &User{
		Email:  "someemaile@idonno.com",
		APIKey: "someapikey",
	}
	testErr = user.Create(ctx)
	if testErr != nil {
		t.Error(testErr)
		log.Println(testErr)
		return
	}
	user, testErr = GetUser(ctx, user.ID)
	if testErr != nil {
		t.Error(testErr)
		log.Println(testErr)
		return
	}
	log.Println("Created user: ", *user)
}

// TestGetUsers tests GetUsers
func TestGetUsers(t *testing.T) {
	users, testErr := GetUsers(ctx)
	if testErr != nil {
		t.Error(testErr)
		log.Println(testErr)
		return
	}
	log.Println("Created user: ", users)
}

// TestGetUser tests GetUser
func TestGetUser(t *testing.T) {
	TestCreate(t)
}

// TestDelete tests Delete
func TestDelete(t *testing.T) {
	user := &User{
		Email:  "someemaile@idonno.com",
		APIKey: "someapikey",
	}
	testErr = user.Create(ctx)
	if testErr != nil {
		t.Error(testErr)
		log.Println(testErr)
		return
	}
	testErr = user.Delete(ctx)
	if testErr != nil {
		t.Error(testErr)
		log.Println(testErr)
		return
	}
	user, testErr = GetUser(ctx, user.ID)
	if testErr == nil {
		testErr = errors.New("Deletion failed. User still exists.")
		t.Error(testErr)
		log.Println(testErr)
		return
	}
}
