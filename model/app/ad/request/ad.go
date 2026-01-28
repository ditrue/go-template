package request

// ?keywords=abc&pos=abc

type AdRequest struct {
	Keywords string `json:"keywords" form:"keywords"`
	Pos      string `json:"pos" form:"pos"`
}
