// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/actions.proto

package managementpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetActionRequest) Validate() error {
	if this.ActionId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ActionId", fmt.Errorf(`value '%v' must not be an empty string`, this.ActionId))
	}
	return nil
}
func (this *GetActionResponse) Validate() error {
	return nil
}
func (this *StartMySQLExplainActionRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.Query == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Query", fmt.Errorf(`value '%v' must not be an empty string`, this.Query))
	}
	return nil
}
func (this *StartMySQLExplainActionResponse) Validate() error {
	return nil
}
func (this *StartMySQLExplainJSONActionRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.Query == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Query", fmt.Errorf(`value '%v' must not be an empty string`, this.Query))
	}
	return nil
}
func (this *StartMySQLExplainJSONActionResponse) Validate() error {
	return nil
}
func (this *StartMySQLExplainTraditionalJSONActionRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.Query == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Query", fmt.Errorf(`value '%v' must not be an empty string`, this.Query))
	}
	return nil
}
func (this *StartMySQLExplainTraditionalJSONActionResponse) Validate() error {
	return nil
}
func (this *StartMySQLShowCreateTableActionRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.TableName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TableName", fmt.Errorf(`value '%v' must not be an empty string`, this.TableName))
	}
	return nil
}
func (this *StartMySQLShowCreateTableActionResponse) Validate() error {
	return nil
}
func (this *StartMySQLShowTableStatusActionRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.TableName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TableName", fmt.Errorf(`value '%v' must not be an empty string`, this.TableName))
	}
	return nil
}
func (this *StartMySQLShowTableStatusActionResponse) Validate() error {
	return nil
}
func (this *StartMySQLShowIndexActionRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.TableName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TableName", fmt.Errorf(`value '%v' must not be an empty string`, this.TableName))
	}
	return nil
}
func (this *StartMySQLShowIndexActionResponse) Validate() error {
	return nil
}
func (this *StartPostgreSQLShowCreateTableActionRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.TableName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TableName", fmt.Errorf(`value '%v' must not be an empty string`, this.TableName))
	}
	return nil
}
func (this *StartPostgreSQLShowCreateTableActionResponse) Validate() error {
	return nil
}
func (this *StartPostgreSQLShowIndexActionRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.TableName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TableName", fmt.Errorf(`value '%v' must not be an empty string`, this.TableName))
	}
	return nil
}
func (this *StartPostgreSQLShowIndexActionResponse) Validate() error {
	return nil
}
func (this *StartMongoDBExplainActionRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.Query == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Query", fmt.Errorf(`value '%v' must not be an empty string`, this.Query))
	}
	return nil
}
func (this *StartMongoDBExplainActionResponse) Validate() error {
	return nil
}
func (this *StartPTSummaryActionRequest) Validate() error {
	return nil
}
func (this *StartPTSummaryActionResponse) Validate() error {
	return nil
}
func (this *StartPTMySQLSummaryActionRequest) Validate() error {
	return nil
}
func (this *StartPTMySQLSummaryActionResponse) Validate() error {
	return nil
}
func (this *StartPTPgSQLSummaryActionRequest) Validate() error {
	return nil
}
func (this *StartPTPgSQLSummaryActionResponse) Validate() error {
	return nil
}

func (this *CancelActionRequest) Validate() error {
	if this.ActionId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ActionId", fmt.Errorf(`value '%v' must not be an empty string`, this.ActionId))
	}
	return nil
}
func (this *CancelActionResponse) Validate() error {
	return nil
}
