package entity

type Branch struct {
	BranchId int    `json:"branch_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
