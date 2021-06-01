package handler

import (
	"FileMgmtMicroservice/restapi/operations/file_mgmt"
	"FileMgmtMicroservice/service"
	"reflect"
	"testing"

	"github.com/go-openapi/runtime/middleware"
)

func TestGetFile(t *testing.T) {
	type args struct {
		params file_mgmt.GetFileParams
	}
	tests := []struct {
		name string
		args args
		want middleware.Responder
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.GetFile(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListFile(t *testing.T) {
	type args struct {
		params file_mgmt.ListFilesParams
	}
	tests := []struct {
		name string
		args args
		want middleware.Responder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.ListFile(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUploadFile(t *testing.T) {
	type args struct {
		params file_mgmt.UploadFileParams
	}
	tests := []struct {
		name string
		args args
		want middleware.Responder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.UploadFile(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
