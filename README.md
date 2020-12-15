# eud_backup

This repository contains a simple app which allows us to backup our databases daily.

## Setup
<li>
    Add environment variables into <code>.env</code> file
</li>
<li>
    Start backup by running <code>go run cmd/backup/main.go database1 database2</code> (change database1 database2 to actual database names (eg. website, central))
</li>