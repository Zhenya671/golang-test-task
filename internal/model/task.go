package model

type Task struct {
	InputData  []interface{} `json:"input_data"`
	OutputData interface{}   `json:"output_data"`
}
