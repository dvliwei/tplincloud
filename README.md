# tplincloud

## tplink商云平台SDK接入

## 项目介绍
tplincloud是一个基于tplink商云2.0 主要实现IPC设备信息采集，视频预览，回放，双向对讲，云台操作，截图录像

#### 方法介绍
```azure
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
```

#### 下载安装
```azure
go get github.com/dvliwei/tplincloud@
```

#### 使用
````azure
	var tpconfig tplinkcloud.TplinkCloudConfig
	tpconfig.Ak = "xxxxxx"
	tpconfig.Sk = "xxxxxx"
	tpconfig.TerminalId = "xxxxxx"
	tplincloudRes := tplinkcloud.NewTplinkCloud(&tpconfig)
	result, err := tplincloudRes.MakeTplinkCloud().GetIpcList()
````