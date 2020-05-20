// Code generated by goa v3.1.2, DO NOT EDIT.
//
// listener HTTP client CLI support package
//
// Command:
// $ goa gen github.com/wondenge/listeners/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	healthc "github.com/wondenge/listeners/gen/http/health/client"
	mpesac "github.com/wondenge/listeners/gen/http/mpesa/client"
	goahttp "goa.design/goa/v3/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `health show
mpesa (account-balance-timeout|account-balance-result|transaction-status-timeout|transaction-status-result|reversal-timeout|reversal-result|b2-c-timeout|b2-c-result|c2-b-validation|c2-b-confirmation)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` health show` + "\n" +
		os.Args[0] + ` mpesa account-balance-timeout --body '{
      "Result": {
         "ConversationId": "236543-276372-2",
         "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
         "ReferenceData": {
            "ReferenceItem": {
               "QueueTimeoutURL": "https://internalsandbox.safaricom.co.ke/mpesa/abresults/v1/submit"
            }
         },
         "ResultCode": 0,
         "ResultDesc": "Initiator information is invalid",
         "ResultParameters": {
            "ResultParameter": [
               {
                  "AccountBalance": "Working Account|KES|46713.00|46713.00|0.00|0.00\u0026Float Account|KES|0.00|0.00|0.00|0.00\u0026Utility Account|KES|49217.00|49217.00|0.00|0.00\u0026Charges Paid Account|KES|-220.00|-220.00|0.00|0.00\u0026Organization Settlement Account|KES|0.00|0.00|0.00|0.00",
                  "BOCompletedTime": "20170728095642"
               },
               {
                  "AccountBalance": "Working Account|KES|46713.00|46713.00|0.00|0.00\u0026Float Account|KES|0.00|0.00|0.00|0.00\u0026Utility Account|KES|49217.00|49217.00|0.00|0.00\u0026Charges Paid Account|KES|-220.00|-220.00|0.00|0.00\u0026Organization Settlement Account|KES|0.00|0.00|0.00|0.00",
                  "BOCompletedTime": "20170728095642"
               }
            ]
         },
         "ResultType": 0,
         "TransactionID": "LHG31AA5TX"
      }
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (endpoint.Endpoint, interface{}, error) {
	var (
		healthFlags = flag.NewFlagSet("health", flag.ContinueOnError)

		healthShowFlags = flag.NewFlagSet("show", flag.ExitOnError)

		mpesaFlags = flag.NewFlagSet("mpesa", flag.ContinueOnError)

		mpesaAccountBalanceTimeoutFlags    = flag.NewFlagSet("account-balance-timeout", flag.ExitOnError)
		mpesaAccountBalanceTimeoutBodyFlag = mpesaAccountBalanceTimeoutFlags.String("body", "REQUIRED", "")

		mpesaAccountBalanceResultFlags    = flag.NewFlagSet("account-balance-result", flag.ExitOnError)
		mpesaAccountBalanceResultBodyFlag = mpesaAccountBalanceResultFlags.String("body", "REQUIRED", "")

		mpesaTransactionStatusTimeoutFlags    = flag.NewFlagSet("transaction-status-timeout", flag.ExitOnError)
		mpesaTransactionStatusTimeoutBodyFlag = mpesaTransactionStatusTimeoutFlags.String("body", "REQUIRED", "")

		mpesaTransactionStatusResultFlags    = flag.NewFlagSet("transaction-status-result", flag.ExitOnError)
		mpesaTransactionStatusResultBodyFlag = mpesaTransactionStatusResultFlags.String("body", "REQUIRED", "")

		mpesaReversalTimeoutFlags    = flag.NewFlagSet("reversal-timeout", flag.ExitOnError)
		mpesaReversalTimeoutBodyFlag = mpesaReversalTimeoutFlags.String("body", "REQUIRED", "")

		mpesaReversalResultFlags    = flag.NewFlagSet("reversal-result", flag.ExitOnError)
		mpesaReversalResultBodyFlag = mpesaReversalResultFlags.String("body", "REQUIRED", "")

		mpesaB2CTimeoutFlags    = flag.NewFlagSet("b2-c-timeout", flag.ExitOnError)
		mpesaB2CTimeoutBodyFlag = mpesaB2CTimeoutFlags.String("body", "REQUIRED", "")

		mpesaB2CResultFlags    = flag.NewFlagSet("b2-c-result", flag.ExitOnError)
		mpesaB2CResultBodyFlag = mpesaB2CResultFlags.String("body", "REQUIRED", "")

		mpesaC2BValidationFlags    = flag.NewFlagSet("c2-b-validation", flag.ExitOnError)
		mpesaC2BValidationBodyFlag = mpesaC2BValidationFlags.String("body", "REQUIRED", "")

		mpesaC2BConfirmationFlags    = flag.NewFlagSet("c2-b-confirmation", flag.ExitOnError)
		mpesaC2BConfirmationBodyFlag = mpesaC2BConfirmationFlags.String("body", "REQUIRED", "")
	)
	healthFlags.Usage = healthUsage
	healthShowFlags.Usage = healthShowUsage

	mpesaFlags.Usage = mpesaUsage
	mpesaAccountBalanceTimeoutFlags.Usage = mpesaAccountBalanceTimeoutUsage
	mpesaAccountBalanceResultFlags.Usage = mpesaAccountBalanceResultUsage
	mpesaTransactionStatusTimeoutFlags.Usage = mpesaTransactionStatusTimeoutUsage
	mpesaTransactionStatusResultFlags.Usage = mpesaTransactionStatusResultUsage
	mpesaReversalTimeoutFlags.Usage = mpesaReversalTimeoutUsage
	mpesaReversalResultFlags.Usage = mpesaReversalResultUsage
	mpesaB2CTimeoutFlags.Usage = mpesaB2CTimeoutUsage
	mpesaB2CResultFlags.Usage = mpesaB2CResultUsage
	mpesaC2BValidationFlags.Usage = mpesaC2BValidationUsage
	mpesaC2BConfirmationFlags.Usage = mpesaC2BConfirmationUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "health":
			svcf = healthFlags
		case "mpesa":
			svcf = mpesaFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "health":
			switch epn {
			case "show":
				epf = healthShowFlags

			}

		case "mpesa":
			switch epn {
			case "account-balance-timeout":
				epf = mpesaAccountBalanceTimeoutFlags

			case "account-balance-result":
				epf = mpesaAccountBalanceResultFlags

			case "transaction-status-timeout":
				epf = mpesaTransactionStatusTimeoutFlags

			case "transaction-status-result":
				epf = mpesaTransactionStatusResultFlags

			case "reversal-timeout":
				epf = mpesaReversalTimeoutFlags

			case "reversal-result":
				epf = mpesaReversalResultFlags

			case "b2-c-timeout":
				epf = mpesaB2CTimeoutFlags

			case "b2-c-result":
				epf = mpesaB2CResultFlags

			case "c2-b-validation":
				epf = mpesaC2BValidationFlags

			case "c2-b-confirmation":
				epf = mpesaC2BConfirmationFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint endpoint.Endpoint
		err      error
	)
	{
		switch svcn {
		case "health":
			c := healthc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "show":
				endpoint = c.Show()
				data = nil
			}
		case "mpesa":
			c := mpesac.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "account-balance-timeout":
				endpoint = c.AccountBalanceTimeout()
				data, err = mpesac.BuildAccountBalanceTimeoutPayload(*mpesaAccountBalanceTimeoutBodyFlag)
			case "account-balance-result":
				endpoint = c.AccountBalanceResultEndpoint()
				data, err = mpesac.BuildAccountBalanceResultEndpointPayload(*mpesaAccountBalanceResultBodyFlag)
			case "transaction-status-timeout":
				endpoint = c.TransactionStatusTimeout()
				data, err = mpesac.BuildTransactionStatusTimeoutPayload(*mpesaTransactionStatusTimeoutBodyFlag)
			case "transaction-status-result":
				endpoint = c.TransactionStatusResultEndpoint()
				data, err = mpesac.BuildTransactionStatusResultEndpointPayload(*mpesaTransactionStatusResultBodyFlag)
			case "reversal-timeout":
				endpoint = c.ReversalTimeout()
				data, err = mpesac.BuildReversalTimeoutPayload(*mpesaReversalTimeoutBodyFlag)
			case "reversal-result":
				endpoint = c.ReversalResultEndpoint()
				data, err = mpesac.BuildReversalResultEndpointPayload(*mpesaReversalResultBodyFlag)
			case "b2-c-timeout":
				endpoint = c.B2CTimeout()
				data, err = mpesac.BuildB2CTimeoutPayload(*mpesaB2CTimeoutBodyFlag)
			case "b2-c-result":
				endpoint = c.B2CResult()
				data, err = mpesac.BuildB2CResultPayload(*mpesaB2CResultBodyFlag)
			case "c2-b-validation":
				endpoint = c.C2BValidation()
				data, err = mpesac.BuildC2BValidationPayload(*mpesaC2BValidationBodyFlag)
			case "c2-b-confirmation":
				endpoint = c.C2BConfirmation()
				data, err = mpesac.BuildC2BConfirmationPayload(*mpesaC2BConfirmationBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// healthUsage displays the usage of the health command and its subcommands.
func healthUsage() {
	fmt.Fprintf(os.Stderr, `Service is the health service interface.
