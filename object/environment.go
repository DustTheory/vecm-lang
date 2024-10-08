package object

import (
	"errors"
	"fmt"
)

type EnvStoreEntry struct {
	Object
	IsConstant bool
	IsExported bool
}

type Environment struct {
	typeStore map[string]Type
	store     map[string]*EnvStoreEntry
	outer     *Environment
}

var GlobalTypes = map[Kind]Kind{
	"char":      Int8Kind,
	"int":       Int64Kind,
	Int8Kind:    Int8Kind,
	Int16Kind:   Int16Kind,
	Int32Kind:   Int32Kind,
	Int64Kind:   Int64Kind,
	UInt8Kind:   UInt8Kind,
	UInt16Kind:  UInt16Kind,
	UInt32Kind:  UInt32Kind,
	UInt64Kind:  Int64Kind,
	Float32Kind: Float32Kind,
	Float64Kind: Float64Kind,
	BooleanKind: BooleanKind,
	NullKind:    NullKind,
	StringKind:  StringKind,
	VoidKind:    VoidKind,
}

func NewEnvironment() *Environment {
	s := make(map[string]*EnvStoreEntry)
	return &Environment{store: s, outer: nil, typeStore: map[string]Type{}}
}

func (e *Environment) GetReference(name string) Object {
	entry, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.GetReference(name)
	}
	if !ok {
		return nil
	}
	return &VariableReference{
		Env:  e,
		Name: name,
		ReferenceType: ReferenceType{
			IsConstantReference: entry.IsConstant,
			ValueType:           entry.Object.Type(),
		},
	}
}

func (e *Environment) Get(name string) Object {
	entry, ok := e.store[name]

	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}
	if !ok {
		return nil
	}
	return entry.Object
}

func (e *Environment) GetObjectType(name string) (Type, bool) {
	globalObjectType, globalObjectTypeExists := GlobalTypes[Kind(name)]
	if globalObjectTypeExists {
		return globalObjectType, true
	}
	objectType, objectTypeExists := e.typeStore[name]
	if !objectTypeExists && e.outer != nil {
		objectType, objectTypeExists = e.outer.GetObjectType(name)
	}
	return objectType, objectTypeExists
}

func (e *Environment) Declare(name string, isConstant bool, val Object) (Reference, error) {
	_, exists := e.store[name]
	if exists {
		return nil, fmt.Errorf("identifier with name %s already exists", name)
	}

	newReference := &VariableReference{e, name, ReferenceType{isConstant, val.Type()}}
	e.store[name] = &EnvStoreEntry{val, isConstant, false}
	return newReference, nil
}

func (e *Environment) Set(name string, val Object) (Object, error) {
	entry, exists := e.store[name]
	if exists && entry.IsConstant {
		return nil, errors.New("cannot assign to const variable")
	}
	e.store[name] = &EnvStoreEntry{val, entry.IsConstant, false}
	return e.store[name].Object, nil
}

func (e *Environment) GetStore() map[string]*EnvStoreEntry {
	return e.store
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
