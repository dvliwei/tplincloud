/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  tplinkcloud
 * @Version: 1.0.0
 * @Date: 2025/03/21 10:05
 * @Update liwei 2025/3/21 10:05
 */

package tplinkcloud

type TplinkCloud struct {
	config *TplinkCloudConfig
}

func NewTplinkCloud(config *TplinkCloudConfig) *TplinkCloud {
	return &TplinkCloud{
		config: config,
	}
}

func (t *TplinkCloud) MakeTplinkCloud() IsTplinkCloud {
	return newTplinkCloudIPCRes(t)
}
