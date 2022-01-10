# sample-operator

## Operator Pods
```
% oc get all -n sample-operator-system
NAME                                      READY   STATUS    RESTARTS   AGE
pod/controller-manager-7d98658766-7t6rk   2/2     Running   0          87s

NAME                                         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
service/controller-manager-metrics-service   ClusterIP   172.21.88.128   <none>        8443/TCP   89s

NAME                                 READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/controller-manager   1/1     1            1           88s

NAME                                            DESIRED   CURRENT   READY   AGE
replicaset.apps/controller-manager-7d98658766   1         1         1       89s
```
### Pod Labels:
```
% oc get pods -n sample-operator-system --show-labels
NAME                                  READY   STATUS    RESTARTS   AGE    LABELS
controller-manager-7d98658766-7t6rk   2/2     Running   0          101s   control-plane=controller-manager,pod-template-hash=7d98658766
```

## Creation of LabelOperator CR:
```
% oc create -f config/samples/multiple_v1alpha1_labeloperator.yaml
labeloperator.multiple.example.com/labeloperator-sample created

% oc get labeloperator.multiple.example.com -A
NAMESPACE                NAME                   AGE
sample-operator-system   labeloperator-sample   6s
```

## Logs:
```
% oc logs -f -n sample-operator-system pod/controller-manager-7d98658766-7t6rk -c manager
I0110 03:39:06.817415       1 request.go:668] Waited for 1.017508574s due to client-side throttling, not priority and fairness, request: GET:https://172.21.0.1:443/apis/snapshot.storage.k8s.io/v1?timeout=32s
2022-01-10T03:39:08.199Z	INFO	controller-runtime.metrics	metrics server is starting to listen	{"addr": "127.0.0.1:8080"}
2022-01-10T03:39:08.200Z	INFO	setup	starting manager
I0110 03:39:08.200889       1 leaderelection.go:243] attempting to acquire leader lease sample-operator-system/f3bbe435.example.com...
2022-01-10T03:39:08.200Z	INFO	controller-runtime.manager	starting metrics server	{"path": "/metrics"}
I0110 03:39:08.335904       1 leaderelection.go:253] successfully acquired lease sample-operator-system/f3bbe435.example.com
2022-01-10T03:39:08.336Z	INFO	controller-runtime.manager.controller.labeloperator	Starting EventSource	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "source": "kind source: /, Kind="}
2022-01-10T03:39:08.336Z	INFO	controller-runtime.manager.controller.labeloperator	Starting Controller	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator"}
2022-01-10T03:39:08.336Z	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"ConfigMap","namespace":"sample-operator-system","name":"f3bbe435.example.com","uid":"61a53410-6c57-4364-b0e5-914677246a76","apiVersion":"v1","resourceVersion":"72397551"}, "reason": "LeaderElection", "message": "controller-manager-7d98658766-7t6rk_7beb4df0-7b56-45ce-bf46-eb5bc4c92cbe became leader"}
2022-01-10T03:39:08.336Z	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"Lease","namespace":"sample-operator-system","name":"f3bbe435.example.com","uid":"0f63139c-3b17-4f40-b37f-33d680528e34","apiVersion":"coordination.k8s.io/v1","resourceVersion":"72397553"}, "reason": "LeaderElection", "message": "controller-manager-7d98658766-7t6rk_7beb4df0-7b56-45ce-bf46-eb5bc4c92cbe became leader"}
2022-01-10T03:39:08.336Z	INFO	controller-runtime.manager.controller.replicasoperator	Starting EventSource	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator", "source": "kind source: /, Kind="}
2022-01-10T03:39:08.336Z	INFO	controller-runtime.manager.controller.replicasoperator	Starting Controller	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator"}
2022-01-10T03:39:08.537Z	INFO	controller-runtime.manager.controller.labeloperator	Starting workers	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "worker count": 1}
2022-01-10T03:39:08.537Z	INFO	controller-runtime.manager.controller.replicasoperator	Starting workers	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator", "worker count": 1}
2022-01-10T03:40:59.361Z	INFO	controller-runtime.manager.controller.labeloperator	In LabelOperator......	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "name": "labeloperator-sample", "namespace": "sample-operator-system", "namespace:": "sample-operator-system"}
2022-01-10T03:40:59.863Z	INFO	controller-runtime.manager.controller.labeloperator	Get PodsList....	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "name": "labeloperator-sample", "namespace": "sample-operator-system"}
2022-01-10T03:40:59.863Z	INFO	controller-runtime.manager.controller.labeloperator	Pod Found	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "name": "labeloperator-sample", "namespace": "sample-operator-system", "name:": "controller-manager-7d98658766-7t6rk"}
2022-01-10T03:40:59.863Z	INFO	controller-runtime.manager.controller.labeloperator	Label Added in Pod	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "name": "labeloperator-sample", "namespace": "sample-operator-system", "name:": "controller-manager-7d98658766-7t6rk", "label": "my-custom-label"}
2022-01-10T03:40:59.922Z	INFO	controller-runtime.manager.controller.labeloperator	Pod Label Added....	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "name": "labeloperator-sample", "namespace": "sample-operator-system", "name": "controller-manager-7d98658766-7t6rk"}
```

