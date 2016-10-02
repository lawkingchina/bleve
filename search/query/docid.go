//  Copyright (c) 2015 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package query

import (
	"github.com/blevesearch/bleve/index"
	"github.com/blevesearch/bleve/mapping"
	"github.com/blevesearch/bleve/search"
	"github.com/blevesearch/bleve/search/searcher"
)

type DocIDQuery struct {
	IDs   []string `json:"ids"`
	Boost *Boost   `json:"boost,omitempty"`
}

// NewDocIDQuery creates a new Query object returning indexed documents among
// the specified set. Combine it with ConjunctionQuery to restrict the scope of
// other queries output.
func NewDocIDQuery(ids []string) *DocIDQuery {
	return &DocIDQuery{
		IDs: ids,
	}
}

func (q *DocIDQuery) SetBoost(b float64) {
	boost := Boost(b)
	q.Boost = &boost
}

func (q *DocIDQuery) Searcher(i index.IndexReader, m mapping.IndexMapping, explain bool) (search.Searcher, error) {
	return searchers.NewDocIDSearcher(i, q.IDs, q.Boost.Value(), explain)
}
