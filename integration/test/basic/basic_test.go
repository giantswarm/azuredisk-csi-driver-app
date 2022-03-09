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

// TestBasic ensures that there is a ready azuredisk-csi-app daemonset.
func TestBasic(t *testing.T) {
	ctx := context.Background()
	var err error

	// Check azuredisk-csi-app deamonset is ready.
	err = checkReadyDaemonset(ctx)
	if err != nil {
		t.Fatalf("could not get azuredisk-csi-app: %v", err)
	}
}

func checkReadyDaemonset(ctx context.Context) error {
	var err error

	l.LogCtx(ctx, "level", "debug", "message", "waiting for ready daemonset")

	o := func() error {
		selector := fmt.Sprintf("%s=%s", "app", app)
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