Usage:
    %s [globalflags] health COMMAND [flags]

COMMAND:
    show: Health check endpoint.

Additional help:
    %s health COMMAND --help
`, os.Args[0], os.Args[0])
}
func healthShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] health show

Health check endpoint.

Example:
    `+os.Args[0]+` health show
`, os.Args[0])
}

// mpesaUsage displays the usage of the mpesa command and its subcommands.
func mpesaUsage() {
	fmt.Fprintf(os.Stderr, `Service is the mpesa service interface.
Usage:
    %s [globalflags] mpesa COMMAND [flags]

COMMAND:
    account-balance-timeout: Account Balance Queue TimeOut URL
    account-balance-result: Account Balance Result URL
    transaction-status-timeout: Transaction Status Queue TimeOut URL
    transaction-status-result: Transaction Status Result URL
    reversal-timeout: Reversal Queue TimeOut URL
    reversal-result: Reversal Result URL
    b2-c-timeout: B2C Queue TimeOut URL
    b2-c-result: B2C Result URL
    c2-b-validation: C2B Validation URL
    c2-b-confirmation: C2B Confirmation URL

Additional help:
    %s mpesa COMMAND --help
`, os.Args[0], os.Args[0])
}
func mpesaAccountBalanceTimeoutUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mpesa account-balance-timeout -body JSON

