package k8s

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"time"
)

const (
	KindDeployment  = "Deployment"
	KindStatefulSet = "StatefulSet"
	KindDaemonSet   = "DaemonSet"
	KindPod         = "Pod"
	KindService     = "Service"
	KindCronJob     = "CronJob"
	KindJob         = "Job"
)

type K8S interface {
	GetNode(nodeName string) (*v1.Node, error)
	GetNodeByIP(nodeIP string) (*v1.Node, error)
	GetPod(namespace string, podName string) (*v1.Pod, error)
	DeletePod(namespace string, podName string) error
	GetPodList(namespace string, labels map[string]string) (*v1.PodList, error)
	GetDeployment(namespace string, deployName string) (*appsv1.Deployment, error)
	DeleteDeployment(namespace string, deployName string) error
	GetDeploymentByPod(pod *v1.Pod) (*appsv1.Deployment, error)
	GetDeploymentList(namespace string, labels map[string]string) (*appsv1.DeploymentList, error)
	GetHPA(namespace string, hpaName string) (*autoscalingv1.HorizontalPodAutoscaler, error)
	GetStatefulSet(namespace string, stsName string) (*appsv1.StatefulSet, error)
	DeleteStatefulSet(namespace string, stsName string) error
	GetStatefulSetList(namespace string, labels map[string]string) (*appsv1.StatefulSetList, error)
	GetDaemonSet(namespace string, dsName string) (*appsv1.DaemonSet, error)
	GetDaemonSetList(namespace string, labels map[string]string) (*appsv1.DaemonSetList, error)
	GetCronJob(namespace string, cronJobName string) (*batchv1.CronJob, error)
	DeleteCronJob(namespace string, cronJobName string) error
	GetCronJobList(namespace string, labels map[string]string) (*batchv1.CronJobList, error)
	GetJob(namespace string, jobName string) (*batchv1.Job, error)
	DeleteJob(namespace string, jobName string) error
	GetJobList(namespace string, labels map[string]string) (*batchv1.JobList, error)
	GetPvc(namespace string, pvcName string) (*v1.PersistentVolumeClaim, error)
	PodExecCmd(pod *v1.Pod, command string) (string, error)
	RestartDeployment(namespace string, deployName string) error
	WithContext(ctx context.Context) *Client
	K8S() *kubernetes.Clientset
}

type Client struct {
	k8s     *kubernetes.Clientset
	config  *rest.Config
	context context.Context
	webURL  *WebURL
}

func getKubeConfig(conf options) string {
	var path string
	if conf.KubeConfigPath != "" {
		path = conf.KubeConfigPath
	} else {
		defaultPath := filepath.Join(".kube", "config")
		if home := homedir.HomeDir(); home != "" {
			defaultPath = filepath.Join(home, ".kube", "config")
		}
		path = defaultPath
	}
	ret := flag.String("kubeConfig", path, "(optional) absolute path to the kubeConfig file")
	flag.Parse()
	return *ret
}

func getRestConfig(conf options) (*rest.Config, error) {
	if conf.OutOfCluster {
		kubeConfig := getKubeConfig(conf)
		c, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return nil, err
		}

		return c, nil
	} else {
		c, err := rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("building client config error: %s", err)
		}
		return c, nil
	}
}

