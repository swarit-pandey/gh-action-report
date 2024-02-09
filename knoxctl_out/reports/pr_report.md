# Report [2024-02-09 06:22:04 UTC]


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
<details>
<summary>Egress Connections</summary>

<table><tr><th>Protocol</th><th>Command</th><th>POD/SVC/IP</th><th>Port</th><th>Namespace</th><th>Type</th></tr><tr><td>
<code>UDP</code>
</td><td>
<code>/usr/local/lib/erlang/erts-13.2.2.2/bin/erlexec</code>
</td><td>
<code>N/A</code>
</td><td>
<code>N/A</code>
</td><td>
<code>N/A</code>
</td><td>
<code>N/A</code>
</td></tr>
<tr><td>

```diff
+ UDP	 
```
</td><td>
<code>/usr/local/lib/erlang/erts-13.2.2.2/bin/erlexec</code>
</td><td>
<code>N/A</code>
</td><td>
<code>N/A</code>
</td><td>
<code>N/A</code>
</td><td>
<code>N/A</code>
</td></tr>
<tr><td>
<code>TCP</code>
</td><td>

```diff
+ /usr/local/lib/erlang/erts-13.2.2.2/bin/erlexec	 
```
</td><td>
<code>127.0.0.1</code>
</td><td>
<code>4369</code>
</td><td>
<code>N/A</code>
</td><td>
<code>N/A</code>
</td></tr>
<tr><td>
<code>TCP</code>
</td><td>
<code>/usr/local/lib/erlang/erts-13.2.2.2/bin/erlexec</code>
</td><td>

```diff
+ 127.0.0.1	 
```
</td><td>
<code>4369</code>
</td><td>
<code>N/A</code>
</td><td>
<code>N/A</code>
</td></tr>
<tr><td>
<code>TCP</code>
</td><td>
<code>/usr/local/lib/erlang/erts-13.2.2.2/bin/erlexec</code>
</td><td>
<code>127.0.0.1</code>
</td><td>

```diff
+ 4369	 
```
</td><td>
<code>N/A</code>
</td><td>
<code>N/A</code>
</td></tr>
<tr><td>

```diff
+ TCP	 
```
</td><td>
<code>/usr/local/lib/erlang/erts-13.2.2.2/bin/erlexec</code>
</td><td>
<code>127.0.0.1</code>
</td><td>
<code>4369</code>
</td><td>
<code>N/A</code>
</td><td>
<code>N/A</code>
</td></tr>
</table>

</details>
<details>
<summary>Bind Connections</summary>

<table><tr><th>Protocol</th><th>Command</th><th>Bind Port</th><th>Bind Address</th></tr><tr><td>
<code>AF_INET</code>
</td><td>
<code>/usr/local/lib/erlang/erts-13.2.2.2/bin/erlexec</code>
</td><td>
<code>N/A</code>
</td><td>

```diff
+ 0.0.0.0	 
```
</td></tr><tr><td>
<code>AF_INET</code>
</td><td>
<code>/usr/local/lib/erlang/erts-13.2.2.2/bin/erlexec</code>
</td><td>
<code>N/A</code>
</td><td>

```diff
+ 0	 
```
</td></tr><tr><td>

```diff
+ AF_INET	 
```
</td><td>
<code>/usr/local/lib/erlang/erts-13.2.2.2/bin/erlexec</code>
</td><td>
<code>N/A</code>
</td><td>
<code>0.0.0.0</code>
</td></tr><tr><td>
<code>AF_INET</code>
</td><td>

```diff
+ /usr/local/lib/erlang/erts-13.2.2.2/bin/erlexec	 
```
</td><td>
<code>N/A</code>
</td><td>
<code>0.0.0.0</code>
</td></tr></table>

</details>
<hr>
