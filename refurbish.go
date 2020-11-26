package main

import (
	"log"
	"time"
	"uc/lib/utils"
	"uc/models"
)

type Refurbish struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		FanStory struct {
			ID             int         `json:"id"`
			Title          string      `json:"title"`
			RoomPattern    string      `json:"roomPattern"`
			CoverImgURL    string      `json:"coverImgUrl"`
			Description    string      `json:"description"`
			PublishDate    int64       `json:"publishDate"`
			ViewNumber     int         `json:"viewNumber"`
			State          int         `json:"state"`
			ToURL          interface{} `json:"toUrl"`
			CreateTime     int64       `json:"createTime"`
			CreateUser     string      `json:"createUser"`
			UpdateTime     int64       `json:"updateTime"`
			UpdateUser     string      `json:"updateUser"`
			Village        string      `json:"village"`
			AvatarURL      string      `json:"avatarUrl"`
			NickName       string      `json:"nickName"`
			Content        string      `json:"content"`
			IndexImgURL    string      `json:"indexImgUrl"`
			PublishDateStr string      `json:"publishDateStr"`
		} `json:"fanStory"`
		FanStoryList []struct {
			ID             int         `json:"id"`
			Title          string      `json:"title"`
			RoomPattern    string      `json:"roomPattern"`
			CoverImgURL    string      `json:"coverImgUrl"`
			Description    string      `json:"description"`
			PublishDate    int64       `json:"publishDate"`
			ViewNumber     int         `json:"viewNumber"`
			State          int         `json:"state"`
			ToURL          interface{} `json:"toUrl"`
			CreateTime     int64       `json:"createTime"`
			CreateUser     string      `json:"createUser"`
			UpdateTime     int64       `json:"updateTime"`
			UpdateUser     string      `json:"updateUser"`
			Village        string      `json:"village"`
			AvatarURL      string      `json:"avatarUrl"`
			NickName       string      `json:"nickName"`
			Content        string      `json:"content"`
			IndexImgURL    string      `json:"indexImgUrl"`
			PublishDateStr string      `json:"publishDateStr"`
		} `json:"fanStoryList"`
	} `json:"data"`
}

type CaseInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CaseDetail struct {
			ID        int    `json:"id"`
			Title     string `json:"title"`
			IndexImg  string `json:"indexImg"`
			ReadCount int    `json:"readCount"`
			Publisher struct {
				ID      interface{} `json:"id"`
				Name    string      `json:"name"`
				HeadImg string      `json:"headImg"`
			} `json:"publisher"`
			StyleID                 int         `json:"styleId"`
			StyleName               string      `json:"styleName"`
			Area                    int         `json:"area"`
			ApartmentLayoutName     string      `json:"apartmentLayoutName"`
			ImgWidth                string      `json:"imgWidth"`
			ImgHeight               string      `json:"imgHeight"`
			PublishContent          string      `json:"publishContent"`
			DesignerID              interface{} `json:"designerId"`
			DesignerImg             interface{} `json:"designerImg"`
			DesignerName            interface{} `json:"designerName"`
			DesignerLabel           interface{} `json:"designerLabel"`
			DesignerPopularityCount interface{} `json:"designerPopularityCount"`
			ArticleImg              bool        `json:"articleImg"`
			CreateTime              int64       `json:"createTime"`
		} `json:"caseDetail"`
		CaseList []struct {
			ID          int         `json:"id"`
			Title       string      `json:"title"`
			IndexImg    string      `json:"indexImg"`
			ReadCount   int         `json:"readCount"`
			PublisherID int         `json:"publisherId"`
			DesignerID  interface{} `json:"designerId"`
			Publisher   struct {
				ID      interface{} `json:"id"`
				Name    string      `json:"name"`
				HeadImg string      `json:"headImg"`
			} `json:"publisher"`
			Source              int         `json:"source"`
			ImgWidth            string      `json:"imgWidth"`
			ImgHeight           string      `json:"imgHeight"`
			DetailLink          interface{} `json:"detailLink"`
			ImgCount            int         `json:"imgCount"`
			ApartmentLayoutName string      `json:"apartmentLayoutName"`
			DesignerImgURL      string      `json:"designerImgUrl"`
			DesignerName        interface{} `json:"designerName"`
			AreaCode            interface{} `json:"areaCode"`
			AreaCodeName        interface{} `json:"areaCodeName"`
			SourceType          int         `json:"sourceType"`
			LinkURL             interface{} `json:"linkUrl"`
		} `json:"caseList"`
	} `json:"data"`
}

func Chuli() {

	for i := 0; i < 700; i++ {
		url := "https://idecoration.ikongjian.com/app/case/detail/formal"
		cb := new(CaseInfo)
		err := utils.HttpPostJson(url, map[string]interface{}{
			"id": i,
		}, &cb)
		if err != nil {
			log.Println("mdzz")
			return
		}
		if cb.Code == 1 {
			m := models.RefurbishCase{
				Title:               cb.Data.CaseDetail.Title,
				Logo:                cb.Data.CaseDetail.IndexImg,
				ShowNum:             cb.Data.CaseDetail.ReadCount,
				RoomPattern:         "",
				CoverImgUrl:         cb.Data.CaseDetail.IndexImg,
				Description:         cb.Data.CaseDetail.Title,
				Status:              1,
				ToPath:              "",
				CreatedUser:         cb.Data.CaseDetail.Publisher.Name,
				UpdateUser:          cb.Data.CaseDetail.Publisher.Name,
				Village:             "",
				AvatarUrl:           cb.Data.CaseDetail.Publisher.HeadImg,
				Nickname:            cb.Data.CaseDetail.Publisher.Name,
				Content:             cb.Data.CaseDetail.PublishContent,
				StyleName:           cb.Data.CaseDetail.StyleName,
				Area:                cb.Data.CaseDetail.Area,
				ApartmentLayoutName: cb.Data.CaseDetail.ApartmentLayoutName,
			}
			models.RefurbishCaseAdd(&m)
			log.Println(cb.Data.CaseDetail.ID)
		}
		time.Sleep(time.Millisecond * 100)
	}

}
