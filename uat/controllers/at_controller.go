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
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"time"
	"uat/pkg/schedule"
	"uat/pkg/spawn"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	uatv1alpha1 "uat/api/v1alpha1"
)

// AtReconciler reconciles a At object
type AtReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=uat.jerseymec.dev,resources=ats,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=uat.jerseymec.dev,resources=ats/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=uat.jerseymec.dev,resources=ats/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the At object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *AtReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	reqLogger := r.Log.WithValues("at", req.NamespacedName)
	reqLogger.Info("=== Reconciling At")
	instance := &uatv1alpha1.At{}
	err := r.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			//object not found, could have been deleted after reconcile request, hence dont requeue
			return ctrl.Result{}, nil
		}
		//error reading the object, requeue the request
		return ctrl.Result{}, err
	}
	//if no phase set , default to Pending
	if instance.Status.Phase == "" {
		instance.Status.Phase = uatv1alpha1.PhasePending
	}
	//state transition PENDING -> RUNNING -> DONE
	switch instance.Status.Phase {
	case uatv1alpha1.PhasePending:
		reqLogger.Info("Phase: PENDING")
		diff, err := schedule.TimeUntillSchedule(instance.Spec.Schedule)
		if err != nil {
			reqLogger.Error(err, "Schedule parsing failure")
			return ctrl.Result{}, err
		}
		reqLogger.Info("Schedule parsing done", "Results: ", fmt.Sprintf("%v", diff))
		if diff > 0 {
			// not yet time to execute, wait until schedule time
			return ctrl.Result{RequeueAfter: diff * time.Second}, nil
		}
		reqLogger.Info("It's time!", "Ready to execute", instance.Spec.Command)
		//change state
		instance.Status.Phase = uatv1alpha1.PhaseRunning

	case uatv1alpha1.PhaseRunning:
		reqLogger.Info("Phase: RUNNING")
		pod := spawn.NewPodForCR(instance)
		err := ctrl.SetControllerReference(instance, pod, r.Scheme)
		if err != nil {
			//requeue with error
			return ctrl.Result{}, err
		}
		query := &corev1.Pod{}
		//check if pod exists
		err = r.Get(context.TODO(), req.NamespacedName, query)
		if err != nil && errors.IsNotFound(err) {
			// pod does not exist so create one
			err = r.Create(context.TODO(), pod)
			if err != nil {
				return ctrl.Result{}, err
			}
			// successfully created a pod
			reqLogger.Info("Pod Create successfully", "name", pod.Name)
			return ctrl.Result{}, nil
		} else if err != nil {
			//requeue with err
			reqLogger.Error(err, "cannot create pod")
			return ctrl.Result{}, err
		} else if query.Status.Phase == corev1.PodFailed || query.Status.Phase == corev1.PodSucceeded {
			//pod already finsihed or errored out
			reqLogger.Info("Container terminated", "reason", query.Status.Reason, "message", query.Status.Message)
			instance.Status.Phase = uatv1alpha1.PhaseDone
		} else {
			// will automatically requeue when pod status changes
			return ctrl.Result{}, nil
		}
	case uatv1alpha1.PhaseDone:
		reqLogger.Info("Phase: DONE")
		//reconcile without requeueing
		return ctrl.Result{}, nil
	default:
		reqLogger.Info("NOP")
		return ctrl.Result{}, nil

	}
	//update status
	err = r.Status().Update(context.TODO(), instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AtReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&uatv1alpha1.At{}).
		Owns(&corev1.Pod{}).
		Complete(r)
}
