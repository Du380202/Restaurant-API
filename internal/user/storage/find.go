package storage

import (
	"context"
	"restaurant/internal/user/model"
)

func (s *userStore) FindUser(
	context context.Context,
	condition map[string]interface{},
	moreInfo ...string,
) (*model.User, error) {
	db := s.db.Begin()

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var data model.User

	if err := db.Where(condition).First(&data).Error; err != nil {
		// if err == gorm.ErrRecordNotFound {
		// 	return nil, err
		// }
		// fmt.Println(err)
		return nil, err
	}

	// fmt.Println(data)

	return &data, nil
}
