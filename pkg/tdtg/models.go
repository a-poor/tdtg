package tdtg

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/teris-io/shortid"
	"golang.org/x/crypto/bcrypt"
)

var (
	genUserID = shortid.MustNew(1, shortid.DefaultABC, 0)
	genListID = shortid.MustNew(1, shortid.DefaultABC, 1)
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	PassHash string `json:"passHash"`
}

func NewUser(name, rawPass string) (*User, error) {
	// Generate a random user ID
	id, err := genUserID.Generate()
	if err != nil {
		return nil, fmt.Errorf("unable to generate shortid: %w", err)
	}

	// Hash the password
	h, err := bcrypt.GenerateFromPassword([]byte(rawPass), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("unable to hash password: %w", err)
	}

	// Return the user
	return &User{
		ID:       "u" + id,
		Username: name,
		PassHash: string(h),
	}, nil
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.ID, validation.Required),
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.PassHash, validation.Required),
	)
}

type ToDoList struct {
	ID        string      `json:"id"`
	Title     string      `json:"title"`
	CreatedBy string      `json:"createdBy"`
	CreatedAt time.Time   `json:"createdAt"`
	Items     []*ToDoItem `json:"items"`
}

func NewToDoList(title, createdBy string) (*ToDoList, error) {
	// Generate a random user ID
	id, err := genListID.Generate()
	if err != nil {
		return nil, fmt.Errorf("unable to generate shortid: %w", err)
	}

	return &ToDoList{
		ID:        "l" + id,
		Title:     title,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}, nil
}

func (l *ToDoList) Validate() error {
	return validation.ValidateStruct(
		l,
		validation.Field(&l.ID, validation.Required),
		validation.Field(&l.Title, validation.Required),
		validation.Field(&l.CreatedBy, validation.Required),
		validation.Field(&l.CreatedAt, validation.Required),
		validation.Field(&l.Items, validation.Each(validation.NotNil)),
	)
}

type ToDoItem struct {
	Text        string    `json:"text"`
	Description string    `json:"description,omitempty"`
	AddedAt     time.Time `json:"addedAt"`
	AddedBy     string    `json:"addedBy"`
	Completed   bool      `json:"completed"`
	CompletedAt time.Time `json:"completedAt,omitempty"`
	CompletedBy string    `json:"completedBy,omitempty"`
}

func NewToDoItem(txt, desc, addBy string) (*ToDoItem, error) {
	return &ToDoItem{
		Text:        txt,
		Description: desc,
	}, nil
}
