package k8s

import (
	"errors"
	"fmt"
	"strings"
)

type WebURL struct {
	baseURL string
}

func NewKubeSphereURL(baseURL string) (*WebURL, error) {
	if len(baseURL) == 0 {
		return nil, errors.New("baseURL is nil")
	}

	return &WebURL{
		baseURL: baseURL,
	}, nil
}

func (u *WebURL) Build() *WebURLBuilder {
	return &WebURLBuilder{
		baseURL: u.baseURL,
		builder: &strings.Builder{},
	}
}

type WebURLBuilder struct {
	baseURL string
	builder *strings.Builder
}

func (b *WebURLBuilder) append(name string, value string) *WebURLBuilder {
	b.builder.WriteString("/")
	b.builder.WriteString(name)
	b.builder.WriteString("/")
	b.builder.WriteString(value)
	return b
}

func (b *WebURLBuilder) Node(node string) *WebURLBuilder {
	return b.append("infrastructure/nodes", node)
}

func (b *WebURLBuilder) Namespace(namespace string) *WebURLBuilder {
	return b.append("projects", namespace)
}

func (b *WebURLBuilder) Volume(pvcName string) *WebURLBuilder {
	return b.append("volumes", pvcName)
}

func (b *WebURLBuilder) Deployment(deployName string) *WebURLBuilder {
	return b.append("deployments", deployName)
}

func (b *WebURLBuilder) StatefulSet(stsName string) *WebURLBuilder {
	return b.append("statefulsets", stsName)
}

func (b *WebURLBuilder) DaemonSet(dsName string) *WebURLBuilder {
	return b.append("daemonsets", dsName)
}

func (b *WebURLBuilder) Service(svcName string) *WebURLBuilder {
	return b.append("services", svcName)
}

func (b *WebURLBuilder) Pod(podName string) *WebURLBuilder {
	return b.append("pods", podName)
}

func (b *WebURLBuilder) Job(jobName string) *WebURLBuilder {
	return b.append("jobs", jobName)
}

func (b *WebURLBuilder) CronJob(cronJobName string) *WebURLBuilder {
	return b.append("cronjobs", cronJobName)
}

func (b *WebURLBuilder) Workload(kind string, value string) *WebURLBuilder {
	switch kind {
	case KindDeployment:
		return b.Deployment(value)
	case KindStatefulSet:
		return b.StatefulSet(value)
	case KindDaemonSet:
		return b.DaemonSet(value)
	case KindService:
		return b.Service(value)
	case KindPod:
		return b.Pod(value)
	case KindJob:
		return b.Job(value)
	case KindCronJob:
		return b.CronJob(value)
	default:
		return b
	}
}

func (b *WebURLBuilder) Reset() {
	b.builder.Reset()
}

func (b *WebURLBuilder) String() string {
	return fmt.Sprintf("%s%s", b.baseURL, b.builder.String())
}
