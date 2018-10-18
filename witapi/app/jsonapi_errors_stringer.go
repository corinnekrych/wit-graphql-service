// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": String functions for JSONAPI Errors - See goasupport/jsonapi_errors_stringer/generator.go
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-wit/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-wit/app
// --version=v1.3.0

package app

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)


// String implements the Stringer interface for JSONAPIErrors
func (mt JSONAPIErrors) String() string {
	if mt.Errors == nil {
		return ""
	}
	res := fmt.Sprintf("%d JSONAPI Error(s):\n", len(mt.Errors))
	for i, e := range mt.Errors {
		res += fmt.Sprintf("[ERROR No. %3d]: %s\n", i, e)
	}
	return res
}

// String implements the Stringer interface for a JSONAPIError
func (ut JSONAPIError) String() string {
	return fmt.Sprintf(`
		Code:    %[1]s
		Detail:  %[2]s
		ID:      %[3]s
		Links:   %[4]s
		Meta:    %[5]s
		Source:  %[6]s
		Status:  %[7]s
		Title:   %[8]s
	`, spew.Sdump(ut.Code),
		spew.Sdump(ut.Detail),
		spew.Sdump(ut.ID),
		spew.Sdump(ut.Links),
		spew.Sdump(ut.Meta),
		spew.Sdump(ut.Source),
		spew.Sdump(ut.Status),
		spew.Sdump(ut.Title))
}
