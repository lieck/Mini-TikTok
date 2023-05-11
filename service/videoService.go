package service

import "time"

type VideoService interface {
	// GetVideoCount 返回用户的投稿数量
	GetVideoCount(id int64) (int64, error)
	
	//brief：根据timeUnix参数，返回按投稿时间倒序的视频列表和最早发布时间。视频数由服务端控制，单次最多30个。
	//@para timeUnix time.Time：最早发布时间，返回视频列表中视频发布时间不得晚于timeUnix。
	//@return []video: 按投稿时间倒序的视频列表
	//@return int64： 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	GetVideos(timeUnix time.Time, user int64) (videos []Video, nextTime int64)
}
