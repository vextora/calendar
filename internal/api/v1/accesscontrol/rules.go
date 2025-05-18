package accesscontrol

func UserCanEditArticle(userID string, articleOwnerID string) bool {
	return userID == articleOwnerID
}
