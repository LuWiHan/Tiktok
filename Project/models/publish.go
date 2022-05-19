// publish 包，该包定义了与 `videos` 表对应的结构体以及视频相关接口请求响应格式
// 创建人：吴润泽
// 创建时间：2022-5-15

package models

import "gorm.io/gorm"

// Video 视频对象，定义了视频的基本信息
type Video struct {
	gorm.Model    `gorm:"embedded"`
	Id            int64  `json:"id"`                                              // 视频 ID
	Author        User   `gorm:"foreignKey:AuthorID;references:ID" json:"author"` // 视频发布者
	AuthorID      int64  `gorm:"column:author_id"`                                // 外键
	PlayUrl       string `gorm:"column:play_url" json:"play_url"`                 // 视频播放地址
	CoverUrl      string `gorm:"column:cover_url" json:"cover_url"`               // 视频封面地址
	FavoriteCount int64  `gorm:"column:favorite_count" json:"favorite_count"`     // 视频点赞总数
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count"`       // 视频评论总数
	IsFavorite    bool   `gorm:"column:is_favorite" json:"is_favorite"`           // 是否已点赞
}

// PublishVideoRequest 投稿视频请求格式
type PublishVideoRequest struct {
	UserID int64  `json:"UserID"`
	Token  string `json:"Token"`
	Data   Video  `json:"Data"` // 视频数据
}

// PublishVideoResponse 投稿视频响应格式
type PublishVideoResponse struct {
}
