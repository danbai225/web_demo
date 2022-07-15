package admin

import (
	"web_demo/internal/pkg/core"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/mysql/key_map"
	"web_demo/pkg/errors"
)

func (s *service) SettingGet(ctx core.Context, key string) (string, error) {
	first, err := key_map.NewQueryBuilder().WhereK(mysql.EqualPredicate, key).First(s.db.GetDb())
	if err != nil {
		return "", errors.DbErr(s.logger, err)
	}
	return first.Val, nil
}
func (s *service) SettingPost(ctx core.Context, key, val string) (err error) {
	_, err = s.SettingGet(ctx, key)
	if err == nil {
		err = s.db.GetDb().Table("key_map").Where("k=?", key).Update("val", val).Error
	} else {
		err = s.db.GetDb().Create(&key_map.KeyMap{
			K:   key,
			Val: val,
		}).Error
	}
	return errors.DbErr(s.logger, err)
}
