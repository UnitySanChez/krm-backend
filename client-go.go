//项目的总路口

package main

import (
	"context"
	"encoding/json"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 删除一个资源
func delPod(clientset *kubernetes.Clientset, namespace, name string) error {
	err := clientset.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	return err
}

// 创建资源
func createNamespace(clientset *kubernetes.Clientset, namespace string) error {
	var newNamespace corev1.Namespace
	newNamespace.Name = namespace
	//newNamespace.ObjectMeta.Name = "teee"
	_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &newNamespace, metav1.CreateOptions{})
	return err
}

func main() {
	//1.初始化config实例
	config, err := clientcmd.BuildConfigFromFlags("", "./config/metakubeconfig") //这里masterUrl为空，因为地址已经在metakubeconfig中配置了
	if err != nil {
		panic(err.Error()) //失败了，后续基本上都无法执行了，所以要panic掉
	}
	//2.创建客户端工具，clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//3.操作集群
	/*pods, err := clientset.CoreV1().Pods("jenkins-k8s").List(context.TODO(), metav1.ListOptions{}) //context.TODO()固定写法 ，  metav1.ListOptions{}查询条件
	//kubectl api-resources 可以查看所有资源类型 CoreV1() 就是v1
	//kubectl api-resources |grep ingress 可以查到ingress是 networking.k8s.io/v1 ，所以就要用下面的那个
	//clientset.NetworkingV1()

	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "List pods failed")
	} else {
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		pod0 := pods.Items[0]
		fmt.Printf("pod0: %v\n", pod0.Spec.Containers) //kubectl -n jenkins-k8s get pod ubuntu-deployment-fb5f54689-8d4s4 -o json 就是这里面展示的东西
	}
	deploy, _ := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	for _, dep := range deploy.Items {
		fmt.Printf("当前资源名字是: %s，namespace是 %s\n", dep.Name, dep.Namespace)
	}
	//namespaces, _ := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})  这都同理

	//GET方法,get的参数比list多了一个name
	detal, err := clientset.CoreV1().Pods("jenkins-k8s").Get(context.TODO(), "jenkins-7c74fc8f59-b7256", metav1.GetOptions{})
	if err != nil {
		fmt.Println("查询失败")
	} else {
		//fmt.Printf("查询成功，pod信息是: %v\n", detal)
		fmt.Printf("pod第一个镜像是：%s\n", detal.Spec.Containers[0].Image)
	}
	nsdetal, _ := clientset.CoreV1().Namespaces().Get(context.TODO(), "jenkins-k8s", metav1.GetOptions{})
	fmt.Println("namespace信息是:", nsdetal)

	//更新操作
	//获取deploy并修改
	deploydetal, _ := clientset.AppsV1().Deployments("dev").Get(context.TODO(), "hello-app", metav1.GetOptions{})
	fmt.Println("查询的名字是:", deploydetal.Name)
	//获取当前label的值，如果label是空会报空指针，需要先初始化一下
	if deploydetal.Labels == nil {
		deploydetal.Labels = make(map[string]string)
	}
	//我通过创建新的label进行修改，还用赋值给deploydetal吗。不需要，他是map类型的
	deploydetal.Labels["newlabel"] = "newlabelvalue"

	//修改副本数
	newReplicas := int32(3)
	deploydetal.Spec.Replicas = &newReplicas
	//修改其他的同理
	_, err = clientset.AppsV1().Deployments("dev").Update(context.TODO(), deploydetal, metav1.UpdateOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "Update deployment failed")
	} else {
		fmt.Println("label,副本更新成功")
	}*/
	//删除一个资源,除了pod外，其余同理
	/*	err = delPod(clientset, "dev", "hello-app-5567884fbb-gdz7z")
		if err != nil {
			logs.Error(map[string]interface{}{"error": err.Error()}, "Delete pod failed")
		}*/
	//创建资源namespace
	/*err = createNamespace(clientset, "dev")
	//创建deploy
	var newDeploy appsv1.Deployment
	newDeploy.Name = "test"
	newDeploy.Namespace = "dev"
	label := map[string]string{}
	label["app"] = "test"
	//此处是deploy的label
	newDeploy.Labels = label
	//此处是deploy下select的label，这里面Selector是一个指针类型，要先初始话，否则会报错，指针类型都有这个问题
	if newDeploy.Spec.Selector == nil {
		newDeploy.Spec.Selector = &metav1.LabelSelector{}
	}
	newDeploy.Spec.Selector.MatchLabels = label
	//此处是要被select选中的pod的标签
	newDeploy.Spec.Template.Labels = label

	//先声明一个容器
	var container corev1.Container
	container.Name = "nginx"
	container.Image = "nginx"
	newDeploy.Spec.Template.Spec.Containers = []corev1.Container{container} // []类型{值} 初始化一个切片，所以可以这面写直接赋予Containers字段
	clientset.AppsV1().Deployments("dev").Create(context.TODO(), &newDeploy, metav1.CreateOptions{})
	//kubectl create deployment nginx --image=nginx  --dry-run=client -oyaml  生成一个创建模板，通过他可以查看具体要什么东西

	*/

	//通过json创建k8s资源
	//kubectl create deployment nginx --image=nginx  --dry-run=client -ojson
	//常用这种方法，前端传来一个json串，然后我们绑定结构体
	deployJson := `{
    "kind": "Deployment",
    "apiVersion": "apps/v1",
    "metadata": {
        "name": "nginx",
        "labels": {
            "app": "nginx"
        }
    },
    "spec": {
        "replicas": 1,
        "selector": {
            "matchLabels": {
                "app": "nginx"
            }
        },
        "template": {
            "metadata": {
                "labels": {
                    "app": "nginx"
                }
            },
            "spec": {
                "containers": [
                    {
                        "name": "nginx",
                        "image": "nginx",
                        "resources": {}
                    }
                ]
            }
        },
        "strategy": {}
    },
    "status": {}
}
`

	var newDeploy appsv1.Deployment
	//把json转换成结构体
	json.Unmarshal([]byte(deployJson), &newDeploy)
	if err != nil {
		fmt.Println("json转换失败", err)
	}
	fmt.Println("json转换成功,详情为", newDeploy)
	clientset.AppsV1().Deployments("dev").Create(context.TODO(), &newDeploy, metav1.CreateOptions{})
}
