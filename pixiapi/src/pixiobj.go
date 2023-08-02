package src

type SpriteEnt struct {
	Name    string  `json:"name"`
	Url     string  `json:"url"`
	Width   float32 `json:"width"`
	Height  float32 `json:"height"`
	X       float32 `json:"x"`
	Y       float32 `json:"y"`
	Visable bool    `json:"visable"`
}
