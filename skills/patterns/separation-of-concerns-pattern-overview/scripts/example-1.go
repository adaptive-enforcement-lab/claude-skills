// Bad: CLI, business logic, and I/O mixed together
func DeployCommand(cmd *cobra.Command, args []string) error {
    // Parsing flags (CLI concern)
    namespace, _ := cmd.Flags().GetString("namespace")
    image, _ := cmd.Flags().GetString("image")

    // Validation (business logic)
    if namespace == "" {
        return fmt.Errorf("namespace required")
    }

    // Kubernetes client creation (infrastructure)
    config, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)
    clientset, _ := kubernetes.NewForConfig(config)

    // Deployment logic (business logic)
    deployment := &appsv1.Deployment{
        ObjectMeta: metav1.ObjectMeta{Name: "app"},
        Spec: appsv1.DeploymentSpec{
            Template: corev1.PodTemplateSpec{
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{{
                        Name:  "app",
                        Image: image,
                    }},
                },
            },
        },
    }

    // API call (infrastructure)
    _, err := clientset.AppsV1().Deployments(namespace).Create(ctx, deployment, metav1.CreateOptions{})
    return err
}