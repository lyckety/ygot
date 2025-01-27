/*
Package ctestschema is a generated package which contains definitions
of structs which represent a YANG schema. The generated schema can be
compressed by a series of transformations (compression was true
in this case).

This package was generated by /usr/local/google/home/wenbli/gocode/src/github.com/openconfig/ygot/genutil/names.go
using the following YANG input files:
  - ../yang/ctestschema.yang

Imported modules were sourced from:
  - ...
*/
package ctestschema

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Binary is a type that is used for fields that have a YANG type of
// binary. It is used such that binary fields can be distinguished from
// leaf-lists of uint8s (which are mapped to []uint8, equivalent to
// []byte in reflection).
type Binary []byte

// YANGEmpty is a type that is used for fields that have a YANG type of
// empty. It is used such that empty fields can be distinguished from boolean fields
// in the generated code.
type YANGEmpty bool

// UnionInt8 is an int8 type assignable to unions of which it is a subtype.
type UnionInt8 int8

// UnionInt16 is an int16 type assignable to unions of which it is a subtype.
type UnionInt16 int16

// UnionInt32 is an int32 type assignable to unions of which it is a subtype.
type UnionInt32 int32

// UnionInt64 is an int64 type assignable to unions of which it is a subtype.
type UnionInt64 int64

// UnionUint8 is a uint8 type assignable to unions of which it is a subtype.
type UnionUint8 uint8

// UnionUint16 is a uint16 type assignable to unions of which it is a subtype.
type UnionUint16 uint16

// UnionUint32 is a uint32 type assignable to unions of which it is a subtype.
type UnionUint32 uint32

// UnionUint64 is a uint64 type assignable to unions of which it is a subtype.
type UnionUint64 uint64

// UnionFloat64 is a float64 type assignable to unions of which it is a subtype.
type UnionFloat64 float64

// UnionString is a string type assignable to unions of which it is a subtype.
type UnionString string

// UnionBool is a bool type assignable to unions of which it is a subtype.
type UnionBool bool

// UnionUnsupported is an interface{} wrapper type for unsupported types. It is
// assignable to unions of which it is a subtype.
type UnionUnsupported struct {
	Value interface{}
}

var (
	SchemaTree map[string]*yang.Entry
	ΛEnumTypes map[string][]reflect.Type
)

func init() {
	var err error
	initΛEnumTypes()
	if SchemaTree, err = UnzipSchema(); err != nil {
		panic("schema error: " + err.Error())
	}
}

// Schema returns the details of the generated schema.
func Schema() (*ytypes.Schema, error) {
	uzp, err := UnzipSchema()
	if err != nil {
		return nil, fmt.Errorf("cannot unzip schema, %v", err)
	}

	return &ytypes.Schema{
		Root:       &Device{},
		SchemaTree: uzp,
		Unmarshal:  Unmarshal,
	}, nil
}

// UnzipSchema unzips the zipped schema and returns a map of yang.Entry nodes,
// keyed by the name of the struct that the yang.Entry describes the schema for.
func UnzipSchema() (map[string]*yang.Entry, error) {
	var schemaTree map[string]*yang.Entry
	var err error
	if schemaTree, err = ygot.GzipToSchema(ySchema); err != nil {
		return nil, fmt.Errorf("could not unzip the schema; %v", err)
	}
	return schemaTree, nil
}

// Unmarshal unmarshals data, which must be RFC7951 JSON format, into
// destStruct, which must be non-nil and the correct GoStruct type. It returns
// an error if the destStruct is not found in the schema or the data cannot be
// unmarshaled. The supplied options (opts) are used to control the behaviour
// of the unmarshal function - for example, determining whether errors are
// thrown for unknown fields in the input JSON.
func Unmarshal(data []byte, destStruct ygot.GoStruct, opts ...ytypes.UnmarshalOpt) error {
	tn := reflect.TypeOf(destStruct).Elem().Name()
	schema, ok := SchemaTree[tn]
	if !ok {
		return fmt.Errorf("could not find schema for type %s", tn)
	}
	var jsonTree interface{}
	if err := json.Unmarshal([]byte(data), &jsonTree); err != nil {
		return err
	}
	return ytypes.Unmarshal(schema, destStruct, jsonTree, opts...)
}