Account Balance Queue TimeOut URL
    -body JSON: 

Example:
    `+os.Args[0]+` mpesa account-balance-timeout --body '{
      "Result": {
         "ConversationId": "236543-276372-2",
         "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
         "ReferenceData": {
            "ReferenceItem": {
               "QueueTimeoutURL": "https://internalsandbox.safaricom.co.ke/mpesa/abresults/v1/submit"
            }
         },
         "ResultCode": 0,
         "ResultDesc": "Initiator information is invalid",
         "ResultParameters": {
            "ResultParameter": [
               {
                  "AccountBalance": "Working Account|KES|46713.00|46713.00|0.00|0.00\u0026Float Account|KES|0.00|0.00|0.00|0.00\u0026Utility Account|KES|49217.00|49217.00|0.00|0.00\u0026Charges Paid Account|KES|-220.00|-220.00|0.00|0.00\u0026Organization Settlement Account|KES|0.00|0.00|0.00|0.00",
                  "BOCompletedTime": "20170728095642"
               },
               {
                  "AccountBalance": "Working Account|KES|46713.00|46713.00|0.00|0.00\u0026Float Account|KES|0.00|0.00|0.00|0.00\u0026Utility Account|KES|49217.00|49217.00|0.00|0.00\u0026Charges Paid Account|KES|-220.00|-220.00|0.00|0.00\u0026Organization Settlement Account|KES|0.00|0.00|0.00|0.00",
                  "BOCompletedTime": "20170728095642"
               }
            ]
         },
         "ResultType": 0,
         "TransactionID": "LHG31AA5TX"
      }
   }'
