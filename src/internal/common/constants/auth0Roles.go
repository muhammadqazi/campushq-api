package constants

type Role struct {
	Key string
	ID  string
}

var Auth0Roles = map[string]Role{
	"SUPER_USER": {
		Key: "SUPER_USER",
		ID:  "rol_dazEPnck0Z7wOooC",
	},
	"STUDENT": {
		Key: "STUDENT",
		ID:  "rol_5TCJpWRjweumpb0m",
	},
	"INSTRUCTOR": {
		Key: "INSTRUCTOR",
		ID:  "rol_SXmuhNKaSE2j8272",
	},
}
