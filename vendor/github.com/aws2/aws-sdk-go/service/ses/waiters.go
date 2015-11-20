// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ses

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

var waiterIdentityExists *waiter.Config

func (c *SES) WaitUntilIdentityExists(input *GetIdentityVerificationAttributesInput) error {
	if waiterIdentityExists == nil {
		waiterIdentityExists = &waiter.Config{
			Operation:   "GetIdentityVerificationAttributes",
			Delay:       3,
			MaxAttempts: 20,
			Acceptors: []waiter.WaitAcceptor{
				{
					State:    "success",
					Matcher:  "pathAll",
					Argument: "VerificationAttributes.*.VerificationStatus",
					Expected: "Success",
				},
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterIdentityExists,
	}
	return w.Wait()
}
