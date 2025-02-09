package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/bearts/nimbus/src/authorizationserver"
	"github.com/bearts/nimbus/src/constellation"
	"github.com/bearts/nimbus/src/cron"
	"github.com/bearts/nimbus/src/docker"
	"github.com/bearts/nimbus/src/market"
	"github.com/bearts/nimbus/src/metrics"
	"github.com/bearts/nimbus/src/storage"
	"github.com/bearts/nimbus/src/utils"
)

func main() {
	utils.Log("------------------------------------------")
	utils.Log("Starting Cosmos-Server version " + GetCosmosVersion())
	utils.Log("------------------------------------------")

	// utils.ReBootstrapContainer = docker.BootstrapContainerFromTags
	utils.PushShieldMetrics = metrics.PushShieldMetrics
	utils.GetContainerIPByName = docker.GetContainerIPByName
	utils.DoesContainerExist = docker.DoesContainerExist
	utils.CheckDockerNetworkMode = docker.CheckDockerNetworkMode

	rand.Seed(time.Now().UnixNano())

	docker.IsInsideContainer()

	LoadConfig()

	utils.RemovePIDFile()

	utils.CheckHostNetwork()

	go CRON()

	docker.ExportDocker()

	docker.DockerListenEvents()

	docker.BootstrapAllContainersFromTags()

	docker.RemoveSelfUpdater()

	go func() {
		time.Sleep(180 * time.Second)
		docker.CheckUpdatesAvailable()
	}()

	version, err := docker.DockerClient.ServerVersion(context.Background())
	if err == nil {
		utils.Log("Docker API version: " + version.APIVersion)
	}

	config := utils.GetMainConfig()

	if !config.NewInstall {
		MigratePre013()
		MigratePre014()

		utils.CheckInternet()

		docker.CheckPuppetDB()

		utils.InitDBBuffers()

		utils.Log("Starting monitoring services...")

		metrics.Init()

		utils.Log("Starting market services...")

		market.Init()

		utils.Log("Starting OpenID services...")

		authorizationserver.Init()

		utils.Log("Starting constellation services...")

		utils.InitFBL()

		constellation.Init()

		storage.InitSnapRAIDConfig()

		// Has to be done last, so scheduler does not re-init
		cron.Init()

		utils.Log("Starting server...")
	}

	StartServer()
}
