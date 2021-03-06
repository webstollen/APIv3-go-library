// Code generated by go-swagger; DO NOT EDIT.

package smtp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// CreateSMTPTemplateReader is a Reader for the CreateSMTPTemplate structure.
type CreateSMTPTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSMTPTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSMTPTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSMTPTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateSMTPTemplateCreated creates a CreateSMTPTemplateCreated with default headers values
func NewCreateSMTPTemplateCreated() *CreateSMTPTemplateCreated {
	return &CreateSMTPTemplateCreated{}
}

/*CreateSMTPTemplateCreated handles this case with default header values.

SMTP template created
*/
type CreateSMTPTemplateCreated struct {
	Payload *models.CreateModel
}

func (o *CreateSMTPTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /smtp/templates][%d] createSmtpTemplateCreated  %+v", 201, o.Payload)
}

func (o *CreateSMTPTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSMTPTemplateBadRequest creates a CreateSMTPTemplateBadRequest with default headers values
func NewCreateSMTPTemplateBadRequest() *CreateSMTPTemplateBadRequest {
	return &CreateSMTPTemplateBadRequest{}
}

/*CreateSMTPTemplateBadRequest handles this case with default header values.

bad request
*/
type CreateSMTPTemplateBadRequest struct {
	Payload *models.ErrorModel
}

func (o *CreateSMTPTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /smtp/templates][%d] createSmtpTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSMTPTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
