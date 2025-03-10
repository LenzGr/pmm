// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agentpb/agent.proto

package agentpb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/rpc/status"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"

	_ "github.com/percona/pmm/api/inventorypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

var (
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *TextFiles) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

func (this *Ping) Validate() error {
	return nil
}

func (this *Pong) Validate() error {
	if this.CurrentTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CurrentTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CurrentTime", err)
		}
	}
	return nil
}

func (this *QANCollectRequest) Validate() error {
	for _, item := range this.MetricsBucket {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MetricsBucket", err)
			}
		}
	}
	return nil
}

func (this *QANCollectResponse) Validate() error {
	return nil
}

func (this *StateChangedRequest) Validate() error {
	return nil
}

func (this *StateChangedResponse) Validate() error {
	return nil
}

func (this *SetStateRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

func (this *SetStateRequest_AgentProcess) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

func (this *SetStateRequest_BuiltinAgent) Validate() error {
	if this.TextFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TextFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TextFiles", err)
		}
	}
	return nil
}

func (this *SetStateResponse) Validate() error {
	return nil
}

func (this *QueryActionValue) Validate() error {
	if oneOfNester, ok := this.GetKind().(*QueryActionValue_Timestamp); ok {
		if oneOfNester.Timestamp != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Timestamp); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
			}
		}
	}
	if oneOfNester, ok := this.GetKind().(*QueryActionValue_Slice); ok {
		if oneOfNester.Slice != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Slice); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Slice", err)
			}
		}
	}
	if oneOfNester, ok := this.GetKind().(*QueryActionValue_Map); ok {
		if oneOfNester.Map != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Map); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Map", err)
			}
		}
	}
	if oneOfNester, ok := this.GetKind().(*QueryActionValue_Binary); ok {
		if oneOfNester.Binary != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Binary); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Binary", err)
			}
		}
	}
	return nil
}

func (this *QueryActionSlice) Validate() error {
	for _, item := range this.Slice {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Slice", err)
			}
		}
	}
	return nil
}

func (this *QueryActionMap) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

func (this *QueryActionBinary) Validate() error {
	return nil
}

func (this *QueryActionResult) Validate() error {
	for _, item := range this.Rows {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rows", err)
			}
		}
	}
	for _, item := range this.Docs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Docs", err)
			}
		}
	}
	return nil
}

func (this *StartActionRequest) Validate() error {
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MysqlExplainParams); ok {
		if oneOfNester.MysqlExplainParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlExplainParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlExplainParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MysqlShowCreateTableParams); ok {
		if oneOfNester.MysqlShowCreateTableParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlShowCreateTableParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlShowCreateTableParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MysqlShowTableStatusParams); ok {
		if oneOfNester.MysqlShowTableStatusParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlShowTableStatusParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlShowTableStatusParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MysqlShowIndexParams); ok {
		if oneOfNester.MysqlShowIndexParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlShowIndexParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlShowIndexParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_PostgresqlShowCreateTableParams); ok {
		if oneOfNester.PostgresqlShowCreateTableParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PostgresqlShowCreateTableParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PostgresqlShowCreateTableParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_PostgresqlShowIndexParams); ok {
		if oneOfNester.PostgresqlShowIndexParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PostgresqlShowIndexParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PostgresqlShowIndexParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MongodbExplainParams); ok {
		if oneOfNester.MongodbExplainParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MongodbExplainParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MongodbExplainParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_PtSummaryParams); ok {
		if oneOfNester.PtSummaryParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PtSummaryParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PtSummaryParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_PtPgSummaryParams); ok {
		if oneOfNester.PtPgSummaryParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PtPgSummaryParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PtPgSummaryParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_PtMongodbSummaryParams); ok {
		if oneOfNester.PtMongodbSummaryParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PtMongodbSummaryParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PtMongodbSummaryParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_PtMysqlSummaryParams); ok {
		if oneOfNester.PtMysqlSummaryParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PtMysqlSummaryParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PtMysqlSummaryParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MysqlQueryShowParams); ok {
		if oneOfNester.MysqlQueryShowParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlQueryShowParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlQueryShowParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MysqlQuerySelectParams); ok {
		if oneOfNester.MysqlQuerySelectParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlQuerySelectParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlQuerySelectParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_PostgresqlQueryShowParams); ok {
		if oneOfNester.PostgresqlQueryShowParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PostgresqlQueryShowParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PostgresqlQueryShowParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_PostgresqlQuerySelectParams); ok {
		if oneOfNester.PostgresqlQuerySelectParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PostgresqlQuerySelectParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PostgresqlQuerySelectParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MongodbQueryGetparameterParams); ok {
		if oneOfNester.MongodbQueryGetparameterParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MongodbQueryGetparameterParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MongodbQueryGetparameterParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MongodbQueryBuildinfoParams); ok {
		if oneOfNester.MongodbQueryBuildinfoParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MongodbQueryBuildinfoParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MongodbQueryBuildinfoParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MongodbQueryGetcmdlineoptsParams); ok {
		if oneOfNester.MongodbQueryGetcmdlineoptsParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MongodbQueryGetcmdlineoptsParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MongodbQueryGetcmdlineoptsParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MongodbQueryReplsetgetstatusParams); ok {
		if oneOfNester.MongodbQueryReplsetgetstatusParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MongodbQueryReplsetgetstatusParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MongodbQueryReplsetgetstatusParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetParams().(*StartActionRequest_MongodbQueryGetdiagnosticdataParams); ok {
		if oneOfNester.MongodbQueryGetdiagnosticdataParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MongodbQueryGetdiagnosticdataParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MongodbQueryGetdiagnosticdataParams", err)
			}
		}
	}
	if this.Timeout != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timeout); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timeout", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MySQLExplainParams) Validate() error {
	if this.TlsFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MySQLShowCreateTableParams) Validate() error {
	if this.TlsFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MySQLShowTableStatusParams) Validate() error {
	if this.TlsFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MySQLShowIndexParams) Validate() error {
	if this.TlsFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_PostgreSQLShowCreateTableParams) Validate() error {
	if this.TlsFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_PostgreSQLShowIndexParams) Validate() error {
	if this.TlsFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MongoDBExplainParams) Validate() error {
	if this.TextFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TextFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TextFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_PTSummaryParams) Validate() error {
	return nil
}

