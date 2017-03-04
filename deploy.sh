#!/bin/bash
echo "Deploying EduNav"
ls -l ./main
sshpass -e scp ./main edunav@edunav.eyskens.me:/opt/edunav/