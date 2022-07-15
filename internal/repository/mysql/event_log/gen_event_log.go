///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package event_log

import (
	"fmt"
	"time"

	"web_demo/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *EventLog {
	return new(EventLog)
}

func NewQueryBuilder() *eventLogQueryBuilder {
	return new(eventLogQueryBuilder)
}

func (t *EventLog) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type eventLogQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *eventLogQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *eventLogQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&EventLog{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *eventLogQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&EventLog{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *eventLogQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&EventLog{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *eventLogQueryBuilder) First(db *gorm.DB) (*EventLog, error) {
	ret := &EventLog{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *eventLogQueryBuilder) QueryOne(db *gorm.DB) (*EventLog, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *eventLogQueryBuilder) QueryAll(db *gorm.DB) ([]*EventLog, error) {
	var ret []*EventLog
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *eventLogQueryBuilder) Limit(limit int) *eventLogQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *eventLogQueryBuilder) Offset(offset int) *eventLogQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *eventLogQueryBuilder) WhereId(p mysql.Predicate, value int64) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereIdIn(value []int64) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereIdNotIn(value []int64) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) OrderById(asc bool) *eventLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *eventLogQueryBuilder) WhereType(p mysql.Predicate, value string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereTypeIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereTypeNotIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) OrderByType(asc bool) *eventLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}

func (qb *eventLogQueryBuilder) WhereDevice(p mysql.Predicate, value string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "device", p),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereDeviceIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "device", "IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereDeviceNotIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "device", "NOT IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) OrderByDevice(asc bool) *eventLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "device "+order)
	return qb
}

func (qb *eventLogQueryBuilder) WhereUsername(p mysql.Predicate, value string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", p),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereUsernameIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", "IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereUsernameNotIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", "NOT IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) OrderByUsername(asc bool) *eventLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "username "+order)
	return qb
}

func (qb *eventLogQueryBuilder) WhereUserId(p mysql.Predicate, value int64) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", p),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereUserIdIn(value []int64) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereUserIdNotIn(value []int64) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) OrderByUserId(asc bool) *eventLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "user_id "+order)
	return qb
}

func (qb *eventLogQueryBuilder) WhereIp(p mysql.Predicate, value string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ip", p),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereIpIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ip", "IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereIpNotIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ip", "NOT IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) OrderByIp(asc bool) *eventLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "ip "+order)
	return qb
}

func (qb *eventLogQueryBuilder) WhereLocation(p mysql.Predicate, value string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "location", p),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereLocationIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "location", "IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereLocationNotIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "location", "NOT IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) OrderByLocation(asc bool) *eventLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "location "+order)
	return qb
}

func (qb *eventLogQueryBuilder) WhereContent(p mysql.Predicate, value string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", p),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereContentIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereContentNotIn(value []string) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "NOT IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) OrderByContent(asc bool) *eventLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "content "+order)
	return qb
}

func (qb *eventLogQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereCreatedAtIn(value []time.Time) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *eventLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *eventLogQueryBuilder) OrderByCreatedAt(asc bool) *eventLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}