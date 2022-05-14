package cotroller

import (
	"strconv"

	"page/service"

	"page/repository"
)

type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type FeedBack struct {
	Flag   int64  `json:"flag"`
	Status string `json:"status"`
}

func QueryPageInfo(topicIdStr string) *PageData {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	pageInfo, err := service.QueryPageInfo(topicId)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: pageInfo,
	}

}

func UpdatePage(post *repository.Post) *FeedBack {
	err := service.UpdatePostIndex(post)
	if err != nil {

		return &FeedBack{
			Flag:   -1,
			Status: err.Error(),
		}
	}
	return &FeedBack{
		Flag:   1,
		Status: "Success",
	}
}