// Device represents the /device YANG schema element.
type Device struct {
	ΛMetadata              []ygot.Annotation                 `path:"@" ygotAnnotation:"true"`
	OrderedList            *OrderedList_OrderedMap           `path:"ordered-lists/ordered-list" module:"ctestschema/ctestschema"`
	ΛOrderedList           []ygot.Annotation                 `path:"ordered-lists/@ordered-list" ygotAnnotation:"true"`
	OrderedMultikeyedList  *OrderedMultikeyedList_OrderedMap `path:"ordered-multikeyed-lists/ordered-multikeyed-list" module:"ctestschema/ctestschema"`
	ΛOrderedMultikeyedList []ygot.Annotation                 `path:"ordered-multikeyed-lists/@ordered-multikeyed-list" ygotAnnotation:"true"`
	UnorderedList          map[string]*UnorderedList         `path:"unordered-lists/unordered-list" module:"ctestschema/ctestschema"`
	ΛUnorderedList         []ygot.Annotation                 `path:"unordered-lists/@unordered-list" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that Device implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*Device) IsYANGGoStruct() {}

// OrderedMultikeyedList_Key represents the key for list OrderedMultikeyedList of element /device.
type OrderedMultikeyedList_Key struct {
	Key1 string `path:"key1"`
	Key2 uint64 `path:"key2"`
}

// IsYANGGoKeyStruct ensures that OrderedMultikeyedList_Key partially implements the
// yang.GoKeyStruct interface. This allows functions that need to
// handle this key struct to identify it as being generated by gogen.
func (OrderedMultikeyedList_Key) IsYANGGoKeyStruct() {}

// ΛListKeyMap returns the values of the OrderedMultikeyedList_Key key struct.
func (t OrderedMultikeyedList_Key) ΛListKeyMap() (map[string]interface{}, error) {
	return map[string]interface{}{
		"key1": t.Key1,
		"key2": t.Key2,
	}, nil
}

// NewUnorderedList creates a new entry in the UnorderedList list of the
// Device struct. The keys of the list are populated from the input
// arguments.
func (t *Device) NewUnorderedList(Key string) (*UnorderedList, error) {

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.UnorderedList == nil {
		t.UnorderedList = make(map[string]*UnorderedList)
	}

	key := Key

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.UnorderedList[key]; ok {
		return nil, fmt.Errorf("duplicate key %v for list UnorderedList", key)
	}

	t.UnorderedList[key] = &UnorderedList{
		Key: &Key,
	}

	return t.UnorderedList[key], nil
}

// RenameUnorderedList renames an entry in the list UnorderedList within
// the Device struct. The entry with key oldK is renamed to newK updating
// the key within the value.
func (t *Device) RenameUnorderedList(oldK, newK string) error {
	if _, ok := t.UnorderedList[newK]; ok {
		return fmt.Errorf("key %v already exists in UnorderedList", newK)
	}

	e, ok := t.UnorderedList[oldK]
	if !ok {
		return fmt.Errorf("key %v not found in UnorderedList", oldK)
	}
	e.Key = &newK

	t.UnorderedList[newK] = e
	delete(t.UnorderedList, oldK)
	return nil
}

// GetOrCreateUnorderedList retrieves the value with the specified keys from
// the receiver Device. If the entry does not exist, then it is created.
// It returns the existing or new list member.
func (t *Device) GetOrCreateUnorderedList(Key string) *UnorderedList {

	key := Key

	if v, ok := t.UnorderedList[key]; ok {
		return v
	}
	// Panic if we receive an error, since we should have retrieved an existing
	// list member. This allows chaining of GetOrCreate methods.
	v, err := t.NewUnorderedList(Key)
	if err != nil {
		panic(fmt.Sprintf("GetOrCreateUnorderedList got unexpected error: %v", err))
	}
	return v
}

// GetUnorderedList retrieves the value with the specified key from
// the UnorderedList map field of Device. If the receiver is nil, or
// the specified key is not present in the list, nil is returned such that Get*
// methods may be safely chained.
func (t *Device) GetUnorderedList(Key string) *UnorderedList {

	if t == nil {
		return nil
	}

	key := Key

	if lm, ok := t.UnorderedList[key]; ok {
		return lm
	}
	return nil
}

// AppendUnorderedList appends the supplied UnorderedList struct to the
// list UnorderedList of Device. If the key value(s) specified in
// the supplied UnorderedList already exist in the list, an error is
// returned.
func (t *Device) AppendUnorderedList(v *UnorderedList) error {
	if v.Key == nil {
		return fmt.Errorf("invalid nil key received for Key")
	}

	key := *v.Key

	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.UnorderedList == nil {
		t.UnorderedList = make(map[string]*UnorderedList)
	}

	if _, ok := t.UnorderedList[key]; ok {
		return fmt.Errorf("duplicate key for list UnorderedList %v", key)
	}

	t.UnorderedList[key] = v
	return nil
}

// AppendNewOrderedList creates a new entry in the OrderedList
// ordered map of the Device struct. The keys of the list are
// populated from the input arguments.
func (s *Device) AppendNewOrderedList(Key string) (*OrderedList, error) {
	if s.OrderedList == nil {
		s.OrderedList = &OrderedList_OrderedMap{}
	}
	return s.OrderedList.AppendNew(Key)
}

// AppendOrderedList appends the supplied OrderedList struct
// to the list OrderedList of Device. If the key value(s)
// specified in the supplied OrderedList already exist in the list, an
// error is returned.
func (s *Device) AppendOrderedList(v *OrderedList) error {
	if s.OrderedList == nil {
		s.OrderedList = &OrderedList_OrderedMap{}
	}
	return s.OrderedList.Append(v)
}

// GetOrderedList retrieves the value with the specified key from the
// OrderedList map field of Device. If the receiver
// is nil, or the specified key is not present in the list, nil is returned
// such that Get* methods may be safely chained.
func (s *Device) GetOrderedList(Key string) *OrderedList {
	key := Key
	return s.OrderedList.Get(key)
}

// DeleteOrderedList deletes the value with the specified keys from
// the receiver Device. If there is no such element, the
// function is a no-op.
func (s *Device) DeleteOrderedList(Key string) bool {
	key := Key
	return s.OrderedList.Delete(key)
}

// OrderedList_OrderedMap is an ordered map that represents the "ordered-by user"
// list elements at /ctestschema/ordered-lists/ordered-list.
type OrderedList_OrderedMap struct {
	keys     []string
	valueMap map[string]*OrderedList
}

// IsYANGOrderedList ensures that OrderedList_OrderedMap implements the
// ygot.GoOrderedList interface.
func (*OrderedList_OrderedMap) IsYANGOrderedList() {}

// init initializes any uninitialized values.
func (o *OrderedList_OrderedMap) init() {
	if o == nil {
		return
	}
	if o.valueMap == nil {
		o.valueMap = map[string]*OrderedList{}
	}
}

// Keys returns a copy of the list's keys.
func (o *OrderedList_OrderedMap) Keys() []string {
	if o == nil {
		return nil
	}
	return append([]string{}, o.keys...)
}

// Values returns the current set of the list's values in order.
func (o *OrderedList_OrderedMap) Values() []*OrderedList {
	if o == nil {
		return nil
	}
	var values []*OrderedList
	for _, key := range o.keys {
		values = append(values, o.valueMap[key])
	}
	return values
}

// Len returns a size of OrderedList_OrderedMap
func (o *OrderedList_OrderedMap) Len() int {
	if o == nil {
		return 0
	}
	return len(o.keys)
}

// Get returns the value corresponding to the key. If the key is not found, nil
// is returned.
func (o *OrderedList_OrderedMap) Get(key string) *OrderedList {
	if o == nil {
		return nil
	}
	val, _ := o.valueMap[key]
	return val
}

// Delete deletes an element.
func (o *OrderedList_OrderedMap) Delete(key string) bool {
	if o == nil {
		return false
	}
	if _, ok := o.valueMap[key]; !ok {
		return false
	}
	for i, k := range o.keys {
		if k == key {
			o.keys = append(o.keys[:i], o.keys[i+1:]...)
			delete(o.valueMap, key)
			return true
		}
	}
	return false
}

// Append appends a OrderedList, returning an error if the key
// already exists in the ordered list or if the key is unspecified.
func (o *OrderedList_OrderedMap) Append(v *OrderedList) error {
	if o == nil {
		return fmt.Errorf("nil ordered map, cannot append OrderedList")
	}
	if v == nil {
		return fmt.Errorf("nil OrderedList")
	}
	if v.Key == nil {
		return fmt.Errorf("invalid nil key received for Key")
	}

	key := *v.Key

	if _, ok := o.valueMap[key]; ok {
		return fmt.Errorf("duplicate key for list Statement %v", key)
	}
	o.keys = append(o.keys, key)
	o.init()
	o.valueMap[key] = v
	return nil
}

// AppendNew creates and appends a new OrderedList, returning the
// newly-initialized v. It returns an error if the v already exists.
func (o *OrderedList_OrderedMap) AppendNew(Key string) (*OrderedList, error) {
	if o == nil {
		return nil, fmt.Errorf("nil ordered map, cannot append OrderedList")
	}
	key := Key

	if _, ok := o.valueMap[key]; ok {
		return nil, fmt.Errorf("duplicate key for list Statement %v", key)
	}
	o.keys = append(o.keys, key)
	newElement := &OrderedList{
		Key: &Key,
	}
	o.init()
	o.valueMap[key] = newElement
	return newElement, nil
}

// AppendNewOrderedMultikeyedList creates a new entry in the OrderedMultikeyedList
// ordered map of the Device struct. The keys of the list are
// populated from the input arguments.
func (s *Device) AppendNewOrderedMultikeyedList(Key1 string, Key2 uint64) (*OrderedMultikeyedList, error) {
	if s.OrderedMultikeyedList == nil {
		s.OrderedMultikeyedList = &OrderedMultikeyedList_OrderedMap{}
	}
	return s.OrderedMultikeyedList.AppendNew(Key1, Key2)
}

// AppendOrderedMultikeyedList appends the supplied OrderedMultikeyedList struct
// to the list OrderedMultikeyedList of Device. If the key value(s)
// specified in the supplied OrderedMultikeyedList already exist in the list, an
// error is returned.
func (s *Device) AppendOrderedMultikeyedList(v *OrderedMultikeyedList) error {
	if s.OrderedMultikeyedList == nil {
		s.OrderedMultikeyedList = &OrderedMultikeyedList_OrderedMap{}
	}
	return s.OrderedMultikeyedList.Append(v)
}

// GetOrderedMultikeyedList retrieves the value with the specified key from the
// OrderedMultikeyedList map field of Device. If the receiver
// is nil, or the specified key is not present in the list, nil is returned
// such that Get* methods may be safely chained.
func (s *Device) GetOrderedMultikeyedList(Key1 string, Key2 uint64) *OrderedMultikeyedList {
	key := OrderedMultikeyedList_Key{
		Key1: Key1,
		Key2: Key2,
	}
	return s.OrderedMultikeyedList.Get(key)
}

// DeleteOrderedMultikeyedList deletes the value with the specified keys from
// the receiver Device. If there is no such element, the
// function is a no-op.
func (s *Device) DeleteOrderedMultikeyedList(Key1 string, Key2 uint64) bool {
	key := OrderedMultikeyedList_Key{
		Key1: Key1,
		Key2: Key2,
	}
	return s.OrderedMultikeyedList.Delete(key)
}

// OrderedMultikeyedList_OrderedMap is an ordered map that represents the "ordered-by user"
// list elements at /ctestschema/ordered-multikeyed-lists/ordered-multikeyed-list.
type OrderedMultikeyedList_OrderedMap struct {
	keys     []OrderedMultikeyedList_Key
	valueMap map[OrderedMultikeyedList_Key]*OrderedMultikeyedList
}

// IsYANGOrderedList ensures that OrderedMultikeyedList_OrderedMap implements the
// ygot.GoOrderedList interface.
func (*OrderedMultikeyedList_OrderedMap) IsYANGOrderedList() {}

// init initializes any uninitialized values.
func (o *OrderedMultikeyedList_OrderedMap) init() {
	if o == nil {
		return
	}
	if o.valueMap == nil {
		o.valueMap = map[OrderedMultikeyedList_Key]*OrderedMultikeyedList{}
	}
}

// Keys returns a copy of the list's keys.
func (o *OrderedMultikeyedList_OrderedMap) Keys() []OrderedMultikeyedList_Key {
	if o == nil {
		return nil
	}
	return append([]OrderedMultikeyedList_Key{}, o.keys...)
}

// Values returns the current set of the list's values in order.
func (o *OrderedMultikeyedList_OrderedMap) Values() []*OrderedMultikeyedList {
	if o == nil {
		return nil
	}
	var values []*OrderedMultikeyedList
	for _, key := range o.keys {
		values = append(values, o.valueMap[key])
	}
	return values
}

// Len returns a size of OrderedMultikeyedList_OrderedMap
func (o *OrderedMultikeyedList_OrderedMap) Len() int {
	if o == nil {
		return 0
	}
	return len(o.keys)
}

// Get returns the value corresponding to the key. If the key is not found, nil
// is returned.
func (o *OrderedMultikeyedList_OrderedMap) Get(key OrderedMultikeyedList_Key) *OrderedMultikeyedList {
	if o == nil {
		return nil
	}
	val, _ := o.valueMap[key]
	return val
}

// Delete deletes an element.
func (o *OrderedMultikeyedList_OrderedMap) Delete(key OrderedMultikeyedList_Key) bool {
	if o == nil {
		return false
	}
	if _, ok := o.valueMap[key]; !ok {
		return false
	}
	for i, k := range o.keys {
		if k == key {
			o.keys = append(o.keys[:i], o.keys[i+1:]...)
			delete(o.valueMap, key)
			return true
		}
	}
	return false
}

// Append appends a OrderedMultikeyedList, returning an error if the key
// already exists in the ordered list or if the key is unspecified.
func (o *OrderedMultikeyedList_OrderedMap) Append(v *OrderedMultikeyedList) error {
	if o == nil {
		return fmt.Errorf("nil ordered map, cannot append OrderedMultikeyedList")
	}
	if v == nil {
		return fmt.Errorf("nil OrderedMultikeyedList")
	}
	if v.Key1 == nil {
		return fmt.Errorf("invalid nil key for Key1")
	}
	if v.Key2 == nil {
		return fmt.Errorf("invalid nil key for Key2")
	}
	key := OrderedMultikeyedList_Key{
		Key1: *v.Key1,
		Key2: *v.Key2,
	}

	if _, ok := o.valueMap[key]; ok {
		return fmt.Errorf("duplicate key for list Statement %v", key)
	}
	o.keys = append(o.keys, key)
	o.init()
	o.valueMap[key] = v
	return nil
}

// AppendNew creates and appends a new OrderedMultikeyedList, returning the
// newly-initialized v. It returns an error if the v already exists.
func (o *OrderedMultikeyedList_OrderedMap) AppendNew(Key1 string, Key2 uint64) (*OrderedMultikeyedList, error) {
	if o == nil {
		return nil, fmt.Errorf("nil ordered map, cannot append OrderedMultikeyedList")
	}
	key := OrderedMultikeyedList_Key{
		Key1: Key1,
		Key2: Key2,
	}

	if _, ok := o.valueMap[key]; ok {
		return nil, fmt.Errorf("duplicate key for list Statement %v", key)
	}
	o.keys = append(o.keys, key)
	newElement := &OrderedMultikeyedList{
		Key1: &Key1,
		Key2: &Key2,
	}
	o.init()
	o.valueMap[key] = newElement
	return newElement, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Device) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["Device"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *Device) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *Device) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of Device.
func (*Device) ΛBelongingModule() string {
	return ""
}

// OrderedList represents the /ctestschema/ordered-lists/ordered-list YANG schema element.
type OrderedList struct {
	ΛMetadata []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	Key       *string           `path:"config/key|key" module:"ctestschema/ctestschema|ctestschema"`
	ΛKey      []ygot.Annotation `path:"config/@key|@key" ygotAnnotation:"true"`
	Value     *string           `path:"config/value" module:"ctestschema/ctestschema"`
	ΛValue    []ygot.Annotation `path:"config/@value" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that OrderedList implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*OrderedList) IsYANGGoStruct() {}

