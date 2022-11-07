// Code generated by goctl. DO NOT EDIT.
package types

type PagerProps struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

type GetPostsReq struct {
	Query    string `form:"query,optional"`
	Type     string `form:"type"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

type GetPostsResp struct {
	List  []*PostFormated `json:"list"`
	Pager Pager           `json:"pager"`
}

type PostContentFormated struct {
	ID      int64  `json:"id"`
	PostID  int64  `json:"post_id"`
	Content string `json:"content"`
	Type    int    `json:"type"`
	Sort    int64  `json:"sort"`
}

type PostFormated struct {
	ID              int64                  `json:"id"`
	UserID          int64                  `json:"user_id"`
	User            *UserFormated          `json:"user"`
	Contents        []*PostContentFormated `json:"contents"`
	CommentCount    int64                  `json:"comment_count"`
	CollectionCount int64                  `json:"collection_count"`
	UpvoteCount     int64                  `json:"upvote_count"`
	Visibility      int64                  `json:"visibility"`
	IsTop           int64                  `json:"is_top"`
	IsEssence       int64                  `json:"is_essence"`
	IsLock          int64                  `json:"is_lock"`
	LatestRepliedOn int64                  `json:"latest_replied_on"`
	CreatedOn       int64                  `json:"created_on"`
	ModifiedOn      int64                  `json:"modified_on"`
	Tags            map[string]int8        `json:"tags"`
	AttachmentPrice int64                  `json:"attachment_price"`
	IPLoc           string                 `json:"ip_loc"`
}

type UserFormated struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	IsAdmin  bool   `json:"is_admin"`
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

type GetTagReq struct {
	Type string `form:"type"`
	Num  int    `form:"num"`
}

type GetTagResp struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data []*TagFormated `json:"data"`
}

type TagFormated struct {
	ID       int64         `json:"id"`
	UserID   int64         `json:"user_id"`
	User     *UserFormated `json:"user"`
	Tag      string        `json:"tag"`
	QuoteNum int64         `json:"quote_num"`
}

type PTag struct {
	Id         int64  `json:"id"`          // 标签ID
	UserId     int64  `json:"user_id"`     // 创建者ID
	Tag        string `json:"tag"`         // 标签名
	QuoteNum   int64  `json:"quote_num"`   // 引用数
	CreatedOn  int64  `json:"created_on"`  // 创建时间
	ModifiedOn int64  `json:"modified_on"` // 修改时间
	DeletedOn  int64  `json:"deleted_on"`  // 删除时间
	IsDel      int64  `json:"is_del"`      // 是否删除 0 为未删除、1 为已删除
}

type UserRegisterReq struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type UserRegisterRes struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}
