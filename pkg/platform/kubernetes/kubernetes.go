package kubernetes

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/platform/platforming"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Kubernetes struct {
	Client    *kubernetes.Clientset
	Namespace string
}

func NewKubernetes(cnf platforming.Config) (d *Kubernetes, err error) {
	c, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	return &Kubernetes{Client: clientset, Namespace: cnf.Kubernetes.Namespace}, nil
}

func (k *Kubernetes) DeployWorkers(ctx context.Context, info worker.Info) error {
	name := info.Label + "-" + info.Topic
	labels := map[string]string{"scoretrak_worker": info.Label}
	path, err := util.GenerateConfigFile(info)
	if err != nil {
		return err
	}
	cEnc, err := util.EncodeConfigFile(path)
	if err != nil {
		return err
	}
	_, err = k.Client.AppsV1().DaemonSets(k.Namespace).Create(ctx,
		&appv1.DaemonSet{ObjectMeta: metav1.ObjectMeta{
			Name: name},
			Spec: appv1.DaemonSetSpec{
				Selector: &metav1.LabelSelector{MatchLabels: labels},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Name:   name,
						Labels: labels,
					},
					Spec: corev1.PodSpec{
						Tolerations:  []corev1.Toleration{{Key: "node-role.kubernetes.io/master", Effect: corev1.TaintEffectPreferNoSchedule}},
						NodeSelector: labels,
						Containers: []corev1.Container{
							{
								Name:    "worker",
								Image:   util.Image,
								Command: []string{"./worker", "-encoded-config", cEnc},
							},
						},
					},
				},
			},
		}, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (k *Kubernetes) RemoveWorkers(ctx context.Context, info worker.Info) error {
	return k.Client.AppsV1().DaemonSets(k.Namespace).Delete(ctx, info.Label+"-"+info.Topic, metav1.DeleteOptions{})
}
