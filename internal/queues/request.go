package queues

type CreateQueueRequest struct {
	Name string `json:"name" binding:"required"`
}

type GetQueuesRequest struct {
	Page     int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
