package amember

import "net/url"

type User struct {
	Id        int    `json:"user_id"`
	FirstName string `json:"name_f"`
	LastName  string `json:"name_l"`

	Nickname  string   `json:"nickname"`   // Makerspace extension
	Fob       string   `json:"fob"`        // Makerspace extension
	FobAccess string   `json:"fob_access"` // Makerspace extension
	Announce  []string `json:"announce"`   // Makerspace extension
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
	return allPages[User](c, "/api/users", query)
}
