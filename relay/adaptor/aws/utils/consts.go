package utils

import (
	"context"
	"slices"
	"strings"

	"github.com/songquanpeng/one-api/common/logger"
)

// CrossRegionInferences is a list of model IDs that support cross-region inference.
//
// https://docs.aws.amazon.com/bedrock/latest/userguide/inference-profiles-support.html
//
// document.querySelectorAll('pre.programlisting code').forEach((e) => {console.log(e.innerHTML)})
var CrossRegionInferences = []string{
	"us.amazon.nova-lite-v1:0",
	"us.amazon.nova-micro-v1:0",
	"us.amazon.nova-pro-v1:0",
	"us.anthropic.claude-3-5-haiku-20241022-v1:0",
	"us.anthropic.claude-3-5-sonnet-20240620-v1:0",
	"us.anthropic.claude-3-5-sonnet-20241022-v2:0",
	"us.anthropic.claude-3-7-sonnet-20250219-v1:0",
	"us.anthropic.claude-3-haiku-20240307-v1:0",
	"us.anthropic.claude-3-opus-20240229-v1:0",
	"us.anthropic.claude-3-sonnet-20240229-v1:0",
	"us.meta.llama3-1-405b-instruct-v1:0",
	"us.meta.llama3-1-70b-instruct-v1:0",
	"us.meta.llama3-1-8b-instruct-v1:0",
	"us.meta.llama3-2-11b-instruct-v1:0",
	"us.meta.llama3-2-1b-instruct-v1:0",
	"us.meta.llama3-2-3b-instruct-v1:0",
	"us.meta.llama3-2-90b-instruct-v1:0",
	"us.meta.llama3-3-70b-instruct-v1:0",
	"us-gov.anthropic.claude-3-5-sonnet-20240620-v1:0",
	"us-gov.anthropic.claude-3-haiku-20240307-v1:0",
	"eu.amazon.nova-lite-v1:0",
	"eu.amazon.nova-micro-v1:0",
	"eu.amazon.nova-pro-v1:0",
	"eu.anthropic.claude-3-5-sonnet-20240620-v1:0",
	"eu.anthropic.claude-3-haiku-20240307-v1:0",
	"eu.anthropic.claude-3-sonnet-20240229-v1:0",
	"eu.meta.llama3-2-1b-instruct-v1:0",
	"eu.meta.llama3-2-3b-instruct-v1:0",
	"apac.amazon.nova-lite-v1:0",
	"apac.amazon.nova-micro-v1:0",
	"apac.amazon.nova-pro-v1:0",
	"apac.anthropic.claude-3-5-sonnet-20240620-v1:0",
	"apac.anthropic.claude-3-5-sonnet-20241022-v2:0",
	"apac.anthropic.claude-3-haiku-20240307-v1:0",
	"apac.anthropic.claude-3-sonnet-20240229-v1:0",
}

// ConvertModelID2CrossRegionProfile converts the model ID to a cross-region profile ID.
func ConvertModelID2CrossRegionProfile(model, region string) string {
	var regionPrefix string
	switch prefix := strings.Split(region, "-")[0]; prefix {
	case "us", "eu":
		regionPrefix = prefix
	case "ap":
		regionPrefix = "apac"
	default:
		// not supported, return original model
		return model
	}

	newModelID := regionPrefix + "." + model
	if slices.Contains(CrossRegionInferences, newModelID) {
		logger.Debugf(context.TODO(), "convert model %s to cross-region profile %s", model, newModelID)
		return newModelID
	}

	// not found, return original model
	return model
}
