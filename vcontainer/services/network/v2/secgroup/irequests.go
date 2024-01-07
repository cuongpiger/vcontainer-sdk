package secgroup

type ICreateOptsBuilder interface {
	GetProjectID() string
	ToRequestBody() interface{}
}

type IDeleteOptsBuilder interface {
	GetProjectID() string
	GetSecgroupUUID() string
}

type IGetOptsBuilder interface {
	GetProjectID() string
	GetSecgroupUUID() string
}