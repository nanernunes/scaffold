package scaffold

import (
	"reflect"

	"gorm.io/gorm"
)

type ORM struct {
	db    *gorm.DB
	Model interface{}
}

type ID struct {
	ID uint `json:"id"`
}

func NewORM(db *gorm.DB, model interface{}) *ORM {
	return &ORM{db: db, Model: model}
}

func (c *ORM) NewInstance() interface{} {
	t := reflect.TypeOf(c.Model)
	return reflect.New(t).Interface()
}

func (c *ORM) NewInstances() interface{} {
	t := reflect.TypeOf(c.Model)
	empty := reflect.MakeSlice(reflect.SliceOf(t), 0, 0)
	return empty.Interface()

}

func (c *ORM) Create(model interface{}) error {
	return c.db.Model(c.Model).Create(model).Error
}

func (c *ORM) FindByID(id uint, model interface{}) error {
	return c.db.First(c.Model, id).Error
}

func (c *ORM) FindByUUID(uuid string, model interface{}) error {
	return c.db.Model(c.Model).Where("uuid = ?", uuid).First(model).Error
}

func (c *ORM) GetAll(models interface{}) error {
	return c.db.Find(models).Error
}

func (c *ORM) Update(id uint, model interface{}) error {
	return c.db.Model(c.Model).Where("id = ?", id).Updates(model).Error
}

func (c *ORM) UpdateByUUID(uuid string, model interface{}) error {
	t := ID{}
	c.FindByUUID(uuid, &t)

	return c.db.Model(c.Model).Where("id = ?", t.ID).Updates(model).Error
}

func (c *ORM) Delete(id uint) error {
	return c.db.Model(c.Model).Where("id = ?", id).Delete(c.NewInstance()).Error
}

func (c *ORM) DeleteByUUID(uuid string) error {
	t := ID{}
	c.FindByUUID(uuid, &t)

	return c.Delete(t.ID)
}