// GetKey retrieves the value of the leaf Key from the OrderedList
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Key is set, it can
// safely use t.GetKey() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Key == nil' before retrieving the leaf's value.
func (t *OrderedList) GetKey() string {
	if t == nil || t.Key == nil {
		return ""
	}
	return *t.Key
}

// GetValue retrieves the value of the leaf Value from the OrderedList
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Value is set, it can
// safely use t.GetValue() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Value == nil' before retrieving the leaf's value.
func (t *OrderedList) GetValue() string {
	if t == nil || t.Value == nil {
		return ""
	}
	return *t.Value
}

// ΛListKeyMap returns the keys of the OrderedList struct, which is a YANG list entry.
func (t *OrderedList) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Key == nil {
		return nil, fmt.Errorf("nil value for key Key")
	}

	return map[string]interface{}{
		"key": *t.Key,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OrderedList) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["OrderedList"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OrderedList) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *OrderedList) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of OrderedList.
func (*OrderedList) ΛBelongingModule() string {
	return "ctestschema"
}

// OrderedMultikeyedList represents the /ctestschema/ordered-multikeyed-lists/ordered-multikeyed-list YANG schema element.
type OrderedMultikeyedList struct {
	ΛMetadata []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	Key1      *string           `path:"config/key1|key1" module:"ctestschema/ctestschema|ctestschema"`
	ΛKey1     []ygot.Annotation `path:"config/@key1|@key1" ygotAnnotation:"true"`
	Key2      *uint64           `path:"config/key2|key2" module:"ctestschema/ctestschema|ctestschema"`
	ΛKey2     []ygot.Annotation `path:"config/@key2|@key2" ygotAnnotation:"true"`
	Value     *string           `path:"config/value" module:"ctestschema/ctestschema"`
	ΛValue    []ygot.Annotation `path:"config/@value" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that OrderedMultikeyedList implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*OrderedMultikeyedList) IsYANGGoStruct() {}

// GetKey1 retrieves the value of the leaf Key1 from the OrderedMultikeyedList
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Key1 is set, it can
// safely use t.GetKey1() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Key1 == nil' before retrieving the leaf's value.
func (t *OrderedMultikeyedList) GetKey1() string {
	if t == nil || t.Key1 == nil {
		return ""
	}
	return *t.Key1
}

// GetKey2 retrieves the value of the leaf Key2 from the OrderedMultikeyedList
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Key2 is set, it can
// safely use t.GetKey2() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Key2 == nil' before retrieving the leaf's value.
func (t *OrderedMultikeyedList) GetKey2() uint64 {
	if t == nil || t.Key2 == nil {
		return 0
	}
	return *t.Key2
}

// GetValue retrieves the value of the leaf Value from the OrderedMultikeyedList
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Value is set, it can
// safely use t.GetValue() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Value == nil' before retrieving the leaf's value.
func (t *OrderedMultikeyedList) GetValue() string {
	if t == nil || t.Value == nil {
		return ""
	}
	return *t.Value
}

// ΛListKeyMap returns the keys of the OrderedMultikeyedList struct, which is a YANG list entry.
func (t *OrderedMultikeyedList) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Key1 == nil {
		return nil, fmt.Errorf("nil value for key Key1")
	}

	if t.Key2 == nil {
		return nil, fmt.Errorf("nil value for key Key2")
	}

	return map[string]interface{}{
		"key1": *t.Key1,
		"key2": *t.Key2,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OrderedMultikeyedList) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["OrderedMultikeyedList"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *OrderedMultikeyedList) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *OrderedMultikeyedList) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of OrderedMultikeyedList.
func (*OrderedMultikeyedList) ΛBelongingModule() string {
	return "ctestschema"
}

// UnorderedList represents the /ctestschema/unordered-lists/unordered-list YANG schema element.
type UnorderedList struct {
	ΛMetadata []ygot.Annotation `path:"@" ygotAnnotation:"true"`
	Key       *string           `path:"config/key|key" module:"ctestschema/ctestschema|ctestschema"`
	ΛKey      []ygot.Annotation `path:"config/@key|@key" ygotAnnotation:"true"`
	Value     *string           `path:"config/value" module:"ctestschema/ctestschema"`
	ΛValue    []ygot.Annotation `path:"config/@value" ygotAnnotation:"true"`
}

// IsYANGGoStruct ensures that UnorderedList implements the yang.GoStruct
// interface. This allows functions that need to handle this struct to
// identify it as being generated by ygen.
func (*UnorderedList) IsYANGGoStruct() {}

// GetKey retrieves the value of the leaf Key from the UnorderedList
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Key is set, it can
// safely use t.GetKey() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Key == nil' before retrieving the leaf's value.
func (t *UnorderedList) GetKey() string {
	if t == nil || t.Key == nil {
		return ""
	}
	return *t.Key
}

// GetValue retrieves the value of the leaf Value from the UnorderedList
// struct. If the field is unset but has a default value in the YANG schema,
// then the default value will be returned.
// Caution should be exercised whilst using this method since when without a
// default value, it will return the Go zero value if the field is explicitly
// unset. If the caller explicitly does not care if Value is set, it can
// safely use t.GetValue() to retrieve the value. In the case that the
// caller has different actions based on whether the leaf is set or unset, it
// should use 'if t.Value == nil' before retrieving the leaf's value.
func (t *UnorderedList) GetValue() string {
	if t == nil || t.Value == nil {
		return ""
	}
	return *t.Value
}

// ΛListKeyMap returns the keys of the UnorderedList struct, which is a YANG list entry.
func (t *UnorderedList) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Key == nil {
		return nil, fmt.Errorf("nil value for key Key")
	}

	return map[string]interface{}{
		"key": *t.Key,
	}, nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *UnorderedList) ΛValidate(opts ...ygot.ValidationOption) error {
	if err := ytypes.Validate(SchemaTree["UnorderedList"], t, opts...); err != nil {
		return err
	}
	return nil
}

// Validate validates s against the YANG schema corresponding to its type.
func (t *UnorderedList) Validate(opts ...ygot.ValidationOption) error {
	return t.ΛValidate(opts...)
}

// ΛEnumTypeMap returns a map, keyed by YANG schema path, of the enumerated types
// that are included in the generated code.
func (t *UnorderedList) ΛEnumTypeMap() map[string][]reflect.Type { return ΛEnumTypes }

// ΛBelongingModule returns the name of the module that defines the namespace
// of UnorderedList.
func (*UnorderedList) ΛBelongingModule() string {
	return "ctestschema"
}

var (
	// ySchema is a byte slice contain a gzip compressed representation of the
	// YANG schema from which the Go code was generated. When uncompressed the
	// contents of the byte slice is a JSON document containing an object, keyed
	// on the name of the generated struct, and containing the JSON marshalled
	// contents of a goyang yang.Entry struct, which defines the schema for the
	// fields within the struct.
	ySchema = []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x5d, 0x4b, 0x73, 0xdb, 0x36,
		0x10, 0xbe, 0xeb, 0x57, 0x70, 0x70, 0x96, 0x63, 0x51, 0x95, 0x1f, 0xd5, 0xcd, 0x8d, 0x9b, 0xe9,
		0x4c, 0xea, 0x36, 0x93, 0x34, 0xbd, 0x74, 0x72, 0x40, 0x68, 0x48, 0xc1, 0x98, 0x22, 0x3d, 0x04,
		0xe8, 0x5a, 0xd3, 0xd1, 0x7f, 0xef, 0x50, 0xa4, 0x18, 0x92, 0xe2, 0x63, 0x01, 0x89, 0xb6, 0xa4,
		0x7c, 0x47, 0x91, 0x0b, 0x02, 0xfb, 0xf8, 0x96, 0xbb, 0x8b, 0x05, 0xf5, 0xdf, 0xc0, 0x71, 0x1c,
		0x87, 0xfd, 0xc1, 0x17, 0x82, 0x4d, 0x1d, 0x76, 0x2f, 0x9e, 0xa4, 0x27, 0xd8, 0x30, 0xbd, 0xfa,
		0x5e, 0x06, 0xf7, 0x6c, 0xea, 0xb8, 0xd9, 0xcf, 0xb7, 0x61, 0x30, 0x93, 0x73, 0x36, 0x75, 0x46,
		0xd9, 0x85, 0x5b, 0x19, 0xb1, 0xa9, 0x93, 0x3e, 0x62, 0x7d, 0x21, 0x8c, 0xee, 0x45, 0x24, 0xee,
		0xcf, 0x7c, 0xa9, 0xb4, 0x2a, 0xdd, 0x2a, 0xcd, 0x52, 0x26, 0x1b, 0x96, 0x89, 0xca, 0x93, 0xe6,
		0x97, 0xab, 0x93, 0xe7, 0x37, 0x3e, 0x44, 0x62, 0x26, 0x9f, 0xb7, 0x26, 0x2b, 0x4d, 0xe8, 0x6d,
		0x4d, 0xb3, 0xbe, 0xfd, 0x29, 0x8c, 0x23, 0x4f, 0xd4, 0x0e, 0x4d, 0x97, 0x22, 0x96, 0xff, 0x86,
		0x51, 0xb2, 0x1a, 0xf6, 0x98, 0xce, 0x32, 0xac, 0x27, 0xfc, 0x8d, 0xab, 0x9b, 0x68, 0x1e, 0x2f,
		0x44, 0xa0, 0xd9, 0xd4, 0xd1, 0x51, 0x2c, 0x1a, 0x08, 0x0b, 0x54, 0xeb, 0x45, 0x6d, 0x51, 0xad,
		0x4a, 0x57, 0x56, 0x15, 0x5e, 0xab, 0x02, 0xaf, 0x15, 0x7c, 0x33, 0x3f, 0x75, 0xf2, 0x6f, 0x62,
		0xa9, 0x5e, 0x0d, 0x9d, 0xea, 0xa0, 0xa8, 0x85, 0xa8, 0x1e, 0xaa, 0x9a, 0x8c, 0xd5, 0x65, 0xac,
		0x36, 0xba, 0xfa, 0xea, 0xd5, 0xd8, 0xa0, 0xce, 0x4e, 0xb5, 0xe6, 0x04, 0xde, 0x46, 0xda, 0x1d,
		0x12, 0xc8, 0x05, 0x9a, 0xd2, 0x77, 0x70, 0xd3, 0xae, 0x62, 0xb2, 0xaa, 0x4d, 0x54, 0x6e, 0xa8,
		0x7a, 0x53, 0x13, 0xb0, 0x36, 0x05, 0x6b, 0x93, 0x30, 0x37, 0x8d, 0x76, 0x13, 0xe9, 0x30, 0x15,
		0xb2, 0xc9, 0xe4, 0x84, 0x0f, 0x62, 0x49, 0x17, 0xdb, 0x46, 0x2b, 0xc9, 0x20, 0x22, 0xdf, 0x99,
		0x11, 0x8d, 0x88, 0xe4, 0x54, 0x63, 0xb2, 0x31, 0x2a, 0x4b, 0xe3, 0xb2, 0x35, 0xb2, 0x9d, 0x8d,
		0x6d, 0x67, 0xa3, 0xb3, 0x37, 0x3e, 0x9a, 0x11, 0x12, 0x8d, 0x31, 0x5f, 0xc6, 0x5f, 0xcb, 0x47,
		0x61, 0xa7, 0x29, 0xa5, 0x23, 0x19, 0xcc, 0x4d, 0x94, 0xb5, 0x71, 0x5e, 0xd7, 0x7b, 0xe5, 0xe0,
		0x26, 0x08, 0x42, 0xcd, 0xb5, 0x0c, 0x03, 0x33, 0x3e, 0x96, 0xf3, 0x50, 0x9f, 0x85, 0xde, 0x99,
		0x17, 0x2e, 0x1e, 0x23, 0xa1, 0x54, 0xf2, 0x72, 0x15, 0x7c, 0x96, 0x3c, 0x84, 0x28, 0xe2, 0xc1,
		0x1e, 0x58, 0x60, 0x4f, 0xdc, 0x8f, 0x85, 0x39, 0xdc, 0xd3, 0x61, 0x00, 0x3c, 0x00, 0x0f, 0xc0,
		0x1f, 0x10, 0xe0, 0x77, 0x0a, 0x10, 0x0c, 0x19, 0x63, 0xca, 0xfb, 0x26, 0x16, 0xfc, 0x91, 0xeb,
		0x6f, 0x89, 0x7a, 0xce, 0x3d, 0x2d, 0x94, 0x4e, 0xaf, 0x9d, 0x97, 0xd2, 0xb5, 0xd2, 0xaf, 0xf3,
		0x2c, 0xd0, 0x1c, 0xd8, 0xf1, 0xd1, 0xc2, 0x03, 0x29, 0x6e, 0x31, 0x88, 0x57, 0x88, 0x6e, 0x0b,
		0xc1, 0xee, 0x21, 0x06, 0xbb, 0x64, 0x37, 0x93, 0x4b, 0x3a, 0xc1, 0x62, 0x24, 0x66, 0x14, 0x69,
		0x6f, 0xfc, 0xca, 0x15, 0x81, 0xf6, 0x43, 0x06, 0x8f, 0x37, 0x6f, 0x32, 0xcb, 0x3f, 0x4f, 0x2c,
		0xaf, 0x07, 0xeb, 0x57, 0x9a, 0x6b, 0x41, 0xb7, 0xff, 0x94, 0x7c, 0xcf, 0xe9, 0xde, 0x18, 0x08,
		0x40, 0xba, 0x87, 0xe8, 0x0f, 0xd1, 0xdf, 0x11, 0x46, 0x7f, 0xc8, 0xa6, 0x80, 0x27, 0xe0, 0xe9,
		0x25, 0xf1, 0x74, 0x04, 0xc9, 0x4a, 0x1a, 0x25, 0xd9, 0x46, 0x6b, 0x46, 0x15, 0xfd, 0xf7, 0xeb,
		0xd7, 0x70, 0xcb, 0x8b, 0x95, 0xfd, 0x2e, 0x95, 0xbe, 0xd1, 0xba, 0xa3, 0xf0, 0x7f, 0x27, 0x83,
		0x5f, 0x7d, 0x91, 0xd8, 0xa6, 0x6a, 0x07, 0x3d, 0xbb, 0xe3, 0xcf, 0x05, 0x4a, 0xf7, 0x7a, 0x32,
		0xb9, 0xbc, 0x9a, 0x4c, 0x46, 0x57, 0x3f, 0x5d, 0x8d, 0x7e, 0xbe, 0xb8, 0x70, 0x2f, 0xdd, 0x8b,
		0x96, 0xc1, 0x7f, 0xa6, 0x62, 0xfa, 0xc5, 0x20, 0xe5, 0x8a, 0x95, 0x88, 0xba, 0x22, 0x4e, 0x03,
		0x8f, 0x50, 0xf4, 0x02, 0x1b, 0xa5, 0x7d, 0xa5, 0x44, 0x25, 0x56, 0xe8, 0x2f, 0x21, 0x7e, 0xcd,
		0x49, 0x0f, 0x31, 0x7c, 0x2e, 0xd4, 0xcf, 0xc9, 0x04, 0xe9, 0xd2, 0x8c, 0x8c, 0x48, 0x3c, 0xeb,
		0x88, 0x9f, 0xc5, 0x81, 0xd2, 0xfc, 0xab, 0xdf, 0x2e, 0xc6, 0xa2, 0xcc, 0xa6, 0xce, 0x3f, 0xad,
		0xdc, 0x18, 0x84, 0xeb, 0x04, 0x25, 0xef, 0x1a, 0xaf, 0x1b, 0x29, 0x7b, 0x7f, 0x31, 0x7b, 0xb7,
		0xd2, 0x09, 0x5e, 0xad, 0xf1, 0xee, 0x17, 0x23, 0x3d, 0x13, 0xbd, 0x9f, 0x9d, 0xd7, 0x63, 0xad,
		0x59, 0x66, 0x14, 0x7b, 0x3a, 0xc8, 0x94, 0x9d, 0xd9, 0x6b, 0xe2, 0x98, 0xea, 0x25, 0xb3, 0x32,
		0xdc, 0xad, 0xee, 0xe0, 0x8b, 0xcc, 0x4f, 0x79, 0x35, 0xdf, 0xe7, 0x2c, 0xcc, 0x97, 0x9b, 0xd1,
		0x22, 0xf6, 0xb5, 0x7c, 0x10, 0x4b, 0x7a, 0x13, 0xc2, 0xd6, 0x08, 0xf4, 0x23, 0x90, 0x35, 0xdc,
		0xd9, 0x8f, 0x50, 0x11, 0x2e, 0xbd, 0x35, 0xa1, 0x3a, 0x10, 0x5d, 0x0a, 0xe8, 0x52, 0x40, 0xe1,
		0xf6, 0x94, 0xcb, 0x56, 0xae, 0x55, 0xdd, 0xca, 0x45, 0xa2, 0x8d, 0x44, 0xfb, 0x84, 0x12, 0xed,
		0x1f, 0xa3, 0x4f, 0xe1, 0x41, 0x2c, 0xc7, 0x56, 0x70, 0x1f, 0x03, 0xee, 0x80, 0xfb, 0xcb, 0xc1,
		0x3d, 0x96, 0x81, 0xbe, 0x9c, 0x58, 0xc0, 0xfd, 0xda, 0x60, 0xc8, 0x47, 0x1e, 0xcc, 0x45, 0x67,
		0xd6, 0x6e, 0x9e, 0xc5, 0xd7, 0x15, 0x93, 0x8c, 0xad, 0x28, 0x1f, 0xfc, 0x77, 0x56, 0x0c, 0x1f,
		0x0d, 0xed, 0xc6, 0xbf, 0x8b, 0xb8, 0x97, 0xf8, 0xab, 0x5b, 0x39, 0x97, 0x5d, 0xc5, 0xac, 0x76,
		0xdd, 0x88, 0x39, 0xd7, 0xf2, 0x29, 0x59, 0xcb, 0x8c, 0xfb, 0x4a, 0x18, 0x3f, 0x65, 0x35, 0xb4,
		0x10, 0x1d, 0x7f, 0xde, 0x5d, 0x74, 0x66, 0x45, 0xb9, 0x63, 0x91, 0xe6, 0xa0, 0x1f, 0xea, 0x2f,
		0x78, 0x5f, 0x7e, 0x5f, 0x06, 0x76, 0xa2, 0xf0, 0xc6, 0x44, 0x80, 0xec, 0xa0, 0xaf, 0xaf, 0x97,
		0xad, 0xb2, 0x6a, 0x05, 0xb4, 0xe9, 0x46, 0xbf, 0xdd, 0x7e, 0xae, 0x51, 0xbb, 0x9f, 0x8b, 0x7e,
		0x3f, 0xf4, 0xfb, 0xbd, 0x4e, 0xbf, 0x9f, 0xdb, 0x13, 0x00, 0xc6, 0x46, 0x00, 0x18, 0x03, 0x00,
		0x00, 0xc0, 0xeb, 0x00, 0x60, 0x8c, 0x8e, 0x57, 0x40, 0x00, 0x5b, 0x07, 0xc8, 0x8c, 0x90, 0x19,
		0x1d, 0x52, 0x66, 0x84, 0xca, 0x3c, 0xd0, 0x84, 0xca, 0x3c, 0x2a, 0xf3, 0xa8, 0xcc, 0xa3, 0x32,
		0x7f, 0x00, 0x95, 0x79, 0x14, 0xbe, 0xf1, 0x42, 0x42, 0x78, 0xf7, 0x92, 0xe1, 0xdd, 0xd1, 0xd6,
		0x95, 0x5f, 0xe5, 0x60, 0x86, 0xeb, 0xb4, 0x44, 0xad, 0x38, 0x9e, 0xd1, 0xec, 0x21, 0x70, 0x3c,
		0x23, 0x7d, 0x00, 0x8e, 0x67, 0xf4, 0x57, 0x60, 0x3a, 0xd1, 0xe3, 0x19, 0x54, 0x8f, 0x68, 0x7c,
		0x68, 0xe3, 0x2e, 0x1f, 0x7f, 0x50, 0xc7, 0x37, 0xb6, 0xce, 0x56, 0x10, 0x4e, 0x72, 0xc4, 0x01,
		0xf1, 0x2b, 0x92, 0x55, 0x42, 0x9c, 0xdb, 0x20, 0xab, 0xb6, 0xf1, 0xdc, 0x46, 0x59, 0xa6, 0xdd,
		0xc7, 0x35, 0x2a, 0xf4, 0x38, 0xa5, 0x81, 0x53, 0x1a, 0x7b, 0x49, 0x9e, 0xb0, 0xd5, 0x82, 0x8f,
		0x8b, 0x20, 0x13, 0x47, 0x26, 0x8e, 0x16, 0x34, 0xf4, 0x9c, 0x02, 0xf0, 0x00, 0x3c, 0x00, 0x7f,
		0x22, 0xb5, 0xc1, 0x4a, 0xd2, 0x56, 0xf9, 0x8d, 0xef, 0x49, 0x22, 0xe0, 0xc5, 0xf7, 0x24, 0xd1,
		0x5d, 0x07, 0x04, 0x20, 0xe5, 0x43, 0x04, 0x88, 0x08, 0xf0, 0x60, 0x22, 0x40, 0x64, 0x54, 0xc0,
		0x13, 0xf0, 0xf4, 0x92, 0x78, 0x3a, 0x92, 0x84, 0x05, 0xdf, 0x94, 0x6c, 0x1a, 0x5c, 0x6c, 0x5a,
		0x08, 0x62, 0xdf, 0x37, 0xd8, 0x8a, 0x6f, 0xee, 0x77, 0xeb, 0x71, 0x8f, 0xb6, 0x5d, 0xd1, 0xf4,
		0x1d, 0xd9, 0xcf, 0x9b, 0x71, 0x07, 0xb1, 0x13, 0x5b, 0xdd, 0x24, 0x6d, 0xda, 0x80, 0x1d, 0x14,
		0xe6, 0x6d, 0x9a, 0x8f, 0x49, 0xf5, 0x36, 0x2f, 0x51, 0x7c, 0x5a, 0x3f, 0x7f, 0xcb, 0x3f, 0x32,
		0xa9, 0xde, 0xf1, 0x07, 0xf1, 0x31, 0x0c, 0xb7, 0x7d, 0x67, 0x75, 0x9d, 0xac, 0x78, 0xab, 0x24,
		0xc3, 0xdb, 0xf4, 0x6f, 0x08, 0xd3, 0x45, 0x0d, 0x56, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x01,
		0x00, 0x00, 0xff, 0xff, 0x86, 0xd6, 0x57, 0x77, 0xa5, 0x70, 0x00, 0x00,
	}
)

// ΛEnumTypes is a map, keyed by a YANG schema path, of the enumerated types that
// correspond with the leaf. The type is represented as a reflect.Type. The naming
// of the map ensures that there are no clashes with valid YANG identifiers.
func initΛEnumTypes() {
	ΛEnumTypes = map[string][]reflect.Type{}
}
