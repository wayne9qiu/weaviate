/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package personalized_infos




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/core/models"
)

/*WeavePersonalizedInfosPatchOK Successful response

swagger:response weavePersonalizedInfosPatchOK
*/
type WeavePersonalizedInfosPatchOK struct {

	// In: body
	Payload *models.PersonalizedInfo `json:"body,omitempty"`
}

// NewWeavePersonalizedInfosPatchOK creates WeavePersonalizedInfosPatchOK with default headers values
func NewWeavePersonalizedInfosPatchOK() *WeavePersonalizedInfosPatchOK {
	return &WeavePersonalizedInfosPatchOK{}
}

// WithPayload adds the payload to the weave personalized infos patch o k response
func (o *WeavePersonalizedInfosPatchOK) WithPayload(payload *models.PersonalizedInfo) *WeavePersonalizedInfosPatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weave personalized infos patch o k response
func (o *WeavePersonalizedInfosPatchOK) SetPayload(payload *models.PersonalizedInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeavePersonalizedInfosPatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}