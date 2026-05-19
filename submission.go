package judge0

import(
	"time"
	"encoding/json"

)

type Submission struct {

	SourceCode string `json:"source_code"`
	LanguageID int `json:"language_id"`

	CompilerOptions string `json:"compiler_options,omitempty"`
	CommandLineArguments string `json:"command_line_arguments,omitempty"`
	Stdin string `json:"stdin,omitempty"`
	ExpectedOutput string `json:"expected_output,omitempty"`
	CPUTimeLimit *float64 `json:"cpu_time_limit,omitempty"`
	CPUExtraTime *float64 `json:"cpu_extra_time,omitempty"`
	WallTimeLimit *float64 `json:"wall_time_limit,omitempty"`
	MemoryLimit *float64 `json:"memory_limit,omitempty"`
	StackLimit	*int `json:"stack_limit,omitempty"`
	MaxProcessesAndOrThreads *int `json:"max_processes_and_or_threads,omitempty"`
	EnablePerProcessAndThreadTimeLimit *bool `json:"enable_per_process_and_thread_time_limit,omitempty"`
	EnablePerProcessAndThreadMemoryLimit *bool `json:"enable_per_process_and_thread_memory_limit,omitempty"`
	MaxFileSize *int `json:"max_file_size,omitempty"`
	RedirectStderrToStdout *bool `json:"redirect_stderr_to_stdout,omitempty"`
	EnableNetwork *bool `json:"enable_network,omitempty"`
	NumberOfRuns *int `json:"number_of_runs,omitempty"`
	AdditionalFiles string `json:"additional_files,omitempty"`//base 64 encoded
	CallbackURL string `json:"callback_url,omitempty"`

	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	CompileOutput string`json:"compile_output"`
	Message string `json:"message"`
	ExitCode int `json:"exit_code"`
	ExitSignal int `json:"exit_signal"`
	Status Status `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	FinishedAt time.Time `json:"finished_at"`
	Token string `json:"token"`
	Time float64 `json:"time"`
	WallTime float64 `json:"wall_time"`
	Memory float64 `json:"memory"`

}



func (sub *Submission) asBody() ([]byte, error){
	dat, err := json.Marshal(sub)
	if err!=nil{
		return nil, err
	}

	return dat, nil
}