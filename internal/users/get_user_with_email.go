package users

func getIndexOfUserFromEmail(email string) int {
	l := len(UsersSlice)
	for i := 0; i < l; i++ {
		if UsersSlice[i].Email == email {
			return i
		}
	}
	return -1
}
