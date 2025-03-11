package utils

import "time"

const (
	HOST_HTTP = "http://localhost:8080"
	HOST_GRPC = ":8002"

	TIME_OUTE = time.Duration(60 * time.Second)
)

const (
	ErrCreatedGrpcClient     = "Failed to create gRPC client"
	ErrExpectGRPCStatusError = "Expected gRPC status error"

	ErrExpectInvalidArgument             = "Expected error code 3 \"InvalidArgument\""
	ErrMessageExpectInvalidArgument      = "Expected error message \"phone not match pattern 8xxxxxxxxxx\""
	ErrMessageExpectInvalidArgumentDebit = "Expected error message \"unavailable operation\""
	ErrExpectUnauthenticated             = "Expected error code 16 \"Unauthenticated\""
	ErrMessageExpectUnauthenticated      = "Expected error message \"not authorized\""
	ErrExpectAlreadyExists               = "Expected error code 6 \"AlreadyExists\""
	ErrMessageExpectAlreadyExists        = "Expected error message \"User allready has 3 accounts\"" // В слове allready одна лишняя l
	ErrMessageExpectAlreadyFullLevel     = "Expected error message \"user allready has FULL level\""
	ErrExpectNotFound                    = "Expected error code 5 \"NotFound\""
	ErrMessageExpectNotFound             = "Expected error message \"account not found\""

	ErrCreatedUserErrorShouldBeNil     = "Error should be nil when creating a user"
	ErrCreatedUserMustNotBeNil         = "Created user must not be nil"
	ErrCreatedUserLevelShouldBeAnon    = "Created user level must be ANON"
	ErrCreatedUserShoulBeNil           = "Created user should be nil"
	ErrCreatedUserErrorMustNotBeNil    = "Error must be not nil when creating a user"
	ErrLoginUserErrorShouldBeNil       = "Error should be nil when login user"
	ErrLoginUserMustNotBeNil           = "Login user must not be nil"
	ErrAuthTokenMustNotBeNil           = "Authorization token must not be nil"
	ErrCreatedAccountErrorShouldBeNil  = "Error should be nil when creating an account"
	ErrCreatedAccountMustNotBeNil      = "Created account must not be nil"
	ErrAccountAmountMustBeEqual        = "Amount on the account must be equal"
	ErrCreatedAccountErrorMustNotBeNil = "Error must be not nil when creating an account"
	ErrCreatedAccountShouldBeNil       = "Created account should be nil"
	ErrUpgradedUserErrorShouldBeNil    = "Error should be nil when upgrade a user"
	ErrGetUserErrorShouldBeNil         = "Error should be nil when getting a user"
	ErrGetUserMustNotBeNil             = "Getting user must not be nil"
	ErrGetUserShouldBeNil              = "Getting user should be nil"
	ErrGetUserErrorMustNotBeNil        = "Error must be not nil when getting user"
	ErrUserIdMustBeEqual               = "User ID must be equal"
	ErrUpgradedUserLevelShouldBeFull   = "Upgraded user level must be FULL"
	ErrUpgradeUserErrorMustNotBeNil    = "Error must be not nil when upgrade a user"
	ErrGetBalanceErrorShouldBeNil      = "Error should be nil when getting an account balance"
	ErrGetBalanceMustNotBeNil          = "Get account balance must not be nil"
	ErrBalanceAmountMustBeEqual        = "Amount on the balance must be equal"
	ErrGetBalanceShouldBeNil           = "Getting account balance should be nil"
	ErrGetBalanceErrorMustNotBeNil     = "Error must be not nil when getting an account balance"
	ErrDebitErrorShouldBeNil           = "Error should be nil when debit to account balance"
	ErrDebitErrorMustNotBeNil          = "Error must be not nil when debit to account"
	ErrCreditErrorShouldBeNil          = "Error should be nil when credit from account balance"
	ErrCreditErrorMustNotBeNil         = "Error must be not nil when credit from account"
)
