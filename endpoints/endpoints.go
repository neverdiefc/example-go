package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/neverdiefc/example-go/service"

	"github.com/neverdiefc/example-go/endpoints/book"
	"github.com/neverdiefc/example-go/endpoints/category"
	"github.com/neverdiefc/example-go/endpoints/lendbook"
	"github.com/neverdiefc/example-go/endpoints/user"
)

// Endpoints .
type Endpoints struct {
	FindUser    endpoint.Endpoint
	FindAllUser endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint

	FindCategory    endpoint.Endpoint
	FindAllCategory endpoint.Endpoint
	CreateCategory  endpoint.Endpoint
	UpdateCategory  endpoint.Endpoint
	DeleteCategory  endpoint.Endpoint

	FindBook    endpoint.Endpoint
	FindAllBook endpoint.Endpoint
	CreateBook  endpoint.Endpoint
	UpdateBook  endpoint.Endpoint
	DeleteBook  endpoint.Endpoint

	FindLendbook    endpoint.Endpoint
	FindAllLendbook endpoint.Endpoint
	CreateLendbook  endpoint.Endpoint
	UpdateLendbook  endpoint.Endpoint
	DeleteLendbook  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		FindUser:    user.MakeFindEndPoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),

		FindCategory:    category.MakeFindEndPoint(s),
		FindAllCategory: category.MakeFindAllEndpoint(s),
		CreateCategory:  category.MakeCreateEndpoint(s),
		UpdateCategory:  category.MakeUpdateEndpoint(s),
		DeleteCategory:  category.MakeDeleteEndpoint(s),

		FindBook:    book.MakeFindEndPoint(s),
		FindAllBook: book.MakeFindAllEndpoint(s),
		CreateBook:  book.MakeCreateEndpoint(s),
		UpdateBook:  book.MakeUpdateEndpoint(s),
		DeleteBook:  book.MakeDeleteEndpoint(s),

		FindLendbook:    lendbook.MakeFindEndPoint(s),
		FindAllLendbook: lendbook.MakeFindAllEndpoint(s),
		CreateLendbook:  lendbook.MakeCreateEndpoint(s),
		UpdateLendbook:  lendbook.MakeUpdateEndpoint(s),
		DeleteLendbook:  lendbook.MakeDeleteEndpoint(s),
	}
}
