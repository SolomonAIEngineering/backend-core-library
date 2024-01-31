package message_definitionv1 // import "github.com/SolomonAIEngineering/backend-core-library/message_queue/generated/message-definition/message_definition/v1"

import (
	context "context"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type DeleteAccountMessageFormatORM struct {
	AuthZeroId  string
	Email       string
	ProfileType string
	UserId      uint64
}

// TableName overrides the default tablename generated by GORM
func (DeleteAccountMessageFormatORM) TableName() string {
	return "delete_account_message_formats"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *DeleteAccountMessageFormat) ToORM(ctx context.Context) (DeleteAccountMessageFormatORM, error) {
	to := DeleteAccountMessageFormatORM{}
	var err error
	if prehook, ok := interface{}(m).(DeleteAccountMessageFormatWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.AuthZeroId = m.AuthZeroId
	to.Email = m.Email
	to.UserId = m.UserId
	to.ProfileType = DeleteAccountMessageFormat_ProfileType_name[int32(m.ProfileType)]
	if posthook, ok := interface{}(m).(DeleteAccountMessageFormatWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *DeleteAccountMessageFormatORM) ToPB(ctx context.Context) (DeleteAccountMessageFormat, error) {
	to := DeleteAccountMessageFormat{}
	var err error
	if prehook, ok := interface{}(m).(DeleteAccountMessageFormatWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.AuthZeroId = m.AuthZeroId
	to.Email = m.Email
	to.UserId = m.UserId
	to.ProfileType = DeleteAccountMessageFormat_ProfileType(DeleteAccountMessageFormat_ProfileType_value[m.ProfileType])
	if posthook, ok := interface{}(m).(DeleteAccountMessageFormatWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type DeleteAccountMessageFormat the arg will be the target, the caller the one being converted from

// DeleteAccountMessageFormatBeforeToORM called before default ToORM code
type DeleteAccountMessageFormatWithBeforeToORM interface {
	BeforeToORM(context.Context, *DeleteAccountMessageFormatORM) error
}

// DeleteAccountMessageFormatAfterToORM called after default ToORM code
type DeleteAccountMessageFormatWithAfterToORM interface {
	AfterToORM(context.Context, *DeleteAccountMessageFormatORM) error
}

// DeleteAccountMessageFormatBeforeToPB called before default ToPB code
type DeleteAccountMessageFormatWithBeforeToPB interface {
	BeforeToPB(context.Context, *DeleteAccountMessageFormat) error
}

// DeleteAccountMessageFormatAfterToPB called after default ToPB code
type DeleteAccountMessageFormatWithAfterToPB interface {
	AfterToPB(context.Context, *DeleteAccountMessageFormat) error
}

// DefaultCreateDeleteAccountMessageFormat executes a basic gorm create call
func DefaultCreateDeleteAccountMessageFormat(ctx context.Context, in *DeleteAccountMessageFormat, db *gorm.DB) (*DeleteAccountMessageFormat, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(DeleteAccountMessageFormatORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(DeleteAccountMessageFormatORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type DeleteAccountMessageFormatORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type DeleteAccountMessageFormatORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

// DefaultApplyFieldMaskDeleteAccountMessageFormat patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskDeleteAccountMessageFormat(ctx context.Context, patchee *DeleteAccountMessageFormat, patcher *DeleteAccountMessageFormat, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*DeleteAccountMessageFormat, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"AuthZeroId" {
			patchee.AuthZeroId = patcher.AuthZeroId
			continue
		}
		if f == prefix+"Email" {
			patchee.Email = patcher.Email
			continue
		}
		if f == prefix+"UserId" {
			patchee.UserId = patcher.UserId
			continue
		}
		if f == prefix+"ProfileType" {
			patchee.ProfileType = patcher.ProfileType
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListDeleteAccountMessageFormat executes a gorm list call
func DefaultListDeleteAccountMessageFormat(ctx context.Context, db *gorm.DB) ([]*DeleteAccountMessageFormat, error) {
	in := DeleteAccountMessageFormat{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(DeleteAccountMessageFormatORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &DeleteAccountMessageFormatORM{}, &DeleteAccountMessageFormat{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(DeleteAccountMessageFormatORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	ormResponse := []DeleteAccountMessageFormatORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(DeleteAccountMessageFormatORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*DeleteAccountMessageFormat{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type DeleteAccountMessageFormatORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type DeleteAccountMessageFormatORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type DeleteAccountMessageFormatORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]DeleteAccountMessageFormatORM) error
}
