//go:build k8srequired
// +build k8srequired

package basic

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestBasic ensures that there is a ready azuredisk-csi-driver-app daemonset.
func TestBasic(t *testing.T) {
	ctx := context.Background()
	var err error

	// Check csi-azuredisk-node deamonset is ready.
	err = checkReadyDaemonset("csi-azuredisk-node", ctx)
	if err != nil {
		t.Fatalf("could not get azuredisk-csi-driver-app: %v", err)
	}

	// Check csi-azuredisk-controller deployment is ready.
	err = checkReadyDeployment("csi-azuredisk-controller", ctx)
	if err != nil {
		t.Fatalf("could not get azuredisk-csi-driver-app: %v", err)
	}

	// Check csi-snapshot-controller deployment is ready.
	err = checkReadyDeployment("csi-snapshot-controller", ctx)
	if err != nil {
		t.Fatalf("could not get azuredisk-csi-driver-app: %v", err)
	}
}

func checkReadyDeployment(name string, ctx context.Context) error {
	var err error

	l.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("waiting for deployment %s to be ready", name))

	o := func() error {
		selector := fmt.Sprintf("%s=%s", "app", name)
		lo := metav1.ListOptions{
			LabelSelector: selector,
		}

		deployments, err := appTest.K8sClient().AppsV1().Deployments(testNamespace).List(ctx, lo)
		if err != nil {
			return microerror.Mask(err)
		} else if len(deployments.Items) == 0 {
			return microerror.Maskf(executionFailedError, "deployment with label %#q in %#q not found", selector, testNamespace)
		}

		ds := deployments.Items[0]

		if ds.Status.ReadyReplicas != *ds.Spec.Replicas {
			return microerror.Maskf(executionFailedError, "deployment %#q want %d replicas %d ready", ds.Name, ds.Status.ReadyReplicas, *ds.Spec.Replicas)
		}

		return nil
	}
	b := backoff.NewConstant(2*time.Minute, 5*time.Second)
	n := backoff.NewNotifier(l, ctx)

	err = backoff.RetryNotify(o, b, n)
	if err != nil {
		return microerror.Mask(err)
	}

	l.LogCtx(ctx, "level", "debug", "message", "daemonset is ready")

	return nil
}

func checkReadyDaemonset(dsName string, ctx context.Context) error {
	var err error

	l.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("waiting for daemonset %s to be ready", dsName))

	o := func() error {
		selector := fmt.Sprintf("%s=%s", "app", dsName)
		lo := metav1.ListOptions{
			LabelSelector: selector,
		}

		daemonsets, err := appTest.K8sClient().AppsV1().DaemonSets(testNamespace).List(ctx, lo)
		if err != nil {
			return microerror.Mask(err)
		} else if len(daemonsets.Items) == 0 {
			return microerror.Maskf(executionFailedError, "daemonset with label %#q in %#q not found", selector, testNamespace)
		}

		ds := daemonsets.Items[0]

		if ds.Status.NumberReady != ds.Status.DesiredNumberScheduled {
			return microerror.Maskf(executionFailedError, "daemonset %#q want %d replicas %d ready", ds.Name, ds.Status.DesiredNumberScheduled, ds.Status.NumberReady)
		}

		return nil
	}
	b := backoff.NewConstant(2*time.Minute, 5*time.Second)
	n := backoff.NewNotifier(l, ctx)

	err = backoff.RetryNotify(o, b, n)
	if err != nil {
		return microerror.Mask(err)
	}

	l.LogCtx(ctx, "level", "debug", "message", "daemonset is ready")

	return nil
}
