package users

func userCredentialsMatch(email string, pwd string) bool {
	l := len(UsersSlice)

	for i := 0; i < l; i++ {
		if UsersSlice[i].Email == email && UsersSlice[i].Password == pwd {
			return true
		}
	}

	return false
}
