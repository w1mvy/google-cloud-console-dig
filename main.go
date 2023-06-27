package main

import (
	"flag"
	"fmt"
	"os/exec"

	"github.com/ktr0731/go-fuzzyfinder"
)

type Item struct {
	Name string
	URL  string
}

func main() {
	var projectID string
	flag.StringVar(&projectID, "p", "", "Google Cloud Project ID")
	flag.Parse()

	if projectID == "" {
		fmt.Println("Error: Project ID is required. Use the '-p' option.")
		return
	}

	items := []Item{
		// App Engine
		{Name: "App Engine", URL: fmt.Sprintf("https://console.cloud.google.com/appengine?project=%s", projectID)},
		{Name: "App Engine - Services", URL: fmt.Sprintf("https://console.cloud.google.com/appengine/services?project=%s", projectID)},

		// Cloud Run
		{Name: "Cloud Run", URL: fmt.Sprintf("https://console.cloud.google.com/run?project=%s", projectID)},
		{Name: "Cloud Run Jobs", URL: fmt.Sprintf("https://console.cloud.google.com/run/jobs?project=%s", projectID)},

		// BigQuery
		{Name: "BigQuery", URL: fmt.Sprintf("https://console.cloud.google.com/bigquery?project=%s", projectID)},

		// Compute Engine
		{Name: "Compute Engine", URL: fmt.Sprintf("https://console.cloud.google.com/compute?project=%s", projectID)},
		{Name: "Compute Engine - VM Instances", URL: fmt.Sprintf("https://console.cloud.google.com/compute/instances?project=%s", projectID)},
		{Name: "Compute Engine - Instance Groups", URL: fmt.Sprintf("https://console.cloud.google.com/compute/instanceGroups?project=%s", projectID)},

		// GKE
		{Name: "GKE", URL: fmt.Sprintf("https://console.cloud.google.com/kubernetes?project=%s", projectID)},
		{Name: "GKE - Clusters", URL: fmt.Sprintf("https://console.cloud.google.com/kubernetes/list?project=%s", projectID)},
		{Name: "GKE - Workloads", URL: fmt.Sprintf("https://console.cloud.google.com/kubernetes/workload?project=%s", projectID)},
		{Name: "GKE - Services & Ingress", URL: fmt.Sprintf("https://console.cloud.google.com/kubernetes/discovery?project=%s", projectID)},
		{Name: "GKE - Applications", URL: fmt.Sprintf("https://console.cloud.google.com/kubernetes/application?project=%s", projectID)},
		{Name: "GKE - Config Maps", URL: fmt.Sprintf("https://console.cloud.google.com/kubernetes/config?project=%s", projectID)},
		{Name: "GKE - Storage", URL: fmt.Sprintf("https://console.cloud.google.com/kubernetes/storage?project=%s", projectID)},

		// Cloud Storage
		{Name: "Cloud Storage", URL: fmt.Sprintf("https://console.cloud.google.com/storage/browser?project=%s", projectID)},

		// Cloud SQL
		{Name: "Cloud SQL", URL: fmt.Sprintf("https://console.cloud.google.com/sql/instances?project=%s", projectID)},

		// Cloud Functions
		{Name: "Cloud Functions", URL: fmt.Sprintf("https://console.cloud.google.com/functions?project=%s", projectID)},

		// IAM & Admin
		{Name: "IAM & Admin", URL: fmt.Sprintf("https://console.cloud.google.com/iam-admin?project=%s", projectID)},
		{Name: "IAM & Admin - IAM", URL: fmt.Sprintf("https://console.cloud.google.com/iam-admin/iam?project=%s", projectID)},
		{Name: "IAM & Admin - Service Accounts", URL: fmt.Sprintf("https://console.cloud.google.com/iam-admin/serviceaccounts?project=%s", projectID)},

		// Cloud Pub/Sub
		{Name: "Cloud Pub/Sub", URL: fmt.Sprintf("https://console.cloud.google.com/cloudpubsub?project=%s", projectID)},
		{Name: "Cloud Pub/Sub - Topics", URL: fmt.Sprintf("https://console.cloud.google.com/cloudpubsub/topic/list?project=%s", projectID)},
		{Name: "Cloud Pub/Sub - Subscriptions", URL: fmt.Sprintf("https://console.cloud.google.com/cloudpubsub/subscription/list?project=%s", projectID)},

		// Cloud Datastore
		{Name: "Cloud Datastore", URL: fmt.Sprintf("https://console.cloud.google.com/datastore?project=%s", projectID)},

		// Cloud Spanner
		{Name: "Cloud Spanner", URL: fmt.Sprintf("https://console.cloud.google.com/spanner?project=%s", projectID)},
		// Cloud Logging
		{Name: "Cloud Logging", URL: fmt.Sprintf("https://console.cloud.google.com/logs?project=%s", projectID)},
		{Name: "Cloud Logging - Logs Explorer", URL: fmt.Sprintf("https://console.cloud.google.com/logs/query?project=%s", projectID)},
		{Name: "Cloud Logging - Logs Router", URL: fmt.Sprintf("https://console.cloud.google.com/logs/router?project=%s", projectID)},
		{Name: "Cloud Logging - Metrics", URL: fmt.Sprintf("https://console.cloud.google.com/logs/metrics?project=%s", projectID)},

		// Cloud Monitoring
		{Name: "Cloud Monitoring", URL: fmt.Sprintf("https://console.cloud.google.com/monitoring?project=%s", projectID)},
		{Name: "Cloud Monitoring - Dashboards", URL: fmt.Sprintf("https://console.cloud.google.com/monitoring/dashboards?project=%s", projectID)},
		{Name: "Cloud Monitoring - Uptime Checks", URL: fmt.Sprintf("https://console.cloud.google.com/monitoring/uptime?project=%s", projectID)},
		{Name: "Cloud Monitoring - Alerting Policies", URL: fmt.Sprintf("https://console.cloud.google.com/monitoring/alerting?project=%s", projectID)},
		{Name: "Cloud Monitoring - Notification Channels", URL: fmt.Sprintf("https://console.cloud.google.com/monitoring/settings/notificationchannels?project=%s", projectID)},

		// Cloud Dataflow
		{Name: "Cloud Dataflow", URL: fmt.Sprintf("https://console.cloud.google.com/dataflow?project=%s", projectID)},

		// Cloud Build
		{Name: "Cloud Build", URL: fmt.Sprintf("https://console.cloud.google.com/cloud-build?project=%s", projectID)},
		{Name: "Cloud Build - History", URL: fmt.Sprintf("https://console.cloud.google.com/cloud-build/builds?project=%s", projectID)},
		{Name: "Cloud Build - Triggers", URL: fmt.Sprintf("https://console.cloud.google.com/cloud-build/triggers?project=%s", projectID)},

		// Cloud Scheduler
		{Name: "Cloud Scheduler", URL: fmt.Sprintf("https://console.cloud.google.com/cloudscheduler?project=%s", projectID)},

		// Cloud Tasks
		{Name: "Cloud Tasks", URL: fmt.Sprintf("https://console.cloud.google.com/cloudtasks?project=%s", projectID)},
		// Support
		{Name: "Support Cases", URL: fmt.Sprintf("https://console.cloud.google.com/support/cases?project=%s", projectID)},

		// Cloud Memorystore
		{Name: "Cloud Memorystore", URL: fmt.Sprintf("https://console.cloud.google.com/memorystore?project=%s", projectID)},

		// Cloud Firestore
		{Name: "Cloud Firestore", URL: fmt.Sprintf("https://console.cloud.google.com/firestore?project=%s", projectID)},

		// Cloud Armor
		{Name: "Cloud Armor", URL: fmt.Sprintf("https://console.cloud.google.com/netsec?project=%s", projectID)},

		// Secret Manager
		{Name: "Secret Manager", URL: fmt.Sprintf("https://console.cloud.google.com/security/secret-manager?project=%s", projectID)},

		// API Gateway
		{Name: "API Gateway", URL: fmt.Sprintf("https://console.cloud.google.com/apis/api-gateway?project=%s", projectID)},

		// Recommendations AI
		{Name: "Recommendations AI", URL: fmt.Sprintf("https://console.cloud.google.com/recommendations-ai?project=%s", projectID)},

		// Cloud Workflow
		{Name: "Cloud Workflow", URL: fmt.Sprintf("https://console.cloud.google.com/workflows?project=%s", projectID)},
	}

	selectedItem, err := selectItem(items)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	openBrowser(selectedItem.URL)
}

func selectItem(items []Item) (Item, error) {
	index, err := fuzzyfinder.Find(items, func(i int) string {
		return items[i].Name
	})

	if err != nil {
		return Item{}, err
	}

	return items[index], nil
}

func openBrowser(url string) {
	cmd := exec.Command("open", url)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
