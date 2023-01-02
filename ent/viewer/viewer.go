// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package viewer

import (
	"context"

	ent "github.com/heshaofeng1991/entgo/ent/gen"
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	Tenant() int64 // Tenant id.
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	T *ent.Tenant
}

func (v UserViewer) Tenant() int64 {
	if v.T != nil {
		return v.T.ID
	}
	return 0
}

type ctxKey struct{}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}
