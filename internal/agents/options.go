package agents

import (
	"context"
	"getidex_api/internal/agents/workflows"
	"getidex_api/internal/chain"
	"getidex_api/internal/conversation"
	"getidex_api/internal/utils"

	Lagents "github.com/starmvp/langchaingo/agents"
	Lcallbacks "github.com/starmvp/langchaingo/callbacks"
	Lchains "github.com/starmvp/langchaingo/chains"
	Lllms "github.com/starmvp/langchaingo/llms"
	Lprompts "github.com/starmvp/langchaingo/prompts"
	Lschema "github.com/starmvp/langchaingo/schema"
	Ltools "github.com/starmvp/langchaingo/tools"
)

type Options struct {
	Options []Lagents.Option

	LLM Lllms.Model

	LangChainAgent   Lagents.Agent
	UseStreamingMode bool

	// LangChainAgent options
	PromptTemplate           Lprompts.PromptTemplate
	PromptPrefix             string
	PromptSuffix             string
	PromptFormatInstructions string
	OutputKey                string

	Ctx              context.Context
	Builder          *chain.ChainBuilder
	Chains           []Lchains.Chain
	Tools            []Ltools.Tool
	Memory           *Lschema.Memory
	CallbacksHandler []Lcallbacks.Handler
	Conversation     conversation.Conversation

	utils.IO

	// ConversationalWorkflowAgent
	Workflows []workflows.Workflow
}

type Option func(*Options)

func WithLangChainOption(o Lagents.Option) Option {
	return func(opts *Options) {
		opts.Options = append(opts.Options, o)
	}
}

func WithLLM(llm Lllms.Model) Option {
	return func(o *Options) {
		o.LLM = llm
	}
}

func WithLangChainAgent(a Lagents.Agent) Option {
	return func(o *Options) {
		o.LangChainAgent = a
	}
}

// for (starmvp/langchaingo/agents/conversational.go)ConversationalAgent
func WithStreamingMode(use bool) Option {
	return func(o *Options) {
		o.UseStreamingMode = use
	}
}

func WithPromptTemplate(t Lprompts.PromptTemplate) Option {
	return func(o *Options) {
		o.PromptTemplate = t
		o.Options = append(o.Options, Lagents.WithPrompt(t))
	}
}

func WithPromptPrefix(prefix string) Option {
	return func(o *Options) {
		o.PromptPrefix = prefix
		o.Options = append(o.Options, Lagents.WithPromptPrefix(prefix))
	}
}

func WithFormatInstructions(instructions string) Option {
	return func(o *Options) {
		o.PromptFormatInstructions = instructions
		o.Options = append(o.Options, Lagents.WithPromptFormatInstructions(instructions))
	}
}

func WithPromptSuffix(suffix string) Option {
	return func(o *Options) {
		o.PromptSuffix = suffix
		o.Options = append(o.Options, Lagents.WithPromptSuffix(suffix))
	}
}

func WithOutputKey(key string) Option {
	return func(o *Options) {
		o.OutputKey = key
		o.Options = append(o.Options, Lagents.WithOutputKey(key))
	}
}

func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		o.Ctx = ctx
	}
}

func WithChainBuilder(b *chain.ChainBuilder) Option {
	return func(o *Options) {
		o.Builder = b
	}
}

func WithChains(c []Lchains.Chain) Option {
	return func(o *Options) {
		o.Chains = append(o.Chains, c...)
	}
}

func WithChain(c Lchains.Chain) Option {
	return func(o *Options) {
		o.Chains = append(o.Chains, c)
	}
}

func WithTools(t []Ltools.Tool) Option {
	return func(o *Options) {
		o.Tools = append(o.Tools, t...)
	}
}

func WithTool(t Ltools.Tool) Option {
	return func(o *Options) {
		o.Tools = append(o.Tools, t)
	}
}

func WithMemory(m Lschema.Memory) Option {
	return func(o *Options) {
		o.Memory = &m
	}
}

func WithCallbacksHandler(h Lcallbacks.Handler) Option {
	return func(o *Options) {
		o.CallbacksHandler = append(o.CallbacksHandler, h)
	}
}

func WithConversation(c conversation.Conversation) Option {
	return func(o *Options) {
		o.Conversation = c
	}
}

func WithIO(io utils.IO) Option {
	return func(o *Options) {
		o.IO = io
	}
}

func WithWorkflows(w []workflows.Workflow) Option {
	return func(o *Options) {
		o.Workflows = w
	}
}