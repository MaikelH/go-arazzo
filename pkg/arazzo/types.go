// Package arazzo provides Go structs for working with Arazzo specification documents.
package arazzo

// Document represents the root Arazzo document structure.
type Document struct {
	Arazzo             string              `json:"arazzo"`
	Info               Info                `json:"info"`
	SourceDescriptions []SourceDescription `json:"sourceDescriptions"`
	Workflows          []Workflow          `json:"workflows"`
	Components         *Components         `json:"components,omitempty"`
	Extensions         map[string]any      `json:"-"`
}

// Info provides metadata about the Arazzo description.
type Info struct {
	Title       string         `json:"title"`
	Summary     string         `json:"summary,omitempty"`
	Description string         `json:"description,omitempty"`
	Version     string         `json:"version"`
	Extensions  map[string]any `json:"-"`
}

// SourceDescription describes a source description (such as an OpenAPI description)
// that will be referenced by one or more workflows.
type SourceDescription struct {
	Name       string         `json:"name"`
	URL        string         `json:"url"`
	Type       string         `json:"type,omitempty"` // "arazzo" or "openapi"
	Extensions map[string]any `json:"-"`
}

// Workflow describes the steps to be taken across one or more APIs to achieve an objective.
type Workflow struct {
	WorkflowID     string            `json:"workflowId"`
	Summary        string            `json:"summary,omitempty"`
	Description    string            `json:"description,omitempty"`
	Inputs         any               `json:"inputs,omitempty"` // JSON Schema object
	DependsOn      []string          `json:"dependsOn,omitempty"`
	Steps          []Step            `json:"steps"`
	SuccessActions []any             `json:"successActions,omitempty"` // SuccessAction or ReusableObject
	FailureActions []any             `json:"failureActions,omitempty"` // FailureAction or ReusableObject
	Outputs        map[string]string `json:"outputs,omitempty"`
	Parameters     []any             `json:"parameters,omitempty"` // Parameter or ReusableObject
	Extensions     map[string]any    `json:"-"`
}

// Step describes a single workflow step which may be a call to an API operation or another workflow.
type Step struct {
	StepID          string            `json:"stepId"`
	Description     string            `json:"description,omitempty"`
	OperationID     string            `json:"operationId,omitempty"`
	OperationPath   string            `json:"operationPath,omitempty"`
	WorkflowID      string            `json:"workflowId,omitempty"`
	Parameters      []any             `json:"parameters,omitempty"` // Parameter or ReusableObject
	RequestBody     *RequestBody      `json:"requestBody,omitempty"`
	SuccessCriteria []Criterion       `json:"successCriteria,omitempty"`
	OnSuccess       []any             `json:"onSuccess,omitempty"` // SuccessAction or ReusableObject
	OnFailure       []any             `json:"onFailure,omitempty"` // FailureAction or ReusableObject
	Outputs         map[string]string `json:"outputs,omitempty"`
	Extensions      map[string]any    `json:"-"`
}

// RequestBody represents the request body to pass to an operation.
type RequestBody struct {
	ContentType  string               `json:"contentType,omitempty"`
	Payload      any                  `json:"payload,omitempty"`
	Replacements []PayloadReplacement `json:"replacements,omitempty"`
	Extensions   map[string]any       `json:"-"`
}

// PayloadReplacement describes a location within a payload and a value to set within the location.
type PayloadReplacement struct {
	Target     string         `json:"target"`
	Value      string         `json:"value"`
	Extensions map[string]any `json:"-"`
}

// Criterion specifies the context, conditions, and condition types for assertions.
type Criterion struct {
	Context    string         `json:"context,omitempty"`
	Condition  string         `json:"condition"`
	Type       string         `json:"type,omitempty"` // "simple", "regex", "jsonpath", "xpath"
	Version    string         `json:"version,omitempty"`
	Extensions map[string]any `json:"-"`
}

// SuccessAction describes an action to take upon success of a workflow step.
type SuccessAction struct {
	Name       string         `json:"name"`
	Type       string         `json:"type"` // "end" or "goto"
	WorkflowID string         `json:"workflowId,omitempty"`
	StepID     string         `json:"stepId,omitempty"`
	Criteria   []Criterion    `json:"criteria,omitempty"`
	Extensions map[string]any `json:"-"`
}

// FailureAction describes an action to take upon failure of a workflow step.
type FailureAction struct {
	Name       string         `json:"name"`
	Type       string         `json:"type"` // "end", "goto", or "retry"
	WorkflowID string         `json:"workflowId,omitempty"`
	StepID     string         `json:"stepId,omitempty"`
	RetryAfter float64        `json:"retryAfter,omitempty"`
	RetryLimit int            `json:"retryLimit,omitempty"`
	Criteria   []Criterion    `json:"criteria,omitempty"`
	Extensions map[string]any `json:"-"`
}

// ReusableObject allows referencing of objects contained within the Components Object.
type ReusableObject struct {
	Reference string `json:"reference"`
	Value     any    `json:"value,omitempty"`
}

// Parameter describes a single step parameter.
type Parameter struct {
	Name       string         `json:"name"`
	In         string         `json:"in,omitempty"` // "path", "query", "header", "cookie"
	Value      any            `json:"value"`
	Extensions map[string]any `json:"-"`
}

// Components holds a set of reusable objects for different aspects of the Arazzo Specification.
type Components struct {
	Inputs         map[string]any           `json:"inputs,omitempty"`
	Parameters     map[string]Parameter     `json:"parameters,omitempty"`
	SuccessActions map[string]SuccessAction `json:"successActions,omitempty"`
	FailureActions map[string]FailureAction `json:"failureActions,omitempty"`
	Extensions     map[string]any           `json:"-"`
}
