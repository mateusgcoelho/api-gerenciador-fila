package reports

type CreateReportRequest struct {
	PersonID     int   `json:"person_id" binding:"required,min=1"`
	ResponsiveID int32 `json:"responsive_id" binding:"required,min=1"`
	QueueID      int32 `json:"queue_id" binding:"required,min=1"`
}

type GetReportsRequest struct {
	Page     int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
