package errors_test

import (
	errors2 "errors"
	"github.com/bushaHQ/httputil/errors"
	"reflect"
	"testing"
)

type logger struct {
}

func (l logger) Error(args ...interface{}) {
	//a:=
}

func TestCoverErr(t *testing.T) {
	type args struct {
		err       error
		defaultTo error
		log       logger
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			"Should use the original value",
			args{errors.New("original"), errors2.New("defaultTo"), logger{}},
			"original",
		},
		{
			"Should use the default to value",
			args{errors2.New("original"), errors2.New("defaultTo"), logger{}},
			"defaultTo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := errors.CoverErr(tt.args.err, tt.args.defaultTo, tt.args.log); err.Error() != tt.wantErr {
				t.Errorf("CoverErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestErrResponse_Error(t *testing.T) {
	type fields struct {
		Err            error
		HTTPStatusCode int
		AppCode        int64
		StatusText     string
		ErrorText      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			"return the value of the error",
			fields{
				Err:            errors2.New("err"),
				HTTPStatusCode: 400,
				AppCode:        0,
				StatusText:     "bad request",
				ErrorText:      "bad request",
			},
			"err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := errors.ErrResponse{
				Err:            tt.fields.Err,
				HTTPStatusCode: tt.fields.HTTPStatusCode,
				AppCode:        tt.fields.AppCode,
				StatusText:     tt.fields.StatusText,
				ErrorText:      tt.fields.ErrorText,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		errText    string
		statusCode []int
	}
	tests := []struct {
		name string
		args args
		want *errors.ErrResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.New(tt.args.errText, tt.args.statusCode...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		err        error
		statusCode []int
	}
	tests := []struct {
		name string
		args args
		want *errors.ErrResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Wrap(tt.args.err, tt.args.statusCode...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
