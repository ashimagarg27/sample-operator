/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	multiplev1alpha1 "github.com/ashimagarg27/sample-operator/api/v1alpha1"
)

// LabelOperatorReconciler reconciles a LabelOperator object
type LabelOperatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=multiple.example.com,resources=labeloperators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=multiple.example.com,resources=labeloperators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=multiple.example.com,resources=labeloperators/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the LabelOperator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile
func (r *LabelOperatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("In LabelOperator......")

	// your logic here
	instance := &multiplev1alpha1.LabelOperator{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if k8serr.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			logger.Info("LabelOperator resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object.
		logger.Error(err, "failed to get LabelOperator resource")
		return ctrl.Result{}, err
	}

	var pod corev1.Pod
	if err := r.Get(ctx, req.NamespacedName, &pod); err != nil {
		if k8serr.IsNotFound(err) {
			logger.Info("Pod not found")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "unable to fetch Pod")
		return ctrl.Result{}, err
	}

	label := instance.Spec.Label

	if pod.Labels == nil {
		pod.Labels = make(map[string]string)
	}

	pod.Labels["updatedLabel/Operator"] = label
	logger.Info("Label Added in Pod", pod.Name, label)

	if err := r.Update(ctx, &pod); err != nil {
		if k8serr.IsNotFound(err) {
			logger.Info("Pod not found...Retrying...")
			return ctrl.Result{Requeue: true}, nil
		}
		logger.Error(err, "unable to update Pod")
		return ctrl.Result{}, err
	}
	logger.Info("Pod Lable Added....")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LabelOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&multiplev1alpha1.LabelOperator{}).
		Complete(r)
}