### Pod Labels:
```
% oc get pods -n sample-operator-system --show-labels
NAME                                  READY   STATUS    RESTARTS   AGE     LABELS
controller-manager-7d98658766-7t6rk   2/2     Running   0          2m18s   control-plane=controller-manager,new-label-by-operator=my-custom-label,pod-template-hash=7d98658766
```

## Creation of ReplicasOperator CR
```
% oc create -f config/samples/multiple_v1alpha1_replicasoperator.yaml 
replicasoperator.multiple.example.com/replicasoperator-sample created

% oc get replicasoperator.multiple.example.com -A
NAMESPACE                NAME                      AGE
sample-operator-system   replicasoperator-sample   6s
```

#### Resulting Pods:
```
% oc get pods -n sample-operator-system --show-labels
NAME                                            READY   STATUS             RESTARTS   AGE   LABELS
controller-manager-7d98658766-7t6rk             2/2     Running            0          3m    control-plane=controller-manager,new-label-by-operator=my-custom-label,pod-template-hash=7d98658766
replicas-operator-deployment-54cbdd955b-4ckz7   0/1     CrashLoopBackOff   1          12s   ReplicasOperator_cr=replicasoperator-sample,app=replicas,pod-template-hash=54cbdd955b
replicas-operator-deployment-54cbdd955b-7wqt6   0/1     CrashLoopBackOff   1          12s   ReplicasOperator_cr=replicasoperator-sample,app=replicas,pod-template-hash=54cbdd955b
replicas-operator-deployment-54cbdd955b-tv5sq   0/1     CrashLoopBackOff   1          12s   ReplicasOperator_cr=replicasoperator-sample,app=replicas,pod-template-hash=54cbdd955b
```

