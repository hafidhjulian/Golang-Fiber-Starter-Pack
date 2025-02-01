package models

import "golang-fiber-starterpack/utils"

func GetExample() utils.Respon {
	var response utils.Respon
	data := "Hello World"

	response.Success = true
	response.Data = data
	response.Message = "Success"
	return response
}

func PostExample(data string) utils.Respon {
	var response utils.Respon

	response.Success = true
	response.Data = data
	response.Message = "Success"
	return response
}
