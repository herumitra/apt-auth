package seeders

import (
	config "github.com/herumitra/apt-auth/config"
	models "github.com/herumitra/apt-auth/models"
)

func UserBranchSeed() {
	userBranch := []models.UserBranch{
		{UserID: "USR01072023004", BranchID: "BRC11122024001"},
		{UserID: "USR01072023002", BranchID: "BRC11122024001"},
		{UserID: "USR01072023004", BranchID: "BRC11122024002"},
	}
	config.DB.Create(&userBranch)
}