`, os.Args[0])
}

func mpesaAccountBalanceResultUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mpesa account-balance-result -body JSON

Account Balance Result URL
    -body JSON: 

Example:
    `+os.Args[0]+` mpesa account-balance-result --body '{
      "Result": {
         "ConversationId": "236543-276372-2",
         "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
         "ReferenceData": {
            "ReferenceItem": {
               "QueueTimeoutURL": "https://internalsandbox.safaricom.co.ke/mpesa/abresults/v1/submit"
            }
         },
         "ResultCode": 0,
         "ResultDesc": "Initiator information is invalid",
         "ResultParameters": {
            "ResultParameter": [
               {
                  "AccountBalance": "Working Account|KES|46713.00|46713.00|0.00|0.00\u0026Float Account|KES|0.00|0.00|0.00|0.00\u0026Utility Account|KES|49217.00|49217.00|0.00|0.00\u0026Charges Paid Account|KES|-220.00|-220.00|0.00|0.00\u0026Organization Settlement Account|KES|0.00|0.00|0.00|0.00",
                  "BOCompletedTime": "20170728095642"
               },
               {
                  "AccountBalance": "Working Account|KES|46713.00|46713.00|0.00|0.00\u0026Float Account|KES|0.00|0.00|0.00|0.00\u0026Utility Account|KES|49217.00|49217.00|0.00|0.00\u0026Charges Paid Account|KES|-220.00|-220.00|0.00|0.00\u0026Organization Settlement Account|KES|0.00|0.00|0.00|0.00",
                  "BOCompletedTime": "20170728095642"
               }
            ]
         },
         "ResultType": 0,
         "TransactionID": "LHG31AA5TX"
      }
   }'
`, os.Args[0])
}

func mpesaTransactionStatusTimeoutUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mpesa transaction-status-timeout -body JSON

Transaction Status Queue TimeOut URL
    -body JSON: 

Example:
    `+os.Args[0]+` mpesa transaction-status-timeout --body '{
      "Result": {
         "ConversationId": "236543-276372-2",
         "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
         "ReferenceData": {
            "ReferenceItem": {
               "Occasion": "Occasion"
            }
         },
         "ResultCode": 0,
         "ResultDesc": "Initiator information is invalid",
         "ResultParameters": {
            "ResultParameter": [
               {
                  "Amount": 10,
                  "ConversationID": "AG_20170727_00004492b1b6d0078fbe",
                  "CreditPartyName": "254708374149 - John Doe",
                  "DebitAccountType": "Utility Account",
                  "DebitPartyCharges": "Fee For B2C Payment|KES|33.00",
                  "DebitPartyName": "600134 - Safaricom157",
                  "FinalisedTime": "20170727101415",
                  "InitiatedTime": "20170727101415",
                  "Originator Conversation ID": "19455-773836-1",
                  "ReasonType": "Salary Payment via API",
                  "ReceiptNo": "LGR919G2AV",
                  "TransactionReason": "Transaction Reason",
                  "TransactionStatus": "Completed"
               },
               {
                  "Amount": 10,
                  "ConversationID": "AG_20170727_00004492b1b6d0078fbe",
                  "CreditPartyName": "254708374149 - John Doe",
                  "DebitAccountType": "Utility Account",
                  "DebitPartyCharges": "Fee For B2C Payment|KES|33.00",
                  "DebitPartyName": "600134 - Safaricom157",
                  "FinalisedTime": "20170727101415",
                  "InitiatedTime": "20170727101415",
                  "Originator Conversation ID": "19455-773836-1",
                  "ReasonType": "Salary Payment via API",
                  "ReceiptNo": "LGR919G2AV",
                  "TransactionReason": "Transaction Reason",
                  "TransactionStatus": "Completed"
               },
               {
                  "Amount": 10,
                  "ConversationID": "AG_20170727_00004492b1b6d0078fbe",
                  "CreditPartyName": "254708374149 - John Doe",
                  "DebitAccountType": "Utility Account",
                  "DebitPartyCharges": "Fee For B2C Payment|KES|33.00",
                  "DebitPartyName": "600134 - Safaricom157",
                  "FinalisedTime": "20170727101415",
                  "InitiatedTime": "20170727101415",
                  "Originator Conversation ID": "19455-773836-1",
                  "ReasonType": "Salary Payment via API",
                  "ReceiptNo": "LGR919G2AV",
                  "TransactionReason": "Transaction Reason",
                  "TransactionStatus": "Completed"
               }
            ]
         },
         "ResultType": 0,
         "TransactionID": "LHG31AA5TX"
      }
   }'
