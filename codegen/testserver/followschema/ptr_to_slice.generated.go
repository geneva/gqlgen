// Code generated by github.com/geneva/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"

	"github.com/geneva/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _PtrToSliceContainer_ptrToSlice(ctx context.Context, field graphql.CollectedField, obj *PtrToSliceContainer) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PtrToSliceContainer_ptrToSlice(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PtrToSlice, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*[]string)
	fc.Result = res
	return ec.marshalOString2ᚖᚕstringᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PtrToSliceContainer_ptrToSlice(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PtrToSliceContainer",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var ptrToSliceContainerImplementors = []string{"PtrToSliceContainer"}

func (ec *executionContext) _PtrToSliceContainer(ctx context.Context, sel ast.SelectionSet, obj *PtrToSliceContainer) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, ptrToSliceContainerImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PtrToSliceContainer")
		case "ptrToSlice":
			out.Values[i] = ec._PtrToSliceContainer_ptrToSlice(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNPtrToSliceContainer2githubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToSliceContainer(ctx context.Context, sel ast.SelectionSet, v PtrToSliceContainer) graphql.Marshaler {
	return ec._PtrToSliceContainer(ctx, sel, &v)
}

func (ec *executionContext) marshalNPtrToSliceContainer2ᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToSliceContainer(ctx context.Context, sel ast.SelectionSet, v *PtrToSliceContainer) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._PtrToSliceContainer(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
