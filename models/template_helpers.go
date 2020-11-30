// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"time"

	"github.com/percona-platform/saas/pkg/alert"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
)

func checkUniqueTemplateName(q *reform.Querier, name string) error {
	if name == "" {
		panic("empty template name")
	}

	template := &Template{Name: name}
	switch err := q.Reload(template); err {
	case nil:
		return status.Errorf(codes.AlreadyExists, "Template with name %q already exists.", name)
	case reform.ErrNoRows:
		return nil
	default:
		return errors.WithStack(err)
	}
}

// FindTemplates returns saved notification rule templates.
func FindTemplates(q *reform.Querier) ([]Template, error) {
	structs, err := q.SelectAllFrom(TemplateTable, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to select notification rule templates")
	}

	templates := make([]Template, len(structs))
	for i, s := range structs {
		c := s.(*Template)

		templates[i] = *c
	}

	return templates, nil
}

// FindTemplateByName finds template by name.
func FindTemplateByName(q *reform.Querier, name string) (*Template, error) {
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty template name.")
	}

	template := &Template{Name: name}
	switch err := q.Reload(template); err {
	case nil:
		return template, nil
	case reform.ErrNoRows:
		return nil, status.Errorf(codes.NotFound, "Template with name %q not found.", name)
	default:
		return nil, errors.WithStack(err)
	}
}

// CreateTemplateParams are params for creating new rule template.
type CreateTemplateParams struct {
	Rule   *alert.Rule
	Source string
}

// CreateTemplate creates rule template.
func CreateTemplate(q *reform.Querier, params *CreateTemplateParams) (*Template, error) {
	rule := params.Rule
	if err := rule.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid rule template: %v.", err)
	}

	if err := checkUniqueTemplateName(q, params.Rule.Name); err != nil {
		return nil, err
	}

	p, err := convertTemplateParams(params.Rule.Params)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid rule template parameters: %v.", err)
	}

	row := &Template{
		Name:     rule.Name,
		Version:  rule.Version,
		Summary:  rule.Summary,
		Tiers:    rule.Tiers,
		Expr:     rule.Expr,
		Params:   p,
		For:      time.Duration(rule.For),
		Severity: rule.Severity.String(),
		Source:   params.Source,
	}

	if err := row.SetLabels(rule.Labels); err != nil {
		return nil, err
	}

	if err := row.SetAnnotations(rule.Annotations); err != nil {
		return nil, err
	}

	err = q.Insert(row)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rule template")
	}

	return row, nil
}

// ChangeTemplateParams is params for changing existing rule template.
type ChangeTemplateParams struct {
	Rule *alert.Rule
}

// ChangeTemplate updates existing rule template.
func ChangeTemplate(q *reform.Querier, params *ChangeTemplateParams) (*Template, error) {
	row, err := FindTemplateByName(q, params.Rule.Name)
	if err != nil {
		return nil, err
	}

	rule := params.Rule
	if err := rule.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid rule template: %v.", err)
	}

	p, err := convertTemplateParams(params.Rule.Params)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid rule template parameters: %v.", err)
	}

	row.Name = rule.Name
	row.Version = rule.Version
	row.Summary = rule.Summary
	row.Tiers = rule.Tiers
	row.Expr = rule.Expr
	row.Params = p
	row.For = time.Duration(rule.For)
	row.Severity = rule.Severity.String()

	if err := row.SetLabels(rule.Labels); err != nil {
		return nil, err
	}

	if err := row.SetAnnotations(rule.Annotations); err != nil {
		return nil, err
	}

	err = q.Update(row)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update rule template")
	}

	return row, nil
}

// RemoveTemplate removes rule template with specified name.
func RemoveTemplate(q *reform.Querier, name string) error {
	if _, err := FindTemplateByName(q, name); err != nil {
		return err
	}

	err := q.Delete(&Template{Name: name})
	if err != nil {
		return errors.Wrap(err, "failed to delete rule template")
	}
	return nil
}

func convertTemplateParams(params []alert.Parameter) (Params, error) {
	res := make(Params, len(params))
	for i, param := range params {
		res[i] = Param{
			Name:    param.Name,
			Summary: param.Summary,
			Unit:    param.Unit,
			Type:    string(param.Type),
		}

		switch param.Type {
		case alert.Float:
			var fp FloatParam
			var err error
			fp.Default, err = param.GetValueForFloat()
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse param value")
			}

			fp.Min, fp.Max, err = param.GetRangeForFloat()
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse param range")
			}

			res[i].FloatParam = &fp
		}
	}

	return res, nil
}
