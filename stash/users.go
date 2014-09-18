package stash

type User struct {
	Username string `"json:username"`
}

type UserResource struct {
	client *Client
}

// Get list of users
func (u *UserResource) Current() (User, error) {
	var user = User{}
	var path = "username"

	if err := u.client.do("GET", "core", path, nil, nil, &user); err != nil {
		return user, err
	}

	return user, nil
}
