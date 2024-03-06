package users

func userExists(email string) bool {
	for i := 0; i < len(UsersSlice); i++ {
		if UsersSlice[i].Email == email {
			return true
		}
	}
	return false
}
