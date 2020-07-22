// Code generated by go-swagger; DO NOT EDIT.

package file_mgmt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewUploadFileParams creates a new UploadFileParams object
// no default values defined in spec.
func NewUploadFileParams() UploadFileParams {

	return UploadFileParams{}
}

// UploadFileParams contains all the bound params for the upload file operation
// typically these are obtained from a http.Request
//
// swagger:parameters uploadFile
type UploadFileParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The file to be uploaded
	  Required: true
	  In: formData
	*/
	InputFile io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUploadFileParams() beforehand.
func (o *UploadFileParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	inputFile, inputFileHeader, err := r.FormFile("inputFile")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "inputFile", err))
	} else if err := o.bindInputFile(inputFile, inputFileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.InputFile = &runtime.File{Data: inputFile, Header: inputFileHeader}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindInputFile binds file parameter InputFile.
//
// The only supported validations on files are MinLength and MaxLength
func (o *UploadFileParams) bindInputFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}
