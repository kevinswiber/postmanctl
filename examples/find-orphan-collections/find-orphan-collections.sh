set -x

# Get all workspace IDs.
postmanctl get ws -o json | jq -r ".[].id" > workspaces.txt

# Get all collection UIDs.
postmanctl get co -o json | jq -r ".[].uid" > collections.txt

# Iterate over all workspaces, and record all associated collections.
cat workspaces.txt | while read workspace_id; do sh -c "postmanctl get ws $workspace_id -o json | jq -r '.[].collections[].uid' >> workspace_collections.txt"; done

# Sort data in files.
sort -o workspace_collections.txt workspace_collections.txt
sort -o collections.txt collections.txt

# Compare files and show only collections that exist in the complete collections list yet not in the list of collections associated with workspaces.
comm -23 collections.txt workspace_collections.txt
