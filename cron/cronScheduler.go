package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/aluis94/terra-pi-server/models"
	"github.com/go-co-op/gocron"
)

// VerbInstr struct
type VerbalInstruction struct {
	Seconds      int
	Days         int
	Months       int
	Hours        int
	Minutes      int
	MonthsString string
	At           string
}

var scriptsDir = "scripts/"

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ExecScript(scriptName string) string {
	var out bytes.Buffer
	var err_out bytes.Buffer
	cmd := exec.Command("python", scriptsDir+scriptName)
	err := cmd.Run()
	cmd.Stdout = &out
	cmd.Stderr = &err_out
	check(err)
	return out.String() + "\n" + err_out.String()
}

func RunCronJobs(jobs *[]models.Job) *gocron.Scheduler {
	//1. Create new scheduler
	s := gocron.NewScheduler(time.UTC)
	//2. Get jobs
	for _, job := range *jobs {
		//3. Schedule jobs
		if job.VerbalInstr != "" {
			setVerbalInstruction(&job, s)
		} else if job.Minute != "" && job.Hour != "" && job.Day_Month != "" && job.Month != "" && job.Day_Week != "" {
			expr := CreateCronExpression(&job)
			fmt.Println(job.ScriptName)
			s.Cron(expr).Tag(job.ScriptName).Do(ExecScript(job.ScriptName))
		}

	}
	//4.
	s.StartAsync()
	return s

}

func setVerbalInstruction(job *models.Job, s *gocron.Scheduler) {
	var vi VerbalInstruction
	if err := json.Unmarshal([]byte(job.VerbalInstr), &vi); err != nil {
		fmt.Println("Error with unmarshal")
	}
	if vi.Seconds != 0 {
		s.Every(vi.Seconds).Seconds().Tag(job.ScriptName).Do(ExecScript(job.ScriptName))
	}
	if vi.Minutes != 0 {
		s.Every(vi.Minutes).Minutes().Tag(job.ScriptName).Do(ExecScript(job.ScriptName))
	}
	if vi.Hours != 0 {
		s.Every(vi.Hours).Hours().Tag(job.ScriptName).Do(ExecScript(job.ScriptName))
	}
	if vi.Days != 0 {
		if vi.Days == 1 && vi.At != "" {
			s.Every(vi.Days).Day().At(vi.At).Tag(job.ScriptName).Do(ExecScript(job.ScriptName))
		} else {
			s.Every(vi.Days).Days().Tag(job.ScriptName).Do(ExecScript(job.ScriptName))
		}

	}
	if vi.Months != 0 {
		var daysM []int
		stringDOM := strings.Split(vi.MonthsString, ",")
		for _, dOM := range stringDOM {
			i, err := strconv.Atoi(dOM)
			check(err)
			daysM = append(daysM, i)

		}
		if vi.Months == 1 && vi.MonthsString != "" && daysM[0] != 0 {
			s.Every(vi.Months).Month(daysM...).Tag(job.ScriptName).Do(ExecScript(job.ScriptName))
		} else {
			s.Every(vi.Months).Months().Tag(job.ScriptName).Do(ExecScript(job.ScriptName))

		}

	}

}

func CreateCronExpression(job *models.Job) string {
	expr := ""
	min := job.Minute
	hour := job.Hour
	day_month := job.Day_Month
	month := job.Month
	day_week := job.Day_Week
	if job.Minute != "" && job.Hour != "" && job.Day_Month != "" && job.Month != "" && job.Day_Week != "" {
		expr = min + " " + hour + " " + day_month + " " + month + " " + day_week
	}
	return expr
}
