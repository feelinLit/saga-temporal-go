package workflow

import "go.temporal.io/sdk/workflow"

type Compensations struct {
	registeredCompensations []any
	arguments               [][]any
}

func (c *Compensations) Add(activity any, arguments ...any) {
	c.registeredCompensations = append(c.registeredCompensations, activity)
	c.arguments = append(c.arguments, arguments)
}

func (c *Compensations) Apply(ctx workflow.Context) {
	for i := len(c.registeredCompensations) - 1; i >= 0; i-- {
		err := workflow.ExecuteActivity(ctx, c.registeredCompensations[i], c.arguments[i]...).Get(ctx, nil)
		if err != nil {
			workflow.GetLogger(ctx).Error("executing compensation failed", "Error", err)
		}
	}
}