## Logs:
```
% oc logs -f -n sample-operator-system pod/controller-manager-7d98658766-7t6rk -c manager
I0110 03:39:06.817415       1 request.go:668] Waited for 1.017508574s due to client-side throttling, not priority and fairness, request: GET:https://172.21.0.1:443/apis/snapshot.storage.k8s.io/v1?timeout=32s
2022-01-10T03:39:08.199Z	INFO	controller-runtime.metrics	metrics server is starting to listen	{"addr": "127.0.0.1:8080"}
2022-01-10T03:39:08.200Z	INFO	setup	starting manager
I0110 03:39:08.200889       1 leaderelection.go:243] attempting to acquire leader lease sample-operator-system/f3bbe435.example.com...
2022-01-10T03:39:08.200Z	INFO	controller-runtime.manager	starting metrics server	{"path": "/metrics"}
I0110 03:39:08.335904       1 leaderelection.go:253] successfully acquired lease sample-operator-system/f3bbe435.example.com
2022-01-10T03:39:08.336Z	INFO	controller-runtime.manager.controller.labeloperator	Starting EventSource	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "source": "kind source: /, Kind="}
2022-01-10T03:39:08.336Z	INFO	controller-runtime.manager.controller.labeloperator	Starting Controller	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator"}
2022-01-10T03:39:08.336Z	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"ConfigMap","namespace":"sample-operator-system","name":"f3bbe435.example.com","uid":"61a53410-6c57-4364-b0e5-914677246a76","apiVersion":"v1","resourceVersion":"72397551"}, "reason": "LeaderElection", "message": "controller-manager-7d98658766-7t6rk_7beb4df0-7b56-45ce-bf46-eb5bc4c92cbe became leader"}
2022-01-10T03:39:08.336Z	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"Lease","namespace":"sample-operator-system","name":"f3bbe435.example.com","uid":"0f63139c-3b17-4f40-b37f-33d680528e34","apiVersion":"coordination.k8s.io/v1","resourceVersion":"72397553"}, "reason": "LeaderElection", "message": "controller-manager-7d98658766-7t6rk_7beb4df0-7b56-45ce-bf46-eb5bc4c92cbe became leader"}
2022-01-10T03:39:08.336Z	INFO	controller-runtime.manager.controller.replicasoperator	Starting EventSource	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator", "source": "kind source: /, Kind="}
2022-01-10T03:39:08.336Z	INFO	controller-runtime.manager.controller.replicasoperator	Starting Controller	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator"}
2022-01-10T03:39:08.537Z	INFO	controller-runtime.manager.controller.labeloperator	Starting workers	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "worker count": 1}
2022-01-10T03:39:08.537Z	INFO	controller-runtime.manager.controller.replicasoperator	Starting workers	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator", "worker count": 1}
2022-01-10T03:40:59.361Z	INFO	controller-runtime.manager.controller.labeloperator	In LabelOperator......	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "name": "labeloperator-sample", "namespace": "sample-operator-system", "namespace:": "sample-operator-system"}
2022-01-10T03:40:59.863Z	INFO	controller-runtime.manager.controller.labeloperator	Get PodsList....	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "name": "labeloperator-sample", "namespace": "sample-operator-system"}
2022-01-10T03:40:59.863Z	INFO	controller-runtime.manager.controller.labeloperator	Pod Found	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "name": "labeloperator-sample", "namespace": "sample-operator-system", "name:": "controller-manager-7d98658766-7t6rk"}
2022-01-10T03:40:59.863Z	INFO	controller-runtime.manager.controller.labeloperator	Label Added in Pod	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "name": "labeloperator-sample", "namespace": "sample-operator-system", "name:": "controller-manager-7d98658766-7t6rk", "label": "my-custom-label"}
2022-01-10T03:40:59.922Z	INFO	controller-runtime.manager.controller.labeloperator	Pod Label Added....	{"reconciler group": "multiple.example.com", "reconciler kind": "LabelOperator", "name": "labeloperator-sample", "namespace": "sample-operator-system", "name": "controller-manager-7d98658766-7t6rk"}
2022-01-10T03:41:47.638Z	INFO	controller-runtime.manager.controller.replicasoperator	In ReplicasOperator......	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator", "name": "replicasoperator-sample", "namespace": "sample-operator-system"}
2022-01-10T03:41:47.939Z	INFO	controller-runtime.manager.controller.replicasoperator	Creating a new Deployment	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator", "name": "replicasoperator-sample", "namespace": "sample-operator-system", "Deployment.Namespace": "sample-operator-system", "Deployment.Name": "replicas-operator-deployment"}
2022-01-10T03:41:47.996Z	INFO	controller-runtime.manager.controller.replicasoperator	Deployment created successfully!!	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator", "name": "replicasoperator-sample", "namespace": "sample-operator-system"}
2022-01-10T03:41:48.002Z	INFO	controller-runtime.manager.controller.replicasoperator	In ReplicasOperator......	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator", "name": "replicasoperator-sample", "namespace": "sample-operator-system"}
2022-01-10T03:41:48.012Z	INFO	controller-runtime.manager.controller.replicasoperator	Deployment has created desired number of replicas of pod	{"reconciler group": "multiple.example.com", "reconciler kind": "ReplicasOperator", "name": "replicasoperator-sample", "namespace": "sample-operator-system"}
```
