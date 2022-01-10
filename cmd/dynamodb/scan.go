package cliaws

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/alexflint/go-arg"
	"github.com/nathants/cli-aws/lib"
)

func init() {
	lib.Commands["dynamodb-scan"] = dynamodbScan
	lib.Args["dynamodb-scan"] = dynamodbScanArgs{}
}

type dynamodbScanArgs struct {
	Table string `arg:"positional"`
}

func (dynamodbScanArgs) Description() string {
	return "\nscan dynamodb table\n"
}

func dynamodbScan() {
	var args dynamodbScanArgs
	arg.MustParse(&args)
	ctx := context.Background()
	var start map[string]*dynamodb.AttributeValue
	for {
		out, err := lib.DynamoDBClient().ScanWithContext(ctx, &dynamodb.ScanInput{
			TableName:         aws.String(args.Table),
			ExclusiveStartKey: start,
		})
		if err != nil {
		    panic(err)
		}
		for _, item := range out.Items {
			val := make(map[string]interface{})
			err := dynamodbattribute.UnmarshalMap(item, &val)
			if err != nil {
			    panic(err)
			}
			bytes, err := json.Marshal(val)
			if err != nil {
			    panic(err)
			}
			fmt.Println(string(bytes))
		}
		if out.LastEvaluatedKey == nil {
			break
		}
		start = out.LastEvaluatedKey
	}
}