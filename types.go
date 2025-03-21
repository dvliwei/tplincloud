/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  type
 * @Version: 1.0.0
 * @Date: 2025/03/20 10:37
 * @Update liwei 2025/3/20 10:37
 */

package tplinkcloud

type (
	TplinkCloudConfig struct {
		Ak         string
		Sk         string
		TerminalId string
	}
)

// 项目下设备列表结构体
type (
	GetDeviceListInProjectApplicationRequest struct {
		Start     int `json:"start"`
		Limit     int `json:"limit"`
		FilterAnd struct {
			HasChild       int      `json:"hasChild"`
			DeviceTypeList []string `json:"deviceTypeList"`
		} `json:"filterAnd"`
	}
	// 定义与 JSON 数据对应的结构体
	DeviceInfo struct {
		QrCode       string `json:"qrCode"`
		DeviceName   string `json:"deviceName"`
		DeviceType   string `json:"deviceType"`
		DeviceStatus int    `json:"deviceStatus"`
		OpenType     int    `json:"openType"`
		OpenStatus   int    `json:"openStatus"`
		DeviceModel  string `json:"deviceModel"`
		IP           string `json:"ip"`
		MAC          string `json:"mac"`
		RegionName   string `json:"regionName"`
		RegionId     string `json:"regionId"`
		ChannelId    int    `json:"channelId"`
	}

	Result struct {
		Total int          `json:"total"`
		List  []DeviceInfo `json:"list"`
	}

	GetDeviceListInProjectApplicationResponseData struct {
		Result    Result `json:"result"`
		ErrorCode int    `json:"error_code"`
	}
)

//1.4.4.4.5设置云台转动、调焦

type (
	MotionCtrlRequest struct {
		QrCode    string `json:"qrCode"`
		ChannelId int    `json:"channelId"`
		Direction int    `json:"direction"` //云台操作类型 0停止，1向右 2右下，3下 4左下，5左 6左上，7上 8右上，9 扫描，10变倍+ ，11变倍-，12调焦+，13调焦-，14调焦停止，15云台复位
		Speed     int    `json:"speed"`     //云台速度大小，取值范围：[1, 7]，默认值7
	}
	MotionCtrlResponse struct {
		ErrorCode int `json:"error_code"`
	}
)

// 1.4.4.4.7控制云台设备回到原点所在位置
type (
	ResetPtzDevicePositionRequest struct {
		QrCode    string `json:"qrCode"`
		ChannelId int    `json:"channelId"`
	}
	ResetPtzDevicePositionResponse struct {
		ErrorCode int `json:"error_code"`
	}
)

type (
	VideoStreamAddressRequest struct {
		QrCode     string `json:"qrCode"`
		Resolution int    `json:"resolution"`
		ClientType string `json:"clientType"`
		ChannelId  int    `json:"channelId"`
		StreamType string `json:"streamType"`
	}

	VideoStreamAddressResponse struct {
		ErrorCode int `json:"error_code"`
		Result    struct {
			SdkStreamUrl string `json:"sdkStreamUrl"`
		} `json:"result"`
	}
)

// 搜索存在回放数据的日期
type (
	Dev struct {
		QrCode    string `json:"qrCode"`
		ChannelId int    `json:"channelId"`
	}
	SearchYearRequest struct {
		DevList []Dev  `json:"devList"`
		Year    string `json:"year"`
	}

	DeviceVideoInfo struct {
		QrCode    string   `json:"qrCode"`
		ChannelId int      `json:"channelId"`
		Dates     []string `json:"dates"`
	}

	SearchYearResult struct {
		DeviceVideoInfoList []DeviceVideoInfo `json:"deviceVideoInfoList"`
	}
	SearchYearFail struct {
		QrCode    string `json:"qrCode"`
		Mac       string `json:"mac"`
		ChannelId int    `json:"channelId"`
		ErrorCode int    `json:"error_code"`
	}
	SearchYearResponse struct {
		Result    SearchYearResult `json:"result"`
		FailList  SearchYearFail   `json:"failList,omitempty"`
		ErrorCode int              `json:"error_code"`
	}
)

type (
	SearchVideoRequest struct {
		QrCode    string `json:"qrCode"`
		ChannelId int    `json:"channelId"`
		SearchDay string `json:"searchDay"` //搜索的日期，yyyyMMDD的格式
		StartIdx  int    `json:"startIdx"`  //录像事件开始索引
		EndIdx    int    `json:"endIdx"`    //录像事件结束索引，endIdx和startIdx差值不能超过100
	}

	SearchVideoResult struct {
		UserId int `json:"userId"` //回放用户id，用于区分并限制同时观看回放视频的用户数
		Videos []struct {
			StartTime  int64 `json:"startTime"`  //开始时间，单位：毫秒
			EndTime    int64 `json:"endTime"`    //结束时间，单位：毫秒
			VideoType  []int `json:"videoType"`  //录像类型
			SourceType int   `json:"sourceType"` //录像存储类型
		} `json:"videos"`
	}
	SearchVideoResponse struct {
		Result    SearchVideoResult `json:"result"`
		ErrorCode int               `json:"error_code"`
	}
)

type (
	SearchHumanDetectionVideoRequest struct {
		QrCode    string `json:"qrCode"`
		ChannelId int    `json:"channelId"`
		SearchDay string `json:"searchDay"` //搜索的日期，yyyyMMDD的格式
	}
	SearchHumanDetectionVideoResponse struct {
		Result    SearchVideoResult `json:"result"`
		ErrorCode int               `json:"error_code"`
	}
)
