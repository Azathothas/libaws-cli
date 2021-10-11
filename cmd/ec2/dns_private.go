package cliaws

import (
	"context"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/nathants/cli-aws/lib"
)

func init() {
	lib.Commands["ec2-dns-private"] = ec2DnsPrivate
	lib.Args["ec2-dns-private"] = ec2DnsPrivateArgs{}
}

type ec2DnsPrivateArgs struct {
	Selectors []string `arg:"positional" help:"instance-id | dns-name | private-dns-name | tag | vpc-id | subnet-id | security-group-id | ip-address | private-ip-address"`
}

func (ec2DnsPrivateArgs) Description() string {
	return "\nlist ec2 private dns\n"
}

func ec2DnsPrivate() {
	var args ec2DnsPrivateArgs
	arg.MustParse(&args)
	ctx := context.Background()
	instances, err := lib.EC2ListInstances(ctx, args.Selectors, ec2.InstanceStateNameRunning)
	if err != nil {
		lib.Logger.Fatal("error: ", err)
	}
	if len(instances) == 0 {
		os.Exit(1)
	}
	for _, instance := range instances {
		if instance.PrivateDnsName != nil {
			fmt.Println(*instance.PrivateDnsName)
		}
	}
}
