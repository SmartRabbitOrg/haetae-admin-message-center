// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-message-center/ent/emaillog"
	"github.com/suyuan32/simple-admin-message-center/ent/predicate"
	"github.com/suyuan32/simple-admin-message-center/ent/smslog"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEmailLog = "EmailLog"
	TypeSmsLog   = "SmsLog"
)

// EmailLogMutation represents an operation that mutates the EmailLog nodes in the graph.
type EmailLogMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	created_at     *time.Time
	updated_at     *time.Time
	target         *string
	subject        *string
	content        *string
	send_status    *uint8
	addsend_status *int8
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*EmailLog, error)
	predicates     []predicate.EmailLog
}

var _ ent.Mutation = (*EmailLogMutation)(nil)

// emaillogOption allows management of the mutation configuration using functional options.
type emaillogOption func(*EmailLogMutation)

// newEmailLogMutation creates new mutation for the EmailLog entity.
func newEmailLogMutation(c config, op Op, opts ...emaillogOption) *EmailLogMutation {
	m := &EmailLogMutation{
		config:        c,
		op:            op,
		typ:           TypeEmailLog,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEmailLogID sets the ID field of the mutation.
func withEmailLogID(id uuid.UUID) emaillogOption {
	return func(m *EmailLogMutation) {
		var (
			err   error
			once  sync.Once
			value *EmailLog
		)
		m.oldValue = func(ctx context.Context) (*EmailLog, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().EmailLog.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEmailLog sets the old EmailLog of the mutation.
func withEmailLog(node *EmailLog) emaillogOption {
	return func(m *EmailLogMutation) {
		m.oldValue = func(context.Context) (*EmailLog, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EmailLogMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EmailLogMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of EmailLog entities.
func (m *EmailLogMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EmailLogMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EmailLogMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().EmailLog.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *EmailLogMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *EmailLogMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the EmailLog entity.
// If the EmailLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailLogMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *EmailLogMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *EmailLogMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *EmailLogMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the EmailLog entity.
// If the EmailLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailLogMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *EmailLogMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetTarget sets the "target" field.
func (m *EmailLogMutation) SetTarget(s string) {
	m.target = &s
}

// Target returns the value of the "target" field in the mutation.
func (m *EmailLogMutation) Target() (r string, exists bool) {
	v := m.target
	if v == nil {
		return
	}
	return *v, true
}

// OldTarget returns the old "target" field's value of the EmailLog entity.
// If the EmailLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailLogMutation) OldTarget(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTarget is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTarget requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTarget: %w", err)
	}
	return oldValue.Target, nil
}

// ResetTarget resets all changes to the "target" field.
func (m *EmailLogMutation) ResetTarget() {
	m.target = nil
}

// SetSubject sets the "subject" field.
func (m *EmailLogMutation) SetSubject(s string) {
	m.subject = &s
}

// Subject returns the value of the "subject" field in the mutation.
func (m *EmailLogMutation) Subject() (r string, exists bool) {
	v := m.subject
	if v == nil {
		return
	}
	return *v, true
}

// OldSubject returns the old "subject" field's value of the EmailLog entity.
// If the EmailLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailLogMutation) OldSubject(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSubject is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSubject requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSubject: %w", err)
	}
	return oldValue.Subject, nil
}

// ResetSubject resets all changes to the "subject" field.
func (m *EmailLogMutation) ResetSubject() {
	m.subject = nil
}

// SetContent sets the "content" field.
func (m *EmailLogMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *EmailLogMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the EmailLog entity.
// If the EmailLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailLogMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *EmailLogMutation) ResetContent() {
	m.content = nil
}

// SetSendStatus sets the "send_status" field.
func (m *EmailLogMutation) SetSendStatus(u uint8) {
	m.send_status = &u
	m.addsend_status = nil
}

// SendStatus returns the value of the "send_status" field in the mutation.
func (m *EmailLogMutation) SendStatus() (r uint8, exists bool) {
	v := m.send_status
	if v == nil {
		return
	}
	return *v, true
}

// OldSendStatus returns the old "send_status" field's value of the EmailLog entity.
// If the EmailLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailLogMutation) OldSendStatus(ctx context.Context) (v uint8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSendStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSendStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSendStatus: %w", err)
	}
	return oldValue.SendStatus, nil
}

// AddSendStatus adds u to the "send_status" field.
func (m *EmailLogMutation) AddSendStatus(u int8) {
	if m.addsend_status != nil {
		*m.addsend_status += u
	} else {
		m.addsend_status = &u
	}
}

// AddedSendStatus returns the value that was added to the "send_status" field in this mutation.
func (m *EmailLogMutation) AddedSendStatus() (r int8, exists bool) {
	v := m.addsend_status
	if v == nil {
		return
	}
	return *v, true
}

// ResetSendStatus resets all changes to the "send_status" field.
func (m *EmailLogMutation) ResetSendStatus() {
	m.send_status = nil
	m.addsend_status = nil
}

// Where appends a list predicates to the EmailLogMutation builder.
func (m *EmailLogMutation) Where(ps ...predicate.EmailLog) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the EmailLogMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *EmailLogMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.EmailLog, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *EmailLogMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *EmailLogMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (EmailLog).
func (m *EmailLogMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EmailLogMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.created_at != nil {
		fields = append(fields, emaillog.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, emaillog.FieldUpdatedAt)
	}
	if m.target != nil {
		fields = append(fields, emaillog.FieldTarget)
	}
	if m.subject != nil {
		fields = append(fields, emaillog.FieldSubject)
	}
	if m.content != nil {
		fields = append(fields, emaillog.FieldContent)
	}
	if m.send_status != nil {
		fields = append(fields, emaillog.FieldSendStatus)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EmailLogMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case emaillog.FieldCreatedAt:
		return m.CreatedAt()
	case emaillog.FieldUpdatedAt:
		return m.UpdatedAt()
	case emaillog.FieldTarget:
		return m.Target()
	case emaillog.FieldSubject:
		return m.Subject()
	case emaillog.FieldContent:
		return m.Content()
	case emaillog.FieldSendStatus:
		return m.SendStatus()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EmailLogMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case emaillog.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case emaillog.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case emaillog.FieldTarget:
		return m.OldTarget(ctx)
	case emaillog.FieldSubject:
		return m.OldSubject(ctx)
	case emaillog.FieldContent:
		return m.OldContent(ctx)
	case emaillog.FieldSendStatus:
		return m.OldSendStatus(ctx)
	}
	return nil, fmt.Errorf("unknown EmailLog field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmailLogMutation) SetField(name string, value ent.Value) error {
	switch name {
	case emaillog.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case emaillog.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case emaillog.FieldTarget:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTarget(v)
		return nil
	case emaillog.FieldSubject:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSubject(v)
		return nil
	case emaillog.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case emaillog.FieldSendStatus:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSendStatus(v)
		return nil
	}
	return fmt.Errorf("unknown EmailLog field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EmailLogMutation) AddedFields() []string {
	var fields []string
	if m.addsend_status != nil {
		fields = append(fields, emaillog.FieldSendStatus)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EmailLogMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case emaillog.FieldSendStatus:
		return m.AddedSendStatus()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmailLogMutation) AddField(name string, value ent.Value) error {
	switch name {
	case emaillog.FieldSendStatus:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSendStatus(v)
		return nil
	}
	return fmt.Errorf("unknown EmailLog numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EmailLogMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EmailLogMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EmailLogMutation) ClearField(name string) error {
	return fmt.Errorf("unknown EmailLog nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EmailLogMutation) ResetField(name string) error {
	switch name {
	case emaillog.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case emaillog.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case emaillog.FieldTarget:
		m.ResetTarget()
		return nil
	case emaillog.FieldSubject:
		m.ResetSubject()
		return nil
	case emaillog.FieldContent:
		m.ResetContent()
		return nil
	case emaillog.FieldSendStatus:
		m.ResetSendStatus()
		return nil
	}
	return fmt.Errorf("unknown EmailLog field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EmailLogMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EmailLogMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EmailLogMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EmailLogMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EmailLogMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EmailLogMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EmailLogMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown EmailLog unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EmailLogMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown EmailLog edge %s", name)
}

// SmsLogMutation represents an operation that mutates the SmsLog nodes in the graph.
type SmsLogMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	created_at     *time.Time
	updated_at     *time.Time
	phone_number   *string
	content        *string
	send_status    *uint8
	addsend_status *int8
	provider       *string
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*SmsLog, error)
	predicates     []predicate.SmsLog
}

var _ ent.Mutation = (*SmsLogMutation)(nil)

// smslogOption allows management of the mutation configuration using functional options.
type smslogOption func(*SmsLogMutation)

// newSmsLogMutation creates new mutation for the SmsLog entity.
func newSmsLogMutation(c config, op Op, opts ...smslogOption) *SmsLogMutation {
	m := &SmsLogMutation{
		config:        c,
		op:            op,
		typ:           TypeSmsLog,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSmsLogID sets the ID field of the mutation.
func withSmsLogID(id uuid.UUID) smslogOption {
	return func(m *SmsLogMutation) {
		var (
			err   error
			once  sync.Once
			value *SmsLog
		)
		m.oldValue = func(ctx context.Context) (*SmsLog, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().SmsLog.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSmsLog sets the old SmsLog of the mutation.
func withSmsLog(node *SmsLog) smslogOption {
	return func(m *SmsLogMutation) {
		m.oldValue = func(context.Context) (*SmsLog, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SmsLogMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SmsLogMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of SmsLog entities.
func (m *SmsLogMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SmsLogMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SmsLogMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().SmsLog.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *SmsLogMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *SmsLogMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the SmsLog entity.
// If the SmsLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsLogMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *SmsLogMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *SmsLogMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *SmsLogMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the SmsLog entity.
// If the SmsLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsLogMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *SmsLogMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetPhoneNumber sets the "phone_number" field.
func (m *SmsLogMutation) SetPhoneNumber(s string) {
	m.phone_number = &s
}

// PhoneNumber returns the value of the "phone_number" field in the mutation.
func (m *SmsLogMutation) PhoneNumber() (r string, exists bool) {
	v := m.phone_number
	if v == nil {
		return
	}
	return *v, true
}

// OldPhoneNumber returns the old "phone_number" field's value of the SmsLog entity.
// If the SmsLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsLogMutation) OldPhoneNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhoneNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhoneNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhoneNumber: %w", err)
	}
	return oldValue.PhoneNumber, nil
}

// ResetPhoneNumber resets all changes to the "phone_number" field.
func (m *SmsLogMutation) ResetPhoneNumber() {
	m.phone_number = nil
}

// SetContent sets the "content" field.
func (m *SmsLogMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *SmsLogMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the SmsLog entity.
// If the SmsLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsLogMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *SmsLogMutation) ResetContent() {
	m.content = nil
}

// SetSendStatus sets the "send_status" field.
func (m *SmsLogMutation) SetSendStatus(u uint8) {
	m.send_status = &u
	m.addsend_status = nil
}

// SendStatus returns the value of the "send_status" field in the mutation.
func (m *SmsLogMutation) SendStatus() (r uint8, exists bool) {
	v := m.send_status
	if v == nil {
		return
	}
	return *v, true
}

// OldSendStatus returns the old "send_status" field's value of the SmsLog entity.
// If the SmsLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsLogMutation) OldSendStatus(ctx context.Context) (v uint8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSendStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSendStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSendStatus: %w", err)
	}
	return oldValue.SendStatus, nil
}

// AddSendStatus adds u to the "send_status" field.
func (m *SmsLogMutation) AddSendStatus(u int8) {
	if m.addsend_status != nil {
		*m.addsend_status += u
	} else {
		m.addsend_status = &u
	}
}

// AddedSendStatus returns the value that was added to the "send_status" field in this mutation.
func (m *SmsLogMutation) AddedSendStatus() (r int8, exists bool) {
	v := m.addsend_status
	if v == nil {
		return
	}
	return *v, true
}

// ResetSendStatus resets all changes to the "send_status" field.
func (m *SmsLogMutation) ResetSendStatus() {
	m.send_status = nil
	m.addsend_status = nil
}

// SetProvider sets the "provider" field.
func (m *SmsLogMutation) SetProvider(s string) {
	m.provider = &s
}

// Provider returns the value of the "provider" field in the mutation.
func (m *SmsLogMutation) Provider() (r string, exists bool) {
	v := m.provider
	if v == nil {
		return
	}
	return *v, true
}

// OldProvider returns the old "provider" field's value of the SmsLog entity.
// If the SmsLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SmsLogMutation) OldProvider(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldProvider is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldProvider requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProvider: %w", err)
	}
	return oldValue.Provider, nil
}

// ResetProvider resets all changes to the "provider" field.
func (m *SmsLogMutation) ResetProvider() {
	m.provider = nil
}

// Where appends a list predicates to the SmsLogMutation builder.
func (m *SmsLogMutation) Where(ps ...predicate.SmsLog) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the SmsLogMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *SmsLogMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.SmsLog, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *SmsLogMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *SmsLogMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (SmsLog).
func (m *SmsLogMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SmsLogMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.created_at != nil {
		fields = append(fields, smslog.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, smslog.FieldUpdatedAt)
	}
	if m.phone_number != nil {
		fields = append(fields, smslog.FieldPhoneNumber)
	}
	if m.content != nil {
		fields = append(fields, smslog.FieldContent)
	}
	if m.send_status != nil {
		fields = append(fields, smslog.FieldSendStatus)
	}
	if m.provider != nil {
		fields = append(fields, smslog.FieldProvider)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SmsLogMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case smslog.FieldCreatedAt:
		return m.CreatedAt()
	case smslog.FieldUpdatedAt:
		return m.UpdatedAt()
	case smslog.FieldPhoneNumber:
		return m.PhoneNumber()
	case smslog.FieldContent:
		return m.Content()
	case smslog.FieldSendStatus:
		return m.SendStatus()
	case smslog.FieldProvider:
		return m.Provider()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SmsLogMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case smslog.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case smslog.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case smslog.FieldPhoneNumber:
		return m.OldPhoneNumber(ctx)
	case smslog.FieldContent:
		return m.OldContent(ctx)
	case smslog.FieldSendStatus:
		return m.OldSendStatus(ctx)
	case smslog.FieldProvider:
		return m.OldProvider(ctx)
	}
	return nil, fmt.Errorf("unknown SmsLog field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SmsLogMutation) SetField(name string, value ent.Value) error {
	switch name {
	case smslog.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case smslog.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case smslog.FieldPhoneNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhoneNumber(v)
		return nil
	case smslog.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case smslog.FieldSendStatus:
		v, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSendStatus(v)
		return nil
	case smslog.FieldProvider:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProvider(v)
		return nil
	}
	return fmt.Errorf("unknown SmsLog field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SmsLogMutation) AddedFields() []string {
	var fields []string
	if m.addsend_status != nil {
		fields = append(fields, smslog.FieldSendStatus)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SmsLogMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case smslog.FieldSendStatus:
		return m.AddedSendStatus()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SmsLogMutation) AddField(name string, value ent.Value) error {
	switch name {
	case smslog.FieldSendStatus:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSendStatus(v)
		return nil
	}
	return fmt.Errorf("unknown SmsLog numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SmsLogMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SmsLogMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SmsLogMutation) ClearField(name string) error {
	return fmt.Errorf("unknown SmsLog nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SmsLogMutation) ResetField(name string) error {
	switch name {
	case smslog.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case smslog.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case smslog.FieldPhoneNumber:
		m.ResetPhoneNumber()
		return nil
	case smslog.FieldContent:
		m.ResetContent()
		return nil
	case smslog.FieldSendStatus:
		m.ResetSendStatus()
		return nil
	case smslog.FieldProvider:
		m.ResetProvider()
		return nil
	}
	return fmt.Errorf("unknown SmsLog field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SmsLogMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SmsLogMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SmsLogMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SmsLogMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SmsLogMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SmsLogMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SmsLogMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown SmsLog unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SmsLogMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown SmsLog edge %s", name)
}
