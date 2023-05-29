package database

import (
	"errors"
	"math"
	"strconv"

	"github.com/rohanshrestha09/todo/configs"
	"github.com/rohanshrestha09/todo/scopes"
	"gorm.io/gorm"
)

func GetByID[T any](paramID string, args GetByIDArgs) (T, error) {

	var data T

	id, err := strconv.Atoi(paramID)

	if err != nil {
		return data, err
	}

	dbScopes := []func(*gorm.DB) *gorm.DB{
		scopes.Exclude(args.Exclude...),
	}

	for k, v := range args.Include {
		dbScopes = append(dbScopes, scopes.Include(k, v...))
	}

	if err := configs.DB.Scopes(dbScopes...).Where(id).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil

}

func Get[T any](args GetArgs[T]) (T, error) {

	var data T

	dbScopes := []func(*gorm.DB) *gorm.DB{
		scopes.Exclude(args.Exclude...),
	}

	for k, v := range args.Include {
		dbScopes = append(dbScopes, scopes.Include(k, v...))
	}

	if err := configs.DB.Scopes(dbScopes...).Where(&args.Filter).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil

}

func GetAll[T any](bindQuery func(any) error, args GetAllArgs[T]) (GetAllResponse[T], *Error) {

	var (
		data  []T
		count int64
	)

	var query Query

	if err := bindQuery(&query); err != nil {
		if err != nil {
			return GetAllResponse[T]{}, &Error{Response{err.Error()}, err}
		}
	}

	query.Search = "%" + query.Search + "%"

	dbScopes := []func(*gorm.DB) *gorm.DB{
		scopes.Exclude(args.Exclude...),
		scopes.Sort(query.Sort, query.Order),
	}

	if args.Pagination {
		dbScopes = append(dbScopes, scopes.Paginate(query.Page, query.Size))
	}

	if args.Search {
		dbScopes = append(dbScopes, scopes.Search(query.Search))
	}

	for k, v := range args.Include {
		dbScopes = append(dbScopes, scopes.Include(k, v...))
	}

	err := configs.DB.
		Scopes(dbScopes...).
		Where(&args.Filter).
		Where(args.MapFilter).
		Find(&data).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error

	if err != nil {
		return GetAllResponse[T]{}, &Error{Response{err.Error()}, err}
	}

	currentPage := query.Page

	totalPage := math.Ceil(float64(count) / float64(query.Size))

	length := len(data)

	response := GetAllResponse[T]{
		Message: "Data Fetched",
		Data:    data,
		Length:  length,
	}

	if args.Pagination {
		response.Count = count
		response.CurrentPage = currentPage
		response.TotalPage = totalPage
	}

	return response, nil

}

func Create[T any](data T) (func(string) Response, *Error) {

	if err := configs.DB.Create(&data).Error; err != nil {
		return nil, &Error{Response{err.Error()}, err}
	}

	return func(message string) Response { return Response{message} }, nil

}

func Update[T any, K map[string]any | any](model T, data K) (func(string) Response, *Error) {

	if err := configs.DB.Model(&model).Updates(&data).Error; err != nil {
		return nil, &Error{Response{err.Error()}, err}
	}

	return func(message string) Response { return Response{message} }, nil

}

func Delete[T any](data T) (func(string) Response, *Error) {

	if err := configs.DB.Unscoped().Delete(&data).Error; err != nil {
		return nil, &Error{Response{err.Error()}, err}
	}

	return func(message string) Response { return Response{message} }, nil

}

func RecordExists[T any](filter *T, ORFilter ...*T) (bool, error) {

	var data T

	err := configs.DB.Where(&filter).Or(&ORFilter).First(&data).Error

	if err != nil {
		return !errors.Is(err, gorm.ErrRecordNotFound), err
	}

	return true, err

}
