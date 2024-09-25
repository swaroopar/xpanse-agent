/*
 * SPDX-License-Identifier: Apache-2.0
 *  SPDX-FileCopyrightText: Huawei Inc.
 */

package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"os"
	"time"
	"xpanse-agent/pkg/logger"
	"xpanse-agent/pkg/poller"
)

func StartPolling(serviceId string, resourceName string, jobFrequency int, xpanseApiEndpoint string) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		os.Exit(1)
	}
	_, err = scheduler.NewJob(gocron.DurationJob(time.Duration(jobFrequency)*time.Second),
		gocron.NewTask(poller.PollXpanseApi, serviceId, resourceName, xpanseApiEndpoint),
		gocron.WithSingletonMode(gocron.LimitMode(1)))

	if err != nil {
		return
	}
	logger.Logger.Info(fmt.Sprintf("scheduler started for serviceId %s and resourceName %s", serviceId, resourceName))

	scheduler.Start()

	select {}
}
