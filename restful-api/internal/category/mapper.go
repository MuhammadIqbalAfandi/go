package category

func ToResponse(category Category) Response {
	return Response{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToResponseList(categories []Category) []Response {
	var list []Response
	for _, category := range categories {
		list = append(list, ToResponse(category))
	}
	return list
}
