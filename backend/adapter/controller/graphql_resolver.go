package controller

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Organization() OrganizationResolver {
	return &organizationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Work() WorkResolver {
	return &workResolver{r}
}

func (r *Resolver) WorkHolder() WorkHolderResolver {
	return &workHolderResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Noop(ctx context.Context, input *NoopInput) (*NoopPayload, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateOrganization(ctx context.Context, input OrganizationInput) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateOrganization(ctx context.Context, input OrganizationInput) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteOrganization(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateWork(ctx context.Context, input WorkInput) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateWork(ctx context.Context, input WorkInput) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteWork(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateWorkHolder(ctx context.Context, input WorkHolderInput) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateWorkHolder(ctx context.Context, input WorkHolderInput) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteWorkHolder(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}

type organizationResolver struct{ *Resolver }

func (r *organizationResolver) UpperOrganization(ctx context.Context, obj *Organization) (*Organization, error) {
	panic("not implemented")
}

func (r *organizationResolver) LowerOrganizations(ctx context.Context, obj *Organization) ([]*Organization, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Node(ctx context.Context, id string) (Node, error) {
	panic("not implemented")
}

func (r *queryResolver) Organization(ctx context.Context, id string) (*Organization, error) {
	panic("not implemented")
}

func (r *queryResolver) Organizations(ctx context.Context, condition *OrganizationCondition) ([]Organization, error) {
	panic("not implemented")
}

func (r *queryResolver) Work(ctx context.Context, id string) (*Work, error) {
	panic("not implemented")
}

func (r *queryResolver) Works(ctx context.Context, condition *WorkCondition) ([]Work, error) {
	panic("not implemented")
}

func (r *queryResolver) WorkHolder(ctx context.Context, id string) (*WorkHolder, error) {
	panic("not implemented")
}

func (r *queryResolver) WorkHolders(ctx context.Context, condition *WorkHolderCondition) ([]WorkHolder, error) {
	panic("not implemented")
}

type workResolver struct{ *Resolver }

func (r *workResolver) WorkHolders(ctx context.Context, obj *Work) ([]*WorkHolder, error) {
	panic("not implemented")
}

type workHolderResolver struct{ *Resolver }

func (r *workHolderResolver) Organizations(ctx context.Context, obj *WorkHolder) ([]*Organization, error) {
	panic("not implemented")
}

func (r *workHolderResolver) HoldWorks(ctx context.Context, obj *WorkHolder) ([]*Work, error) {
	panic("not implemented")
}
