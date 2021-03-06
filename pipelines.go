package transformer

import (
	"expvar"
	"log"
	//"sort"
)

type PipelineStage struct {
	Name        string
	Transformer Transformer
	Reader      StoreReader
	Writer      StoreWriter
}

var stagesDone *expvar.Int
var currentStage *expvar.String

func init() {
	stagesDone = expvar.NewInt("StagesComplete")
	currentStage = expvar.NewString("CurrentStage")
}

func RunPipeline(stages []PipelineStage, skipStages int) {
	for idx, stage := range stages[skipStages:] {
		currentStage.Set(stage.Name)
		log.Printf("Running pipeline stage %v (%v)", idx+skipStages, stage.Name)
		RunTransformer(stage.Transformer, stage.Reader, stage.Writer)
		stagesDone.Add(1)
	}
	log.Printf("Done all pipelines.")
}
