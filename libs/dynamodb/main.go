package dynamodb

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Repository struct {
	conn      *dynamodb.Client
	tableName string
}

var tables = map[string]string{
	"local":       "koge-tb-dsv",
	"development": "koge-tb-dsv",
	"production":  "koge-tb",
}

func NewDynamoDB(env string) (*Repository, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("sa-east-1"),
	)
	if err != nil {
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg)
	return &Repository{
		conn:      svc,
		tableName: tables[env],
	}, nil
}

func (r *Repository) Put(ctx context.Context, v map[string]string) error {
	item, err := attributevalue.MarshalMap(v)
	if err != nil {
		return err
	}

	item["pk"] = makeKey("PK", v["teste"])
	item["sk"] = makeKey("SK", v["test"])

	_, err = r.conn.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	})

	return err
}

func makeKey(keys ...string) *types.AttributeValueMemberS {
	return &types.AttributeValueMemberS{Value: strings.Join(keys, "#")}
}
