package dtos

type CreateTutorialRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
}

type TutorialResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
	CreatedAt   string `json:"created_at"`
}