func (this *StartActionRequest_PTPgSummaryParams) Validate() error {
	return nil
}

func (this *StartActionRequest_PTMongoDBSummaryParams) Validate() error {
	return nil
}

func (this *StartActionRequest_PTMySQLSummaryParams) Validate() error {
	return nil
}

func (this *StartActionRequest_MySQLQueryShowParams) Validate() error {
	if this.TlsFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MySQLQuerySelectParams) Validate() error {
	if this.TlsFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_PostgreSQLQueryShowParams) Validate() error {
	if this.TlsFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_PostgreSQLQuerySelectParams) Validate() error {
	if this.TlsFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MongoDBQueryGetParameterParams) Validate() error {
	if this.TextFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TextFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TextFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MongoDBQueryBuildInfoParams) Validate() error {
	if this.TextFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TextFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TextFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MongoDBQueryGetCmdLineOptsParams) Validate() error {
	if this.TextFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TextFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TextFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MongoDBQueryReplSetGetStatusParams) Validate() error {
	if this.TextFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TextFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TextFiles", err)
		}
	}
	return nil
}

func (this *StartActionRequest_MongoDBQueryGetDiagnosticDataParams) Validate() error {
	if this.TextFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TextFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TextFiles", err)
		}
	}
	return nil
}

func (this *StartActionResponse) Validate() error {
	return nil
}

func (this *StopActionRequest) Validate() error {
	return nil
}

func (this *StopActionResponse) Validate() error {
	return nil
}

func (this *ActionResultRequest) Validate() error {
	return nil
}

func (this *ActionResultResponse) Validate() error {
	return nil
}

func (this *PBMSwitchPITRRequest) Validate() error {
	if this.TextFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TextFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TextFiles", err)
		}
	}
	return nil
}

func (this *PBMSwitchPITRResponse) Validate() error {
	return nil
}

func (this *CheckConnectionRequest) Validate() error {
	if this.Timeout != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timeout); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timeout", err)
		}
	}
	if this.TextFiles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TextFiles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TextFiles", err)
		}
	}
	return nil
}

func (this *CheckConnectionResponse) Validate() error {
	if this.Stats != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Stats); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Stats", err)
		}
	}
	return nil
}

func (this *CheckConnectionResponse_Stats) Validate() error {
	return nil
}

func (this *JobStatusRequest) Validate() error {
	return nil
}

func (this *JobStatusResponse) Validate() error {
	return nil
}

func (this *S3LocationConfig) Validate() error {
	return nil
}

