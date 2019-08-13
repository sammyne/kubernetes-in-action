# Head-Ups for macOS
- Install Docker Desktop for mac, not by brew  
- Install minikube with 
    ```bash
    curl -LO https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-hyperkit \
        && sudo install -o root -g wheel -m 4755 docker-machine-driver-hyperkit /usr/local/bin/
    brew install minikube
    ```

# DEPRECATION NOTICES 

|                 old | new                     |
| ------------------: | :---------------------- |
| `--generate=run/v1` | `--generate=run-pod/v1` |