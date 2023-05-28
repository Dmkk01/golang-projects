package user

import (
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
)

func TestMain(m *testing.M) {
	m.Run()

	os.Remove(dbPath)
}

func cleanDb(b *testing.B) {

	os.Remove(dbPath)

	u := &User{
		ID:   bson.NewObjectId(),
		Name: "Test",
		Role: "Tester",
	}
	err := u.Save()

	if err != nil {
		b.Fatalf("Error saving user: %s", err)
	}

	b.ResetTimer()
}

func BenchmarkCreate(b *testing.B) {

	cleanDb(b)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "Test_" + strconv.Itoa(i),
			Role: "Tester",
		}
		b.StartTimer()
		err := u.Save()

		if err != nil {
			b.Fatalf("Error saving user: %s", err)
		}
	}
}

func BenchmarkCRUD(b *testing.B) {

	os.Remove(dbPath)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "Test User",
			Role: "Tester",
		}

		err := u.Save()

		if err != nil {
			b.Fatalf("Error saving user: %s", err)
		}

		_, err = One(u.ID)

		if err != nil {
			b.Fatalf("Error reading user: %s", err)
		}

		u.Role = "Admin"

		err = u.Save()

		if err != nil {
			b.Fatalf("Error saving user: %s", err)
		}

		_, err = One(u.ID)

		if err != nil {
			b.Fatalf("Error reading user: %s", err)
		}
		err = Delete(u.ID)

		if err != nil {
			b.Fatalf("Error deleting user: %s", err)
		}
	}

}

func TestCRUD(t *testing.T) {
	t.Log("Testing CRUD operations")

	u := &User{
		ID:   bson.NewObjectId(),
		Name: "Test User",
		Role: "Tester",
	}

	err := u.Save()

	if err != nil {
		t.Fatalf("Error saving user: %s", err)
	}

	t.Log("Read")

	u2, err := One(u.ID)

	if err != nil {
		t.Fatalf("Error reading user: %s", err)
	}

	if !reflect.DeepEqual(u, u2) {
		t.Fatalf("Read user is not the same as the written one")
	}

	t.Log("Update")

	u.Role = "Admin"

	err = u.Save()

	if err != nil {
		t.Fatalf("Error saving user: %s", err)
	}

	u3, err := One(u.ID)

	if err != nil {
		t.Fatalf("Error reading user: %s", err)
	}

	if !reflect.DeepEqual(u3, u) {
		t.Fatalf("Read user is not the same as the written one")
	}

	t.Log("Delete")
	err = Delete(u.ID)

	if err != nil {
		t.Fatalf("Error deleting user: %s", err)
	}

	_, err = One(u.ID)

	if err == nil {
		t.Fatalf("User was not deleted")
	}

	if err != storm.ErrNotFound {
		t.Fatalf("Error deleting user: %s", err)
	}

	t.Log("Read All")

	u2.ID = bson.NewObjectId()
	u3.ID = bson.NewObjectId()

	err = u.Save()

	if err != nil {
		t.Fatalf("Error saving user: %s", err)
	}
	err = u2.Save()

	if err != nil {
		t.Fatalf("Error saving user: %s", err)
	}
	err = u3.Save()

	if err != nil {
		t.Fatalf("Error saving user: %s", err)
	}

	users, err := All()

	if err != nil {
		t.Fatalf("Error reading users: %s", err)
	}

	if len(users) != 3 {
		t.Fatalf("Wrong number of users returned")
	}

}