func (this *StartJobRequest) Validate() error {
	if this.Timeout != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timeout); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timeout", err)
		}
	}
	if oneOfNester, ok := this.GetJob().(*StartJobRequest_MysqlBackup); ok {
		if oneOfNester.MysqlBackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlBackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlBackup", err)
			}
		}
	}
	if oneOfNester, ok := this.GetJob().(*StartJobRequest_MysqlRestoreBackup); ok {
		if oneOfNester.MysqlRestoreBackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlRestoreBackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlRestoreBackup", err)
			}
		}
	}
	if oneOfNester, ok := this.GetJob().(*StartJobRequest_MongodbBackup); ok {
		if oneOfNester.MongodbBackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MongodbBackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MongodbBackup", err)
			}
		}
	}
	if oneOfNester, ok := this.GetJob().(*StartJobRequest_MongodbRestoreBackup); ok {
		if oneOfNester.MongodbRestoreBackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MongodbRestoreBackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MongodbRestoreBackup", err)
			}
		}
	}
	return nil
}

func (this *StartJobRequest_MySQLBackup) Validate() error {
	if oneOfNester, ok := this.GetLocationConfig().(*StartJobRequest_MySQLBackup_S3Config); ok {
		if oneOfNester.S3Config != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.S3Config); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("S3Config", err)
			}
		}
	}
	return nil
}

func (this *StartJobRequest_MySQLRestoreBackup) Validate() error {
	if oneOfNester, ok := this.GetLocationConfig().(*StartJobRequest_MySQLRestoreBackup_S3Config); ok {
		if oneOfNester.S3Config != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.S3Config); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("S3Config", err)
			}
		}
	}
	return nil
}

func (this *StartJobRequest_MongoDBBackup) Validate() error {
	if oneOfNester, ok := this.GetLocationConfig().(*StartJobRequest_MongoDBBackup_S3Config); ok {
		if oneOfNester.S3Config != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.S3Config); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("S3Config", err)
			}
		}
	}
	return nil
}

func (this *StartJobRequest_MongoDBRestoreBackup) Validate() error {
	if this.PitrTimestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PitrTimestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PitrTimestamp", err)
		}
	}
	if oneOfNester, ok := this.GetLocationConfig().(*StartJobRequest_MongoDBRestoreBackup_S3Config); ok {
		if oneOfNester.S3Config != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.S3Config); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("S3Config", err)
			}
		}
	}
	return nil
}

func (this *StartJobResponse) Validate() error {
	return nil
}

func (this *StopJobRequest) Validate() error {
	return nil
}

func (this *StopJobResponse) Validate() error {
	return nil
}

func (this *JobResult) Validate() error {
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	if oneOfNester, ok := this.GetResult().(*JobResult_Error_); ok {
		if oneOfNester.Error != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Error); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResult().(*JobResult_MysqlBackup); ok {
		if oneOfNester.MysqlBackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlBackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlBackup", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResult().(*JobResult_MysqlRestoreBackup); ok {
		if oneOfNester.MysqlRestoreBackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlRestoreBackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlRestoreBackup", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResult().(*JobResult_MongodbBackup); ok {
		if oneOfNester.MongodbBackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MongodbBackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MongodbBackup", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResult().(*JobResult_MongodbRestoreBackup); ok {
		if oneOfNester.MongodbRestoreBackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MongodbRestoreBackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MongodbRestoreBackup", err)
			}
		}
	}
	return nil
}

func (this *JobResult_Error) Validate() error {
	return nil
}

func (this *JobResult_MongoDBBackup) Validate() error {
	return nil
}

func (this *JobResult_MySQLBackup) Validate() error {
	return nil
}

func (this *JobResult_MySQLRestoreBackup) Validate() error {
	return nil
}

func (this *JobResult_MongoDBRestoreBackup) Validate() error {
	return nil
}

func (this *JobProgress) Validate() error {
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	if oneOfNester, ok := this.GetResult().(*JobProgress_MysqlBackup); ok {
		if oneOfNester.MysqlBackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlBackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlBackup", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResult().(*JobProgress_MysqlRestoreBackup); ok {
		if oneOfNester.MysqlRestoreBackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MysqlRestoreBackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MysqlRestoreBackup", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResult().(*JobProgress_Logs_); ok {
		if oneOfNester.Logs != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Logs); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Logs", err)
			}
		}
	}
	return nil
}

func (this *JobProgress_MySQLBackup) Validate() error {
	return nil
}

func (this *JobProgress_MySQLRestoreBackup) Validate() error {
	return nil
}

func (this *JobProgress_Logs) Validate() error {
	return nil
}

func (this *GetVersionsRequest) Validate() error {
	for _, item := range this.Softwares {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Softwares", err)
			}
		}
	}
	return nil
}

func (this *GetVersionsRequest_MySQLd) Validate() error {
	return nil
}

func (this *GetVersionsRequest_Xtrabackup) Validate() error {
	return nil
}