func New(opts ...Option) (*Client, error) {
	conf := setDefault()
	for _, o := range opts {
		if o != nil {
			o(&conf)
		}
	}
	config, err := getRestConfig(conf)
	if err != nil {
		return nil, fmt.Errorf("get k8s client config error: %s", err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("creating k8s client error: %s", err)
	}

	ksURL, err := NewKubeSphereURL(conf.WebBaseURL)
	if err != nil {

	}

	return &Client{
		k8s:     client,
		config:  config,
		context: context.TODO(),
		webURL:  ksURL,
	}, nil
}

func (c *Client) WebURL() *WebURLBuilder {
	return c.webURL.Build()
}

type PodStatus struct {
	Phase        string
	RestartCount int32
}

// labelsToString 将标签的 map 转化为 Kubernetes 兼容的标签选择器字符串
func labelsToString(labels map[string]string) string {
	labelSelector := ""
	for key, value := range labels {
		if labelSelector != "" {
			labelSelector += ","
		}
		labelSelector += key + "=" + value
	}
	return labelSelector
}

func (s *PodStatus) IsRunning() bool {
	return s.Phase == string(v1.PodRunning)
}

func (c *Client) GetNode(nodeName string) (*v1.Node, error) {
	return c.k8s.CoreV1().Nodes().Get(c.context, nodeName, metav1.GetOptions{})
}

func (c *Client) GetNodeByIP(nodeIP string) (*v1.Node, error) {
	nodes, err := c.k8s.CoreV1().Nodes().List(c.context, metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	for _, node := range nodes.Items {
		for _, addr := range node.Status.Addresses {
			if addr.Type == v1.NodeInternalIP && addr.Address == nodeIP {
				return &node, nil
			}
		}
	}

	return nil, fmt.Errorf("node(ip=%s) not found", nodeIP)
}

func (c *Client) GetJob(namespace string, jobName string) (*batchv1.Job, error) {
	return c.k8s.BatchV1().Jobs(namespace).Get(c.context, jobName, metav1.GetOptions{})
}

func (c *Client) DeleteJob(namespace string, jobName string) error {
	return c.k8s.BatchV1().Jobs(namespace).Delete(c.context, jobName, metav1.DeleteOptions{})
}

func (c *Client) GetJobList(namespace string, labels map[string]string) (*batchv1.JobList, error) {
	return c.k8s.BatchV1().Jobs(namespace).List(c.context, metav1.ListOptions{LabelSelector: labelsToString(labels)})
}

func (c *Client) GetCronJob(namespace string, cronJobName string) (*batchv1.CronJob, error) {
	return c.k8s.BatchV1().CronJobs(namespace).Get(c.context, cronJobName, metav1.GetOptions{})
}

func (c *Client) DeleteCronJob(namespace string, cronJobName string) error {
	return c.k8s.BatchV1().CronJobs(namespace).Delete(c.context, cronJobName, metav1.DeleteOptions{})
}

func (c *Client) GetCronJobList(namespace string, labels map[string]string) (*batchv1.CronJobList, error) {
	return c.k8s.BatchV1().CronJobs(namespace).List(c.context, metav1.ListOptions{LabelSelector: labelsToString(labels)})
}

func (c *Client) GetCronJobByJob(job *batchv1.Job) (*batchv1.CronJob, error) {
	for _, v := range job.OwnerReferences {
		if v.Kind == "CronJob" {
			cronJob, err := c.GetCronJob(v.Name, job.Namespace)

			if err != nil {
				return nil, err
			}

			return cronJob, nil
		}
	}

	return nil, fmt.Errorf("cronjob not found by job: %s/%s", job.Namespace, job.Name)
}

func (c *Client) GetDaemonSet(namespace string, dsName string) (*appsv1.DaemonSet, error) {
	return c.k8s.AppsV1().DaemonSets(namespace).Get(c.context, dsName, metav1.GetOptions{})
}

func (c *Client) GetDaemonSetList(namespace string, labels map[string]string) (*appsv1.DaemonSetList, error) {
	return c.k8s.AppsV1().DaemonSets(namespace).List(c.context, metav1.ListOptions{LabelSelector: labelsToString(labels)})
}

func (c *Client) GetStatefulSet(namespace string, stsName string) (*appsv1.StatefulSet, error) {
	return c.k8s.AppsV1().StatefulSets(namespace).Get(c.context, stsName, metav1.GetOptions{})
}

func (c *Client) DeleteStatefulSet(namespace string, stsName string) error {
	return c.k8s.AppsV1().StatefulSets(namespace).Delete(c.context, stsName, metav1.DeleteOptions{})
}

func (c *Client) GetStatefulSetList(namespace string, labels map[string]string) (*appsv1.StatefulSetList, error) {
	return c.k8s.AppsV1().StatefulSets(namespace).List(c.context, metav1.ListOptions{LabelSelector: labelsToString(labels)})
}

func (c *Client) GetPod(namespace string, podName string) (*v1.Pod, error) {
	return c.k8s.CoreV1().Pods(namespace).Get(c.context, podName, metav1.GetOptions{})
}

func (c *Client) DeletePod(namespace string, podName string) error {
	return c.k8s.CoreV1().Pods(namespace).Delete(c.context, podName, metav1.DeleteOptions{})
}

func (c *Client) GetPodList(namespace string, labels map[string]string) (*v1.PodList, error) {
	return c.k8s.CoreV1().Pods(namespace).List(c.context, metav1.ListOptions{LabelSelector: labelsToString(labels)})
}

func (c *Client) GetDeployment(namespace string, deployName string) (*appsv1.Deployment, error) {
	return c.k8s.AppsV1().Deployments(namespace).Get(c.context, deployName, metav1.GetOptions{})
}

func (c *Client) DeleteDeployment(namespace string, deployName string) error {
	return c.k8s.AppsV1().Deployments(namespace).Delete(c.context, deployName, metav1.DeleteOptions{})
}

func (c *Client) GetDeploymentList(namespace string, labels map[string]string) (*appsv1.DeploymentList, error) {
	return c.k8s.AppsV1().Deployments(namespace).List(c.context, metav1.ListOptions{LabelSelector: labelsToString(labels)})
}

func (c *Client) GetDeploymentByPod(pod *v1.Pod) (*appsv1.Deployment, error) {
	for _, v := range pod.OwnerReferences {
		if v.Kind == "ReplicaSet" {
			rs, err := c.k8s.AppsV1().ReplicaSets(pod.Namespace).Get(c.context, v.Name, metav1.GetOptions{})

			if err != nil {
				return nil, fmt.Errorf("get deployment by pod: %s/%s error:%s", pod.Namespace, pod.Name, err)
			}

			for _, d := range rs.OwnerReferences {
				if d.Kind == "Deployment" {
					deploy, err := c.GetDeployment(pod.Namespace, d.Name)

					if err != nil {
						return nil, fmt.Errorf("get deployment by pod: %s/%s error:%s", pod.Namespace, pod.Name, err)
					}

					return deploy, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("deployment not found by pod: %s/%s", pod.Namespace, pod.Name)
}

func (c *Client) GetHPA(namespace string, hpaName string) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	return c.k8s.AutoscalingV1().HorizontalPodAutoscalers(namespace).Get(c.context, hpaName, metav1.GetOptions{})
}

func (c *Client) GetPvc(namespace string, pvcName string) (*v1.PersistentVolumeClaim, error) {
	return c.k8s.CoreV1().PersistentVolumeClaims(namespace).Get(c.context, pvcName, metav1.GetOptions{})
}

func (c *Client) PodExecCmd(pod *v1.Pod, command string) (string, error) {
	request := c.k8s.CoreV1().RESTClient().Post().Resource("pods").Namespace(pod.Namespace).Name(pod.Name).SubResource("exec")
	dumpCmds := []string{
		"bash",
		"-c",
		command,
	}
	execOptions := &v1.PodExecOptions{
		Command: dumpCmds,
		Stdin:   false,
		Stdout:  true,
		Stderr:  true,
		TTY:     true,
	}
	request.VersionedParams(execOptions, scheme.ParameterCodec)
	var stdout, stderr bytes.Buffer
	executor, err := remotecommand.NewSPDYExecutor(c.config, "POST", request.URL())
	if err != nil {
		return "", err
	}

	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		return "", err
	}

	if errStr := stderr.String(); errStr != "" {
		return "", fmt.Errorf(errStr)
	}

	return stdout.String(), nil
}

func (c *Client) RestartDeployment(namespace string, deployName string) error {
	deploy := c.k8s.AppsV1().Deployments(namespace)

	json := fmt.Sprintf(`
	{
		"spec":
		{
			"template":
			{
				"metadata":
				{
					"annotations":
					{
						"mon.benlai.cloud/restartedAt": "%s"
					}
				}
			}
		}
	}`, time.Now().String())

	if _, err := deploy.Patch(c.context, deployName, types.StrategicMergePatchType, []byte(json), metav1.PatchOptions{}); err != nil {
		return err
	}

	return nil
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.context = ctx
	return c
}

func (c *Client) K8S() *kubernetes.Clientset {
	return c.k8s
}
