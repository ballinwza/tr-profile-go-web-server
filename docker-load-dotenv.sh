#!/bin/bash

# Run chmod +x load_env.sh ก่อนเพื่อให้permission to run shell script
# ตรวจสอบว่ามีไฟล์ .env หรือไม่
if [ -f .env ]; then
    echo "Reading .env file..."

    # อ่านไฟล์ .env และแสดงข้อมูลบรรทัดทีละบรรทัด
    while IFS= read -r line || [[ -n "$line" ]]; do
        # Skip empty lines and lines starting with '#' (comments)
        if [[ -z "$line" || "$line" =~ ^# ]]; then
            continue
        fi

        # Check if the line contains a valid environment variable (KEY=value)
        if [[ "$line" =~ ^[A-Za-z_][A-Za-z0-9_]*= ]]; then
            export "$line"
            echo "Exported: $line"
        else
            echo "Skipping invalid line: $line"
        fi
    done < .env

    echo "Finished reading .env file."
else
    echo ".env file not found."
fi