# Report [2024-02-08 17:06:37 UTC]


### Workload Information 
>**Cluster**: default  
>**Namespace**: cache  
>**Type/Name**: deployment/redis  

<details>
<summary>Process/File Summary</summary>

<table><tr><th>Source Path</th><th>Destination Path</th><th>Status</th></tr><tr><td>
<code>/var/lib/rancher/k3s/data/bf3548384eaabb3435bf08112f1b0cba1afc5add6a6f2f2372aa2906a598fd04/bin/containerd-shim-runc-v2</code>
</td><td>

```diff
- /usr/local/bin/not_redis_server	 
```
</td><td>Allow</td></tr><tr><td>
<code>/var/lib/rancher/k3s/data/bf3548384eaabb3435bf08112f1b0cba1afc5add6a6f2f2372aa2906a598fd04/bin/containerd-shim-runc-v2</code>
</td><td>

```diff
+ /usr/local/bin/redis-server	 
```
</td><td>Allow</td></tr></table>

</details>
<hr>

### Workload Information 
>**Cluster**: default  
>**Namespace**: accuknox-agents  
>**Type/Name**: deployment/discovery-engine  

<details>
<summary>Process/File Summary</summary>

<table><tr><th>Source Path</th><th>Destination Path</th><th>Status</th></tr><tr><td>

```diff
- /usr/bin/not_dash	 
```
</td><td>
<code>/usr/bin/basename</code>
</td><td>Allow</td></tr><tr><td>

```diff
+ /usr/bin/dash	 
```
</td><td>
<code>/usr/bin/basename</code>
</td><td>Allow</td></tr></table>

</details>
<hr>

### Workload Information 
>**Cluster**: default  
>**Namespace**: default  
>**Type/Name**: deployment/nginx  

<details>
<summary>Process/File Summary</summary>

<table><tr><th>Source Path</th><th>Destination Path</th><th>Status</th></tr><tr><td>
<code></code>
</td><td>

```diff
- /usr/bin/not_cut	 
```
</td><td>Allow</td></tr><tr><td>
<code></code>
</td><td>

```diff
+ /usr/bin/cut	 
```
</td><td>Allow</td></tr></table>

</details>
<hr>
