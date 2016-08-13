package ai

import (
	"app/battle/context"
	"app/battle/event"
	battleAPI "app/typhenapi/type/submarine/battle"
)

type ai struct {
	ctx       *context.Context
	navigator *navigator
}

func newAI(ctx *context.Context) *ai {
	return &ai{
		ctx:       ctx,
		navigator: new(navigator),
	}
}

// Overridable methods.
func (a *ai) Update(submarine context.Actor) {}

func (a *ai) accelerateActor(actor context.Actor, dir float64) {
	actor.Event().Emit(event.AccelerationRequest, &battleAPI.AccelerationRequestObject{Direction: dir})
}

func (a *ai) brakeActor(actor context.Actor, dir float64) {
	actor.Event().Emit(event.BrakeRequest, &battleAPI.BrakeRequestObject{Direction: dir})
}
