# Tools needed

1. download and install minikube
   from [https://minikube.sigs.k8s.io/docs/start/](https://minikube.sigs.k8s.io/docs/start/)

# Machine Preparation

You will need to edit the local DNS resolver. On Linux / MacOS it's straight forward, you will have to figure it out on
your own on Windows

1. run ```minikube ip```
2. copy/paste the IP address you get.
3. edit /etc/hosts like so: ```sudo nano /etc/hosts```
4. add an entry at the end of /etc/hosts ```192.168.99.109  api.k8s.local```
   - if you are on a mac m1, then add ```127.0.0.1  api.k8s.local```
5. make sure to use the correct IP address
6. save and exit
7. run ```minikube addons enable ingress```

# Running the backend locally

1. run ```minikube start``` in case minikube is not already running
   - if you are on a mac m1, then in a new tab run `minikube tunnel` 
2. go to the location where you cloned the `k8s-tutorial` repository
3. apply all yaml files under `aggregator/deployment/kubernetes`, `movies/deployment/kubernetes`
   , `weather/deployment/kubernetes` and `k8s/ingress.yaml` using for example
```kubectl apply -f k8s/ingress.yaml```
4. once the build has completed open your browser and go to ```http://api.k8s.local/aggregator/all```
