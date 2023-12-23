package api

type JobResponse struct {
	By    string `json:"by"`
	ID    int    `json:"id"`
	Score int    `json:"score"`
	Text  string `json:"text"`
	Time  int    `json:"time"`
	Title string `json:"title"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}
