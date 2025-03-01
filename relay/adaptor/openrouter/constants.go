package openrouter

var ModelList = []string{
	"chatgpt-4o-latest",
	"gpt-3.5-turbo",
	"gpt-3.5-turbo-0125",
	"gpt-3.5-turbo-0613",
	"gpt-3.5-turbo-1106",
	"gpt-3.5-turbo-16k",
	"gpt-3.5-turbo-instruct",
	"gpt-4",
	"gpt-4-0314",
	"gpt-4-1106-preview",
	"gpt-4-32k",
	"gpt-4-32k-0314",
	"gpt-4-turbo",
	"gpt-4-turbo-preview",
	"gpt-4o",
	"gpt-4o-2024-05-13",
	"gpt-4o-2024-08-06",
	"gpt-4o-2024-11-20",
	"gpt-4o-mini",
	"gpt-4o-mini-2024-07-18",
	"gpt-4o:extended",
	"o1",
	"o1-mini",
	"o1-mini-2024-09-12",
	"o1-preview",
	"o1-preview-2024-09-12",
	"o3-mini",
	"o3-mini-high",
	"gpt-4.5-preview",
	"anthracite-org/magnum-v2-72b",
	"anthracite-org/magnum-v4-72b",
	"anthropic/claude-2",
	"anthropic/claude-2.0",
	"anthropic/claude-2.0:beta",
	"anthropic/claude-2.1",
	"anthropic/claude-2.1:beta",
	"anthropic/claude-2:beta",
	"anthropic/claude-3-haiku",
	"anthropic/claude-3-haiku:beta",
	"anthropic/claude-3-opus",
	"anthropic/claude-3-opus:beta",
	"anthropic/claude-3-sonnet",
	"anthropic/claude-3-sonnet:beta",
	"anthropic/claude-3.5-haiku",
	"anthropic/claude-3.5-haiku-20241022",
	"anthropic/claude-3.5-haiku-20241022:beta",
	"anthropic/claude-3.5-haiku:beta",
	"anthropic/claude-3.5-sonnet",
	"anthropic/claude-3.5-sonnet-20240620",
	"anthropic/claude-3.5-sonnet-20240620:beta",
	"anthropic/claude-3.5-sonnet:beta",
	"google/gemini-2.0-flash-001",
	"google/gemini-2.0-flash-exp:free",
	"google/gemini-2.0-flash-lite-preview-02-05:free",
	"google/gemini-2.0-flash-thinking-exp-1219:free",
	"google/gemini-2.0-flash-thinking-exp:free",
	"google/gemini-2.0-pro-exp-02-05:free",
	"google/gemini-exp-1206:free",
	"google/gemini-flash-1.5",
	"google/gemini-flash-1.5-8b",
	"google/gemini-flash-1.5-8b-exp",
	"google/gemini-pro",
	"google/gemini-pro-1.5",
	"google/gemini-pro-vision",
	"google/gemma-2-27b-it",
	"google/gemma-2-9b-it",
	"google/gemma-2-9b-it:free",
	"google/gemma-7b-it",
	"google/learnlm-1.5-pro-experimental:free",
	"google/palm-2-chat-bison",
	"google/palm-2-chat-bison-32k",
	"google/palm-2-codechat-bison",
	"google/palm-2-codechat-bison-32k",
	"deepseek/deepseek-chat",
	"deepseek/deepseek-chat-v2.5",
	"deepseek/deepseek-chat:free",
	"deepseek/deepseek-r1",
	"deepseek/deepseek-r1-distill-llama-70b",
	"deepseek/deepseek-r1-distill-llama-70b:free",
	"deepseek/deepseek-r1-distill-llama-8b",
	"deepseek/deepseek-r1-distill-qwen-1.5b",
	"deepseek/deepseek-r1-distill-qwen-14b",
	"deepseek/deepseek-r1-distill-qwen-32b",
	"deepseek/deepseek-r1:free",
	"qwen/qvq-72b-preview",
	"qwen/qwen-2-72b-instruct",
	"qwen/qwen-2-7b-instruct",
	"qwen/qwen-2-7b-instruct:free",
	"qwen/qwen-2-vl-72b-instruct",
	"qwen/qwen-2-vl-7b-instruct",
	"qwen/qwen-2.5-72b-instruct",
	"qwen/qwen-2.5-7b-instruct",
	"qwen/qwen-2.5-coder-32b-instruct",
	"qwen/qwen-max",
	"qwen/qwen-plus",
	"qwen/qwen-turbo",
	"qwen/qwen-vl-plus:free",
	"qwen/qwen2.5-vl-72b-instruct:free",
	"qwen/qwq-32b-preview",
	"meta-llama/llama-2-13b-chat",
	"meta-llama/llama-2-70b-chat",
	"meta-llama/llama-3-70b-instruct",
	"meta-llama/llama-3-8b-instruct",
	"meta-llama/llama-3-8b-instruct:free",
	"meta-llama/llama-3.1-405b",
	"meta-llama/llama-3.1-405b-instruct",
	"meta-llama/llama-3.1-70b-instruct",
	"meta-llama/llama-3.1-8b-instruct",
	"meta-llama/llama-3.2-11b-vision-instruct",
	"meta-llama/llama-3.2-11b-vision-instruct:free",
	"meta-llama/llama-3.2-1b-instruct",
	"meta-llama/llama-3.2-3b-instruct",
	"meta-llama/llama-3.2-90b-vision-instruct",
	"meta-llama/llama-3.3-70b-instruct",
	"meta-llama/llama-3.3-70b-instruct:free",
	"meta-llama/llama-guard-2-8b",
	"x-ai/grok-2-1212",
	"x-ai/grok-2-vision-1212",
	"x-ai/grok-beta",
	"x-ai/grok-vision-beta",
	"microsoft/phi-3-medium-128k-instruct",
	"microsoft/phi-3-medium-128k-instruct:free",
	"microsoft/phi-3-mini-128k-instruct",
	"microsoft/phi-3-mini-128k-instruct:free",
	"microsoft/phi-3.5-mini-128k-instruct",
	"microsoft/phi-4",
	"microsoft/wizardlm-2-7b",
	"microsoft/wizardlm-2-8x22b",
	"minimax/minimax-01",
	"mistralai/codestral-2501",
	"mistralai/codestral-mamba",
	"mistralai/ministral-3b",
	"mistralai/ministral-8b",
	"mistralai/mistral-7b-instruct",
	"mistralai/mistral-7b-instruct-v0.1",
	"mistralai/mistral-7b-instruct-v0.3",
	"mistralai/mistral-7b-instruct:free",
	"mistralai/mistral-large",
	"mistralai/mistral-large-2407",
	"mistralai/mistral-large-2411",
	"mistralai/mistral-medium",
	"mistralai/mistral-nemo",
	"mistralai/mistral-nemo:free",
	"mistralai/mistral-small",
	"mistralai/mistral-small-24b-instruct-2501",
	"mistralai/mistral-small-24b-instruct-2501:free",
	"mistralai/mistral-tiny",
	"mistralai/mixtral-8x22b-instruct",
	"mistralai/mixtral-8x7b",
	"mistralai/mixtral-8x7b-instruct",
	"mistralai/pixtral-12b",
	"mistralai/pixtral-large-2411",
	"perplexity/llama-3.1-sonar-huge-128k-online",
	"perplexity/llama-3.1-sonar-large-128k-chat",
	"perplexity/llama-3.1-sonar-large-128k-online",
	"perplexity/llama-3.1-sonar-small-128k-chat",
	"perplexity/llama-3.1-sonar-small-128k-online",
	"perplexity/sonar",
	"perplexity/sonar-reasoning",
	"01-ai/yi-large",
	"aetherwiing/mn-starcannon-12b",
	"ai21/jamba-1-5-large",
	"ai21/jamba-1-5-mini",
	"ai21/jamba-instruct",
	"aion-labs/aion-1.0",
	"aion-labs/aion-1.0-mini",
	"aion-labs/aion-rp-llama-3.1-8b",
	"allenai/llama-3.1-tulu-3-405b",
	"alpindale/goliath-120b",
	"alpindale/magnum-72b",
	"amazon/nova-lite-v1",
	"amazon/nova-micro-v1",
	"amazon/nova-pro-v1",
	"cognitivecomputations/dolphin-mixtral-8x22b",
	"cognitivecomputations/dolphin-mixtral-8x7b",
	"cohere/command",
	"cohere/command-r",
	"cohere/command-r-03-2024",
	"cohere/command-r-08-2024",
	"cohere/command-r-plus",
	"cohere/command-r-plus-04-2024",
	"cohere/command-r-plus-08-2024",
	"cohere/command-r7b-12-2024",
	"databricks/dbrx-instruct",
	"eva-unit-01/eva-llama-3.33-70b",
	"eva-unit-01/eva-qwen-2.5-32b",
	"eva-unit-01/eva-qwen-2.5-72b",
	"gryphe/mythomax-l2-13b",
	"gryphe/mythomax-l2-13b:free",
	"huggingfaceh4/zephyr-7b-beta:free",
	"infermatic/mn-inferor-12b",
	"inflection/inflection-3-pi",
	"inflection/inflection-3-productivity",
	"jondurbin/airoboros-l2-70b",
	"liquid/lfm-3b",
	"liquid/lfm-40b",
	"liquid/lfm-7b",
	"mancer/weaver",
	"neversleep/llama-3-lumimaid-70b",
	"neversleep/llama-3-lumimaid-8b",
	"neversleep/llama-3-lumimaid-8b:extended",
	"neversleep/llama-3.1-lumimaid-70b",
	"neversleep/llama-3.1-lumimaid-8b",
	"neversleep/noromaid-20b",
	"nothingiisreal/mn-celeste-12b",
	"nousresearch/hermes-2-pro-llama-3-8b",
	"nousresearch/hermes-3-llama-3.1-405b",
	"nousresearch/hermes-3-llama-3.1-70b",
	"nousresearch/nous-hermes-2-mixtral-8x7b-dpo",
	"nousresearch/nous-hermes-llama2-13b",
	"nvidia/llama-3.1-nemotron-70b-instruct",
	"nvidia/llama-3.1-nemotron-70b-instruct:free",
	"openchat/openchat-7b",
	"openchat/openchat-7b:free",
	"openrouter/auto",
	"pygmalionai/mythalion-13b",
	"raifle/sorcererlm-8x22b",
	"sao10k/fimbulvetr-11b-v2",
	"sao10k/l3-euryale-70b",
	"sao10k/l3-lunaris-8b",
	"sao10k/l3.1-70b-hanami-x1",
	"sao10k/l3.1-euryale-70b",
	"sao10k/l3.3-euryale-70b",
	"sophosympatheia/midnight-rose-70b",
	"sophosympatheia/rogue-rose-103b-v0.2:free",
	"teknium/openhermes-2.5-mistral-7b",
	"thedrummer/rocinante-12b",
	"thedrummer/unslopnemo-12b",
	"undi95/remm-slerp-l2-13b",
	"undi95/toppy-m-7b",
	"undi95/toppy-m-7b:free",
	"xwin-lm/xwin-lm-70b",
}
