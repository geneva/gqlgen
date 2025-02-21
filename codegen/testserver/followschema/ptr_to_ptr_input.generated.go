// Code generated by github.com/geneva/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"fmt"
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

func (ec *executionContext) _PtrToPtrInner_key(ctx context.Context, field graphql.CollectedField, obj *PtrToPtrInner) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PtrToPtrInner_key(ctx, field)
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
		return obj.Key, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PtrToPtrInner_key(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PtrToPtrInner",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _PtrToPtrInner_value(ctx context.Context, field graphql.CollectedField, obj *PtrToPtrInner) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PtrToPtrInner_value(ctx, field)
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
		return obj.Value, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PtrToPtrInner_value(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PtrToPtrInner",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _PtrToPtrOuter_name(ctx context.Context, field graphql.CollectedField, obj *PtrToPtrOuter) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PtrToPtrOuter_name(ctx, field)
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
		return obj.Name, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PtrToPtrOuter_name(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PtrToPtrOuter",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _PtrToPtrOuter_inner(ctx context.Context, field graphql.CollectedField, obj *PtrToPtrOuter) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PtrToPtrOuter_inner(ctx, field)
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
		return obj.Inner, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*PtrToPtrInner)
	fc.Result = res
	return ec.marshalOPtrToPtrInner2ᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PtrToPtrOuter_inner(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PtrToPtrOuter",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "key":
				return ec.fieldContext_PtrToPtrInner_key(ctx, field)
			case "value":
				return ec.fieldContext_PtrToPtrInner_value(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PtrToPtrInner", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _PtrToPtrOuter_stupidInner(ctx context.Context, field graphql.CollectedField, obj *PtrToPtrOuter) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PtrToPtrOuter_stupidInner(ctx, field)
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
		return obj.StupidInner, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*******PtrToPtrInner)
	fc.Result = res
	return ec.marshalOPtrToPtrInner2ᚖᚖᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PtrToPtrOuter_stupidInner(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PtrToPtrOuter",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "key":
				return ec.fieldContext_PtrToPtrInner_key(ctx, field)
			case "value":
				return ec.fieldContext_PtrToPtrInner_value(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PtrToPtrInner", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputUpdatePtrToPtrInner(ctx context.Context, obj interface{}) (UpdatePtrToPtrInner, error) {
	var it UpdatePtrToPtrInner
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"key", "value"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "key":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("key"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Key = data
		case "value":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("value"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Value = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputUpdatePtrToPtrOuter(ctx context.Context, obj interface{}) (UpdatePtrToPtrOuter, error) {
	var it UpdatePtrToPtrOuter
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"name", "inner", "stupidInner"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "name":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		case "inner":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("inner"))
			data, err := ec.unmarshalOUpdatePtrToPtrInner2ᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx, v)
			if err != nil {
				return it, err
			}
			it.Inner = data
		case "stupidInner":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("stupidInner"))
			data, err := ec.unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx, v)
			if err != nil {
				return it, err
			}
			it.StupidInner = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var ptrToPtrInnerImplementors = []string{"PtrToPtrInner"}

func (ec *executionContext) _PtrToPtrInner(ctx context.Context, sel ast.SelectionSet, obj *PtrToPtrInner) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, ptrToPtrInnerImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PtrToPtrInner")
		case "key":
			out.Values[i] = ec._PtrToPtrInner_key(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "value":
			out.Values[i] = ec._PtrToPtrInner_value(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
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

var ptrToPtrOuterImplementors = []string{"PtrToPtrOuter"}

func (ec *executionContext) _PtrToPtrOuter(ctx context.Context, sel ast.SelectionSet, obj *PtrToPtrOuter) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, ptrToPtrOuterImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PtrToPtrOuter")
		case "name":
			out.Values[i] = ec._PtrToPtrOuter_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "inner":
			out.Values[i] = ec._PtrToPtrOuter_inner(ctx, field, obj)
		case "stupidInner":
			out.Values[i] = ec._PtrToPtrOuter_stupidInner(ctx, field, obj)
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

func (ec *executionContext) marshalNPtrToPtrOuter2githubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrOuter(ctx context.Context, sel ast.SelectionSet, v PtrToPtrOuter) graphql.Marshaler {
	return ec._PtrToPtrOuter(ctx, sel, &v)
}

func (ec *executionContext) marshalNPtrToPtrOuter2ᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrOuter(ctx context.Context, sel ast.SelectionSet, v *PtrToPtrOuter) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._PtrToPtrOuter(ctx, sel, v)
}

func (ec *executionContext) unmarshalNUpdatePtrToPtrOuter2githubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrOuter(ctx context.Context, v interface{}) (UpdatePtrToPtrOuter, error) {
	res, err := ec.unmarshalInputUpdatePtrToPtrOuter(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOPtrToPtrInner2ᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx context.Context, sel ast.SelectionSet, v *PtrToPtrInner) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._PtrToPtrInner(ctx, sel, v)
}

func (ec *executionContext) marshalOPtrToPtrInner2ᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx context.Context, sel ast.SelectionSet, v **PtrToPtrInner) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOPtrToPtrInner2ᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx, sel, *v)
}

func (ec *executionContext) marshalOPtrToPtrInner2ᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx context.Context, sel ast.SelectionSet, v ***PtrToPtrInner) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOPtrToPtrInner2ᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx, sel, *v)
}

