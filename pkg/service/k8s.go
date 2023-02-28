package service

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

func InitK8sClient() (*kubernetes.Clientset, error) {
	var kubeConfig string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeConfig = ""
	}
	// use the current context in kubeConfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Println("get kubectl config failed", err.Error())
		return nil, err
	}
	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("get k8s client failed: ", err.Error())
		return nil, err
	}
	return clientSet, nil
}

// Create Deployment
func CreateDeployment(client *kubernetes.Clientset, username, deploymentName string, container []corev1.Container) error {
	//get deployments by namespace
	deploymentsClient := client.AppsV1().Deployments(username)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println("get deployment list failed: ", err.Error())
		return err
	}
	namespaceExist := false
	for _, deployment := range list.Items {
		if deployment.Name == deploymentName {
			namespaceExist = true
			break
		}
	}
	if !namespaceExist {
		deployment := &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name: deploymentName,
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: int32Ptr(3),
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": deploymentName,
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app": deploymentName,
						},
					},
					Spec: corev1.PodSpec{
						Containers: container,
					},
				},
			},
		}
		_, err = deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
		if err != nil {
			log.Println("create deployment failed", err.Error())
			return err
		}
	}
	return nil
}

// create namespace
func CreateNamespace(client *kubernetes.Clientset, username string) error {
	namespaceList, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println("get namespace list failed: ", err.Error())
		return err
	}
	namespaceExist := false
	for _, ns := range namespaceList.Items {
		if ns.Name == username {
			namespaceExist = true
			break
		}
	}
	if !namespaceExist {
		namespace := corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: username,
			},
		}
		_, err = client.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
		if err != nil {
			log.Println("create namespace failed: ", err.Error())
			return err
		}
	}
	return nil
}

// create service
func CreateService(client *kubernetes.Clientset, username, serviceName string, ports []corev1.ServicePort) error {
	//get services by namespace
	serviceClient := client.CoreV1().Services(username)
	list, err := serviceClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println("get services list failed: ", err.Error())
		return err
	}
	serviceExist := false
	for _, service := range list.Items {
		if service.Name == serviceName {
			serviceExist = true
			break
		}
	}
	if !serviceExist {
		// Create the service spec
		serviceSpec := &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: serviceName,
			},
			Spec: corev1.ServiceSpec{
				Selector: map[string]string{"app": serviceName},
				Type:     corev1.ServiceTypeNodePort,
				Ports:    ports,
			},
		}
		_, err = serviceClient.Create(context.TODO(), serviceSpec, metav1.CreateOptions{})
		if err != nil {
			log.Println("create service failed: ", err.Error())
			return err
		}
	}
	return nil
}

func int32Ptr(i int32) *int32 { return &i }
