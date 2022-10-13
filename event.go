package warmup

// https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-content-structure.html

type User struct {
	PrincipalId string `json:"principalId"`
}

type Bucket struct {
	Arn           string `json:"arn"`
	Name          string `json:"name"`          // name of the S3 bucket
	OwnerIdentity User   `json:"ownerIdentity"` // Amazon-customer-ID-of-the-bucket-owner
}

type Object struct {
	Etag      string `json:"etag"`      // MD5 hash for each uploaded part of the file, concatenate the hashes into a single binary string and calculate the MD5 hash of that result.
	Key       string `json:"key"`       // object-key
	Sequencer string `json:"sequencer"` // a string representation of a hexadecimal value used to determine event sequence, only used with PUTs and DELETEs
	Size      int    `json:"size"`      // object-size in bytes
	VersionId string `json:"versionId"` // object version if bucket is versioning-enabled, otherwise null
}

type S3 struct {
	Bucket        Bucket `json:"bucket"`
	Configuration string `json:"configurationId"` // ID found in the bucket notification configuration
	Object        Object `json:"object"`
	Version       string `json:"s3SchemaVersion"` // major and minor version in the form <major>.<minor>
}

type Request struct {
	SourceIPAddress string `json:"sourceIPAddress"` // ip-address-where-request-came-from
}

type Record struct {
	AwsRegion    string  `json:"awsRegion"`
	EventName    string  `json:"eventName"` // "ObjectCreated:Put" or "ObjectRemoved:Delete"
	EventTime    string  `json:"eventTime"` // The time, in ISO-8601 format, for example, 1970-01-01T00:00:00.000Z, when Amazon S3 finished processing the request
	EventSource  string  `json:"eventSource"`
	Request      Request `json:"requestParameters"`
	S3           S3      `json:"s3"` // Amazon-customer-ID-of-the-user-who-caused-the-event
	UserIdentity User    `json:"userIdentity"`
	Version      string  `json:"eventVersion"` // major and minor version in the form <major>.<minor>
}

type LambdaEvent struct {
	Records []Record `json:"Records"`
}