`, os.Args[0])
}

func mpesaTransactionStatusResultUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mpesa transaction-status-result -body JSON

Transaction Status Result URL
    -body JSON: 

Example:
    `+os.Args[0]+` mpesa transaction-status-result --body '{
      "Result": {
         "ConversationId": "236543-276372-2",
         "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
         "ReferenceData": {
            "ReferenceItem": {
               "Occasion": "Occasion"
            }
         },
         "ResultCode": 0,
         "ResultDesc": "Initiator information is invalid",
         "ResultParameters": {
            "ResultParameter": [
               {
                  "Amount": 10,
                  "ConversationID": "AG_20170727_00004492b1b6d0078fbe",
                  "CreditPartyName": "254708374149 - John Doe",
                  "DebitAccountType": "Utility Account",
                  "DebitPartyCharges": "Fee For B2C Payment|KES|33.00",
                  "DebitPartyName": "600134 - Safaricom157",
                  "FinalisedTime": "20170727101415",
                  "InitiatedTime": "20170727101415",
                  "Originator Conversation ID": "19455-773836-1",
                  "ReasonType": "Salary Payment via API",
                  "ReceiptNo": "LGR919G2AV",
                  "TransactionReason": "Transaction Reason",
                  "TransactionStatus": "Completed"
               },
               {
                  "Amount": 10,
                  "ConversationID": "AG_20170727_00004492b1b6d0078fbe",
                  "CreditPartyName": "254708374149 - John Doe",
                  "DebitAccountType": "Utility Account",
                  "DebitPartyCharges": "Fee For B2C Payment|KES|33.00",
                  "DebitPartyName": "600134 - Safaricom157",
                  "FinalisedTime": "20170727101415",
                  "InitiatedTime": "20170727101415",
                  "Originator Conversation ID": "19455-773836-1",
                  "ReasonType": "Salary Payment via API",
                  "ReceiptNo": "LGR919G2AV",
                  "TransactionReason": "Transaction Reason",
                  "TransactionStatus": "Completed"
               },
               {
                  "Amount": 10,
                  "ConversationID": "AG_20170727_00004492b1b6d0078fbe",
                  "CreditPartyName": "254708374149 - John Doe",
                  "DebitAccountType": "Utility Account",
                  "DebitPartyCharges": "Fee For B2C Payment|KES|33.00",
                  "DebitPartyName": "600134 - Safaricom157",
                  "FinalisedTime": "20170727101415",
                  "InitiatedTime": "20170727101415",
                  "Originator Conversation ID": "19455-773836-1",
                  "ReasonType": "Salary Payment via API",
                  "ReceiptNo": "LGR919G2AV",
                  "TransactionReason": "Transaction Reason",
                  "TransactionStatus": "Completed"
               },
               {
                  "Amount": 10,
                  "ConversationID": "AG_20170727_00004492b1b6d0078fbe",
                  "CreditPartyName": "254708374149 - John Doe",
                  "DebitAccountType": "Utility Account",
                  "DebitPartyCharges": "Fee For B2C Payment|KES|33.00",
                  "DebitPartyName": "600134 - Safaricom157",
                  "FinalisedTime": "20170727101415",
                  "InitiatedTime": "20170727101415",
                  "Originator Conversation ID": "19455-773836-1",
                  "ReasonType": "Salary Payment via API",
                  "ReceiptNo": "LGR919G2AV",
                  "TransactionReason": "Transaction Reason",
                  "TransactionStatus": "Completed"
               }
            ]
         },
         "ResultType": 0,
         "TransactionID": "LHG31AA5TX"
      }
   }'
`, os.Args[0])
}

func mpesaReversalTimeoutUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mpesa reversal-timeout -body JSON

Reversal Queue TimeOut URL
    -body JSON: 

Example:
    `+os.Args[0]+` mpesa reversal-timeout --body '{
      "Result": {
         "ConversationId": "236543-276372-2",
         "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
         "ReferenceData": {
            "ReferenceItem": {
               "QueueTimeoutURL": "https://internalsandbox.safaricom.co.ke/mpesa/reversalresults/v1/submit"
            }
         },
         "ResultCode": 0,
         "ResultDesc": "Initiator information is invalid",
         "ResultType": 0,
         "TransactionID": "LHG31AA5TX"
      }
   }'
`, os.Args[0])
}

func mpesaReversalResultUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mpesa reversal-result -body JSON

Reversal Result URL
    -body JSON: 

Example:
    `+os.Args[0]+` mpesa reversal-result --body '{
      "Result": {
         "ConversationId": "236543-276372-2",
         "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
         "ReferenceData": {
            "ReferenceItem": {
               "QueueTimeoutURL": "https://internalsandbox.safaricom.co.ke/mpesa/reversalresults/v1/submit"
            }
         },
         "ResultCode": 0,
         "ResultDesc": "Initiator information is invalid",
         "ResultType": 0,
         "TransactionID": "LHG31AA5TX"
      }
   }'
`, os.Args[0])
}

func mpesaB2CTimeoutUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mpesa b2-c-timeout -body JSON

B2C Queue TimeOut URL
    -body JSON: 

Example:
    `+os.Args[0]+` mpesa b2-c-timeout --body '{
      "ReferenceData": {
         "ReferenceItem": {
            "QueueTimeoutURL": "https://internalsandbox.safaricom.co.ke/mpesa/b2cresults/v1/submit"
         }
      },
      "Result": {
         "ConversationId": "236543-276372-2",
         "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
         "ResultCode": 0,
         "ResultDesc": "Initiator information is invalid",
         "ResultType": 0,
         "TransactionID": "LHG31AA5TX"
      },
      "ResultParameters": {
         "ResultParameter": [
            {
               "B2CChargesPaidAccountAvailableFunds": 236543.9,
               "B2CRecipientIsRegisteredCustomer": "Y",
               "B2CUtilityAccountAvailableFunds": 23654.5,
               "B2CWorkingAccountAvailableFunds": 2000,
               "ReceiverPartyPublicName": "254722000000 - Safaricom PLC",
               "TransactionAmount": 8000,
               "TransactionCompletedDateTime": "17.07.2017 10:54:57",
               "TransactionReceipt": "LGH3197RIB"
            },
            {
               "B2CChargesPaidAccountAvailableFunds": 236543.9,
               "B2CRecipientIsRegisteredCustomer": "Y",
               "B2CUtilityAccountAvailableFunds": 23654.5,
               "B2CWorkingAccountAvailableFunds": 2000,
               "ReceiverPartyPublicName": "254722000000 - Safaricom PLC",
               "TransactionAmount": 8000,
               "TransactionCompletedDateTime": "17.07.2017 10:54:57",
               "TransactionReceipt": "LGH3197RIB"
            }
         ]
      }
   }'
