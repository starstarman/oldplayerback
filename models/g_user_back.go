package models

type GUserBack struct {
	Userid    string `xorm:"not null VARCHAR(20)"`
	Partnerid int    `xorm:"comment('合作商Id') INT(11)"`
}
