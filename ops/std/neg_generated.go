package stdops

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

import (
	"context"
	"runtime/trace"

	gctx "gorgonia.org/gorgonia/internal/context"
	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// neg is a elementwise negation.
type negOp struct{ unop }

// String implements fmt.Stringer.
func (op negOp) String() string { return "Neg" }

// Do performs elementwise negation.
func (op negOp) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := gctx.Handle(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Neg(a, tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise negation but with a preallocated return value.
// PreallocDo allows add to implement ops.PreallocOp.
func (op negOp) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := gctx.Handle(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Neg(a, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// DiffWRT returns {true} for neg
func (op negOp) DiffWRT(inputs int) []bool { return onetrue }