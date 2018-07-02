/*
 * Copyright (C) 2018 eeonevision
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package state

import (
	"errors"
)

//go:generate msgp

// Transition struct keeps transition related fields.
type Transition struct {
	ID                  string  `msg:"_id" json:"_id" mapstructure:"_id" bson:"_id"`
	AdvertiserAccountID string  `msg:"advertiser_account_id" json:"advertiser_account_id" mapstructure:"advertiser_account_id" bson:"advertiser_account_id"`
	AffiliateAccountID  string  `msg:"affiliate_account_id" json:"affiliate_account_id" mapstructure:"affiliate_account_id" bson:"affiliate_account_id"`
	ClickID             string  `msg:"click_id" json:"click_id" mapstructure:"click_id" bson:"click_id"`
	StreamID            string  `msg:"stream_id" json:"stream_id" mapstructure:"stream_id" bson:"stream_id"`
	OfferID             string  `msg:"offer_id" json:"offer_id" mapstructure:"offer_id" bson:"offer_id"`
	CreatedAt           float64 `msg:"created_at" json:"created_at" mapstructure:"created_at" bson:"created_at"`
	ExpiresIn           int64   `msg:"expires_in" json:"expires_in" mapstructure:"expires_in" bson:"expires_in"`
}

const transitionsCollection = "transitions"

// AddTransition method adds new transition to state if it not exists.
func (s *State) AddTransition(transition *Transition) error {
	if s.HasTransition(transition.ID) {
		return errors.New("Transition exists")
	}
	return s.SetTransition(transition)
}

// SetTransition method inserts new transition in state without any checks.
func (s *State) SetTransition(transition *Transition) error {
	return s.DB.C(transitionsCollection).Insert(transition)
}

// HasTransition method checks exists conversion in state or not exists.
func (s *State) HasTransition(id string) bool {
	if res, _ := s.GetTransition(id); res != nil {
		return true
	}
	return false
}

// GetTransition method gets transition from state by given id.
func (s *State) GetTransition(id string) (*Transition, error) {
	var result *Transition
	return result, s.DB.C(transitionsCollection).FindId(id).One(&result)
}

// ListTransitions method returns list of all transitions from state.
func (s *State) ListTransitions() (result []*Transition, err error) {
	return result, s.DB.C(transitionsCollection).Find(nil).All(&result)
}

// SearchTransitions method finds conversions in state using mongodb query language.
func (s *State) SearchTransitions(query interface{}, limit, offset int) (result []*Transition, err error) {
	return result, s.DB.C(transitionsCollection).Find(query).Skip(offset).Limit(limit).All(&result)
}
