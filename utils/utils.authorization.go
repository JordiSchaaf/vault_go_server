package utils

import (
	"fmt"
	"vault/server/models"
)

func IsUserAuthorized(issuerId string) bool {
	fmt.Println("Ik kom hier")
	issuer := models.User{}
	res := DB.Where("id = ?", issuerId).First(&issuer)
	fmt.Println(res)
	if res.Error != nil {
		return false
	}
	// TODO Logica om user op te zoeken in een permission tabel
	if issuer.PermissionLevel <= 2 {
		return true
	}
	return false
}