func (ec *executionContext) marshalOPtrToPtrInner2ᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx context.Context, sel ast.SelectionSet, v ****PtrToPtrInner) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOPtrToPtrInner2ᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx, sel, *v)
}

func (ec *executionContext) marshalOPtrToPtrInner2ᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx context.Context, sel ast.SelectionSet, v *****PtrToPtrInner) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOPtrToPtrInner2ᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx, sel, *v)
}

func (ec *executionContext) marshalOPtrToPtrInner2ᚖᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx context.Context, sel ast.SelectionSet, v ******PtrToPtrInner) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOPtrToPtrInner2ᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx, sel, *v)
}

func (ec *executionContext) marshalOPtrToPtrInner2ᚖᚖᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx context.Context, sel ast.SelectionSet, v *******PtrToPtrInner) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOPtrToPtrInner2ᚖᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrInner(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOUpdatePtrToPtrInner2ᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx context.Context, v interface{}) (*UpdatePtrToPtrInner, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputUpdatePtrToPtrInner(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOUpdatePtrToPtrInner2ᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx context.Context, v interface{}) (**UpdatePtrToPtrInner, error) {
	var pres *UpdatePtrToPtrInner
	if v != nil {
		res, err := ec.unmarshalOUpdatePtrToPtrInner2ᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx, v)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		pres = res
	}
	return &pres, nil
}

func (ec *executionContext) unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx context.Context, v interface{}) (***UpdatePtrToPtrInner, error) {
	var pres **UpdatePtrToPtrInner
	if v != nil {
		res, err := ec.unmarshalOUpdatePtrToPtrInner2ᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx, v)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		pres = res
	}
	return &pres, nil
}

func (ec *executionContext) unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx context.Context, v interface{}) (****UpdatePtrToPtrInner, error) {
	var pres ***UpdatePtrToPtrInner
	if v != nil {
		res, err := ec.unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx, v)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		pres = res
	}
	return &pres, nil
}

func (ec *executionContext) unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx context.Context, v interface{}) (*****UpdatePtrToPtrInner, error) {
	var pres ****UpdatePtrToPtrInner
	if v != nil {
		res, err := ec.unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx, v)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		pres = res
	}
	return &pres, nil
}

func (ec *executionContext) unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx context.Context, v interface{}) (******UpdatePtrToPtrInner, error) {
	var pres *****UpdatePtrToPtrInner
	if v != nil {
		res, err := ec.unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx, v)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		pres = res
	}
	return &pres, nil
}

func (ec *executionContext) unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx context.Context, v interface{}) (*******UpdatePtrToPtrInner, error) {
	var pres ******UpdatePtrToPtrInner
	if v != nil {
		res, err := ec.unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx, v)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		pres = res
	}
	return &pres, nil
}

func (ec *executionContext) unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx context.Context, v interface{}) (********UpdatePtrToPtrInner, error) {
	var pres *******UpdatePtrToPtrInner
	if v != nil {
		res, err := ec.unmarshalOUpdatePtrToPtrInner2ᚖᚖᚖᚖᚖᚖᚖgithubᚗcomᚋgenevaᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrInner(ctx, v)
		if err != nil {
			return nil, graphql.ErrorOnPath(ctx, err)
		}
		pres = res
	}
	return &pres, nil
}

// endregion ***************************** type.gotpl *****************************
