package CRUD

import (
	"cronService/Database"
	"cronService/Models"
	"time"
)
// simply creating cj what all inputs do we get from user?
//and whats getting passed and autherization thing needs to be updated.

func CreateCronjob(CronJob *Models.CronJob, username string) (err error)   {

	if err := Database.DB.Create(CronJob).Error; err != nil {
		return err
	}
	var user Models.User
	if err := GetUserFromUserAuth(username, &user); err!= nil{
		return err
	}
	Database.DB.Model(CronJob).Update(Models.CronJob{UserId:user.Id, CreatedAt: time.Now(), OrganizationId:user.OrganisationID})
	return nil
}
