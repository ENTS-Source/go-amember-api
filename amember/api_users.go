package amember

import (
	"net/url"
	"time"
)

type User struct {
	Id        int        `json:"user_id"`
	FirstName string     `json:"name_f"`
	LastName  string     `json:"name_l"`
	Nested    UserNested `json:"nested"`

	Nickname  string   `json:"nickname"`   // Makerspace extension
	Fob       string   `json:"fob"`        // Makerspace extension
	FobAccess string   `json:"fob_access"` // Makerspace extension
	Announce  []string `json:"announce"`   // Makerspace extension
}

type UserNested struct {
	Access []UserAccess `json:"access,flow"`
}

type UserAccess struct {
	// Should be an int, but it's a string
	ProductId string `json:"product_id"`
	// Format: YYYY-MM-DD
	BeginDate string `json:"begin_date"`
	// Format: YYYY-MM-DD
	EndDate string `json:"expire_date"`
}

func (u *User) Name() string {
	name := u.FirstName + " " + string(u.LastName[0]) + "."
	if u.Nickname != "" {
		return u.Nickname
	}
	return name
}

// FindUsersByFob
// Makerspace extension
func (c *Client) FindUsersByFob(fob string) ([]User, error) {
	query := url.Values{}
	query.Set("_filter[fob]", fob)
	query.Set("_nested[]", "access")
	return allPages[User](c, "/api/users", query)
}

func (a UserAccess) IsActiveOn(date time.Time) (bool, error) {
	start, err := time.Parse("2006-01-02", a.BeginDate)
	if err != nil {
		return false, err
	}
	end, err := time.Parse("2006-01-02", a.EndDate)
	if err != nil {
		return false, err
	}
	return date.After(start) && (date.Before(end) || date.Equal(end)), nil
}
