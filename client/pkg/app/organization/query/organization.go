package query

import (
	"context"
	"github.com/TemaStatham/TaskService/client/pkg/app/organization/model"
)

type OrganizationQueryInterface interface {
	GetOrganization(ctx context.Context, orgID uint64) (model.OrganizationModel, error)
	GetOrganizationsByUserID(ctx context.Context, userID uint64) ([]model.OrganizationModel, error)
}

type ClientOrganizationInterface interface {
	GetOrganization(ctx context.Context, orgID uint64) (model.OrganizationModel, error)
	GetOrganizationsByUserID(ctx context.Context, userID uint64) ([]model.OrganizationModel, error)
}

type OrganizationQuery struct {
	client ClientOrganizationInterface
}

func NewOrganization(client ClientOrganizationInterface) *OrganizationQuery {
	return &OrganizationQuery{
		client: client,
	}
}

func (o *OrganizationQuery) GetOrganization(ctx context.Context, orgID uint64) (model.OrganizationModel, error) {
	return o.client.GetOrganization(ctx, orgID)
}

func (o *OrganizationQuery) GetOrganizationsByUserID(ctx context.Context, userID uint64) ([]model.OrganizationModel, error) {
	return o.client.GetOrganizationsByUserID(ctx, userID)
}
