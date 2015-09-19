package kurs

func GetCurrency(bank string) Data {
	response := Data{}

	switch bank{
		case "bi":
			response = processBI()
		case "mandiri":
			response = processMandiri()	
		default:
			response = Data{} 
	}

	return response
}