# Hands-On with Kubernetes Secrets

## Introduction

On [Day 40](day40.md) we gave an introduction to Kubernetes secrets after dabbling a little with them on [Day 39](day39.md)., we discussed the importance of Kubernetes secrets and how they help secure sensitive information. Now, let's get hands-on with some practical scenarios to demonstrate how you can create and use secrets in your Kubernetes environment.

## Scenario 1: Creating and Using a Simple Opaque Secret

For these scenarios we are going to need that minikube cluster again, You can see those instructions on [Day 39](day39.md) to get a cluster up and running. 

### Step-by-Step Guide:

1. **Create a Secret:**
    ```bash
    kubectl create secret generic my-secret --from-literal=username=myuser --from-literal=password=mypassword
    ```

2. **Access the Secret in a Pod:**
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: secret-test
    spec:
      containers:
      - name: mycontainer
        image: nginx
        env:
        - name: USERNAME
          valueFrom:
            secretKeyRef:
              name: my-secret
              key: username
        - name: PASSWORD
          valueFrom:
            secretKeyRef:
              name: my-secret
              key: password
      restartPolicy: Never
    ```

3. **Deploy the Pod:**
    ```bash
    kubectl apply -f secret-pod.yaml
    ```

## Scenario 2: Using Secrets for TLS Certificates

### Step-by-Step Guide:

1. **Create a TLS Secret:**
    ```bash
    kubectl create secret tls tls-secret --cert=path/to/tls.crt --key=path/to/tls.key
    ```

2. **Configure a Pod to Use the TLS Secret:**
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: tls-pod
    spec:
      containers:
      - name: mycontainer
        image: nginx
        volumeMounts:
        - name: tls-volume
          mountPath: "/etc/nginx/ssl"
          readOnly: true
      volumes:
      - name: tls-volume
        secret:
          secretName: tls-secret
      restartPolicy: Never
    ```

3. **Deploy the Pod:**
    ```bash
    kubectl apply -f tls-pod.yaml
    ```

## Scenario 3: Managing Secrets with Environment Variables

### Step-by-Step Guide:

1. **Create a Secret:**
    ```bash
    kubectl create secret generic db-secret --from-literal=db_username=dbuser --from-literal=db_password=dbpass
    ```

2. **Use the Secret as Environment Variables in a Pod:**
    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: env-pod
    spec:
      containers:
      - name: mycontainer
        image: mysql
        env:
        - name: DB_USERNAME
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: db_username
        - name: DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: db_password
      restartPolicy: Never
    ```

3. **Deploy the Pod:**
    ```bash
    kubectl apply -f env-pod.yaml
    ```

## Best Practices for Managing Kubernetes Secrets

1. **Regularly Rotate Secrets**: Change your secrets periodically to reduce the risk of compromise.
2. **Use RBAC to Control Access**: Restrict access to secrets using Kubernetes Role-Based Access Control (RBAC).
3. **Enable Encryption at Rest**: Ensure your etcd database is configured to encrypt secrets at rest for added security.

## Conclusion

In this post, we've explored practical scenarios for creating and using Kubernetes secrets. By following these steps, you can securely manage sensitive information in your Kubernetes clusters. Remember to follow best practices to keep your secrets safe and secure. Happy Kubernetes-ing!


This wraps up the Secrets Management section, [Day 42](day42.md) We can get into into some programming language learning specifically around Python. 