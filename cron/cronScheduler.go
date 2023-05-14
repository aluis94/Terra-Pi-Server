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

var scriptsDir = ""

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkScriptErr(err error, out bytes.Buffer, out_err bytes.Buffer) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(out.String())
		fmt.Println(out_err.String())
	}
}

func ExecScript(scriptName string) string {
	var out bytes.Buffer
	var err_out bytes.Buffer
	cmd := exec.Command("python", scriptsDir+scriptName)
	cmd.Stdout = &out
	cmd.Stderr = &err_out
	err := cmd.Run()

	checkScriptErr(err, out, err_out)
	return out.String() + "\n" + err_out.String()
}

func RunCronJobs(jobs *[]models.Job) (*gocron.Scheduler, []*gocron.Job) {
	//1. Create new scheduler
	s := gocron.NewScheduler(time.UTC)
	//2. Get jobs
	var tempjobs []*gocron.Job
	var tempjob *gocron.Job

	for _, job := range *jobs {
		//3. Schedule jobs
		if job.VerbalInstr != "" {
			tempjob = setVerbalInstruction(&job, s)

		} else if job.Minute != "" && job.Hour != "" && job.Day_Month != "" && job.Month != "" && job.Day_Week != "" {
			expr := CreateCronExpression(&job)
			fmt.Println(job.ScriptName)
			tempjob, _ = s.Cron(expr).Tag(job.ScriptName).Do(ExecScript(job.ScriptName))

		}
		tempjobs = append(tempjobs, tempjob)
	}
	//4.

	return s, tempjobs

}

func setVerbalInstruction(job *models.Job, s *gocron.Scheduler) *gocron.Job {
	var vi VerbalInstruction
	var myjob *gocron.Job
	var myerr error
	if err := json.Unmarshal([]byte(job.VerbalInstr), &vi); err != nil {
		fmt.Println("Error with unmarshal")
	}
	if vi.Seconds != 0 {
		fmt.Println("Scheduled every ", vi.Seconds, " seconds ", " scriptname: ", job.ScriptName)
		myjob, myerr = s.Every(vi.Seconds).Seconds().Tag(job.ScriptName).Do(func() { ExecScript(job.ScriptName) })

	}
	if vi.Minutes != 0 {
		fmt.Println("Scheduled every ", vi.Minutes, " minutes", " scriptname: ", job.ScriptName)
		myjob, myerr = s.Every(vi.Minutes).Minutes().Tag(job.ScriptName).Do(func() { ExecScript(job.ScriptName) })
	}
	if vi.Hours != 0 {
		fmt.Println("Scheduled every ", vi.Hours, " hours", " scriptname: ", job.ScriptName)
		myjob, myerr = s.Every(vi.Hours).Hours().Tag(job.ScriptName).Do(func() { ExecScript(job.ScriptName) })
	}
	if vi.Days != 0 {
		fmt.Println("Scheduled every ", vi.Days, " days", " scriptname: ", job.ScriptName)
		if vi.Days == 1 && vi.At != "" {
			myjob, myerr = s.Every(vi.Days).Day().At(vi.At).Tag(job.ScriptName).Do(func() { ExecScript(job.ScriptName) })
		} else {
			myjob, myerr = s.Every(vi.Days).Days().Tag(job.ScriptName).Do(func() { ExecScript(job.ScriptName) })
		}

	}
	if vi.Months != 0 {
		fmt.Println("Scheduled every ", vi.Months, " months", " scripname: ", job.ScriptName)
		var daysM []int
		stringDOM := strings.Split(vi.MonthsString, ",")
		for _, dOM := range stringDOM {
			i, err := strconv.Atoi(dOM)
			check(err)
			daysM = append(daysM, i)

		}
		if vi.Months == 1 && vi.MonthsString != "" && daysM[0] != 0 {
			myjob, myerr = s.Every(vi.Months).Month(daysM...).Tag(job.ScriptName).Do(func() { ExecScript(job.ScriptName) })
		} else {
			myjob, myerr = s.Every(vi.Months).Months().Tag(job.ScriptName).Do(func() { ExecScript(job.ScriptName) })

		}

	}
	if myerr != nil {
		fmt.Println(myerr)
	}
	return myjob
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
