```
postmanctl on  master via 🐹 v1.14
➜ USER_ID=$(postmanctl get user -o jsonpath="{.id}") && \
➜ postmanctl get monitors -o jsonpath="{[?(@.owner == '$USER_ID')].id}" | \
➜ xargs postmanctl get monitor -o json
```
