// aws-sms-send
//
// Copyright ©2016-2019 Chris Wedgwood
//
// License: GPL3; https://www.gnu.org/licenses/gpl-3.0.en.html

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var (
	fAuthRegion    = flag.String("authregion", "us-west-2", "AWS Region to use")
	fAuthFile      = flag.String("authfile", "", "[optional] path to file containing credentials")
	fAuthProfile   = flag.String("authprofile", "", "[optional] name of AWS profile to use")
	fDisableSSL    = flag.Bool("disablessl", false, "Disable SSL")
	fSenderID      = flag.String("senderid", "", "SenderID (if applicable)")
	fSubject       = flag.String("subject", "", "Subject (if applicable)")
	fTransactional = flag.Bool("transactional", false, "Transactional or Promotional (default)")
	fMaxPrice      = flag.String("maxprice", "0.10", "Maximum price")
	fDryRun        = flag.Bool("dryrun", false, "Dry-run; show what we would do...")
	fVerbose       = flag.Bool("verbose", false, "Be verbose")
	fDebug         = flag.Bool("debug", false, "Be even more verbose")
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 2 {
		const errMsg = `
  Specify exactly two non-flag arguments:
    the phone number in E.164 format
    and the message
`
		fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(errMsg))
		os.Exit(1)
	}

	if *fDebug {
		*fVerbose = true
	}

	awsCfg := &aws.Config{
		Region:      fAuthRegion,
		Credentials: credentials.NewSharedCredentials(*fAuthFile, *fAuthProfile),
		DisableSSL:  aws.Bool(*fDisableSSL)}
	snsSvc := sns.New(session.New(), awsCfg)

	// Build message to send
	ma := make(map[string]*sns.MessageAttributeValue)
	if *fTransactional {
		ma["AWS.SNS.SMS.SMSType"] = &sns.MessageAttributeValue{DataType: aws.String("String"), StringValue: aws.String("Transactional")}
	} else {
		ma["AWS.SNS.SMS.SMSType"] = &sns.MessageAttributeValue{DataType: aws.String("String"), StringValue: aws.String("Promotional")}
	}
	if *fMaxPrice != "" {
		ma["AWS.SNS.SMS.MaxPrice"] = &sns.MessageAttributeValue{DataType: aws.String("Number"), StringValue: fMaxPrice}
	}
	if *fSenderID != "" {
		ma["AWS.SNS.SMS.SenderID"] = &sns.MessageAttributeValue{DataType: aws.String("String"), StringValue: fSenderID}
	}
	inp := &sns.PublishInput{
		PhoneNumber:       aws.String(flag.Args()[0]),
		Message:           aws.String(flag.Args()[1]),
		MessageAttributes: ma}
	if *fSubject != "" {
		inp.Subject = fSubject
	}

	// send an SMS message
	var pr *sns.PublishOutput
	var err error
	if !*fDryRun {
		debugPrint("Sending", inp)
		pr, err = snsSvc.Publish(inp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR sending message: %v\n", err)
			os.Exit(2)
		}
		debugPrint("Got", pr)
	}

	if *fVerbose && !*fDryRun {
		fmt.Printf("Sent with MessageID: %s\n", *pr.MessageId)
	}

}

func debugPrint(msg, o interface{}) {
	if !*fDebug {
		return
	}
	d, _ := json.MarshalIndent(o, "[DEBUG] ", "    ")
	fmt.Printf("[DEBUG] %s: %s\n", msg, d)
}
