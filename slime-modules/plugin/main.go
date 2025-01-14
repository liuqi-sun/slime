/*


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

package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	ctrl "sigs.k8s.io/controller-runtime"
	"slime.io/slime/slime-framework/apis/networking/v1alpha3"
	"slime.io/slime/slime-framework/bootstrap"
	"slime.io/slime/slime-framework/util"
	microserviceslimeiov1alpha1 "slime.io/slime/slime-modules/plugin/api/v1alpha1"
	"slime.io/slime/slime-modules/plugin/controllers"
	// +kubebuilder:scaffold:imports
)

var scheme = runtime.NewScheme()

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	_ = microserviceslimeiov1alpha1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
	_ = v1alpha3.AddToScheme(scheme)
}

func main() {
	util.SetLog()

	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		Port:               9443,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "9487b5c0.my.domain",
	})
	if err != nil {
		log.Errorf("unable to start manager,%+v", err)
		os.Exit(1)
	}

	if err = (&controllers.PluginManagerReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		log.Errorf("unable to create pluginManager controller, %+v", err)
		os.Exit(1)
	}
	if err = (&controllers.EnvoyPluginReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		log.Errorf("unable to create EnvoyPlugin controller, %+v", err)
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	go bootstrap.HealthCheckStart()

	log.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		log.Errorf("problem running manager,%+v", err)
		os.Exit(1)
	}
}