func (this *GetVersionsRequest_Xbcloud) Validate() error {
	return nil
}

func (this *GetVersionsRequest_Qpress) Validate() error {
	return nil
}

func (this *GetVersionsRequest_Software) Validate() error {
	if oneOfNester, ok := this.GetSoftware().(*GetVersionsRequest_Software_Mysqld); ok {
		if oneOfNester.Mysqld != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Mysqld); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mysqld", err)
			}
		}
	}
	if oneOfNester, ok := this.GetSoftware().(*GetVersionsRequest_Software_Xtrabackup); ok {
		if oneOfNester.Xtrabackup != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Xtrabackup); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Xtrabackup", err)
			}
		}
	}
	if oneOfNester, ok := this.GetSoftware().(*GetVersionsRequest_Software_Xbcloud); ok {
		if oneOfNester.Xbcloud != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Xbcloud); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Xbcloud", err)
			}
		}
	}
	if oneOfNester, ok := this.GetSoftware().(*GetVersionsRequest_Software_Qpress); ok {
		if oneOfNester.Qpress != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Qpress); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Qpress", err)
			}
		}
	}
	return nil
}

func (this *GetVersionsResponse) Validate() error {
	for _, item := range this.Versions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Versions", err)
			}
		}
	}
	return nil
}

func (this *GetVersionsResponse_Version) Validate() error {
	return nil
}

func (this *AgentMessage) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_Ping); ok {
		if oneOfNester.Ping != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ping); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ping", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_StateChanged); ok {
		if oneOfNester.StateChanged != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StateChanged); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StateChanged", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_QanCollect); ok {
		if oneOfNester.QanCollect != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.QanCollect); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("QanCollect", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_ActionResult); ok {
		if oneOfNester.ActionResult != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ActionResult); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ActionResult", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_Pong); ok {
		if oneOfNester.Pong != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Pong); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Pong", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_SetState); ok {
		if oneOfNester.SetState != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SetState); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SetState", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_StartAction); ok {
		if oneOfNester.StartAction != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StartAction); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StartAction", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_StopAction); ok {
		if oneOfNester.StopAction != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StopAction); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StopAction", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_CheckConnection); ok {
		if oneOfNester.CheckConnection != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.CheckConnection); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CheckConnection", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_StartJob); ok {
		if oneOfNester.StartJob != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StartJob); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StartJob", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_StopJob); ok {
		if oneOfNester.StopJob != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StopJob); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StopJob", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_JobStatus); ok {
		if oneOfNester.JobStatus != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.JobStatus); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("JobStatus", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_JobResult); ok {
		if oneOfNester.JobResult != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.JobResult); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("JobResult", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_JobProgress); ok {
		if oneOfNester.JobProgress != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.JobProgress); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("JobProgress", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_GetVersions); ok {
		if oneOfNester.GetVersions != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.GetVersions); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("GetVersions", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_PbmSwitchPitr); ok {
		if oneOfNester.PbmSwitchPitr != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PbmSwitchPitr); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PbmSwitchPitr", err)
			}
		}
	}
	return nil
}

func (this *ServerMessage) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_Pong); ok {
		if oneOfNester.Pong != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Pong); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Pong", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_StateChanged); ok {
		if oneOfNester.StateChanged != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StateChanged); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StateChanged", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_QanCollect); ok {
		if oneOfNester.QanCollect != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.QanCollect); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("QanCollect", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_ActionResult); ok {
		if oneOfNester.ActionResult != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ActionResult); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ActionResult", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_Ping); ok {
		if oneOfNester.Ping != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ping); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ping", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_SetState); ok {
		if oneOfNester.SetState != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SetState); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SetState", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_StartAction); ok {
		if oneOfNester.StartAction != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StartAction); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StartAction", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_StopAction); ok {
		if oneOfNester.StopAction != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StopAction); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StopAction", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_CheckConnection); ok {
		if oneOfNester.CheckConnection != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.CheckConnection); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CheckConnection", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_StartJob); ok {
		if oneOfNester.StartJob != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StartJob); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StartJob", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_StopJob); ok {
		if oneOfNester.StopJob != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StopJob); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StopJob", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_JobStatus); ok {
		if oneOfNester.JobStatus != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.JobStatus); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("JobStatus", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_GetVersions); ok {
		if oneOfNester.GetVersions != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.GetVersions); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("GetVersions", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_PbmSwitchPitr); ok {
		if oneOfNester.PbmSwitchPitr != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PbmSwitchPitr); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PbmSwitchPitr", err)
			}
		}
	}
	return nil
}
