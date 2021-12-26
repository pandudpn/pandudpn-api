package jobpresent

import (
	"context"
	"net/http"
	"time"

	"pandudpn/api/src/model"
	"pandudpn/api/src/utils/response"
)

type jobResponse struct {
	Id           int        `json:"id"`
	Office       string     `json:"office"`
	StartAt      time.Time  `json:"startAt"`
	EndAt        *time.Time `json:"endAt"`
	Description  string     `json:"description"`
	StillWorking bool       `json:"stillWorking"`
}

func Response(ctx context.Context, value interface{}) response.OutputResponseInterface {
	if err, ok := value.(error); ok {
		return response.Errors(ctx, http.StatusInternalServerError, err.Error(), err)
	}

	jobExp := value.([]*model.JobExperiences)
	jobs := createJobsResponse(jobExp)
	return response.Success(ctx, http.StatusOK, jobs)
}

func createJobsResponse(jobs []*model.JobExperiences) []*jobResponse {
	var jobsres = make([]*jobResponse, 0)

	for _, job := range jobs {
		jobres := &jobResponse{
			Id:           job.Id,
			Description:  job.Descrtiption,
			EndAt:        job.EndAt,
			Office:       job.Office,
			StartAt:      job.StartAt,
			StillWorking: job.StillWorking,
		}

		jobsres = append(jobsres, jobres)
	}

	return jobsres
}