@echo off

git add .
git commit
git push origin develop

git checkout master
git merge develop

for /f %%a in ('git describe --tags --abbrev^=0 origin/master') do (
  set LATEST_TAG=%%a
)

if defined LATEST_TAG (
  set /p Tag="Enter Tag (Latest: %LATEST_TAG%): "
) else (
  set /p Tag="Enter Tag (E.g: v1.2.3): "
)

git tag %Tag%
git push origin master --tags

git checkout develop
