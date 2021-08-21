# eud_backup

This repository contains a simple app which allows us to backup our databases daily.

## Setup
<li>
    Renam <copy>.env.example</code> to <code>.env</code> and change the environment variables.
</li>
<li>
    Start the app by running the following command - <code>go run cmd/backup/main.go</code>.
</li>
<li>
    To decrypt the file, change directory to <code>scripts</code>. Run <code>go run decrypt_file.go FILE_PATH</code> (change FILE_PATH to the binary file eg. <code>go run decrypt_file.go ../database_backup.bin</code>).
</li>