`, os.Args[0])
}

func mpesaB2CResultUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mpesa b2-c-result -body JSON

B2C Result URL
    -body JSON: 

Example:
    `+os.Args[0]+` mpesa b2-c-result --body '{
      "ReferenceData": {
         "ReferenceItem": {
            "QueueTimeoutURL": "https://internalsandbox.safaricom.co.ke/mpesa/b2cresults/v1/submit"
         }
      },
      "Result": {
         "ConversationId": "236543-276372-2",
         "OriginatorConversationId": "AG_2376487236_126732989KJHJKH",
         "ResultCode": 0,
         "ResultDesc": "Initiator information is invalid",
         "ResultType": 0,
         "TransactionID": "LHG31AA5TX"
      },
      "ResultParameters": {
         "ResultParameter": [
            {
               "B2CChargesPaidAccountAvailableFunds": 236543.9,
               "B2CRecipientIsRegisteredCustomer": "Y",
               "B2CUtilityAccountAvailableFunds": 23654.5,
               "B2CWorkingAccountAvailableFunds": 2000,
               "ReceiverPartyPublicName": "254722000000 - Safaricom PLC",
               "TransactionAmount": 8000,
               "TransactionCompletedDateTime": "17.07.2017 10:54:57",
               "TransactionReceipt": "LGH3197RIB"
            },
            {
               "B2CChargesPaidAccountAvailableFunds": 236543.9,
               "B2CRecipientIsRegisteredCustomer": "Y",
               "B2CUtilityAccountAvailableFunds": 23654.5,
               "B2CWorkingAccountAvailableFunds": 2000,
               "ReceiverPartyPublicName": "254722000000 - Safaricom PLC",
               "TransactionAmount": 8000,
               "TransactionCompletedDateTime": "17.07.2017 10:54:57",
               "TransactionReceipt": "LGH3197RIB"
            },
            {
               "B2CChargesPaidAccountAvailableFunds": 236543.9,
               "B2CRecipientIsRegisteredCustomer": "Y",
               "B2CUtilityAccountAvailableFunds": 23654.5,
               "B2CWorkingAccountAvailableFunds": 2000,
               "ReceiverPartyPublicName": "254722000000 - Safaricom PLC",
               "TransactionAmount": 8000,
               "TransactionCompletedDateTime": "17.07.2017 10:54:57",
               "TransactionReceipt": "LGH3197RIB"
            }
         ]
      }
   }'
`, os.Args[0])
}

func mpesaC2BValidationUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mpesa c2-b-validation -body JSON

C2B Validation URL
    -body JSON: 

Example:
    `+os.Args[0]+` mpesa c2-b-validation --body '{
      "BillRefNumber": "ele",
      "BusinessShortCode": 654321,
      "FirstName": "John",
      "InvoiceNumber": "Animi et.",
      "LastName": "Jane",
      "MSISDN": 7807928622680112375,
      "MiddleName": "Doe",
      "OrgAccountBalance": 30671,
      "ThirdPartyTransID": "Ratione minima rerum voluptatem hic molestiae expedita.",
      "TransAmount": 100,
      "TransID": "LHG31AA5TX",
      "TransTime": "20180713154301",
      "TransactionType": "Buy Goods"
   }'
`, os.Args[0])
}

func mpesaC2BConfirmationUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mpesa c2-b-confirmation -body JSON

C2B Confirmation URL
    -body JSON: 

Example:
    `+os.Args[0]+` mpesa c2-b-confirmation --body '{
      "BillRefNumber": "3wz",
      "BusinessShortCode": 654321,
      "FirstName": "John",
      "InvoiceNumber": "Ut voluptatem dolores qui reprehenderit.",
      "LastName": "Jane",
      "MSISDN": 3226968445935365337,
      "MiddleName": "Doe",
      "OrgAccountBalance": 30671,
      "ThirdPartyTransID": "Animi qui aliquam occaecati expedita at.",
      "TransAmount": 100,
      "TransID": "LHG31AA5TX",
      "TransTime": "20180713154301",
      "TransactionType": "Pay Bill"
   }'
`, os.Args[0])
}
