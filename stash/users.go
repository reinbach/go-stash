package stash

type User struct {
	Username string `"json:username"`
}

type UserResource struct {
	client *Client
}

func (r *UserResource) Current() (*User, error) {
	user := User{}
	if err := r.client.do("GET", "/", nil, nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
