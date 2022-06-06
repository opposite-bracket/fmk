package main

const packageInit = `package {{.PackageName}}`

const apiFile = `package {{.PackageName}}

// mockgen -source=accounts/api.go -destination=accounts/mocks/mock_api.go  -self_package IApi

type IApi interface {
	Create(c *fmk.Context) error
	List(c *fmk.Context) error
	update(c *fmk.Context) error
	delete(c *fmk.Context) error
}

var Endpoint IApi = &endpoint{}

type endpoint struct{}

func (m *endpoint) Create(c *fmk.Context) error {

	var doc Model
	if err := c.ValidateBody(doc); err != nil {
		return err
	}

	if err := Service.Create(*c.TenantDoc(), doc); err != nil {
		return err
	}

	return nil
}

func (m *endpoint) List(c *fmk.Context) error {

	var filter ModelQuery
	if err := c.ValidateQuery(filter); err != nil {
		return err
	}

	if err := Service.List(*c.TenantDoc(), filter); err != nil {
		return err
	}

	return nil
}

func (m *endpoint) update(c *fmk.Context) error {

	var filter ModelQuery
	if err := c.ValidateQuery(filter); err != nil {
		return err
	}

	var doc Model
	if err := c.ValidateBody(filter); err != nil {
		return err
	}

	if err := Service.update(*c.TenantDoc(), filter, doc); err != nil {
		return err
	}

	return nil
}

func (m *endpoint) delete(c *fmk.Context) error {

	var filter ModelQuery
	if err := c.ValidateQuery(filter); err != nil {
		return err
	}

	if err := Service.delete(*c.TenantDoc(), filter); err != nil {
		return err
	}

	return nil
}
`

const apiTestFile = `package {{.PackageName}}`

const serviceFile = `package {{.PackageName}}

import "github.com/opposite-bracket/fmk"

// mockgen -source=accounts/service.go -destination=accounts/mocks/mock_service.go  -self_package IService

type IService interface {
	Create(tenant fmk.TenantDoc, doc Model) error
	List(tenant fmk.TenantDoc, filter ModelQuery) error
	update(tenant fmk.TenantDoc, filter ModelQuery, update Model) error
	delete(tenant fmk.TenantDoc, filter ModelQuery) error
}

var Service IService = &service{}

type service struct{}

func (s *service) Create(tenant fmk.TenantDoc, doc Model) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) List(tenant fmk.TenantDoc, filter ModelQuery) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) update(tenant fmk.TenantDoc, filter ModelQuery, update Model) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) delete(tenant fmk.TenantDoc, filter ModelQuery) error {
	//TODO implement me
	panic("implement me")
}
`

const serviceTestFile = `package {{.PackageName}}`

const modelFile = `package {{.PackageName}}

// mockgen -source=accounts/model.go -destination=accounts/mocks/mock_model.go  -self_package IModel

type IModel interface {
	Save() error
	FindByFilter(filter Model) error
	update(filter Model, doc Model, id string) error
	delete(filter Model) error
}

var Doc IModel = &Model{}

type ModelQuery struct{}
type ModelSort struct{}
type ModelPagination struct{}
type Model struct{}

func (m *Model) Save() error {
	//TODO implement me
	panic("implement me")
}

func (m *Model) FindByFilter(filter Model) error {
	//TODO implement me
	panic("implement me")
}

func (m *Model) update(filter Model, doc Model, id string) error {
	//TODO implement me
	panic("implement me")
}

func (m *Model) delete(filter Model) error {
	//TODO implement me
	panic("implement me")
}
`

const modelTestFile = `package {{.PackageName}}`
