package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var secretname string
var secretnamespace string

var clientset *kubernetes.Clientset
var cfgpath string
var version = "development" // will be overwritten by CI using ldflags
var getVersion bool

func main() {
	parseArgs()

	connectToCluster()

	printSecretValue()
}

func connectToCluster() {
	cfgpath = filepath.Join(os.Getenv("HOME"), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", cfgpath)
	if err != nil {
		panic(err.Error())
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func parseArgs() {
	flag.StringVarP(&secretnamespace, "namespace", "n", "", "The namespace of the secret. Defaults to the current namespace")
	flag.BoolVar(&getVersion, "version", false, "Print out the current version")

	flag.Parse()

	if getVersion == true {
		fmt.Println(version)
		os.Exit(0)
	}

	secretname = flag.Arg(0)
}

func printSecretValue() {
	if secretnamespace == "" {
		// Get the current namespace
		kubeconfig, err := ioutil.ReadFile(cfgpath)

		if err != nil {
			panic(err.Error())
		}

		clientconfig, err := clientcmd.NewClientConfigFromBytes(kubeconfig)

		if err != nil {
			panic(err.Error())
		}

		secretnamespace, _, err = clientconfig.Namespace()
		if err != nil {
			panic(err.Error())
		}
	}

	secret, err := clientset.CoreV1().Secrets(secretnamespace).Get(secretname, metav1.GetOptions{})

	if err != nil {
		panic(err.Error())
	}

	for key, value := range secret.Data {
		fmt.Printf(key + ": " + string(value) + "\n")
	}
}
