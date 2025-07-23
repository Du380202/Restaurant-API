package rstorage

import (
	"context"
	"restaurant/common"
	rmodel "restaurant/internal/restaurant/model"
)

func (s *restaurantStore) ListDataWithCondition(
	context context.Context,
	filter *rmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]rmodel.Restaurant, error) {
	var listData []rmodel.Restaurant

	db := s.db.Table(rmodel.Restaurant{}.TableName()).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&listData).Error; err != nil {
		return nil, err
	}

	if len(listData) > 0 {
		last := listData[len(listData) -1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return listData, nil

}
