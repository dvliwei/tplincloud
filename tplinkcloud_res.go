/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  tplinkcloudipcv2_res
 * @Version: 1.0.0
 * @Date: 2025/03/20 09:57
 * @Update liwei 2025/3/20 09:57
 */

package tplinkcloud

import (
	"encoding/json"
	"fmt"
)

type tplinkCloudIPCRes struct {
	TplinkCloudBase
}

func newTplinkCloudIPCRes(t *TplinkCloud) *tplinkCloudIPCRes {
	return &tplinkCloudIPCRes{
		TplinkCloudBase{
			Ak:         t.config.Ak,
			Sk:         t.config.Sk,
			TerminalId: t.config.TerminalId,
		},
	}
}

// GetIpcListV2 安防设备2.0接入
func (res *tplinkCloudIPCRes) GetIpcList() ([]DeviceInfo, error) {
	var (
		request GetDeviceListInProjectApplicationRequest
	)

	request.Start = 0
	request.Limit = 100
	request.FilterAnd.HasChild = 1
	request.FilterAnd.DeviceTypeList = []string{"SURVEILLANCECAMERA"}
	res.TplinkCloudBase.Path = "/tums/open/deviceManager/v1/getDeviceListInProjectApplication"
	res.TplinkCloudBase.Payload = request
	//首次请求
	respBytes, err := res.PostRequest()
	if err != nil {
		return nil, err
	}
	var response GetDeviceListInProjectApplicationResponseData
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return nil, err
	}
	result := make([]DeviceInfo, 0, response.Result.Total)
	result = append(result, response.Result.List...)
	total := response.Result.Total
	fmt.Println(total)
	if total > request.Limit {
		pages := (total + request.Limit - 1) / request.Limit
		for i := 1; i < pages; i++ {
			request.Start = i * request.Limit
			res.TplinkCloudBase.Payload = request
			respBytes, err := res.PostRequest()
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(respBytes, &response); err != nil {
				return nil, err
			}
			result = append(result, response.Result.List...)
		}
	}
	return result, nil
}

// 1.4.4.4.5设置云台转动、调焦
func (res *tplinkCloudIPCRes) MotionCtrl(qrCode string, direction int, speed int) error {
	var (
		err      error
		request  MotionCtrlRequest
		response MotionCtrlResponse
	)
	request.QrCode = qrCode
	request.ChannelId = 1
	request.Direction = direction
	request.Speed = speed
	res.TplinkCloudBase.Path = "/vms/open/deviceConfig/v1/motionCtrl"
	res.TplinkCloudBase.Payload = request
	respBytes, err := res.PostRequest()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return err
	}
	if response.ErrorCode != 0 {
		return fmt.Errorf("motionCtrl error: %d", response.ErrorCode)
	}
	return nil
}

func (res *tplinkCloudIPCRes) ResetPtzDevicePosition(qrCode string) error {
	var (
		err      error
		request  ResetPtzDevicePositionRequest
		response ResetPtzDevicePositionResponse
	)
	request.QrCode = qrCode
	request.ChannelId = 1
	res.TplinkCloudBase.Path = "/vms/open/deviceConfig/v1/resetPtzDevicePosition"
	res.TplinkCloudBase.Payload = request
	respBytes, err := res.PostRequest()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return err
	}
	if response.ErrorCode != 0 {
		return fmt.Errorf("resetPtzDevicePosition error: %d", response.ErrorCode)
	}
	return nil
}

// 获取TP-LINK私有协议音视频流访问地址
func (res *tplinkCloudIPCRes) RequestStreamUrl(qrCode string, clientType string, streamType string, resolution int) (*string, error) {
	var (
		err      error
		request  VideoStreamAddressRequest
		response VideoStreamAddressResponse
	)
	request.QrCode = qrCode
	request.ClientType = clientType
	request.StreamType = streamType
	request.ChannelId = 1
	request.Resolution = resolution
	res.TplinkCloudBase.Path = "/vms/open/webServer/v1/requestStreamUrl"
	res.TplinkCloudBase.Payload = request
	respBytes, err := res.PostRequest()
	fmt.Println(string(respBytes))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return nil, err
	}
	if response.ErrorCode != 0 {
		return nil, fmt.Errorf("requestStreamUrl error: %d", response.ErrorCode)
	}
	return &response.Result.SdkStreamUrl, nil
}

// 搜索存在回放数据的日期
func (res *tplinkCloudIPCRes) SearchYear(qrCodes []Dev, year string) (*SearchYearResult, error) {
	var (
		err      error
		request  SearchYearRequest
		response SearchYearResponse
	)
	request.DevList = qrCodes
	request.Year = year
	res.TplinkCloudBase.Path = "/vms/open/webServer/v1/searchYear"
	res.TplinkCloudBase.Payload = request
	respBytes, err := res.PostRequest()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return nil, err
	}
	if response.ErrorCode != 0 {
		return nil, fmt.Errorf("searchYear error: %d", response.ErrorCode)
	}
	return &response.Result, nil
}

func (res *tplinkCloudIPCRes) SearchVideo(qrCode string, searchDay string, startIdx int, endIdx int) (*SearchVideoResult, error) {
	var (
		err      error
		request  SearchVideoRequest
		response SearchVideoResponse
	)
	request.QrCode = qrCode
	request.ChannelId = 1
	request.SearchDay = searchDay
	request.StartIdx = startIdx
	request.EndIdx = endIdx
	res.TplinkCloudBase.Path = "/vms/open/webServer/v2/searchVideo"
	res.TplinkCloudBase.Payload = request
	respBytes, err := res.PostRequest()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return nil, err
	}
	if response.ErrorCode != 0 {
		return nil, fmt.Errorf("searchVideo error: %d", response.ErrorCode)
	}
	return &response.Result, nil
}

// 获取选定日期所有的人形录像数据，应该先调用1.4.5.7 searchYear接口获取存在录像的日期信息。
func (res *tplinkCloudIPCRes) SearchHumanDetectionVideo(qrCode string, searchDay string) (*SearchVideoResult, error) {
	var (
		err      error
		request  SearchHumanDetectionVideoRequest
		response SearchHumanDetectionVideoResponse
	)
	request.QrCode = qrCode
	request.SearchDay = searchDay
	res.TplinkCloudBase.Path = "/vms/open/webServer/v2/searchHumanDetectionVideo"
	res.TplinkCloudBase.Payload = request
	respBytes, err := res.PostRequest()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return nil, err
	}
	if response.ErrorCode != 0 {
		return nil, fmt.Errorf("searchHumanDetectionVideo error: %d", response.ErrorCode)
	}
	return &response.Result, nil
}
