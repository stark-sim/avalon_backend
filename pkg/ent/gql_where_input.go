// Code generated by ent, DO NOT EDIT.

package ent

import (
	"avalon_backend/pkg/ent/card"
	"avalon_backend/pkg/ent/predicate"
	"errors"
	"fmt"
	"time"
)

// CardWhereInput represents a where input for filtering Card queries.
type CardWhereInput struct {
	Predicates []predicate.Card  `json:"-"`
	Not        *CardWhereInput   `json:"not,omitempty"`
	Or         []*CardWhereInput `json:"or,omitempty"`
	And        []*CardWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int64  `json:"id,omitempty"`
	IDNEQ   *int64  `json:"idNEQ,omitempty"`
	IDIn    []int64 `json:"idIn,omitempty"`
	IDNotIn []int64 `json:"idNotIn,omitempty"`
	IDGT    *int64  `json:"idGT,omitempty"`
	IDGTE   *int64  `json:"idGTE,omitempty"`
	IDLT    *int64  `json:"idLT,omitempty"`
	IDLTE   *int64  `json:"idLTE,omitempty"`

	// "created_by" field predicates.
	CreatedBy      *int64  `json:"createdBy,omitempty"`
	CreatedByNEQ   *int64  `json:"createdByNEQ,omitempty"`
	CreatedByIn    []int64 `json:"createdByIn,omitempty"`
	CreatedByNotIn []int64 `json:"createdByNotIn,omitempty"`
	CreatedByGT    *int64  `json:"createdByGT,omitempty"`
	CreatedByGTE   *int64  `json:"createdByGTE,omitempty"`
	CreatedByLT    *int64  `json:"createdByLT,omitempty"`
	CreatedByLTE   *int64  `json:"createdByLTE,omitempty"`

	// "updated_by" field predicates.
	UpdatedBy      *int64  `json:"updatedBy,omitempty"`
	UpdatedByNEQ   *int64  `json:"updatedByNEQ,omitempty"`
	UpdatedByIn    []int64 `json:"updatedByIn,omitempty"`
	UpdatedByNotIn []int64 `json:"updatedByNotIn,omitempty"`
	UpdatedByGT    *int64  `json:"updatedByGT,omitempty"`
	UpdatedByGTE   *int64  `json:"updatedByGTE,omitempty"`
	UpdatedByLT    *int64  `json:"updatedByLT,omitempty"`
	UpdatedByLTE   *int64  `json:"updatedByLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "deleted_at" field predicates.
	DeletedAt      *time.Time  `json:"deletedAt,omitempty"`
	DeletedAtNEQ   *time.Time  `json:"deletedAtNEQ,omitempty"`
	DeletedAtIn    []time.Time `json:"deletedAtIn,omitempty"`
	DeletedAtNotIn []time.Time `json:"deletedAtNotIn,omitempty"`
	DeletedAtGT    *time.Time  `json:"deletedAtGT,omitempty"`
	DeletedAtGTE   *time.Time  `json:"deletedAtGTE,omitempty"`
	DeletedAtLT    *time.Time  `json:"deletedAtLT,omitempty"`
	DeletedAtLTE   *time.Time  `json:"deletedAtLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *CardWhereInput) AddPredicates(predicates ...predicate.Card) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the CardWhereInput filter on the CardQuery builder.
func (i *CardWhereInput) Filter(q *CardQuery) (*CardQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyCardWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyCardWhereInput is returned in case the CardWhereInput is empty.
var ErrEmptyCardWhereInput = errors.New("ent: empty predicate CardWhereInput")

// P returns a predicate for filtering cards.
// An error is returned if the input is empty or invalid.
func (i *CardWhereInput) P() (predicate.Card, error) {
	var predicates []predicate.Card
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, card.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Card, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, card.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Card, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, card.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, card.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, card.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, card.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, card.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, card.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, card.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, card.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, card.IDLTE(*i.IDLTE))
	}
	if i.CreatedBy != nil {
		predicates = append(predicates, card.CreatedByEQ(*i.CreatedBy))
	}
	if i.CreatedByNEQ != nil {
		predicates = append(predicates, card.CreatedByNEQ(*i.CreatedByNEQ))
	}
	if len(i.CreatedByIn) > 0 {
		predicates = append(predicates, card.CreatedByIn(i.CreatedByIn...))
	}
	if len(i.CreatedByNotIn) > 0 {
		predicates = append(predicates, card.CreatedByNotIn(i.CreatedByNotIn...))
	}
	if i.CreatedByGT != nil {
		predicates = append(predicates, card.CreatedByGT(*i.CreatedByGT))
	}
	if i.CreatedByGTE != nil {
		predicates = append(predicates, card.CreatedByGTE(*i.CreatedByGTE))
	}
	if i.CreatedByLT != nil {
		predicates = append(predicates, card.CreatedByLT(*i.CreatedByLT))
	}
	if i.CreatedByLTE != nil {
		predicates = append(predicates, card.CreatedByLTE(*i.CreatedByLTE))
	}
	if i.UpdatedBy != nil {
		predicates = append(predicates, card.UpdatedByEQ(*i.UpdatedBy))
	}
	if i.UpdatedByNEQ != nil {
		predicates = append(predicates, card.UpdatedByNEQ(*i.UpdatedByNEQ))
	}
	if len(i.UpdatedByIn) > 0 {
		predicates = append(predicates, card.UpdatedByIn(i.UpdatedByIn...))
	}
	if len(i.UpdatedByNotIn) > 0 {
		predicates = append(predicates, card.UpdatedByNotIn(i.UpdatedByNotIn...))
	}
	if i.UpdatedByGT != nil {
		predicates = append(predicates, card.UpdatedByGT(*i.UpdatedByGT))
	}
	if i.UpdatedByGTE != nil {
		predicates = append(predicates, card.UpdatedByGTE(*i.UpdatedByGTE))
	}
	if i.UpdatedByLT != nil {
		predicates = append(predicates, card.UpdatedByLT(*i.UpdatedByLT))
	}
	if i.UpdatedByLTE != nil {
		predicates = append(predicates, card.UpdatedByLTE(*i.UpdatedByLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, card.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, card.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, card.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, card.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, card.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, card.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, card.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, card.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, card.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, card.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, card.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, card.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, card.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, card.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, card.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, card.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.DeletedAt != nil {
		predicates = append(predicates, card.DeletedAtEQ(*i.DeletedAt))
	}
	if i.DeletedAtNEQ != nil {
		predicates = append(predicates, card.DeletedAtNEQ(*i.DeletedAtNEQ))
	}
	if len(i.DeletedAtIn) > 0 {
		predicates = append(predicates, card.DeletedAtIn(i.DeletedAtIn...))
	}
	if len(i.DeletedAtNotIn) > 0 {
		predicates = append(predicates, card.DeletedAtNotIn(i.DeletedAtNotIn...))
	}
	if i.DeletedAtGT != nil {
		predicates = append(predicates, card.DeletedAtGT(*i.DeletedAtGT))
	}
	if i.DeletedAtGTE != nil {
		predicates = append(predicates, card.DeletedAtGTE(*i.DeletedAtGTE))
	}
	if i.DeletedAtLT != nil {
		predicates = append(predicates, card.DeletedAtLT(*i.DeletedAtLT))
	}
	if i.DeletedAtLTE != nil {
		predicates = append(predicates, card.DeletedAtLTE(*i.DeletedAtLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, card.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, card.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, card.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, card.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, card.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, card.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, card.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, card.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, card.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, card.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, card.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, card.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, card.NameContainsFold(*i.NameContainsFold))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyCardWhereInput
	case 1:
		return predicates[0], nil
	default:
		return card.And(predicates...), nil
	}
}
