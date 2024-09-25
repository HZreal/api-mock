package vo

import "gin-init/model/entity"

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: user.vo.go
 * @Description:
 */

type ApiDetailInfo entity.ApiModel

// type ApiDetailInfo struct {
// 	Id     uint                 `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
// 	Name   string               `json:"name" gorm:"column:name"`
// 	Url    string               `json:"url" gorm:"column:url"`
// 	Method string               `json:"method" gorm:"column:method"`
// 	Params []entity.ParamStruct `json:"params" gorm:"column:params;serializer:json"`
// }
