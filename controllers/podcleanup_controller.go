/*
Copyright 2023.

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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	podcleanerv1 "github.com/example/k8s-pod-cleaner/api/v1"
)

// PodCleanupReconciler reconciles a PodCleanup object
type PodCleanupReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=podcleaner.example.com,resources=podcleanups,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=podcleaner.example.com,resources=podcleanups/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=podcleaner.example.com,resources=podcleanups/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=pods,verbs=list;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PodCleanup object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
// Reconcile cleans up pods not in running status
func (r *PodCleanupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx).WithValues("podcleanup", req.NamespacedName)
	log.Info("Reconciling At")

	// Get the PodCleanup instance
	podCleanup := &podcleanerv1.PodCleanup{}
	err := r.Get(ctx, req.NamespacedName, podCleanup)
	if err != nil {
		log.Error(err, "unable to fetch PodCleanup")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Get all the Pods in the same namespace as the PodCleanup instance
	podList := &corev1.PodList{}
	listOpts := []client.ListOption{
		client.InNamespace(req.Namespace),
	}
	if err := r.List(ctx, podList, listOpts...); err != nil {
		log.Error(err, "unable to list Pods")
		return ctrl.Result{}, err
	}

	// Iterate over the Pods and delete the ones not in running status
	for _, pod := range podList.Items {
		if pod.Status.Phase != corev1.PodRunning {
			log.Info("Deleting pod", "name", pod.Name)
			err := r.Delete(ctx, &pod)
			if err != nil {
				log.Error(err, "unable to delete Pod", "name", pod.Name)
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodCleanupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&podcleanerv1.PodCleanup{}).
		Complete(r)
}
