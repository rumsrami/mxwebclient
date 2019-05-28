package secure

// // Token Create a new token object, specifying signing method and the claims
// func Token(hmacsecret string, claim *jwt.MapClaims) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
// 	tokenString, err := token.SignedString(hmacsecret)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }
