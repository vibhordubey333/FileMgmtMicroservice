package service

import (
	"FileMgmtMicroservice/restapi/operations/file_mgmt"
	"FileMgmtMicroservice/service"
	"reflect"
	"testing"

	"github.com/go-openapi/runtime/middleware"
)

func TestListFile(t *testing.T) {
	type args struct {
		params file_mgmt.ListFilesParams
	}
	tests := []struct {
		name string
		args args
		want middleware.Responder
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.ListFile(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
