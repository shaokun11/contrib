// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"entgo.io/contrib/entoas/internal/oastypes"
)

// The OASTypesFunc type is an adapter to allow the use of ordinary
// function as OASTypes mutator.
type OASTypesFunc func(context.Context, *oastypes.OASTypesMutation) (oastypes.Value, error)

// Mutate calls f(ctx, m).
func (f OASTypesFunc) Mutate(ctx context.Context, m oastypes.Mutation) (oastypes.Value, error) {
	mv, ok := m.(*oastypes.OASTypesMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *oastypes.OASTypesMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, oastypes.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m oastypes.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m oastypes.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m oastypes.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op oastypes.Op) Condition {
	return func(_ context.Context, m oastypes.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m oastypes.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m oastypes.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m oastypes.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk oastypes.Hook, cond Condition) oastypes.Hook {
	return func(next oastypes.Mutator) oastypes.Mutator {
		return oastypes.MutateFunc(func(ctx context.Context, m oastypes.Mutation) (oastypes.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, oastypes.Delete|oastypes.Create)
func On(hk oastypes.Hook, op oastypes.Op) oastypes.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, oastypes.Update|oastypes.UpdateOne)
func Unless(hk oastypes.Hook, op oastypes.Op) oastypes.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) oastypes.Hook {
	return func(oastypes.Mutator) oastypes.Mutator {
		return oastypes.MutateFunc(func(context.Context, oastypes.Mutation) (oastypes.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []oastypes.Hook {
//		return []oastypes.Hook{
//			Reject(oastypes.Delete|oastypes.Update),
//		}
//	}
func Reject(op oastypes.Op) oastypes.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []oastypes.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...oastypes.Hook) Chain {
	return Chain{append([]oastypes.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() oastypes.Hook {
	return func(mutator oastypes.Mutator) oastypes.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...oastypes.Hook) Chain {
	newHooks := make([]oastypes.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
