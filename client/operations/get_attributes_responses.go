// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// GetAttributesReader is a Reader for the GetAttributes structure.
type GetAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAttributesOK creates a GetAttributesOK with default headers values
func NewGetAttributesOK() *GetAttributesOK {
	return &GetAttributesOK{}
}

/*GetAttributesOK handles this case with default header values.

Attributes listed
*/
type GetAttributesOK struct {
	Payload *models.GetAttributes
}

func (o *GetAttributesOK) Error() string {
	return fmt.Sprintf("[GET /contacts/attributes][%d] getAttributesOK  %+v", 200, o.Payload)
}

func (o *GetAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAttributes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
