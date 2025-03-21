/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  tplinkcloud_fac
 * @Version: 1.0.0
 * @Date: 2025/03/21 10:06
 * @Update liwei 2025/3/21 10:06
 */

package tplinkcloud

type TplinkCloudFactory interface {
	MakeTplinkCloud() IsTplinkCloud
}

type IsTplinkCloud interface {
	//安防设备2.0接入
	//1.4.3.3.3获取项目型应用下设备列表
	GetIpcList() ([]DeviceInfo, error)

	//1.4.4.4.5设置云台转动、调焦
	MotionCtrl(qrCode string, direction int, speed int) error

	//1.4.4.4.7控制云台设备回到原点所在位置
	ResetPtzDevicePosition(qrCode string) error

	//1.4.5.3获取音视频流访问地址
	RequestStreamUrl(qrCode string, clientType string, streamType string, resolution int) (*string, error)

	//1.4.5.7搜索存在回放录像的日期
	SearchYear(qrCodes []Dev, year string) (*SearchYearResult, error)

	//1.4.5.8搜索当天录像数据
	SearchVideo(qrCode string, searchDay string, startIdx int, endIdx int) (*SearchVideoResult, error)

	//1.4.5.9搜索当天所有人形侦测录像数据
	SearchHumanDetectionVideo(qrCode string, searchDay string) (*SearchVideoResult, error)

	//1.4.9.2提交设备抓图任务
	SubmitCaptureImageTask(qrCode string, imageType int, expireDays int, playbackTime string, imageId string) (string, error)

	//1.4.9.5查询任务详情
	GetTaskInfo(taskId string) (*GetTaskInfoResponse, error)

	//1.4.9.6分页查询任务文件列表
	GetTaskFilePage(taskId string, pageNum int, pageSize int) (*GetTaskFilePageResponse, error)
}
