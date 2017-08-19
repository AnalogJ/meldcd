package engine

import (
	"capsulecd/pkg/config"
	"capsulecd/pkg/errors"
	"capsulecd/pkg/pipeline"
	"capsulecd/pkg/scm"
	"fmt"
)

func Create(engineType string, pipelineData *pipeline.Data, configImpl config.Interface, sourceImpl scm.Interface) (Interface, error) {

	switch engineType {
	case "chef":
		eng := new(engineChef)
		if err := eng.Init(pipelineData, configImpl, sourceImpl); err != nil {
			return nil, err
		}
		return eng, nil
	case "generic":
		eng := new(engineGeneric)
		if err := eng.Init(pipelineData, configImpl, sourceImpl); err != nil {
			return nil, err
		}
		return eng, nil
	case "golang":
		eng := new(engineGolang)
		if err := eng.Init(pipelineData, configImpl, sourceImpl); err != nil {
			return nil, err
		}
		return eng, nil
	case "node":
		eng := new(engineNode)
		if err := eng.Init(pipelineData, configImpl, sourceImpl); err != nil {
			return nil, err
		}
		return eng, nil
	case "python":
		eng := new(enginePython)
		if err := eng.Init(pipelineData, configImpl, sourceImpl); err != nil {
			return nil, err
		}
		return eng, nil
	case "ruby":
		eng := new(engineRuby)
		if err := eng.Init(pipelineData, configImpl, sourceImpl); err != nil {
			return nil, err
		}
		return eng, nil
	default:
		return nil, errors.EngineUnspecifiedError(fmt.Sprintf("Unknown Engine Type: %s", engineType))
	}
}
