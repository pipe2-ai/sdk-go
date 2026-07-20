# pipe2-ai/sdk-go

Official Go SDK for [Pipe2.ai](https://pipe2.ai) — run AI media pipelines
(video generation, image generation, text-to-speech, video reframing) and track
their results from Go.

## Install

```bash
go get github.com/pipe2-ai/sdk-go
```

## Authenticate

`NewClient` takes a bearer token — a personal access token from your
[Pipe2.ai account](https://pipe2.ai), or a user JWT. It defaults to the
production endpoint `https://api.pipe2.ai/v1/graphql`; pass a second argument to
point somewhere else.

```go
import pipe2 "github.com/pipe2-ai/sdk-go"

client := pipe2.NewClient(os.Getenv("PIPE2_TOKEN"))
```

## Run a pipeline

Pipeline inputs are pipeline-specific, so they are passed as raw JSON. Browse
the available pipelines and their input schemas at
[pipe2.ai](https://pipe2.ai).

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	pipe2 "github.com/pipe2-ai/sdk-go"
)

func main() {
	ctx := context.Background()
	client := pipe2.NewClient(os.Getenv("PIPE2_TOKEN"))

	input := json.RawMessage(`{"prompt": "a red fox in falling snow, cinematic"}`)

	run, err := pipe2.RunPipeline(ctx, client, "video-generator", input)
	if err != nil {
		panic(err)
	}
	fmt.Println("run:", run.Run_pipeline.Run_id)

	// Runs are asynchronous — fetch the current status by run id, and repeat
	// until it reaches a terminal state.
	status, err := pipe2.GetPipelineRun(ctx, client, run.Run_pipeline.Run_id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status: %+v\n", status.Pipeline_runs_by_pk)
}
```

## Estimate cost before running

Runs are billed in credits. `EstimatePipelineCost` returns the price for a given
pipeline and input without starting anything.

```go
est, err := pipe2.EstimatePipelineCost(ctx, client, "video-generator", input)
```

## Uploading assets

Pipelines that take an image, audio, or video reference need the file uploaded
first. Use `RequestUpload` for ordinary files, or `RequestMultipartUpload` /
`CompleteMultipartUpload` / `AbortMultipartUpload` for large ones.

## Related

- [Pipe2.ai](https://pipe2.ai) — pipeline catalog, pricing, and docs
- [pipe2-cli](https://github.com/pipe2-ai/pipe2-cli) — agent-native command line tool
- [`@pipe2-ai/sdk`](https://www.npmjs.com/package/@pipe2-ai/sdk) — TypeScript client
