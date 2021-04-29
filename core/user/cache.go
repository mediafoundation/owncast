package user

var _userIdCache = map[string]*User{}
var _userAccessTokenCache = map[string]*User{}

func getCachedIdUser(id string) *User {
	return _userIdCache[id]
}

func setCachedIdUser(id string, user *User) {
	_userIdCache[id] = user
}

func getCachedAccessTokenUser(id string) *User {
	return _userAccessTokenCache[id]
}

func setCachedAccessTokenUser(id string, user *User) {
	_userAccessTokenCache[id] = user
}
